package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Show all entries in the database",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		data, err := d.List()

		if err != nil {
			fmt.Println("Ошибка вывода данных:", err)
			os.Exit(1)
		}

		for key, value := range data {
			fmt.Println(key + "=" + value)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
