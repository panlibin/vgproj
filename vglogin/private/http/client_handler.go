package http

import (
	"encoding/json"
	"net/http"
	"strconv"
	ec "vgproj/common/define/err_code"
	"vgproj/vglogin/public"
	iaccount "vgproj/vglogin/public/account"
)

func handleRegister(w http.ResponseWriter, pReq *http.Request) {
	var errCode int32 = ec.Unknown

	for {
		strName := pReq.FormValue("account_name")
		strPwd := pReq.FormValue("password")
		strTime := pReq.FormValue("time")
		strSign := pReq.FormValue("sign")

		reqTime, err := strconv.ParseInt(strTime, 10, 64)
		if err != nil {
			errCode = ec.InvalidParam
			break
		}
		if !checkTime(reqTime) {
			errCode = ec.RequestTimeout
			break
		}

		if !checkSign([]string{strName, strPwd, strTime}, public.Server.GetClientKey(), strSign) {
			errCode = ec.InvalidSign
			break
		}

		errCode = public.Server.GetAccountManager().Register(strName, strPwd)

		break
	}

	rsp := make(map[string]interface{}, 1)
	rsp["code"] = errCode

	rspBuf, _ := json.Marshal(rsp)
	w.Write(rspBuf)
}

func handleLogin(w http.ResponseWriter, pReq *http.Request) {
	var errCode int32 = ec.Unknown
	var pAccount iaccount.IAccount

	rsp := make(map[string]interface{}, 8)

	for {
		strLoginType := pReq.FormValue("login_type")
		strName := pReq.FormValue("account_name")
		strPwd := pReq.FormValue("password")
		strTime := pReq.FormValue("time")
		strSign := pReq.FormValue("sign")

		reqTime, err := strconv.ParseInt(strTime, 10, 64)
		if err != nil {
			errCode = ec.InvalidParam
			break
		}
		if !checkTime(reqTime) {
			errCode = ec.RequestTimeout
			break
		}

		if !checkSign([]string{strLoginType, strName, strPwd, strTime}, public.Server.GetClientKey(), strSign) {
			errCode = ec.InvalidSign
			break
		}

		loginType, err := strconv.Atoi(strLoginType)
		if err != nil {
			errCode = ec.InvalidParam
			break
		}

		pAccount, errCode = public.Server.GetAccountManager().Login(int32(loginType), strName, strPwd)
		if errCode == ec.Success && pAccount != nil {
			if pAccount.Lock() != nil {
				errCode = ec.Unknown
				break
			}
			defer pAccount.Unlock()

			if pAccount.IsBan() {
				errCode = ec.AccountBanned
				rsp["is_ban"] = pAccount.IsBan()
				rsp["ban_type"] = pAccount.GetBanType()
				rsp["ban_duration"] = pAccount.GetBanDuration()
				rsp["ban_ts"] = pAccount.GetBanTs()
			} else {
				rsp["account_id"] = pAccount.GetId()
				rsp["token"] = pAccount.GetToken()
				rsp["expire_ts"] = pAccount.GetTokenExpireTs()
			}
		}

		break
	}

	rsp["code"] = errCode

	rspBuf, _ := json.Marshal(rsp)
	w.Write(rspBuf)
}
