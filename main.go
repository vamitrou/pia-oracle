package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func trigger(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Triggered..")
	go importData()
}

func callback(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "callback")
	go exportData()
}

func importData() {
	fmt.Printf("importing data\n")
	odb := new(OracleDB)
	//defer odb.Close()
	odb.ReadConfig("db_conf.json")
	fmt.Println(odb.Conf.Host)
	odb.GetData()
}

func exportData() {
	time.Sleep(10 * time.Second)
	fmt.Printf("exporting data\n")
}

func main() {
	fmt.Println("Server started")

	rest := new(Rester)
	rest.ReadConfig("rest_conf.json")

	http.HandleFunc("/trigger", trigger)
	http.HandleFunc("/callback", callback)
	http.ListenAndServe(":8001", nil)
}
