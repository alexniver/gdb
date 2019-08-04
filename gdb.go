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
	需要实现2个方法
	Path() 方法代表存储的路径名
	TableName() 方法代表存储的文件名
 */
type Gdb interface {
	Key() string
}

func Init(dirs ...string) {
	dbPath = filepath.Join(dirs...)
}

func Save(gdb Gdb) error {
	mux.Lock()
	defer mux.Unlock()

	return save(gdb)
}
// 保存
func save(gdb Gdb) error {
	if dbPath == "" {
		return ErrorNeedInit
	}
	buf, err := gtools.EncodeGZip(gdb)
	typeName := gtools.TypeName(reflect.TypeOf(gdb))
	path := filepath.Join(dbPath, typeName)
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
func SaveAll(gdbArr []Gdb) error {
	mux.Lock()
	defer mux.Unlock()

	for _, gdb := range gdbArr {
		err := save(gdb)
		if nil != err {
			return err
		}
	}
	return nil
}

// One 根据key和type, 获取对应的数据, 返回值为interface{}, 获取到之后需要通过例如v.(*Foo)来进行转换
func One(key string, t reflect.Type) (interface{}, error) {
	mux.Lock()
	defer mux.Unlock()

	return one(key, t)
}

func one(key string, t reflect.Type) (interface{}, error) {
	if dbPath == "" {
		return nil, ErrorNeedInit
	}

	gdbType := reflect.TypeOf((*Gdb)(nil)).Elem()
	if !t.Implements(gdbType) {
		return nil, ErrorTypeNotGdb
	}

	vPath := filepath.Join(dbPath, gtools.TypeName(t), key)
	buf, err := ioutil.ReadFile(vPath)
	if nil != err {
		return nil, err
	}

	v := reflect.New(t).Interface()
	err = gtools.DecodeUnGzip(buf, &v)
	if nil != err {
		return nil, err
	}
	return v, nil
}


// All 返回所有t的数据
func All(t reflect.Type) ([]interface{}, error) {
	mux.Lock()
	defer mux.Unlock()

	// type check
	gdbType := reflect.TypeOf((*Gdb)(nil)).Elem()
	if !t.Implements(gdbType) {
		return nil, ErrorTypeNotGdb
	}

	var result []interface{}
	files, err := ioutil.ReadDir(filepath.Join(dbPath, gtools.TypeName(t)))
	if nil != err {
		return nil, err
	}
	for _, f := range files {
		v, err := one(f.Name(), t)
		if nil != err {
			return nil, err
		}
		result = append(result, v)
	}
	return result, nil
}

// 根据key和t 删除对应的数据
func Del(key string, t reflect.Type) error {
	mux.Lock()
	defer mux.Unlock()

	if dbPath == "" {
		return ErrorNeedInit
	}

	return os.Remove(filepath.Join(dbPath, gtools.TypeName(t), key))
}


