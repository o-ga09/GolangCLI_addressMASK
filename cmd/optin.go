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
	"os"
	"strconv"
	"strings"

	//"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/spf13/cobra"
	"github.com/xuri/excelize/v2"
)

// optinCmd represents the optin command
var optinCmd = &cobra.Command{
	Use:   "optin",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		runCmdOptin_out_Count(args)
	},
}

func init() {
	rootCmd.AddCommand(optinCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// optinCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// optinCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func runCmdOptin_out_Count(args []string) {
	fin, err := excelize.OpenFile(args[0])
	if err != nil {
		fmt.Println("file format is mattched")
		os.Exit(1)
	}

	rows, err := fin.GetRows("オプトアウト分布")
	if err != nil {
		fmt.Println("sheet name is invallid")
		os.Exit(1)
	}

	row_index := 0
	is_found := 0
	for _, row := range rows {
		search_query := row[6]
		if row_index < 1 {
			row_index++
			continue
		}
		for _, searched_col := range rows {
			if strings.Compare(search_query, searched_col[3]) == 0 {
				cellname, _ := excelize.CoordinatesToCellName(9, row_index+1)
				fin.SetCellValue("オプトアウト分布", cellname, searched_col[4])

				cnt1, _ := strconv.ParseFloat(row[7], 64)
				cnt2, _ := strconv.ParseFloat(row[4], 64)
				optin := cnt1 - cnt2
				cellname, _ = excelize.CoordinatesToCellName(10, row_index+1)
				fin.SetCellValue("オプトアウト分布", cellname, optin)
				is_found = 1
				break
			}
		}
		if is_found == 0 {
			cellname, _ := excelize.CoordinatesToCellName(9, row_index+1)
			fin.SetCellValue("オプトアウト分布", cellname, 0)
			cellname, _ = excelize.CoordinatesToCellName(10, row_index+1)
			fin.SetCellValue("オプトアウト分布", cellname, row[7])
		}
		is_found = 0
		row_index++
	}

	err = fin.SaveAs("./Book1.xlsx")
	if err != nil {
		fmt.Println(err)
	}
}
