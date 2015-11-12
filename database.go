package main

import (
	"database/sql"
	"fmt"
	"github.com/famz/SetLocale"
	"github.com/golang/protobuf/proto"
	"github.com/vamitrou/pia-oracle/config"
	"github.com/vamitrou/pia-oracle/protobuf"
	"os"
	"time"

	_ "github.com/vamitrou/go-oci8"
)

func getDSN(conf config.DatabaseConf) string {
	fmt.Println(fmt.Sprintf("%s/%s@%s:%d/%s",
		conf.Username,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.ServiceName))
	return ""
}

func GetData(conf config.DatabaseConf) {
	SetLocale.SetLocale(SetLocale.LC_ALL, "de_DE")
	dsn := getDSN(conf)
	db, err := sql.Open("oci8", dsn)
	check(err)

	defer db.Close()

	query := `select * from (select * from %[1]s.%[2]s where glb_oe_id=4043 and (
	       (invstgt_strt_dt is not NULL) or (clm_rgstr_dttm >= (select min(invstgt_strt_dt)
	             from %[1]s.%[2]s where invstgt_strt_dt is not NULL)) ) and
		              (load_date in (select max(load_date) as load_date from %[1]s.%[2]s))) where rownum<2`
	query = fmt.Sprintf(query, conf.Schema, conf.Table)
	SelectDBData(db, query)
}

func SelectDBData(db *sql.DB, query string) {
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
