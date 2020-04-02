package recharge

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"errors"
	"time"
)

var (
	ErrGooglePlayMissingParam        = errors.New("[GooglePlay] missing parameter!")
	ErrGooglePlayProductIdMismatch   = errors.New("[GooglePlay] product id mismatch!")
	ErrGooglePlayPackageNameMismatch = errors.New("[GooglePlay] package name mismatch!")
	ErrGooglePlayDuplicateOrder      = errors.New("[GooglePlay] duplicate order!")
)

type googlePlay struct {
	*platform
	publicKey   *rsa.PublicKey
	packageName string
}

func newGooglePlay(sp *sdkParam) *googlePlay {
	decodedKey, err := base64.StdEncoding.DecodeString(sp.keys[1])
	if nil != err {
		return nil
	}
	pubInterface, err := x509.ParsePKIXPublicKey(decodedKey)
	if nil != err {
		return nil
	}
	return &googlePlay{
		platform:    newPlatform(),
		publicKey:   pubInterface.(*rsa.PublicKey),
		packageName: sp.keys[0],
	}
}

type GooglePlayIAPData struct {
	OrderId          string `json:"orderId"`
	PackageName      string `json:"packageName"`
	ProductId        string `json:"productId"`
	PurchaseState    int32  `json:"purchaseState"`
	DeveloperPayload string `json:"developerPayload"`
}

func (gp *googlePlay) verifyAndCreateOrder(currency string, amount int64, pfProductId string, localProductId int32, accountId int64, serverId int32, playerId int64, jsonParams []byte) (*order, error) {
	params := map[string]string{}
	if err := json.Unmarshal(jsonParams, &params); err != nil {
		return nil, err
	}

	iapData, exist := params["iapdata"]
	if !exist {
		return nil, ErrGooglePlayMissingParam
	}
	sign, exist := params["sign"]
	if !exist {
		return nil, ErrGooglePlayMissingParam
	}
	ext, exist := params["ext"]
	if !exist {
		return nil, ErrGooglePlayMissingParam
	}
	// purchaseToken, exist := params["purchase_token"]
	// if !exist {
	// 	return nil, ErrGooglePlayMissingParam
	// }

	byteIapData := []byte(iapData)
	hashedData := sha1.Sum(byteIapData)
	decodedSign, err := base64.StdEncoding.DecodeString(sign)
	if err != nil {
		return nil, err
	}
	err = rsa.VerifyPKCS1v15(gp.publicKey, crypto.SHA1, hashedData[:], decodedSign)
	if err != nil {
		return nil, err
	}

	pIAPData := &GooglePlayIAPData{}
	err = json.Unmarshal(byteIapData, pIAPData)
	if err != nil {
		return nil, err
	}

	if pIAPData.ProductId != pfProductId {
		return nil, ErrGooglePlayProductIdMismatch
	}
	if pIAPData.PackageName != gp.packageName {
		return nil, ErrGooglePlayPackageNameMismatch
	}

	o := &order{
		localOrderId:   genOrderId(),
		pfId:           platformGooglePlay,
		pfOrderId:      ext,
		receiveDate:    time.Now(),
		source:         "GooglePlay",
		currency:       currency,
		amount:         amount,
		pfProductId:    pfProductId,
		localProductId: localProductId,
		accountId:      accountId,
		serverId:       serverId,
		playerId:       playerId,
		status:         orderStatusWaitDeliver,
		sandbox:        0,
	}

	o.mtx.Lock()
	defer o.mtx.Unlock()
	if !gp.insertOrder(o) {
		return nil, ErrGooglePlayDuplicateOrder
	}
	if err = o.insert(); err != nil {
		gp.removeOrder(o)
		return nil, err
	}

	return o, nil
}
