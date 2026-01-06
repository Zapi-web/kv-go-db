package storage

import (
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

	if _, ok := d.storage[key]; !ok {
		file, err := os.OpenFile(d.filepath, os.O_APPEND|os.O_WRONLY, 0666)

		if err != nil {
			return fmt.Errorf("Ошибка открытия файла %w:", err)
		}
		defer file.Close()

		file.WriteString(key + "=" + value + "\n")
		d.storage[key] = value

		return nil
	}

	fmt.Println("Такой ключ уже существует")
	return nil
}

func (d *Database) Get(key string) string {
	d.mu.RLock()
	defer d.mu.RUnlock()
	value, ok := d.storage[key]

	if !ok {
		return "Не найдено"
	}

	return value
}

func (d *Database) List() {
	d.mu.RLock()
	defer d.mu.RUnlock()

	for key, value := range d.storage {
		fmt.Println(key + "=" + value)
	}
}
