package main

import (
	"fmt"

	"github.com/haichaom/golang-geeke/excercise/errorwrapper/service"
	"github.com/pkg/errors"
)

func main() {
	record, err := service.GetRecordByID("fake_invalid_id")
	if err != nil {
		fmt.Printf("error: %T %v\n", errors.Cause(err), errors.Cause((err)))
		fmt.Printf("Trackback: %+v\n", err)
	} else {
		fmt.Printf("Record: %v", record)
	}
}
