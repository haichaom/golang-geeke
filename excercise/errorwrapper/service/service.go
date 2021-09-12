package service

import (
	"github.com/haichaom/golang-geeke/excercise/errorwrapper/dao"
	"github.com/pkg/errors"
)

func GetRecordByID(id string) (*dao.Record, error) {
	record, err := dao.GetDBRecord(id)
	if err == nil {
		return record, nil
	} else {
		return nil, errors.Wrap(err, "service::getRecordByID failed with error!")
	}
}
