package gdb

import (
	"github.com/alexniver/gtools"
	"io/ioutil"
	"os"
	"path/filepath"
)

/*
	Gdb 接口
	需要实现2个方法
	Path() 方法代表存储的路径名
	TableName() 方法代表存储的文件名
 */
type Gdb interface {
	Path() string
	TableName() string
}

func Save(gdb Gdb) error {
	buf, err := gtools.EncodeGZip(gdb)
	if nil != err {
		return err
	}
	if _, err := os.Stat(gdb.Path()); os.IsNotExist(err) {
		err := os.MkdirAll(gdb.Path(), os.ModePerm)
		if nil != err {
			return err
		}
	}

	vPath := filepath.Join(gdb.Path(), gdb.TableName())
	err = ioutil.WriteFile(vPath, buf, os.ModePerm)
	if nil != err {
		return err
	}
	return nil
}

func SaveAll(gdbArr []Gdb) error {
	for _, gdb := range gdbArr {
		err := Save(gdb)
		if nil != err {
			return err
		}
	}
	return nil
}

func One(path, tableName string, v interface{}) error{
	vPath := filepath.Join(path, tableName)
	buf, err := ioutil.ReadFile(vPath)
	if nil != err {
		return err
	}

	err = gtools.DecodeUnGzip(buf, v)
	if nil != err {
		return err
	}
	return nil
}


func AllTableName(path string) ([]string, error) {
	var result []string
	files, err := ioutil.ReadDir(path)
	if nil != err {
		return nil, err
	}
	for _, f := range files {
		result = append(result, f.Name())
	}
	return result, nil
}


func Del(path, tableName string) error {
	return os.Remove(filepath.Join(path, tableName))
}


