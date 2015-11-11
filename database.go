package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/famz/SetLocale"
	"github.com/golang/protobuf/proto"
	"github.com/vamitrou/pia-oracle/protobuf"
	"io/ioutil"
	"os"
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

	query := `select * from %[1]s.%[2]s where glb_oe_id=4043 and (
	       (invstgt_strt_dt is not NULL) or (clm_rgstr_dttm >= (select min(invstgt_strt_dt)
	             from %[1]s.%[2]s where invstgt_strt_dt is not NULL)) ) and
		              (load_date in (select max(load_date) as load_date from %[1]s.%[2]s))`
	query = fmt.Sprintf(query, odb.Conf.Schema, odb.Conf.Table)
	odb.SelectDBData(db, query)
}

func (odb OracleDB) SelectDBData(db *sql.DB, query string) {
	defer timeTrack(time.Now(), "get oracle data")

	rows, err := db.Query(query)
	check(err)
	defer rows.Close()

	columns, err := rows.Columns()
	check(err)

	values := make([]interface{}, len(columns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	f, err := os.Create("data/results.txt.0")
	check(err)

	c := 0
	claims := new(protoclaim.ProtoListClaim)

	for rows.Next() {
		c += 1
		err = rows.Scan(scanArgs...)
		check(err)

		var m = make(map[string]interface{})
		for i, colName := range columns {
			m[colName] = values[i]
		}

		claim := ClaimForMap(m)
		claims.Claims = append(claims.Claims, claim)

		if c%3000 == 0 {
			proto_bytes, err := proto.Marshal(claims)
			check(err)
			f.Write(proto_bytes)
			f.Sync()
			f.Close()

			f, err = os.Create(fmt.Sprintf("data/results.txt.%d", c))
			check(err)
			claims = new(protoclaim.ProtoListClaim)
		}
	}

	if rows.Err() != nil {
		fmt.Println(rows.Err())
	}

	proto_bytes, err := proto.Marshal(claims)
	check(err)

	f, err = os.Create(fmt.Sprintf("data/results.txt.%d", c))
	check(err)
	f.Write(proto_bytes)
	f.Sync()
	f.Close()

	fmt.Println(fmt.Sprintf("total count: %d", c))
}
