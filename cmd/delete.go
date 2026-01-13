package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [key]",
	Short: "Remove key (Append-only)",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		key := args[0]

		err := d.Delete(key)

		if err != nil {
			fmt.Println("ошибка удаления:", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
