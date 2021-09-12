package dao

import (
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
)

type Record struct {
	id    string
	value string
}

func GetDBRecord(id string) (record *Record, err error) {
	if id == "fake_invalid_id" {
		return nil, errors.Wrap(sql.ErrNoRows, fmt.Sprintf("No record found"))
	} else {
		return &Record{id, "valid_value"}, nil
	}
}
