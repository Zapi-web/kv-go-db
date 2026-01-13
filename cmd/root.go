package cmd

import (
	"db/storage"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

var d *storage.Database

var rootCmd = &cobra.Command{
	Use:   "db",
	Short: "Simple KV database",

	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		var err error

		err = godotenv.Load(".env")

		if err != nil {
			fmt.Println("Ошибка загрузки файла .env")
			os.Exit(1)
		}

		fileP := os.Getenv("FILEPATH")

		d, err = storage.Init(fileP)
		if err != nil {
			fmt.Println("Ошибка инициализации БД:", err)
			os.Exit(1)
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
