package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"github.com/famz/SetLocale"
	"github.com/linkedin/goavro"
	"github.com/vamitrou/pia-oracle/pialog"
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
	if err != nil {
		pialog.Error(err)
		return
	}
	SelectDBData(db, conf.Database.Query)
}

//func PushData(scores *protoclaim.ProtoListScore) {
func PushScores(scores []interface{}) {
	SetLocale.SetLocale(SetLocale.LC_ALL, "de_DE")
	dsn := getDSN()
	db, err := sql.Open("oci8", dsn)
	check(err)

	defer db.Close()
	tx, err := db.Begin()
	check(err)
	stmt := PrepareStatement(tx, conf.Database.QueryOut)
	/*for _, score := range scores.Scores {
		ExecuteInsert(stmt, score)
	}*/
	fails_count := 0
	for _, score := range scores {
		if s, ok := score.(map[string]interface{}); ok {
			err = ExecuteScoreInsert(stmt, s)
			if err != nil {
				fails_count += 1
			}
		} else {
			pialog.Error("Not a valid score:\n", score)
		}
	}
	tx.Commit()
	pialog.Info("Score failures count:", fails_count)
}

func PushVarIMP(var_imps []interface{}) {
	SetLocale.SetLocale(SetLocale.LC_ALL, "de_DE")
	dsn := getDSN()
	db, err := sql.Open("oci8", dsn)
	check(err)
	defer db.Close()

	tx, err := db.Begin()
	check(err)
	stmt := PrepareStatement(tx, conf.Database.QueryOutImp)
	fails_count := 0
	for _, var_imp := range var_imps {
		if s, ok := var_imp.(map[string]interface{}); ok {
			err = ExecuteImpInsert(stmt, s)
			if err != nil {
				fails_count += 1
			}
		} else {
			pialog.Error("Not a valid var_imp:\n", var_imp)
		}
	}
	tx.Commit()
	pialog.Info("IMP failures count:", fails_count)
}

func SelectDBData(db *sql.DB, query string) {
	defer timeTrack(time.Now(), "SelectDBData")

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
	var claims []interface{}
	outerSchema, innerSchema, codec := LoadAvroSchema("conf/claims.json", "conf/claim.json")

	for rows.Next() {
		c += 1
		err = rows.Scan(scanArgs...)
		check(err)

		m := make(map[string]interface{})
		claim, err := goavro.NewRecord(innerSchema)
		check(err)
		for i, colName := range columns {
			if val, ok := values[i].(time.Time); ok {
				claim.Set(colName, val.Unix())
				m[colName] = val.Unix()
			} else {
				claim.Set(colName, values[i])
				m[colName] = values[i]
			}
		}

		//j, _ := json.Marshal(m)
		//fmt.Println(string(j))
		//ioutil.WriteFile("sample.json", j, 0644)
		claims = append(claims, claim)

		if len(claims) == 3000 {

			claims_avro, err := goavro.NewRecord(outerSchema)
			check(err)
			claims_avro.Set("claims", claims)
			buf := new(bytes.Buffer)
			err = codec.Encode(buf, claims_avro)
			go Post(conf.Rest.PredictionEndpoint, buf.Bytes(), conf.Rest.AppHeader)
			claims = make([]interface{}, 0)
		}
	}

	if rows.Err() != nil {
		pialog.Error(rows.Err())
	}

	claims_avro, err := goavro.NewRecord(outerSchema)
	check(err)
	claims_avro.Set("claims", claims)

	buf := new(bytes.Buffer)
	err = codec.Encode(buf, claims_avro)
	check(err)

	//ioutil.WriteFile("avro_data/results.last", buf.Bytes(), 0644)

	/*	actual := buf.Bytes()
		dec, err := codec.Decode(bytes.NewReader(actual))
		check(err)
		fmt.Println(dec) */

	go Post(conf.Rest.PredictionEndpoint, buf.Bytes(), conf.Rest.AppHeader)

	pialog.Info("Total input records:", c)
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
