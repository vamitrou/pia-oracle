package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/vamitrou/pia-oracle/config"
	"github.com/vamitrou/pia-oracle/protobuf"
	"io"
	"io/ioutil"
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
	if r.Method != "GET" {
		io.WriteString(w, "Method not allowed.")
		return
	}
	io.WriteString(w, "callback\n")

	body, err := ioutil.ReadAll(r.Body)
	check(err)
	//fmt.Println(body)
	fmt.Println(len(body))

	go exportData(body)
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
	protoscore := &protoclaim.ProtoListScore{}
	err := proto.Unmarshal(data, protoscore)
	check(err)
	PushData(protoscore)
	//fmt.Println(protoscore)
}

func main() {
	fmt.Println("Server started")

	conf = config.GetConfig()

	http.HandleFunc("/trigger", trigger)
	http.HandleFunc("/callback", callback)
	http.HandleFunc("/predict", predict)
	http.ListenAndServe(fmt.Sprintf("%s:%d", conf.Local.Listen, conf.Local.Port), nil)
}
