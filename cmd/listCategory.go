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

	"github.com/spf13/cobra"
	"hawk.wie.gg/db"
)

var (
	id bool
)

// listCategoryCmd represents the list command
var listCategoryCmd = &cobra.Command{
	Use:     "list",
	Short:   "List your categories",
	Long:    `List you categories. Append --id to access the database ID as well.`,
	Aliases: []string{"l", "ls"},
	RunE: func(cmd *cobra.Command, args []string) error {
		cats, err := db.GetAllCategories()
		if err != nil {
			return err
		}

		for _, s := range cats {
			if id {
				fmt.Printf("%s\t%s \n", s.Id, s.Name)
			} else {
				fmt.Println(s.Name)
			}
		}

		return nil
	},
}

func init() {
	categoryCmd.AddCommand(listCategoryCmd)

	listCategoryCmd.Flags().BoolVar(&id, "id", false, "print ids")
}
