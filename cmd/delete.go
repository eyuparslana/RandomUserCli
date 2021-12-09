package cmd

import (
	"fmt"
	"user-cli/user"

	"github.com/spf13/cobra"
)

var userId uint64
var all bool

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete user(s)",
	Run: func(cmd *cobra.Command, args []string) {
		log, err := user.DeleteUser(userId, all)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Fprintf(cmd.OutOrStdout(), log)
		}
	},
}

func NewDeleteCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "delete",
		Short: "delete user(s)",
		Run: func(cmd *cobra.Command, args []string) {
			log, err := user.DeleteUser(userId, all)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Fprintf(cmd.OutOrStdout(), log)
			}
		},
	}
}

func init() {
	deleteCmd.Flags().Uint64VarP(&userId, "userid", "u", 0, "id of the user to be deleted.")
	deleteCmd.Flags().BoolVarP(&all, "all", "a", false, "deletes all users.")
	rootCmd.AddCommand(deleteCmd)
}
