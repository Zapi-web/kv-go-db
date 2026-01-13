package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get [key]",
	Short: "Get value by key",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		key := args[0]

		value, err := d.Get(key)

		if err != nil {
			fmt.Println("Ошибка получения:", err)
			os.Exit(1)
		}

		fmt.Println(value)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
