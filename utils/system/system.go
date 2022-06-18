package system

import (
	"errors"
	"os"
	"strings"
)

func SafeRemove(filename string) (err error) {
	if !IsFile(filename) {
		return errors.New("不能删除文件夹")
	}
	if !strings.Contains(filename, "/tmp/") {
		return errors.New("文件不在tmp目录")
	}
	return os.Remove(filename)
}

// IsDir 判断所给路径是否为文件夹
func IsDir(path string) bool {
	if s, err := os.Stat(path); err != nil {
		return false
	} else {
		return s.IsDir()
	}
}

// IsFile 判断所给路径是否为文件
func IsFile(path string) bool {
	return !IsDir(path)
}
