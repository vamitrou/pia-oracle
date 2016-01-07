package main

import (
	"bytes"
	"container/list"
	"fmt"
	"github.com/vamitrou/pia-oracle/pialog"
	"io/ioutil"
	"net/http"
)

func BatchData(l list.List) {
	for e := l.Front(); e != nil; e = e.Next() {
		// do something with e.Value
	}
}

func Post(url string, data []byte, appHeader string) {
	//var jsonStr = []byte(`{"test": "test"}`)
	//req, err := http.NewRequest("POST", r.Conf.PredictEndpoint, bytes.NewBuffer(jsonStr))
	url = fmt.Sprintf("%s?callback=http://%s:%d/callback", url, conf.Local.Hostname, conf.Local.Port)
	req, err := http.NewRequest("POST", url, bytes.NewReader(data))
	if check_with_abort(err, false) {
		return
	}
	req.Header.Set("Content-Type", "avro/binary")
	req.Header.Set("Application", appHeader)

	client := &http.Client{}
	resp, err := client.Do(req)
	if check_with_abort(err, false) {
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if check_with_abort(err, false) {
		pialog.Error(err)
		return
	}
	pialog.Info("Server replied:", string(body), resp.Status)
}
