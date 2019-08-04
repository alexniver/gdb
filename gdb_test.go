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

type Table2 struct {
	Id int
	Name string
}

func (t Table2) Key() string {
	return strconv.Itoa(t.Id)
}




func TestSave(t *testing.T) {
	Init(string(filepath.Separator), "tmp", "gdbtest")
	t1 := &Table1{
		Id:1,
		Name:"a table, type of Table1",
	}
	err := Save(t1)
	if nil != err {
		t.Errorf("Save got a error, err : %+v", err)
	}

	t2 := Table2{
		Id:1,
		Name:"a table, type of Table2",
	}
	err = Save(t2)
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
			Name: "table" + strconv.Itoa(i) + "type Table1",
		}
		gdbArr = append(gdbArr, t1)
	}


	for i := 0; i < 10; i ++ {
		t1 := Table2{
			Id:i,
			Name: "table" + strconv.Itoa(i) + "type Table2",
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
	fmt.Printf("%+v\n", t1.(*Table1))

	t2, err := One("2", reflect.TypeOf((*Table2)(nil)))
	if nil != err {
		t.Errorf("TestOne error %+v", err)
	}
	fmt.Printf("%+v\n", t2.(*Table2))
}

func TestAll(t *testing.T) {
	Init(string(filepath.Separator), "tmp", "gdbtest")

	allInter, err := All(reflect.TypeOf((*Table1)(nil)))
	if nil != err {
		t.Errorf("TestAll error %+v", err)
	}
	for _, v := range allInter {
		fmt.Println(*(v.(*Table1)))
	}

	allInter, err = All(reflect.TypeOf((*Table2)(nil)))
	if nil != err {
		t.Errorf("TestAll error %+v", err)
	}
	for _, v := range allInter {
		fmt.Println(*(v.(*Table2)))
	}
}

func TestDel(t *testing.T) {
	Init(string(filepath.Separator), "tmp", "gdbtest")
	err := Del("1", reflect.TypeOf((*Table1)(nil)))
	if nil != err {
		t.Errorf("TestDel error %+v", err)
	}
	err = Del("1", reflect.TypeOf((*Table2)(nil)))
	if nil != err {
		t.Errorf("TestDel error %+v", err)
	}
}