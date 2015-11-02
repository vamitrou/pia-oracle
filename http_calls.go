package main

import (
	"bytes"
	"container/list"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Rester struct {
	Conf RestConf
}

type RestConf struct {
	PredictEndpoint string `json:"predict_endpoint"`
}

func (r *Rester) ReadConfig(configPath string) {
	dat, err := ioutil.ReadFile(configPath)
	check(err)
	err = json.Unmarshal(dat, &r.Conf)
	check(err)
}

func (r Rester) BatchData(l list.List) {
	for e := l.Front(); e != nil; e = e.Next() {
		// do something with e.Value
	}
}

func (r Rester) Post(data []byte) {
	var jsonStr = []byte(`{"test": "test"}`)
	req, err := http.NewRequest("POST", r.Conf.PredictEndpoint, bytes.NewBuffer(jsonStr))
	check(err)
	req.Header.Set("Content-Type", "application/json")

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
