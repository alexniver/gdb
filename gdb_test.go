package gdb

import (
	"fmt"
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

func TestSaveAll(t *testing.T) {
	var gdbArr []Gdb

	for i := 0; i < 10; i ++ {
		t1 := &Table1{
			Id:i,
			Name: "table" + strconv.Itoa(i),
		}
		gdbArr = append(gdbArr, t1)
	}

	err := SaveAll(gdbArr)
	if nil != err {
		t.Errorf("TestSaveAll error %+v", err)
	}
}


func TestOne(t *testing.T) {
	var t1 Table1
	err := One(t1.Path(), "1", &t1)
	if nil != err {
		t.Errorf("TestOne error")
	}
	fmt.Printf("+%v", t1)
}

func TestAllTableName(t *testing.T) {
	var t1 Table1
	names, err := AllTableName(t1.Path())
	if nil != err {
		t.Errorf("TestAllTableName error %+v", err)
	}
	fmt.Printf("%+v", names)
}

func TestDel(t *testing.T) {
	err := Del(filepath.Join(BasePath, "table1"), "1")
	if nil != err {
		t.Errorf("TestDel error %+v", err)
	}
}