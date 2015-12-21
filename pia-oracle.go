package main

import (
	"fmt"
	"io/ioutil"
	//"github.com/golang/protobuf/proto"
	"github.com/vamitrou/pia-oracle/config"
	//"github.com/vamitrou/pia-oracle/protobuf"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

var conf *config.PiaConf = nil

func trigger(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		io.WriteString(w, "Method not allowed.")
		return
	}
	io.WriteString(w, "Triggered..")
	fmt.Println(time.Now())
	go importData()
}

func callback(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		io.WriteString(w, "Method not allowed.")
		return
	}
	io.WriteString(w, "callback\n")

	body, err := ioutil.ReadAll(r.Body)
	check(err)
	//fmt.Println(string(body))
	fmt.Println(len(body))

	go exportData(body)
	fmt.Println("-> callback")
}

func predict(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "dummy predict\n")
}

func importData() {
	fmt.Printf("importing data\n")
	GetData()
}

func exportData(data []byte) {
	fmt.Printf("exporting data\n")
	//protoscore := &protoclaim.ProtoListScore{}
	//err := proto.Unmarshal(data, protoscore)
	//check(err)
	//PushData(protoscore)

	var j map[string]interface{}
	err := json.Unmarshal(data, &j)
	check(err)
	if scores, ok := j["Score"].([]interface{}); ok {
		PushScores(scores)
	} else {
		fmt.Println("Not valid scores array")
	}
	if var_imps, ok := j["var_imp"].([]interface{}); ok {
		PushVarIMP(var_imps)
	} else {
		fmt.Println("Not valid var_imps array")
	}
	//fmt.Println(protoscore)
	fmt.Println("done")
	fmt.Println(time.Now())
}

func main() {
	fmt.Println("Server started")

	conf = config.GetConfig()

	http.HandleFunc("/trigger", trigger)
	http.HandleFunc("/callback", callback)
	http.HandleFunc("/predict", predict)
	http.ListenAndServe(fmt.Sprintf("%s:%d", conf.Local.Listen, conf.Local.Port), nil)
}
