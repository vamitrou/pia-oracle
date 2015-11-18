package main

import (
	"bytes"
	"container/list"
	"fmt"
	"io/ioutil"
	"net/http"
)

func BatchData(l list.List) {
	for e := l.Front(); e != nil; e = e.Next() {
		// do something with e.Value
	}
}

func Post(url string, data []byte) {
	//var jsonStr = []byte(`{"test": "test"}`)
	//req, err := http.NewRequest("POST", r.Conf.PredictEndpoint, bytes.NewBuffer(jsonStr))
	req, err := http.NewRequest("POST", url, bytes.NewReader(data))
	check(err)
	req.Header.Set("Content-Type", "application/protobuf")

	client := &http.Client{}
	resp, err := client.Do(req)
	check(err)

	defer resp.Body.Close()
	fmt.Println("response status:", resp.Status)
	fmt.Println("response headers:", resp.Header)
	body, err := ioutil.ReadAll(resp.Body)
	check(err)
	fmt.Println("response body:", string(body))
}
