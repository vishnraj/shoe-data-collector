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

	"github.com/spf13/cobra"
)

// collectCmd represents the collect command
var collectCmd = &cobra.Command{
	Use:   "collect",
	Short: "Starts collector(s) on given source(s)",
	Long:  `When collect is run, we may collect on a provided source or we can specify collecting from any supported sources.`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		f := cmd.Flags()
		v, err := f.GetString("source")
		if err != nil {
			return err
		} else if v == "" {
			return fmt.Errorf("We require a non-empty source for collecting shoe data")
		}

		v, err = f.GetString("type")
		if err != nil {
			return err
		} else if v == "" {
			return fmt.Errorf("We require a non-empty shoe type for collecting shoe data")
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(collectCmd)
	collectCmd.PersistentFlags().StringP("source", "s", "", "Specify source")
	collectCmd.PersistentFlags().StringP("type", "t", "", "Specify shoe type")
}
