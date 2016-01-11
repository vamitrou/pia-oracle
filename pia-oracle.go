package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/vamitrou/pia-oracle/config"
	"github.com/vamitrou/pia-oracle/pialog"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

var conf *config.PiaConf = nil

func trigger(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
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
	go exportData(body)
}

func predict(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "dummy predict\n")
}

func exportData(data []byte) {
	pialog.Info("Callback received with payload size:", len(data), "-> Will export data to Oracle DB.")
	defer timeTrack(time.Now(), "Data export")

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
}

func main() {
	version := 0.1

	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	confpath := flag.String("config-dir", fmt.Sprintf("%s/conf", dir), "Config file path")
	flag.Parse()

	config.Path = *confpath

	pialog.InitializeLogging()
	pialog.Info("Starting pia-oracle version:", version)
	pialog.Info("Loading config from:", *confpath)

	var err error
	conf, err = config.GetConfig(fmt.Sprintf("%s/pia-oracle.toml", *confpath))
	if err != nil {
		pialog.Error(err)
		pialog.Info("Exiting")
		return
	}

	pialog.Info("Server started:", fmt.Sprintf("%s:%d", conf.Local.Listen, conf.Local.Port))

	http.HandleFunc("/trigger", trigger)
	http.HandleFunc("/callback", callback)
	http.HandleFunc("/predict", predict)
	err = http.ListenAndServe(fmt.Sprintf("%s:%d", conf.Local.Listen, conf.Local.Port), nil)
	if err != nil {
		pialog.Error(err)
	}
	pialog.Info("Exiting")
}
