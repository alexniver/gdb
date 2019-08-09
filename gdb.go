package gdb

import (
	"errors"
	"github.com/alexniver/gtools"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"sync"
)

var dbPath string
var mux sync.Mutex

var (
	ErrorNeedInit = errors.New("gdb need Init")
	ErrorTypeNotGdb = errors.New("t need implements interface Gdb")
)


/*
	Gdb 接口
	需要实现1个方法
	Key() struct 存储的最终文件名,需要保证唯一性
 */
type Gdb interface {
	Key() string
}

// Init dirs, 设置要目录，以逗号分开
func Init(dirs ...string) {
	dbPath = filepath.Join(dirs...)
}

// Save 保存
func Save(gdb Gdb, subPath ...string) error {
	mux.Lock()
	defer mux.Unlock()

	return save(gdb, subPath...)
}
// 保存
func save(gdb Gdb, subPath ...string) error {
	if dbPath == "" {
		return ErrorNeedInit
	}
	buf, err := gtools.EncodeGZip(gdb)
	typeName := gtools.TypeName(reflect.TypeOf(gdb))
	sPath := filepath.Join(subPath...)
	path := filepath.Join(dbPath, sPath, typeName)
	if nil != err {
		return err
	}
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, os.ModePerm)
		if nil != err {
			return err
		}
	}

	vPath := filepath.Join(path, gdb.Key())
	err = ioutil.WriteFile(vPath, buf, os.ModePerm)
	if nil != err {
		return err
	}
	return nil
}

// SaveAll 保存所有数据
func SaveAll(gdbArr []Gdb, subPath ...string) error {
	mux.Lock()
	defer mux.Unlock()

	for _, gdb := range gdbArr {
		err := save(gdb, subPath...)
		if nil != err {
			return err
		}
	}
	return nil
}

// One 根据key和type, 获取对应的数据, 返回值为interface{}, 获取到之后需要通过例如v.(*Foo)来进行转换
func One(key string, t reflect.Type, subPath ...string) (interface{}, error) {
	mux.Lock()
	defer mux.Unlock()

	return one(key, t, subPath...)
}

func one(key string, t reflect.Type, subPath ...string) (interface{}, error) {
	if dbPath == "" {
		return nil, ErrorNeedInit
	}

	gdbType := reflect.TypeOf((*Gdb)(nil)).Elem()
	if !t.Implements(gdbType) {
		return nil, ErrorTypeNotGdb
	}

	sPath := filepath.Join(subPath...)
	vPath := filepath.Join(dbPath, sPath, gtools.TypeName(t), key)
	buf, err := ioutil.ReadFile(vPath)
	if nil != err {
		return nil, err
	}

	var finalT = t
	if t.Kind() == reflect.Ptr {
		finalT = t.Elem()
	}

	v := reflect.New(finalT).Interface()
	err = gtools.DecodeUnGzip(buf, v)
	if nil != err {
		return nil, err
	}
	return v, nil
}


// All 返回所有t的数据
func All(t reflect.Type, subPath ...string) ([]interface{}, error) {
	mux.Lock()
	defer mux.Unlock()

	// type check
	gdbType := reflect.TypeOf((*Gdb)(nil)).Elem()
	if !t.Implements(gdbType) {
		return nil, ErrorTypeNotGdb
	}

	var result []interface{}
	sPath := filepath.Join(subPath...)
	path := filepath.Join(dbPath, sPath, gtools.TypeName(t))
	files, err := ioutil.ReadDir(path)
	if nil != err {
		return nil, err
	}
	for _, f := range files {
		v, err := one(f.Name(), t, subPath...)
		if nil != err {
			return nil, err
		}
		result = append(result, v)
	}
	return result, nil
}

// 根据key和t 删除对应的数据
func Del(key string, t reflect.Type, subPath ...string) error {
	mux.Lock()
	defer mux.Unlock()

	if dbPath == "" {
		return ErrorNeedInit
	}

	sPath := filepath.Join(subPath...)
	path := filepath.Join(dbPath, sPath, gtools.TypeName(t), key)

	return os.Remove(path)
}


