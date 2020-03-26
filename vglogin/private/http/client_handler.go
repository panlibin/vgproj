package http

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	ec "vgproj/common/define/err_code"
	"vgproj/vglogin/public"
	iaccount "vgproj/vglogin/public/account"

	"github.com/panlibin/virgo/util/vgtime"
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
				curTs := vgtime.Now()
				token := sha256.Sum256([]byte(fmt.Sprintf("ts=%d&auth_key=%s&account_id=%d", curTs, public.Server.GetAuthKey(), pAccount.GetId())))
				rsp["token"] = base64.URLEncoding.EncodeToString(token[:])
				rsp["ts"] = curTs
				mapCharacter := pAccount.GetCharacters()
				cList := make(map[int32]interface{}, len(mapCharacter))
				for serverId, pCharacter := range mapCharacter {
					cList[serverId] = map[string]interface{}{
						"id":        pCharacter.GetPlayerId(),
						"name":      pCharacter.GetName(),
						"combat":    pCharacter.GetCombat(),
						"update_ts": pCharacter.GetUpdateTs(),
					}
				}
				rsp["characters"] = cList
			}
		}

		break
	}

	rsp["code"] = errCode

	rspBuf, _ := json.Marshal(rsp)
	w.Write(rspBuf)
}

func handleServerList(w http.ResponseWriter, pReq *http.Request) {
	var errCode int32 = ec.Unknown

	rsp := make(map[string]interface{}, 8)

	for {
		// strTime := pReq.FormValue("time")
		// strSign := pReq.FormValue("sign")

		// reqTime, err := strconv.ParseInt(strTime, 10, 64)
		// if err != nil {
		// 	errCode = ec.InvalidParam
		// 	break
		// }
		// if !checkTime(reqTime) {
		// 	errCode = ec.RequestTimeout
		// 	break
		// }

		// if !checkSign([]string{strTime}, public.Server.GetClientKey(), strSign) {
		// 	errCode = ec.InvalidSign
		// 	break
		// }

		pGameServerManager := public.Server.GetGameServerManager()
		mapServer := pGameServerManager.GrabServerList()
		defer pGameServerManager.ReleaseServerList()
		rsp["server"] = mapServer

		errCode = ec.Success

		break
	}

	rsp["code"] = errCode
	rspBuf, _ := json.Marshal(rsp)

	w.Write(rspBuf)
}
