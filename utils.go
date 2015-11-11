package main

import (
	"fmt"
	"time"
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
