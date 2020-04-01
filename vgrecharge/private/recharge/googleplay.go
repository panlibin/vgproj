package recharge

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"errors"
)

var (
	ErrGooglePlayMissingParam        = errors.New("[GooglePlay] missing parameter!")
	ErrGooglePlayProductIdMismatch   = errors.New("[GooglePlay] product id mismatch!")
	ErrGooglePlayPackageNameMismatch = errors.New("[GooglePlay] package name mismatch!")
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

func (gp *googlePlay) verify(accountId int64, serverId int32, playerId int64, pfProductId string, jsonParams []byte) error {
	params := map[string]string{}
	err := json.Unmarshal(jsonParams, &params)
	if err != nil {
		return err
	}

	iapData, exist := params["iapdata"]
	if !exist {
		return ErrGooglePlayMissingParam
	}
	sign, exist := params["sign"]
	if !exist {
		return ErrGooglePlayMissingParam
	}
	// ext, exist := params["ext"]
	// if !exist {
	// 	return ErrGooglePlayMissingParam
	// }
	// purchaseToken, exist := params["purchase_token"]
	// if !exist {
	// 	return ErrGooglePlayMissingParam
	// }

	byteIapData := []byte(iapData)
	hashedData := sha1.Sum(byteIapData)
	decodedSign, err := base64.StdEncoding.DecodeString(sign)
	if err != nil {
		return err
	}
	err = rsa.VerifyPKCS1v15(gp.publicKey, crypto.SHA1, hashedData[:], decodedSign)
	if err != nil {
		return err
	}

	pIAPData := &GooglePlayIAPData{}
	err = json.Unmarshal(byteIapData, pIAPData)
	if err != nil {
		return err
	}

	if pIAPData.ProductId != pfProductId {
		return ErrGooglePlayProductIdMismatch
	}
	if pIAPData.PackageName != gp.packageName {
		return ErrGooglePlayPackageNameMismatch
	}

	return nil
}
