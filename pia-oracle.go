package main

import (
	"fmt"
	"io/ioutil"
	//"github.com/golang/protobuf/proto"
	"github.com/vamitrou/pia-oracle/config"
	"github.com/vamitrou/pia-oracle/pialog"
	//"github.com/vamitrou/pia-oracle/protobuf"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

var conf *config.PiaConf = nil

func trigger(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		io.WriteString(w, "Method not allowed.\n")
		return
	}
	io.WriteString(w, "OK")
	pialog.Info("New trigger -", r.Host)
	go GetData()
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
	//fmt.Println(len(body))

	go exportData(body)
}

func predict(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "dummy predict\n")
}

func exportData(data []byte) {
	pialog.Info("Callback received with payload size:", len(data), "-> Exporting data to Oracle DB.")
	timeTrack(time.Now(), "Data export")
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
		pialog.Error("Not valid scores array")
	}

	if var_imps, ok := j["var_imp"].([]interface{}); ok {
		PushVarIMP(var_imps)
	} else {
		pialog.Error("Not valid var_imps array")
	}
	//fmt.Println(protoscore)
}

func main() {
	version := 0.1

	conf = config.GetConfig()

	pialog.InitializeLogging()
	pialog.Info("Starting pia-oracle version:", version)
	pialog.Info("Server started:", fmt.Sprintf("%s:%d", conf.Local.Listen, conf.Local.Port))

	http.HandleFunc("/trigger", trigger)
	http.HandleFunc("/callback", callback)
	http.HandleFunc("/predict", predict)
	http.ListenAndServe(fmt.Sprintf("%s:%d", conf.Local.Listen, conf.Local.Port), nil)
}
