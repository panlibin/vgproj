package main

import (
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
	"vgproj/dummycli/client"

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
	for i := 0; i < client.GlobalConfig.Count; i++ {
		// pClient := client.NewClient(iaccount.LoginTypeCustom, fmt.Sprintf("%s%04d", client.GlobalConfig.Prefix, i), "ababcc", &wg)
		pClient := client.NewClient(10001+int64(i), &wg)
		// pClient := client.NewClient(iaccount.LoginTypeCustom, "test0115", "ababcc", &wg)
		go pClient.Run()
	}
	time.Sleep(time.Millisecond * time.Duration(client.GlobalConfig.Interval))

	wg.Wait()
	logger.Debug(time.Now().Sub(tm))
	logger.DefaultLogger.Flush()
}
