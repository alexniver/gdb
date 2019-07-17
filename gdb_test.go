package gdb

import (
	"path/filepath"
	"strconv"
	"testing"
)

var BasePath = "/tmp/gdbtest/"

type Table1 struct {
	Id int
	Name string
}

func (t *Table1) Path() string {
	return filepath.Join(BasePath, "table1")
}

func (t *Table1) TableName() string {
	return strconv.Itoa(t.Id)
}



func TestSave(t *testing.T) {
	t1 := &Table1{
		Id:1,
		Name:"a table",
	}
	err := Save(t1)
	if nil != err {
		t.Errorf("Save got a error, err : %+v", err)
	}
}
