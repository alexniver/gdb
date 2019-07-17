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

func All(path string) []Gdb {
	return nil
}

func AllTableName(path string) []string {
	return nil
}

func One(path, tableName string) Gdb {
	return nil
}

func Del(path, tableName string) error {
	return nil
}


