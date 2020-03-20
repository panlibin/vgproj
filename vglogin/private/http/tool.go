package http

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
	"vgproj/vglogin/public"

	"github.com/panlibin/virgo/util/vgtime"
)

func checkTime(ts int64) bool {
	if !public.Server.CheckTime() {
		return true
	}
	return ts+300000 > vgtime.Now()
}

func checkSign(params []string, key string, sign string) bool {
	strSrc := fmt.Sprintf("%s%s", strings.Join(params, ""), key)
	sum := md5.Sum([]byte(strSrc))
	checkSum := hex.EncodeToString(append(sum[:]))
	return sign == checkSum
}
