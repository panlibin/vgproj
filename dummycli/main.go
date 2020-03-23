package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"vgproj/dummycli/client"
	iaccount "vgproj/vglogin/public/account"
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

	for i := 0; i < 1; i++ {
		pClient := client.NewClient(iaccount.LoginTypeCustom, fmt.Sprintf("test%04d", i), "ababcc", &wg)
		//pClient := client.NewClient(account.LoginType_Custom, "test0000", "ababcc", &wg)
		go pClient.Run()
	}

	wg.Wait()
}
