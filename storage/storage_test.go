package storage

import (
	"maps"
	"path/filepath"
	"testing"
)

func makeTempDatabase(t *testing.T) *Database {
	t.Helper()
	path := filepath.Join(t.TempDir(), "data.db")
	d, err := Init(path)

	if err != nil {
		t.Fatalf("Не удалось инициализировать базу данных, %s", err)
	}

	return d
}

func TestSet(t *testing.T) {
	d := makeTempDatabase(t)

	tests := []struct {
		name          string
		key, value    string
		expectedError string
	}{
		{"Обычное сохранение", "1234", "4321", ""},
		{"Попытка записи такого же ключа", "1234", "4567", "Такой ключ уже существует"},
		{"Пустой ключ", "", "1234", "Ошибка: пустой ввод"},
		{"Пустое значение", "1234", "", "Ошибка: пустой ввод"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := d.Set(tt.key, tt.value)

			if tt.expectedError != "" {
				if err == nil || err.Error() != tt.expectedError {
					t.Fatalf("Ожидал ошибку %s, получил %s", tt.expectedError, err)
				}
				return
			}

			if err != nil {
				t.Fatalf("Не ожидал ошибки, но получил: %s", err)
			}

			getRes, _ := d.Get(tt.key)
			if getRes != tt.value {
				t.Errorf("После Set ожидал значение %s, но Get вернул %s", tt.value, getRes)
			}
		})
	}
}

func TestGet(t *testing.T) {
	tests := []struct {
		name          string
		key           string
		expectedValue string
		expectedError string
	}{
		{"Обычное получение", "1234", "4321", ""},
		{"Несуществующий ключ", "9876", "", "Ошибка: не найдено"},
		{"Пустой ключ", "", "", "Ошибка: пустой ввод"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := makeTempDatabase(t)

			if tt.expectedValue != "" {
				d.Set(tt.key, tt.expectedValue)
			}

			res, err := d.Get(tt.key)

			if tt.expectedError != "" {
				if err == nil || err.Error() != tt.expectedError {
					t.Fatalf("Ожидал ошибку %s, получил %s", tt.expectedError, err)
				}
				return
			}

			if err != nil {
				t.Fatalf("Не ожидал ошибку, но получил %s", err)
			}

			if res != tt.expectedValue {
				t.Errorf("Ожидал %s, но получил %s", tt.expectedValue, res)
			}
		})
	}
}

func TestList(t *testing.T) {
	tests := []struct {
		name           string
		keys           []string
		values         []string
		expectedResult map[string]string
		expectedError  string
	}{
		{"Обычный вывод", []string{"a", "b", "c"}, []string{"1", "2", "3"}, map[string]string{"a": "1", "b": "2", "c": "3"}, ""},
		{"Пустая бд", []string{}, []string{}, map[string]string{}, "Ошибка: база данных пуста"},
		{"попытка дубликата", []string{"a", "b", "a"}, []string{"1", "2", "3"}, map[string]string{"a": "1", "b": "2"}, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := makeTempDatabase(t)

			if tt.keys != nil && tt.values != nil {
				for i, _ := range tt.keys {
					d.Set(tt.keys[i], tt.values[i])
				}
			}

			res, err := d.List()

			if tt.expectedError != "" {
				if err == nil || err.Error() != tt.expectedError {
					t.Fatalf("Ожидал ошибку %s, получил %s", tt.expectedError, err)
				}
				return
			}

			if err != nil {
				t.Fatalf("Не ожидал ошибку, но получил %s", err)
			}

			if !maps.Equal(tt.expectedResult, res) {
				t.Errorf("Ожидал %v, получил %v", tt.expectedResult, res)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	tests := []struct {
		name           string
		keys           []string
		values         []string
		deletedKey     string
		expectedResult map[string]string
		expectedError  string
	}{
		{"Обычный вывод", []string{"a", "b", "c"}, []string{"1", "2", "3"}, "a", map[string]string{"b": "2", "c": "3"}, ""},
		{"Пустой ввод", []string{}, []string{}, "", map[string]string{}, "Ошибка: пустой ввод"},
		{"Пустая бд", []string{}, []string{}, "a", map[string]string{}, "Ошибка: не найдено"},
		{"Несуществующий ключ", []string{}, []string{}, "a", map[string]string{}, "Ошибка: не найдено"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := makeTempDatabase(t)

			if tt.keys != nil && tt.values != nil {
				for i, _ := range tt.keys {
					d.Set(tt.keys[i], tt.values[i])
				}
			}

			err := d.Delete(tt.deletedKey)

			if tt.expectedError != "" {
				if err == nil || err.Error() != tt.expectedError {
					t.Fatalf("Ожидал ошибку %s, получил %s", tt.expectedError, err)
				}
				return
			}

			if err != nil {
				t.Fatalf("Не ожидал ошибку но получил %s", err)
			}

			newDB, err := Init(d.filepath)
			if err != nil {
				t.Fatalf("Не удалось инициализировать базу заново: %s", err)
			}

			res, _ := newDB.List()

			if !maps.Equal(tt.expectedResult, res) {
				t.Errorf("Ожидал %v, получил %v", tt.expectedResult, res)
			}
		})
	}
}
