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


type Parent struct {
	Id int
	Name string
}

func (p *Parent) Key() string {
	return strconv.Itoa(p.Id)
}

type Child struct {
	Id int
	PId int
	Name string
}

func (c *Child) Key() string {
	return strconv.Itoa(c.Id)
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


	// test subPath,
	p := &Parent{
		Id:1,
		Name: "parent",
	}

	c := &Child{
		Id:1,
		Name:"child",
		PId:1,
	}

	err = Save(p, strconv.Itoa(p.Id))
	if nil != err {
		t.Error(err)
	}

	err = Save(c, strconv.Itoa(c.PId))
	if nil != err {
		t.Error(err)
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

	// test subPath,
	p1 := &Parent{
		Id:1,
		Name: "parent",
	}

	p2 := &Parent{
		Id:2,
		Name: "parent",
	}


	err = Save(p1, strconv.Itoa(p1.Id))
	if nil != err {
		t.Error(err)
	}


	err = Save(p2, strconv.Itoa(p2.Id))
	if nil != err {
		t.Error(err)
	}

	var child1Arr []Gdb
	for i := 0; i < 10; i ++ {
		c := &Child{
			Id: i,
			Name: "c" + strconv.Itoa(i),
			PId:p1.Id,
		}
		child1Arr = append(child1Arr, c)
	}

	err = SaveAll(child1Arr, strconv.Itoa(p1.Id))

	var child2Arr []Gdb
	for i := 0; i < 10; i ++ {
		c := &Child{
			Id: i,
			Name: "c" + strconv.Itoa(i),
			PId:p2.Id,
		}
		child2Arr = append(child2Arr, c)
	}

	err = SaveAll(child2Arr, strconv.Itoa(p2.Id))
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


	// test subPath
	pIdStr := "1"
	p1, err := One(pIdStr, reflect.TypeOf((*Parent)(nil)), pIdStr)
	if nil != err {
		t.Errorf("TestOne error %+v", err)
	}
	fmt.Printf("%+v\n", p1.(*Parent))
	c2, err := One("2", reflect.TypeOf((*Child)(nil)), pIdStr)
	if nil != err {
		t.Errorf("TestOne error %+v", err)
	}
	fmt.Printf("%+v\n", c2.(*Child))


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

	// test subPath

	p1IdStr := "1"
	allInter, err = All(reflect.TypeOf((*Child)(nil)), p1IdStr)
	if nil != err {
		t.Errorf("TestAll error %+v", err)
	}
	for _, v := range allInter {
		fmt.Println(*(v.(*Child)))
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

	// test subPath

	pIdStr := "1"
	err = Del("1", reflect.TypeOf((*Child)(nil)), pIdStr)
	if nil != err {
		t.Errorf("TestDel error %+v", err)
	}

}


func TestAllSubPath(t *testing.T) {
	Init(string(filepath.Separator), "tmp", "gdbtest")
	subPaths, err := AllSubPath()
	if nil != err {
		t.Errorf("Test AllSubPath error %+v", err)
	}

	fmt.Println(subPaths)
}

