package main

import (
	"fmt"
	"io"
	"net/http"
)

func trigger(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Triggered..")
	importData()
}

func callback(w http.ResponseWriter, r *http.Request) {
	exportData()
}

func importData() {
	fmt.Printf("importing data\n")
}

func exportData() {
	fmt.Printf("exporting data\n")
}

func main() {
	odb := new(OracleDB)
	odb.ReadConfig("db_conf.json")
	fmt.Println(odb.Conf.Host)
	odb.GetData()

	rest := new(Rester)
	rest.ReadConfig("rest_conf.json")

	http.HandleFunc("/trigger", trigger)
	http.HandleFunc("/callback", callback)
	http.ListenAndServe(":8001", nil)
}
