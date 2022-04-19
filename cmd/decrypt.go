package cmd

import (
	"bufio"
	"encrypt/utils"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var (
	dOutputPath string
)

var decryptCmd = &cobra.Command{
	Use:     "decrypt",
	Short:   "Decryption a file",
	Long:    `Decryption a file with key.`,
	Aliases: []string{"d", "de"},
	Run:     decryptFile,
	Example: `  # Decrypt file
  aes d ~/secret.key /data/photos/1.en -o /data/photos/1.jpg`,
}

func init() {
	RootCmd.AddCommand(decryptCmd)
	decryptCmd.Flags().StringVarP(&dOutputPath, "output", "o", "", "the decrypted file output path")
}

func decryptFile(cmd *cobra.Command, args []string) {
	// 校验参数
	if len(args) != 2 {
		fmt.Println(`parameters are missing
run:
  aes decrypt -h
get help.`)
		os.Exit(1)
		return
	}

	key := args[0]
	fmt.Printf("[I] key: %s\n", key)
	keyBytes := utils.ReadKey(key)

	file := args[1]
	fmt.Printf("[I] encrypt file: %s\n", file)
	fileExist := utils.FileIsExist(file)
	if !fileExist {
		fmt.Printf("The file [%v] is not exist\n", file)
		os.Exit(1)
		return
	}
	fmt.Printf("[I] output: %s\n", dOutputPath)

	fmt.Println("[I] Start decryption file...")
	f, err := os.Open(file)
	if err != nil {
		fmt.Printf("[E] open file error: %v\n", err.Error())
		os.Exit(1)
	}
	defer f.Close()

	info, _ := f.Stat()
	fileLength := info.Size()
	fmt.Printf("[I] file length: %d\n", fileLength)
	// 每100MB进行一次解密
	maxLen := 1024 * 1024 * 100
	var forNum int64 = 0
	getLen := fileLength

	if fileLength > int64(maxLen) {
		getLen = int64(maxLen)
		forNum = fileLength / int64(maxLen)
		fmt.Printf("[I] decryption %d times\n", forNum+1)
	}

	// 输出文件名
	if dOutputPath == "" {
		dOutputPath = file[0:strings.LastIndex(file, ".")]
	}
	// 判断是否已存在解密文件，防止追加覆盖
	if utils.FileIsExist(dOutputPath) {
		fmt.Printf("[E] output file %s exist.\n", dOutputPath)
		os.Exit(1)
	}
	// 解密后存储的文件
	ff, err := os.OpenFile(dOutputPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("[E] decrypt file write error: %v\n", err.Error())
		os.Exit(1)
	}
	defer ff.Close()

	// 循环解密，并写入文件
	for i := 0; i < int(forNum+1); i++ {
		array := make([]byte, getLen)
		num, err := f.Read(array)
		if err != nil {
			fmt.Printf("[E] file read error: %v\n", err.Error())
			os.Exit(1)
		}
		getByte, err := utils.AesDecrypt(array[:num], keyBytes)
		if err != nil {
			fmt.Printf("[E] encryption error: %v\n", err.Error())
			os.Exit(1)
		}
		// 写入
		buf := bufio.NewWriter(ff)
		buf.Write(getByte)
		buf.Flush()
	}
	ffInfo, _ := ff.Stat()
	fmt.Printf("[I] success, decryptFile: %s, size: %v Byte\n", ffInfo.Name(), ffInfo.Size())
}
