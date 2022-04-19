package utils

import (
	"fmt"
	"os"
	"sync"
)

// FileIsExist 检查文件是否存在
func FileIsExist(path string) bool {
	lock := sync.RWMutex{}
	lock.Lock()
	defer lock.Unlock()
	f, err := os.Open(path)
	if err != nil {
		return false
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)
	return true
}

// CreatePathIfNotExist 检查目录是否存在，如果不存在则创建
func CreatePathIfNotExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	// 如果不存在
	if os.IsNotExist(err) {
		// 创建目录
		err = os.MkdirAll(path, 0777)
		if err != nil {
			fmt.Printf("创建目录失败，path: %s, error: %v", path, err)
			return false, err
		}
	}
	return true, nil
}
