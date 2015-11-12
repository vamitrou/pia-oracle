package main

import (
	"fmt"
	"github.com/vamitrou/pia-oracle/config"
	"io"
	"net/http"
	"time"
)

var conf *config.PiaConf = nil

func trigger(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Triggered..")
	fmt.Println(time.Now())
	go importData()
}

func callback(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "callback")
	go exportData()
}

func importData() {
	fmt.Printf("importing data\n")
	GetData(conf.Database)
}

func exportData() {
	time.Sleep(10 * time.Second)
	fmt.Printf("exporting data\n")
}

func main() {
	fmt.Println("Server started")

	conf = config.GetConfig()
	fmt.Println(conf)

	rest := new(Rester)
	rest.ReadConfig("conf/rest_conf.json")

	http.HandleFunc("/trigger", trigger)
	http.HandleFunc("/callback", callback)
	http.ListenAndServe(fmt.Sprintf("%s:%d", conf.Local.Listen, conf.Local.Port), nil)
}
