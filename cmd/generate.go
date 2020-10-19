/*
Package cmd defines commands for app
Copyright © 2020 Vishnu Rajendran vishnraj@umich.edu

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
	"shoe-data-collector/collectors"

	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a JSON file containing the output of collected data",
	Long:  `For the given shoe source(s) this will take any collected data and write it to the provided file.`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		f := cmd.Flags()
		v, err := f.GetString("outfile")
		if err != nil {
			return err
		} else if v == "" {
			return fmt.Errorf("We require a non-empty JSON filename to write to")
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		f := cmd.Flags()
		shoeSource, _ := f.GetString("source")
		shoeType, _ := f.GetString("type")
		outfile, _ := f.GetString("outfile")

		return collectors.GenerateShoeData(shoeSource, shoeType, outfile)
	},
}

func init() {
	collectCmd.AddCommand(generateCmd)
	generateCmd.Flags().StringP("outfile", "o", "", "JSON file to write to")
}
