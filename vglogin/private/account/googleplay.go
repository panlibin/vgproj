package account

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	logger "github.com/panlibin/vglog"
)

type GooglePlayLoginRsp struct {
	ErrorDescription string `json:"error_description"`
	Aud              string `json:"aud"`
	Sub              string `json:"sub"`
	Exp              string `json:"exp"`
}

type GooglePlay struct {
	appId string
}

func NewGooglePlay() *GooglePlay {
	pObj := new(GooglePlay)
	return pObj
}

func (gp *GooglePlay) SetPlatformParam(appId string) {
	gp.appId = appId
}

func (gp *GooglePlay) login(token string) (accountName string, err error) {
	for {
		strUrl := fmt.Sprintf("https://www.googleapis.com/oauth2/v3/tokeninfo?id_token=%v", token)
		var resp *http.Response
		if resp, err = http.Get(strUrl); err != nil {
			break
		}

		var body []byte
		body, err = ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			break
		}

		rsp := GooglePlayLoginRsp{}
		if err = json.Unmarshal(body, &rsp); err != nil {
			break
		}

		if rsp.ErrorDescription != "" {
			err = errors.New(rsp.ErrorDescription)
			break
		}

		expireTs, err := strconv.ParseInt(rsp.Exp, 10, 64)
		if err != nil {
			break
		}
		if expireTs <= time.Now().Unix() {
			err = errors.New("token expire")
			break
		}

		if !strings.Contains(gp.appId, rsp.Aud) {
			err = errors.New("app id not match")
			break
		}

		if rsp.Sub == "" {
			err = errors.New("rsp.Sub == \"\"")
			break
		}

		accountName = rsp.Sub

		break
	}

	if err != nil {
		logger.Errorf("[GooglePlay] 登录失败: %v", err)
	}

	return
}
