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
	rows, err := db.Query("SELECT GLB_OE_ID, CLM_RK, ABT_DT_ZERO FROM ( select * from V33_GDWHANLT.FRAUD_ABT ) where ROWNUM < 5")
	check(err)
	defer rows.Close()

	/*fmt.Println(rows.Columns())

	for rows.Next() {
		fmt.Println("next")
		var f1 int
		rows.Scan(&f1)
		fmt.Println(f1)
	}*/

	// Get column names
	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	// Make a slice for the values
	values := make([]sql.RawBytes, len(columns))

	// rows.Scan wants '[]interface{}' as an argument, so we must copy the
	// references into such a slice
	// See http://code.google.com/p/go-wiki/wiki/InterfaceSlice for details
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	// Fetch rows
	for rows.Next() {
		// get RawBytes from data
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		// Now do something with the data.
		// Here we just print each column as a string.
		var value string
		for i, col := range values {
			// Here we can check if the value is nil (NULL value)
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			fmt.Println(columns[i], ": ", value)
		}
		fmt.Println("-----------------------------------")
	}

}
