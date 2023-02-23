package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "addressMask",
	Short: "メールアドレスマスク化コマンド",
	Long:`ファイル内の文字レルをメールアドレスマスク化するプログラム`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}