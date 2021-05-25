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
	"strings"
	"time"

	"github.com/spf13/cobra"
	"hawk.wie.gg/db"
	"hawk.wie.gg/exec"
)

var (
	in  string
	out string
	cat string
)

// entryCmd represents the entry command
var entryCmd = &cobra.Command{
	Use:   "entry",
	Short: "Create a time recording",
	Long: `Create a new time recording and add it to the last used category (unless otherwise specified).
	You can provide start and end times with time strings or natural language (eg. 10 minutes ago).`,
	Aliases: []string{"e"},
	Args:    cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if cat == "" {
			cat = db.GetVal("cat")
		} else {
			db.SetVal("cat", cat)
		}

		entry, err := exec.CreateEntry(strings.Join(args, " "), in, out)
		if err != nil {
			return err
		}

		if err := db.AddEntry(entry, cat); err != nil {
			return err
		}

		fmt.Printf("Started tracking \"%s\" on \"%s\"", strings.Join(args, " "), cat)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(entryCmd)

	entryCmd.Flags().StringVarP(&in, "in", "i", time.Now().UTC().Format("2006-01-02T15:04:05-0700"), "start time")
	entryCmd.Flags().StringVarP(&out, "out", "o", "", "end time")
	entryCmd.Flags().StringVarP(&cat, "category", "c", "", "category to add to")

}
