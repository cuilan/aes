package cmd

import (
	"bufio"
	"encrypt/utils"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	eOutputPath string
)

var encryptCmd = &cobra.Command{
	Use:     "encrypt",
	Short:   "Encryption a file",
	Long:    `Encryption a file with key.`,
	Aliases: []string{"e", "en"},
	Run:     encryptFile,
	Example: `  # Encrypt file
  aes e ~/secret.key /data/photos/1.png -o /data/photos`,
}

func init() {
	RootCmd.AddCommand(encryptCmd)
	encryptCmd.Flags().StringVarP(&eOutputPath, "output", "o", "", "the encrypted file output path")
}

func encryptFile(cmd *cobra.Command, args []string) {
	// 校验参数
	if len(args) != 2 {
		fmt.Println(`parameters are missing
run:
  aes encrypt -h
get help.`)
		os.Exit(1)
		return
	}

	key := args[0]
	fmt.Printf("[I] key: %s\n", key)
	keyBytes := utils.ReadKey(key)

	file := args[1]
	fmt.Printf("[I] file: %s\n", file)
	fileExist := utils.FileIsExist(file)
	if !fileExist {
		fmt.Printf("The file [%v] is not exist.\n", file)
		os.Exit(1)
		return
	}
	fmt.Printf("[I] output: %s\n", eOutputPath)

	fmt.Println("[I] Start encryption file...")
	f, err := os.Open(file)
	if err != nil {
		fmt.Printf("[E] open file error: %v", err.Error())
		os.Exit(1)
	}
	defer f.Close()

	info, _ := f.Stat()
	fileLength := info.Size()
	fmt.Printf("[I] file length: %d\n", fileLength)
	// 每100MB进行一次加密
	maxLen := 1024 * 1024 * 100
	var forNum int64 = 0
	getLen := fileLength

	if fileLength > int64(maxLen) {
		getLen = int64(maxLen)
		forNum = fileLength / int64(maxLen)
		fmt.Printf("[I] encryption %d times\n", forNum+1)
	}

	// 输出文件名
	if eOutputPath == "" {
		eOutputPath = file + ".encrypt"
	}
	// 判断是否已存在加密文件，防止追加覆盖
	if utils.FileIsExist(eOutputPath) {
		fmt.Printf("[E] output file %s exist\n", eOutputPath)
		os.Exit(1)
	}
	// 加密后存储的文件
	ff, err := os.OpenFile(eOutputPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("[E] encrypt file write error: %v\n", err.Error())
		os.Exit(1)
	}
	defer ff.Close()

	// 循环加密，并写入文件
	for i := 0; i < int(forNum+1); i++ {
		array := make([]byte, getLen)
		num, err := f.Read(array)
		if err != nil {
			fmt.Printf("[E] file read error: %v\n", err.Error())
			os.Exit(1)
		}
		getByte, err := utils.AesEncrypt(array[:num], keyBytes)
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
	fmt.Printf("[I] success, encryptFile: %s, size: %v Byte\n", ffInfo.Name(), ffInfo.Size())
}
