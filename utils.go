package main

import (
	"database/sql"
	"fmt"
	"github.com/linkedin/goavro"
	"github.com/vamitrou/pia-oracle/pialog"
	"io/ioutil"
	"reflect"
	"time"

	_ "github.com/vamitrou/go-oci8"
)

func check(e error) {
	check_with_abort(e, true)
}

func check_with_abort(e error, abort bool) bool {
	if e != nil {
		if abort {
			panic(e)
		} else {
			fmt.Println(e)
			return true
		}
	}
	return false
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	pialog.Info(name, "duration:", elapsed)
}

func PrepareStatement(tx *sql.Tx, q string) *sql.Stmt {
	stmt, err := tx.Prepare(q)
	check(err)
	return stmt
}

func TimeStampToOraDate(ts int64) string {
	return fmt.Sprintf("TO_DATE('19700101','yyyymmdd') + (%d/24/60/60)", ts)
}

func AssertFloat(value interface{}) float64 {
	if value == nil {
		return 0
	}
	if v, ok := value.(float64); ok {
		return v
	} else {
		fmt.Println("++++")
		fmt.Println(reflect.TypeOf(value))
		fmt.Println("++++")
		panic("AssertFloat: not a float")
	}
}

func AssertString(value interface{}) string {
	if value == nil {
		return ""
	}
	if v, ok := value.(string); ok {
		return v
	} else {
		fmt.Println("++++")
		fmt.Println(reflect.TypeOf(value))
		fmt.Println("++++")
		panic("AssertString: not a string")
	}
}

func AssertTime(value interface{}) int64 {
	if value == nil {
		return 0
	}
	if v, ok := value.(time.Time); ok {
		return v.Unix()
	} else {
		panic("AssertTime: not a time.Time")
	}
}

func LoadAvroSchema(outerFile string, innerFile string) (goavro.RecordSetter, goavro.RecordSetter, goavro.Codec) {
	dat, err := ioutil.ReadFile(innerFile)
	check(err)
	innerSchemaStr := string(dat)

	dat2, err := ioutil.ReadFile(outerFile)
	check(err)
	outerSchemaStr := fmt.Sprintf(string(dat2), innerSchemaStr)

	outerSchema := goavro.RecordSchema(outerSchemaStr)
	innerSchema := goavro.RecordSchema(innerSchemaStr)
	codec, err := goavro.NewCodec(outerSchemaStr)
	check(err)
	return outerSchema, innerSchema, codec
}

func GetAvroFields(record *goavro.Record, object string) []string {
	schema, _ := record.GetFieldSchema(object)
	items := schema.(map[string]interface{})["items"]
	fields := items.(map[string]interface{})["fields"].([]interface{})
	ret_fields := make([]string, len(fields))
	for _, field := range fields {
		f := field.(map[string]interface{})["name"].(string)
		ret_fields = append(ret_fields, f)
	}
	return ret_fields
}
