package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
	"vgproj/dummycli/client"
	iaccount "vgproj/vglogin/public/account"

	logger "github.com/panlibin/vglog"
)

func main() {
	cfg, err := ioutil.ReadFile("./config.json")
	if err != nil {
		return
	}
	err = json.Unmarshal(cfg, &client.GlobalConfig)
	if err != nil {
		return
	}

	http.DefaultClient.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	wg := sync.WaitGroup{}

	tm := time.Now()
	for i := 0; i < 100; i++ {
		for j := 0; j < 10; j++ {
			pClient := client.NewClient(iaccount.LoginTypeCustom, fmt.Sprintf("test%04d", i*10+j), "ababcc", &wg)
			//pClient := client.NewClient(account.LoginType_Custom, "test0000", "ababcc", &wg)
			go pClient.Run()
		}
		time.Sleep(time.Millisecond * 10)
	}

	wg.Wait()
	logger.Debug(time.Now().Sub(tm))
}
