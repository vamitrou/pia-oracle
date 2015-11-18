package main

import (
	"database/sql"
	"fmt"
	"github.com/famz/SetLocale"
	"github.com/golang/protobuf/proto"
	"github.com/vamitrou/pia-oracle/protobuf"
	//	"os"
	"time"

	_ "github.com/vamitrou/go-oci8"
)

func getDSN() string {
	return fmt.Sprintf("%s/%s@%s:%d/%s",
		conf.Database.Username,
		conf.Database.Password,
		conf.Database.Host,
		conf.Database.Port,
		conf.Database.ServiceName)
}

func GetData() {
	SetLocale.SetLocale(SetLocale.LC_ALL, "de_DE")
	dsn := getDSN()
	db, err := sql.Open("oci8", dsn)
	check(err)

	defer db.Close()
	SelectDBData(db, conf.Database.Query)
}

func PushData(scores *protoclaim.ProtoListScore) {
	SetLocale.SetLocale(SetLocale.LC_ALL, "de_DE")
	dsn := getDSN()
	db, err := sql.Open("oci8", dsn)
	check(err)

	defer db.Close()
	tx, err := db.Begin()
	check(err)
	stmt := PrepareStatement(tx, conf.Database.QueryOut)
	for _, score := range scores.Scores {
		ExecuteInsert(stmt, score)
	}
	tx.Commit()

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
			go Post(conf.Rest.PredictionEndpoint, proto_bytes)

			claims = new(protoclaim.ProtoListClaim)
		}
	}

	if rows.Err() != nil {
		fmt.Println(rows.Err())
	}

	proto_bytes, err := proto.Marshal(claims)
	check(err)

	fmt.Println(len(claims.Claims))
	go Post(conf.Rest.PredictionEndpoint, proto_bytes)
	fmt.Println(fmt.Sprintf("total count: %d", c))
	/*score_proto := &protoclaim.ProtoListScore{}
	for i := 0; i < 10; i++ {
		sc := &protoclaim.ProtoListScore_ProtoScore{
			GLB_OE_ID:    proto.Int64(int64(i)),
			CLM_BUS_ID:   proto.String("bus_id"),
			SCORE:        proto.Float64(34.232),
			MODEL:        proto.String("model"),
			CREATE_DT:    proto.Int64(1447671925),
			CREATE_DT_TS: proto.Int64(1447671925),
		}
		score_proto.Scores = append(score_proto.Scores, sc)
	}
	proto_bytes, err = proto.Marshal(score_proto)
	check(err)
	go Post(conf.Rest.PredictionEndpoint, proto_bytes)*/

}
