package account

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	logger "github.com/panlibin/vglog"
)

type FacebookLoginRspError struct {
	Message string `json:"message"`
}

type FacebookLoginRspData struct {
	AppId       string                 `json:"app_id"`
	Application string                 `json:"application"`
	ExpiresAt   int64                  `json:"expires_at"`
	IsValid     bool                   `json:"is_valid"`
	IssuedAt    int64                  `json:"issued_at"`
	UserId      string                 `json:"user_id"`
	Error       *FacebookLoginRspError `json:"error"`
}

type FacebookLoginRsp struct {
	Data  *FacebookLoginRspData  `json:"data"`
	Error *FacebookLoginRspError `json:"error"`
}

type Facebook struct {
	appId  string
	appKey string
}

func NewFacebook() *Facebook {
	pObj := new(Facebook)
	return pObj
}

func (fb *Facebook) SetPlatformParam(appId string, appKey string) {
	fb.appId = appId
	fb.appKey = appKey
}

func (fb *Facebook) login(token string) (accountName string, err error) {
	for {
		strUrl := fmt.Sprintf("https://graph.facebook.com/v3.2/debug_token?input_token=%v&access_token=%v|%v", token, fb.appId, fb.appKey)
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

		rsp := FacebookLoginRsp{}
		if err = json.Unmarshal(body, &rsp); err != nil {
			break
		}

		if rsp.Error != nil {
			err = errors.New(rsp.Error.Message)
			break
		}

		if rsp.Data == nil {
			err = errors.New("rsp.Data = nil")
			break
		}

		if rsp.Data.Error != nil {
			err = errors.New(rsp.Data.Error.Message)
			break
		}

		if !rsp.Data.IsValid {
			err = errors.New("rsp.Data.Is_valid = false")
			break
		}

		if rsp.Data.AppId != fb.appId {
			err = errors.New("rsp.Data.App_id != fb.appId")
			break
		}

		if rsp.Data.ExpiresAt <= time.Now().Unix() {
			err = errors.New("token expire")
			break
		}

		if rsp.Data.UserId == "" {
			err = errors.New("rsp.Data.UserId == \"\"")
			break
		}

		accountName = rsp.Data.UserId

		break
	}

	if err != nil {
		logger.Errorf("[Facebook] 登录失败: %v", err)
	}
	return
}
