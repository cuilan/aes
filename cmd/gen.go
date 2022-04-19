package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"math/rand"
	"os"
	"time"
)

const BASE = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!@#$^&*()_+~[]{};,.<>?"

var genCmd = &cobra.Command{
	Use:     "generate",
	Short:   "Generate an AES key",
	Long:    "Generate an AES key, and save to path.",
	Aliases: []string{"g", "ge", "gen"},
	Run:     generateKey,
	Example: `  # Generate an AES key to save path with Linux/MacOS
  aes gen /tmp/keys/
  aes g ~/data

  # Generate an AES key to save path with Windows
  aes gen D://data/
  aes g D://data`,
}

func init() {
	RootCmd.AddCommand(genCmd)
}

// generateKey 16,24,32位字符串，分别对应 AES-128，AES-192，AES-256 加密算法
func generateKey(cmd *cobra.Command, args []string) {
	fmt.Println("[I] Generate an AES key...")
	key := ""
	// 随机种子
	rand.Seed(time.Now().Unix())
	// 生成 32 个 [0, 83) 范围的伪随机数。
	for i := 0; i < 32; i++ {
		key = key + string(BASE[rand.Intn(83)])
	}
	fmt.Println(key)
	keyPath := time.Now().Format("20060102150405") + ".key"
	if len(args) != 0 {
		keyPath = args[0] + "/" + keyPath
	}
	f, err := os.OpenFile(keyPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("[E] key file write error: %v", err.Error())
		os.Exit(1)
		return
	}
	defer f.Close()
	// 写入
	buf := bufio.NewWriter(f)
	buf.WriteString(key)
	buf.Flush()
}
