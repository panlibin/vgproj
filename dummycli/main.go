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
	if client.GlobalConfig.ClientType == "Game" {
		for i := 0; i < client.GlobalConfig.Count; i++ {
			// pClient := client.NewClient(iaccount.LoginTypeCustom, fmt.Sprintf("%s%04d", client.GlobalConfig.Prefix, i), "ababcc", &wg)
			pClient := client.NewGameClient(10001+int64(i), &wg)
			// pClient := client.NewClient(iaccount.LoginTypeCustom, "test0115", "ababcc", &wg)
			go pClient.Run()
		}
		if client.GlobalConfig.Interval > 0 {
			time.Sleep(time.Millisecond * time.Duration(client.GlobalConfig.Interval))
		}
	} else {
		for i := 0; i < client.GlobalConfig.Count; i++ {
			pClient := client.NewClient(iaccount.LoginTypeCustom, fmt.Sprintf("%s%04d", client.GlobalConfig.Prefix, i), "ababcc", &wg)
			go pClient.Run()
		}
		if client.GlobalConfig.Interval > 0 {
			time.Sleep(time.Millisecond * time.Duration(client.GlobalConfig.Interval))
		}
	}

	wg.Wait()
	logger.Debug(time.Now().Sub(tm))
	logger.DefaultLogger.Flush()
}
