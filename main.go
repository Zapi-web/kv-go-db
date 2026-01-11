package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"db/storage"

	"github.com/joho/godotenv"
)

func CommandEnter(d *storage.Database) error {
	reader := bufio.NewReader(os.Stdin)
	for {
		command, err := reader.ReadString('\n')

		if err != nil {
			return fmt.Errorf("Ошибка чтения вввода: %w", err)
		}

		command = strings.TrimSpace(command)

		sliceCommand := strings.Fields(command)

		if len(sliceCommand) == 0 {
			continue
		}

		switch sliceCommand[0] {
		case "SET":
			if len(sliceCommand) < 3 {
				fmt.Println("Ошибка, для команды SET нужно ввести 3 значения")
				continue
			}

			err := d.Set(sliceCommand[1], sliceCommand[2])

			if err != nil {
				fmt.Println("Ошибка записи,", err)
			}
		case "GET":
			if len(sliceCommand) < 2 {
				fmt.Println("Ошибка, для команды GET нужно ввести 2 значения")
				continue
			}

			value, err := d.Get(sliceCommand[1])

			if err != nil {
				fmt.Println("Ошибка получения:", err)
				continue
			}

			fmt.Println(value)
		case "LIST":
			data, err := d.List()

			if err != nil {
				fmt.Println("Ошибка вывода:", err)
			}

			for key, value := range data {
				fmt.Println(key + "=" + value)
			}
		case "EXIT":
			return nil
		default:
			fmt.Println("Команда не поддерживается")
		}
	}
}

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Не удалось загрузить .env файл")
	}

	filepath := os.Getenv("FILEPATH")
	database, err := storage.Init(filepath)

	if err != nil {
		log.Fatal(err)
	}

	err = CommandEnter(database)

	if err != nil {
		log.Fatal(err)
	}
}
