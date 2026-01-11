package storage

import (
	"errors"
	"fmt"
	"os"
	"sync"
)

type Database struct {
	storage  map[string]string
	mu       sync.RWMutex
	filepath string
}

func (d *Database) Set(key string, value string) error {
	d.mu.Lock()
	defer d.mu.Unlock()

	if key == "" || value == "" {
		return errors.New("Ошибка: пустой ввод")
	}

	if _, ok := d.storage[key]; ok {
		return errors.New("Такой ключ уже существует")
	}

	file, err := os.OpenFile(d.filepath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)

	if err != nil {
		return fmt.Errorf("не удалось открыть или создать файл: %w", err)
	}

	defer file.Close()

	file.WriteString(key + "=" + value + "\n")
	d.storage[key] = value

	return nil
}

func (d *Database) Get(key string) (string, error) {
	d.mu.RLock()
	defer d.mu.RUnlock()

	if key == "" {
		return "", errors.New("Ошибка: пустой ввод")
	}

	value, ok := d.storage[key]

	if !ok {
		return "", errors.New("Ошибка: не найдено")
	}

	return value, nil
}

func (d *Database) List() (map[string]string, error) {
	d.mu.RLock()
	defer d.mu.RUnlock()

	if len(d.storage) == 0 {
		return nil, errors.New("Ошибка: база данных пуста")
	}

	var storageCopy = make(map[string]string)

	for key, value := range d.storage {
		storageCopy[key] = value
	}

	return storageCopy, nil
}
