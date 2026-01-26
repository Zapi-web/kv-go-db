package storage

import (
	"errors"
	"fmt"
	"os"
	"sync"
	"time"

	"go.uber.org/zap"
)

type Database struct {
	storage  map[string]string
	mu       sync.RWMutex
	file     *os.File
	filepath string
	logger   *zap.Logger
}

func (d *Database) Set(key string, value string) error {
	start := time.Now()

	d.mu.Lock()
	defer d.mu.Unlock()

	if key == "" || value == "" {
		return errors.New("ошибка: пустой ввод")
	}

	if _, ok := d.storage[key]; ok {
		return errors.New("такой ключ уже существует")
	}

	_, err := d.file.WriteString(key + "=" + value + "\n")

	if err != nil {
		return fmt.Errorf("ошибка записи в файл %w", err)
	}

	err = d.file.Sync()

	if err != nil {
		return fmt.Errorf("ошибка синхронизации диска %w", err)
	}

	d.storage[key] = value
	d.logger.Debug("Запись добавлена", zap.String("key", key),
		zap.Duration("Duration", time.Since(start)),
	)
	return nil
}

func (d *Database) Get(key string) (string, error) {
	d.mu.RLock()
	defer d.mu.RUnlock()

	if key == "" {
		return "", errors.New("ошибка: пустой ввод")
	}

	value, ok := d.storage[key]

	if !ok {
		return "", errors.New("ошибка: не найдено")
	}

	return value, nil
}

func (d *Database) List() (map[string]string, error) {
	d.mu.RLock()
	defer d.mu.RUnlock()

	if len(d.storage) == 0 {
		return nil, errors.New("ошибка: база данных пуста")
	}

	var storageCopy = make(map[string]string)

	for key, value := range d.storage {
		storageCopy[key] = value
	}

	return storageCopy, nil
}

func (d *Database) Delete(key string) error {
	start := time.Now()

	d.mu.Lock()
	defer d.mu.Unlock()

	if key == "" {
		return errors.New("ошибка: пустой ввод")
	}

	if _, ok := d.storage[key]; !ok {
		return errors.New("ошибка: не найдено")
	}

	delete(d.storage, key)

	_, err := d.file.WriteString(key + "=__TOMBSTONE__" + "\n")

	if err != nil {
		return fmt.Errorf("ошибка записи TOMBSTONE метки %w", err)
	}

	err = d.file.Sync()

	if err != nil {
		return fmt.Errorf("ошибка синхронизации диска %w", err)
	}

	d.logger.Debug("Запись удалена",
		zap.String("key", key),
		zap.Duration("duration", time.Since(start)),
	)

	return nil
}

func (d *Database) Close() error {
	d.logger.Info("database connection closed", zap.String("path", d.filepath))
	if d.file != nil {
		return d.file.Close()
	}
	return nil
}
