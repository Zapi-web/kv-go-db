package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set [key] [value]",
	Short: "Save the value by key",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		key := args[0]
		value := args[1]

		err := d.Set(key, value)

		if err != nil {
			fmt.Println("Ошибка записи:", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(setCmd)
}
