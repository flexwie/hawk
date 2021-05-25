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
	"plugin"
	"runtime"

	"github.com/spf13/cobra"
	"hawk.wie.gg/db"
	"hawk.wie.gg/exec"
)

var (
	format string
	pretty bool
)

type Formatter interface {
	Format(in string)
}

// listCmd represents the list command
var listEntryCmd = &cobra.Command{
	Use:   "list",
	Short: "Prints all entries",
	RunE: func(cmd *cobra.Command, args []string) error {
		entries, err := db.GetAllEntries()
		if err != nil {
			return err
		}

		switch format {
		case "table":
			table, err := exec.PrintTable(entries)
			if err != nil {
				cobra.CheckErr(err)
			}
			fmt.Println(table)
			break

		case "json":
			json, err := exec.PrintJSON(entries, pretty)
			if err != nil {
				cobra.CheckErr(err)
			}
			fmt.Println(json)
			break

		default:
			if runtime.GOOS == "windows" {
				return errors.New("Plugins are not supported on windows. Please use the provided formatters 'table' or 'json'.")
			}

			p, err := plugin.Open(format)
			if err != nil {
				return err
			}

			symGreeter, err := p.Lookup("Formatter")
			if err != nil {
				return err
			}

			var greeter Formatter
			greeter, ok := symGreeter.(Formatter)
			if !ok {
				fmt.Println("unexpected type from module symbol")
			}

			greeter.Format("test")
		}

		return nil
	},
}

func init() {
	entryCmd.AddCommand(listEntryCmd)
	listEntryCmd.Flags().StringVarP(&format, "format", "f", "table", "specify the formatter")
	listEntryCmd.Flags().BoolVar(&pretty, "pretty", false, "pretty print json")
}
