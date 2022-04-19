package utils

import (
	"fmt"
	"os"
)

// ReadKey 读取key文件
func ReadKey(key string) []byte {
	keyExist := FileIsExist(key)
	if !keyExist {
		fmt.Printf("The key [%v] is not exist.\n", key)
		os.Exit(1)
	}

	kFile, err := os.Open(key)
	if err != nil {
		fmt.Printf("[E] open key file error: %v", err.Error())
		os.Exit(1)
	}
	defer kFile.Close()

	// read key
	var tmp = make([]byte, 128)
	n, err := kFile.Read(tmp[:])
	if err != nil {
		fmt.Printf("[E] read key file error: %v\n", err.Error())
	}
	fmt.Printf("[I] read key file %d bytes.\n", n)
	return tmp[0:n]
}
