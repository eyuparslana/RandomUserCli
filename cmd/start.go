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
	"user-cli/user"

	"github.com/spf13/cobra"
)

var Params user.Params

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		user.CreateUsers(Params, true)
	},
}

func NewStartCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "start",
		Short: "command of create user(s)",
		Run: func(cmd *cobra.Command, args []string) {
			user.CreateUsers(Params, true)
			fmt.Fprintf(cmd.OutOrStdout(), "%s user(s) added to db.", Params.Count)
		},
	}
}

func init() {
	startCmd.Flags().StringVarP(&Params.Count, "count", "c", "500", "the count of the users to be created.")
	startCmd.Flags().StringVarP(&Params.Gender, "gender", "g", "", "the gender of the user to be created")
	startCmd.Flags().StringVarP(&Params.Nationality, "nat", "n", "", "the nationality of the user to be created")
	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
