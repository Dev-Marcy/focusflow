package main

import (
	"encoding/json"
	"errors"
	"os"
	"sync"
	"time"
)

const dataFilePath = "focusflow_data.json"

type Task struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
}

type Habit struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Streak     int       `json:"streak"`
	LastCheck  time.Time `json:"last_check"`
	CreatedAt  time.Time `json:"created_at"`
}

type StorageData struct {
	Tasks  []Task  `json:"tasks"`
	Habits []Habit `json:"habits"`
}

type Storage struct {
	mu sync.Mutex
}

func NewStorage() *Storage {
	return &Storage{}
}

func (s *Storage) Load() (*StorageData, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, err := os.Stat(dataFilePath); errors.Is(err, os.ErrNotExist) {
		return &StorageData{
			Tasks:  []Task{},
			Habits: []Habit{},
		}, nil
	}

	bytes, err := os.ReadFile(dataFilePath)
	if err != nil {
		return nil, err
	}

	var data StorageData
	if err := json.Unmarshal(bytes, &data); err != nil {
		return nil, err
	}

	return &data, nil
}

func (s *Storage) Save(data *StorageData) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	bytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(dataFilePath, bytes, 0644)
}
