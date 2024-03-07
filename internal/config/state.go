package config

import (
	"encoding/json"
	"os"
	"time"
)

// Config структура для хранения последнего ID из внешней базы данных
type State struct {
	LastModTime time.Time
}

var configState *State

func CheckStateFile(path string) error {
	if _, err := os.Stat(path); err == nil {
		return nil
	} else {
		return err
	}
}

// SaveState Фиксация состояния во внешнем файле
// json
func SaveState(path string, tm time.Time) error {
	// Get JSON bytes for slice.
	b, _ := json.Marshal(State{LastModTime: tm})

	// Write entire JSON file.
	err := os.WriteFile(path, b, 0644)
	if err != nil {
		return err
	}
	return nil
}

// ReadState Чтение состояния
func ReadState(path string) (time.Time, error) {
	file, _ := os.ReadFile(path)

	err := json.Unmarshal([]byte(file), &configState)
	if err != nil {

		return time.Now(), err
	}

	return configState.LastModTime, nil
}

// возвращает дескриптор объекта DB
func GetState() *State {
	return configState
}
