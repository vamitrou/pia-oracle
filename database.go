package main

import (
	"bufio"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/famz/SetLocale"
	"io/ioutil"
	"os"
	//	"strconv"
	"time"

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
	Schema      string `json:"schema"`
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
	SetLocale.SetLocale(SetLocale.LC_ALL, "de_DE")
	dsn := odb.getDSN()
	db, err := sql.Open("oci8", dsn)
	check(err)

	defer db.Close()

	f, err := os.Open("columns.txt")
	defer f.Close()
	check(err)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		query := fmt.Sprintf("SELECT %s FROM ( select * from %s.%s) where ROWNUM < 5",
			scanner.Text(), odb.Conf.Schema, odb.Conf.Table)
		odb.testSelect(db, query)
	}

}

func (odb OracleDB) testSelect(db *sql.DB, query string) {
	defer timeTrack(time.Now(), "get oracle data")

	//query := fmt.Sprintf("select * from %s.%s",
	//query := fmt.Sprintf("SELECT * FROM ( select * from %s.%s) where ROWNUM < 5",
	//	odb.Conf.Schema, odb.Conf.Table)
	//fmt.Println(query)
	rows, err := db.Query(query)
	// rows, err := db.Query("SELECT * FROM ( select * from V33_GDWHANLT.FRAUD_ABT ) where ROWNUM < 5")
	check(err)
	defer rows.Close()

	columns, err := rows.Columns()
	check(err)

	// rows.Scan wants '[]interface{}' as an argument, so we must copy the
	// references into such a slice
	// See http://code.google.com/p/go-wiki/wiki/InterfaceSlice for details
	values := make([]interface{}, len(columns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	// data := make([]map[string]interface{}, 0)

	// Fetch rows
	for rows.Next() {
		// fmt.Println("new row")
		err = rows.Scan(scanArgs...)
		check(err)
		//fmt.Println(values)
		// get RawBytes from data
		/*	err = rows.Scan(scanArgs...)
			check(err)

			var m = make(map[string]interface{})
			i := 0
			for _, colName := range columns {
				strVal, _ := values[i].(string)
				num, err := strconv.ParseFloat(strVal, 32)
				if err != nil {
					m[colName] = values[i]
				} else {
					m[colName] = num
				}
				i += 1
			}

			data = append(data, m) */
	}

	if rows.Err() != nil {
		fmt.Println(rows.Err())
	}

	//fmt.Println(l.Len())
	/*defer timeTrack(time.Now(), "marshal and write data")
	json_bytes, err := json.Marshal(data)
	check(err)
	f, err := os.Create("results.txt")
	defer f.Close()
	check(err)
	f.WriteString(string(json_bytes))
	// fmt.Println(string(json_bytes))*/
}
