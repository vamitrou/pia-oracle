package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"

	_ "github.com/mattn/go-oci8"
)

type OracleDB struct {
	Conf DBConf
}

type DBConf struct {
	Host        string `json:"host"`
	Port        int    `json:"port"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	ServiceName string `json:"service_name"`
	Table       string `json:"table"`
}

func (odb *OracleDB) ReadConfig(configPath string) {
	dat, err := ioutil.ReadFile(configPath)
	check(err)
	err = json.Unmarshal(dat, &odb.Conf)
	check(err)
}

func (odb OracleDB) getDSN() string {
	return fmt.Sprintf("%s/%s@%s:%d/%s",
		odb.Conf.Username,
		odb.Conf.Password,
		odb.Conf.Host,
		odb.Conf.Port,
		odb.Conf.ServiceName)
}

func (odb OracleDB) GetData() {
	dsn := odb.getDSN()
	db, err := sql.Open("oci8", dsn)
	check(err)

	defer db.Close()
	testSelect(db)
}

func testSelect(db *sql.DB) {
	fmt.Println("test select")
	rows, err := db.Query("select * from TEMPTABLE")
	check(err)
	defer rows.Close()

	for rows.Next() {
		var f1 int
		rows.Scan(&f1)
		fmt.Println(f1)
	}
}
