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

var UserFilter user.Filter

// getUserCmd represents the getUser command
var getUserCmd = &cobra.Command{
	Use:   "getUser",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		user.FilterUser(UserFilter, false)
	},
}

func NewGetUserCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "getUser",
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, args []string) {
			log := user.FilterUser(UserFilter, false)
			fmt.Fprintf(cmd.OutOrStdout(), log)
		},
	}
}


func init() {
	getUserCmd.Flags().StringVarP(&UserFilter.Gender, "gender", "g", "", "filter user by gender.")
	getUserCmd.Flags().IntVarP(&UserFilter.Age, "age", "a", 0, "filter user by age.")
	getUserCmd.Flags().StringVarP(&UserFilter.AgeOperator, "operator", "o", "", "rules of the age filter. (gt, lt, gte, lte, e)")
	getUserCmd.Flags().StringVarP(&UserFilter.Nationality, "nat", "n", "", "filter user by nationality.")
	getUserCmd.Flags().Uint64VarP(&UserFilter.UserId, "userid", "u", 0, "the userID to be retrieved.")
	rootCmd.AddCommand(getUserCmd)
}
