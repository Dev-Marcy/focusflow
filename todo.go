package main

import (
	"fmt"
	"time"
)

type TodoManager struct {
	storage *Storage
}

func NewTodoManager(s *Storage) *TodoManager {
	return &TodoManager{storage: s}
}

func (m *TodoManager) Add(title string) error {
	data, err := m.storage.Load()
	if err != nil {
		return err
	}

	nextID := 1
	for _, t := range data.Tasks {
		if t.ID >= nextID {
			nextID = t.ID + 1
		}
	}

	newTask := Task{
		ID:        nextID,
		Title:     title,
		Completed: false,
		CreatedAt: time.Now(),
	}

	data.Tasks = append(data.Tasks, newTask)
	if err := m.storage.Save(data); err != nil {
		return err
	}

	fmt.Printf("Added task #%d: \"%s\"\n", newTask.ID, newTask.Title)
	return nil
}

func (m *TodoManager) List() error {
	data, err := m.storage.Load()
	if err != nil {
		return err
	}

	if len(data.Tasks) == 0 {
		fmt.Println("No tasks found. Add one with: focusflow todo add \"Task name\"")
		return nil
	}

	fmt.Println("\n Your To-Do List:")
	fmt.Println("--------------------------------------------------")
	for _, t := range data.Tasks {
		status := "[ ]"
		if t.Completed {
			status = "[\u2713]"
		}
		fmt.Printf("%s %d. %s\n", status, t.ID, t.Title)
	}
	fmt.Println()
	return nil
}

func (m *TodoManager) Complete(id int) error {
	data, err := m.storage.Load()
	if err != nil {
		return err
	}

	found := false
	for i, t := range data.Tasks {
		if t.ID == id {
			data.Tasks[i].Completed = true
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("task #%d not found", id)
	}

	if err := m.storage.Save(data); err != nil {
		return err
	}

	fmt.Printf("Completed task #%d!\n", id)
	return nil
}

func (m *TodoManager) Delete(id int) error {
	data, err := m.storage.Load()
	if err != nil {
		return err
	}

	updated := make([]Task, 0, len(data.Tasks))
	found := false
	for _, t := range data.Tasks {
		if t.ID == id {
			found = true
			continue
		}
		updated = append(updated, t)
	}

	if !found {
		return fmt.Errorf("task #%d not found", id)
	}

	data.Tasks = updated
	if err := m.storage.Save(data); err != nil {
		return err
	}

	fmt.Printf("Deleted task #%d.\n", id)
	return nil
}
