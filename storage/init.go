package storage

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"go.uber.org/zap"
)

func Init(filepath string, l *zap.Logger) (*Database, error) {
	start := time.Now()

	var newDatabase Database = Database{
		storage:  make(map[string]string),
		filepath: filepath,
		logger:   l,
	}

	_, err := os.Stat(filepath)

	if errors.Is(err, os.ErrNotExist) {
		newDatabase.file, err = os.Create(filepath)

		if err != nil {
			return nil, fmt.Errorf("Ошибка создания файла: %w", err)
		}

		newDatabase.filepath = filepath

		return &newDatabase, nil
	}

	newDatabase.file, err = os.OpenFile(filepath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)

	if err != nil {
		return &newDatabase, fmt.Errorf("Ошибка открытия файла: %w", err)
	}

	var totalLines, tombstones int

	scanner := bufio.NewScanner(newDatabase.file)

	for scanner.Scan() {
		line := scanner.Text()
		lineSplited := strings.SplitN(line, "=", 2)

		totalLines++
		if len(lineSplited) < 2 {
			newDatabase.logger.Warn("corrupted line",
				zap.String("line", line),
				zap.Int("line_number", totalLines),
			)
			continue
		}

		if lineSplited[1] == "__TOMBSTONE__" {
			delete(newDatabase.storage, lineSplited[0])
			tombstones++
			continue
		}

		newDatabase.storage[lineSplited[0]] = lineSplited[1]
	}

	newDatabase.filepath = filepath
	duration := time.Since(start)

	newDatabase.logger.Info("start info",
		zap.Int("total lines", totalLines),
		zap.Int("active keys", len(newDatabase.storage)),
		zap.Int("tombstones", tombstones),
		zap.Duration("duration", duration),
	)

	return &newDatabase, nil
}
