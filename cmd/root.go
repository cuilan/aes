package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

const version = "1.0"

var RootCmd = &cobra.Command{
	Use:     "aes",
	Short:   "aes tool",
	Long:    `Help you encryption decryption file with AES, or generate a key.`,
	Version: version,
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
