/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"github.com/Ad3bay0c/week4Curriculum/model"
	"strconv"

	"github.com/spf13/cobra"
)

// DoneCmd represents the Done command
var DoneCmd = &cobra.Command{
	Use:   "Done",
	Short: "Check a Task as done e.g Done 1",
	Long: `To Check a Task that is done`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 0 {
			index, err := strconv.Atoi(args[0])
			if err != nil {
				panic(err)
			}
			model.Lst.Done(index)
		} else {
			fmt.Println("Please Enter the Index No you Want to Check")
		}

		//fmt.Println("Done called")
	},
}

func init() {
	rootCmd.AddCommand(DoneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// DoneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// DoneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
