package storage

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func Init(filepath string) (*Database, error) {
	var newDatabase Database = Database{
		storage:  make(map[string]string),
		filepath: filepath,
	}

	_, err := os.Stat(filepath)

	if errors.Is(err, os.ErrNotExist) {
		f, err := os.Create(filepath)

		if err != nil {
			return nil, fmt.Errorf("Ошибка создания файла: %w", err)
		}
		f.Close()

		newDatabase.filepath = filepath

		return &newDatabase, nil
	}

	file, err := os.OpenFile(filepath, os.O_RDONLY, 0666)

	if err != nil {
		return &newDatabase, fmt.Errorf("Ошибка открытия файла: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		lineSplited := strings.SplitN(line, "=", 2)

		if len(lineSplited) < 2 {
			continue
		}

		if lineSplited[1] == "__TOMBSTONE__" {
			delete(newDatabase.storage, lineSplited[0])
			continue
		}

		newDatabase.storage[lineSplited[0]] = lineSplited[1]
	}

	newDatabase.filepath = filepath

	return &newDatabase, nil
}
