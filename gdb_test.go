package gdb

import (
	"fmt"
	"path/filepath"
	"reflect"
	"strconv"
	"testing"
)


type Table1 struct {
	Id int
	Name string
}

func (t *Table1) Key() string {
	return strconv.Itoa(t.Id)
}



func TestSave(t *testing.T) {
	Init(string(filepath.Separator), "tmp", "gdbtest")
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
	Init(string(filepath.Separator), "tmp", "gdbtest")

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
	Init(string(filepath.Separator), "tmp", "gdbtest")
	t1, err := One("2", reflect.TypeOf((*Table1)(nil)))
	if nil != err {
		t.Errorf("TestOne error %+v", err)
	}
	fmt.Printf("%+v\n", *t1.(**Table1))
}

func TestAll(t *testing.T) {
	Init(string(filepath.Separator), "tmp", "gdbtest")

	allInter, err := All(reflect.TypeOf((*Table1)(nil)))
	if nil != err {
		t.Errorf("TestAll error %+v", err)
	}
	for _, v := range allInter {
		vv := v.(**Table1)
		fmt.Println(*vv)
	}
}

func TestDel(t *testing.T) {
	Init(string(filepath.Separator), "tmp", "gdbtest")
	err := Del("1", reflect.TypeOf((*Table1)(nil)))
	if nil != err {
		t.Errorf("TestDel error %+v", err)
	}
}