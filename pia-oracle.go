package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

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
	odb := new(OracleDB)
	odb.ReadConfig("conf/db_conf.json")
	fmt.Println(odb.Conf.Host)
	odb.GetData()
}

func exportData() {
	time.Sleep(10 * time.Second)
	fmt.Printf("exporting data\n")
}

func main() {
	go func() { fmt.Println("yourself") }()
	fmt.Println("Server started")

	rest := new(Rester)
	rest.ReadConfig("conf/rest_conf.json")

	http.HandleFunc("/trigger", trigger)
	http.HandleFunc("/callback", callback)
	http.ListenAndServe(":8001", nil)
}
