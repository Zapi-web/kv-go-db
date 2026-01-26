package cmd

import (
	"fmt"
	"os"

	"github.com/Zapi-web/kv-go-db/internal/logger"
	"github.com/Zapi-web/kv-go-db/storage"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
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

		logLevel := os.Getenv("LOG_LEVEL")
		logger.Init(logLevel)

		logger.Log.Info("Логгер инициализирован", zap.String("level", logLevel))

		fileP := os.Getenv("FILEPATH")

		d, err = storage.Init(fileP, logger.Log)
		if err != nil {
			logger.Log.Error("Ошибка инициализации БД", zap.Error(err), zap.String("path", fileP))
			os.Exit(1)
		}
	},
}

func Execute() {
	defer func() {
		if d != nil {
			d.Close()
		}
		err := logger.Log.Sync()

		if err != nil {
			fmt.Printf("Ошибка синхронизации логера %v", err)
		}
	}()

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
