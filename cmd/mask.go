/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"

	"github.com/spf13/cobra"
)

// maskCmd represents the mask command
var MaskCmd = &cobra.Command{
	Use:   "mask",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		runCmd_ReplaceSring(os.Args)
	},
}

func runCmd_ReplaceSring(args []string) {
	//引数チェック
	if len(os.Args) < 3 {
		fmt.Println("対象フォルダをしてください")
		os.Exit(1)
	}
	//フォルダを開く
	foldername := args[2]
	files ,_:= ioutil.ReadDir(foldername)

	//フォルダに存在するファイルを取得する
	for _,filename := range files {
		bytes, err := ioutil.ReadFile(foldername + "/" + filename.Name())
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		//読み取った文字列を文字列に変換
		contents := string(bytes)
		regstr := regexp.MustCompile("[\\w\\-._]+@")
		contents = regstr.ReplaceAllString(contents,"XXX@")
		fmt.Println(contents)
		//ファイルに書き込む
		err = ioutil.WriteFile(foldername + filename.Name(),[]byte(contents),os.ModePerm)
		if err != nil {
			fmt.Println(err)
		}
	}


}

func init() {
	rootCmd.AddCommand(MaskCmd)
}