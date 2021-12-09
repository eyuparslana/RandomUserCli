package cmd

import (
	"fmt"
	"user-cli/user"

	"github.com/spf13/cobra"
)

var AddParams user.Params

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add random users to list",
	Run: func(cmd *cobra.Command, args []string) {
		user.CreateUsers(AddParams, false)
	},
}

func NewAddCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "add",
		Short: "add random users to list",
		Run: func(cmd *cobra.Command, args []string) {
			user.CreateUsers(AddParams, false)
			fmt.Fprintf(cmd.OutOrStdout(), "%s %s %s user(s) added to db.", AddParams.Count, AddParams.Nationality, AddParams.Gender)
		},
	}
}

func init() {
	addCmd.Flags().StringVarP(&AddParams.Gender, "gender", "g", "", "Specifies which gender the user will be generated.")
	addCmd.Flags().StringVarP(&AddParams.Nationality, "nat", "n", "", "Specifies from which country the user will be generated.")
	addCmd.Flags().StringVarP(&AddParams.Count, "count", "c", "500", "Specifies how many users will be generated.")
	rootCmd.AddCommand(addCmd)
}
