package main

import (
	"database/sql"
	"fmt"
	"reflect"
	"time"

	_ "github.com/vamitrou/go-oci8"
)

func check(e error) {
	check_with_abort(e, true)
}

func check_with_abort(e error, abort bool) {
	if e != nil {
		if abort {
			panic(e)
		} else {
			fmt.Println(e)
		}
	}
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s\n", name, elapsed)
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
