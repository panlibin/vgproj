package recharge

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	appleIAPVerifyUrl        = "https://buy.itunes.apple.com/verifyReceipt"
	appleIAPVerifyUrlSandbox = "https://sandbox.itunes.apple.com/verifyReceipt"
)

var (
	ErrAppleMissingParam          = errors.New("[Apple] missing parameter!")
	ErrAppleReceiptIsNil          = errors.New("[Apple] response receipt is nil!")
	ErrAppleBundleIdMismatch      = errors.New("[Apple] bundle id mismatch!")
	ErrAppleTransactionIdMismatch = errors.New("[Apple] transaction id mismatch!")
	ErrAppleProductIdMismatch     = errors.New("[Apple] product id mismatch!")
	ErrAppleDuplicateOrder        = errors.New("[Apple] duplicate order!")
)

type appleIAP struct {
	*platform
	appId string
}

func newAppleIAP(sp *sdkParam) *appleIAP {
	return &appleIAP{
		platform: newPlatform(),
		appId:    sp.appId,
	}
}

type IapResult struct {
	Status  int         `json:"status"`
	Receipt *IapReceipt `json:"receipt"`
}

type IapReceipt struct {
	OriginalPurchaseDatePst string `json:"original_purchase_date_pst"`
	PurchaseDateMs          string `json:"purchase_date_ms"`
	UniqueIdentifier        string `json:"unique_identifier"`
	OriginalTransactionId   string `json:"original_transaction_id"`
	Bvrs                    string `json:"bvrs"`
	TransactionId           string `json:"transaction_id"`
	Quantity                string `json:"quantity"`
	UniqueVendorIdentifier  string `json:"unique_vendor_identifier"`
	ItemId                  string `json:"item_id"`
	OriginalPurchaseDate    string `json:"original_purchase_date"`
	IsInIntroOfferPeriod    string `json:"is_in_intro_offer_period"`
	ProductId               string `json:"product_id"`
	PurchaseDate            string `json:"purchase_date"`
	IsTrialPeriod           string `json:"is_trial_period"`
	PurchaseDatePst         string `json:"purchase_date_pst"`
	Bid                     string `json:"bid"`
	OriginalPurchaseDateMs  string `json:"original_purchase_date_ms"`

	BundleId string      `json:"bundle_id"`
	InApp    []*IapInApp `json:"in_app"`
}

type IapInApp struct {
	Quantity                string `json:"quantity"`
	ProductId               string `json:"product_id"`
	TransactionId           string `json:"transaction_id"`
	OriginalTransactionId   string `json:"original_transaction_id"`
	PurchaseDate            string `json:"purchase_date"`
	PurchaseDateMs          string `json:"purchase_date_ms"`
	PurchaseDatePst         string `json:"purchase_date_pst"`
	OriginalPurchaseDate    string `json:"original_purchase_date"`
	OriginalPurchaseDateMs  string `json:"original_purchase_date_ms"`
	OriginalPurchaseDatePst string `json:"original_purchase_date_pst"`
	IsTrialPeriod           string `json:"is_trial_period"`
}

func (iap *appleIAP) verifyAndCreateOrder(currency string, amount int64, pfProductId string, localProductId int32, accountId int64, serverId int32, playerId int64, jsonParams []byte) (*order, error) {
	params := map[string]string{}
	if err := json.Unmarshal(jsonParams, &params); err != nil {
		return nil, err
	}

	receipt, exist := params["receipt"]
	if !exist {
		return nil, ErrAppleMissingParam
	}
	transactionId, exist := params["transaction_id"]
	if !exist {
		return nil, ErrAppleMissingParam
	}

	verifyData := map[string]string{}
	verifyData["receipt-data"] = receipt
	verifyBody, err := json.Marshal(&verifyData)
	if err != nil {
		return nil, err
	}

	var sandbox int32 = 0
	verifyResult, err := iap.requestVerify(appleIAPVerifyUrl, verifyBody)
	if err != nil {
		return nil, err
	}

	// 如果是沙盒账号的,则切换地址重新请求
	if verifyResult.Status == 21007 {
		verifyResult, err = iap.requestVerify(appleIAPVerifyUrlSandbox, verifyBody)
		if err != nil {
			return nil, err
		}
		sandbox = 1
	}

	// 状态码 描述
	// 21000 App Store无法读取你提供的JSON数据
	// 21002 收据数据不符合格式
	// 21003 收据无法被验证
	// 21004 你提供的共享密钥和账户的共享密钥不一致
	// 21005 收据服务器当前不可用
	// 21006 收据是有效的，但订阅服务已经过期。当收到这个信息时，解码后的收据信息也包含在返回内容中
	// 21007 收据信息是测试用（sandbox），但却被发送到产品环境中验证
	// 21008 收据信息是产品环境中使用，但却被发送到测试环境中验证
	if verifyResult.Status != 0 {
		return nil, fmt.Errorf("[Apple] verify fail! receipt=%s, transaction_id=%s, status=%d", receipt, transactionId, verifyResult.Status)
	}
	if verifyResult.Receipt == nil {
		return nil, ErrAppleReceiptIsNil
	}

	if len(verifyResult.Receipt.InApp) == 0 {
		if verifyResult.Receipt.Bid != iap.appId {
			return nil, ErrAppleBundleIdMismatch
		}
		if verifyResult.Receipt.TransactionId != transactionId {
			return nil, ErrAppleTransactionIdMismatch
		}
		if verifyResult.Receipt.ProductId != pfProductId {
			return nil, ErrAppleProductIdMismatch
		}
	} else {
		if verifyResult.Receipt.BundleId != iap.appId {
			return nil, ErrAppleBundleIdMismatch
		}
		found := false
		for _, pInAppData := range verifyResult.Receipt.InApp {
			if pInAppData.TransactionId == transactionId {
				if pInAppData.ProductId != pfProductId {
					return nil, ErrAppleProductIdMismatch
				}
				found = true
				break
			}
		}
		if !found {
			return nil, ErrAppleTransactionIdMismatch
		}
	}

	o := &order{
		localOrderId:   genOrderId(),
		pfId:           platformApple,
		pfOrderId:      transactionId,
		receiveDate:    time.Now(),
		source:         "Apple",
		currency:       currency,
		amount:         amount,
		pfProductId:    pfProductId,
		localProductId: localProductId,
		accountId:      accountId,
		serverId:       serverId,
		playerId:       playerId,
		status:         orderStatusWaitDeliver,
		sandbox:        sandbox,
	}

	o.mtx.Lock()
	defer o.mtx.Unlock()
	if !iap.insertOrder(o) {
		return nil, ErrAppleDuplicateOrder
	}
	if err = o.insert(); err != nil {
		iap.removeOrder(o)
		return nil, err
	}

	return o, nil
}

func (iap *appleIAP) requestVerify(url string, reqBody []byte) (*IapResult, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Encoding", "utf-8")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	rspBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	verifyResult := IapResult{}
	err = json.Unmarshal(rspBody, &verifyResult)
	if err != nil {
		return nil, err
	}

	return &verifyResult, nil
}
