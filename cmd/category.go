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
	"errors"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"hawk.wie.gg/db"
)

var name string

// categoryCmd represents the category command
var categoryCmd = &cobra.Command{
	Use:     "category",
	Short:   "Create a new category",
	Aliases: []string{"c"},
	Long: `Create a new category.
	Categories are used to group your time recordings and every recording needs to have a category assigned to it`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if name == "" && len(args) == 0 {
			return errors.New("Please provide a name.")
		}

		if name != "" {
			if err := db.AddCategory(name); err != nil {
				return err
			}
			db.SetVal("cat", name)
			fmt.Printf("Added category %s", name)
		} else {
			if err := db.AddCategory(strings.Join(args, " ")); err != nil {
				return err
			}
			db.SetVal("cat", strings.Join(args, " "))
			fmt.Printf("Added category %s", strings.Join(args, " "))
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(categoryCmd)
	categoryCmd.Flags().StringVarP(&name, "name", "n", "", "category name")
}
