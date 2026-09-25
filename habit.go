package main

import (
	"fmt"
	"time"
)

type HabitManager struct {
	storage *Storage
}

func NewHabitManager(s *Storage) *HabitManager {
	return &HabitManager{storage: s}
}

func (m *HabitManager) Add(name string) error {
	data, err := m.storage.Load()
	if err != nil {
		return err
	}

	nextID := 1
	for _, h := range data.Habits {
		if h.ID >= nextID {
			nextID = h.ID + 1
		}
	}

	newHabit := Habit{
		ID:        nextID,
		Name:      name,
		Streak:    0,
		CreatedAt: time.Now(),
	}

	data.Habits = append(data.Habits, newHabit)
	if err := m.storage.Save(data); err != nil {
		return err
	}

	fmt.Printf("Started tracking habit #%d: \"%s\"\n", newHabit.ID, newHabit.Name)
	return nil
}

func (m *HabitManager) List() error {
	data, err := m.storage.Load()
	if err != nil {
		return err
	}

	if len(data.Habits) == 0 {
		fmt.Println("No habits tracked. Add one with: focusflow habit add \"Habit name\"")
		return nil
	}

	now := time.Now()
	fmt.Println("\n Daily Habit Tracker:")
	fmt.Println("--------------------------------------------------")
	for _, h := range data.Habits {
		checkedToday := false
		if !h.LastCheck.IsZero() {
			y1, m1, d1 := h.LastCheck.Date()
			y2, m2, d2 := now.Date()
			if y1 == y2 && m1 == m2 && d1 == d2 {
				checkedToday = true
			}
		}

		status := "Pending Today"
		if checkedToday {
			status = "Completed Today"
		}

		fmt.Printf("[%d] %s | Streak: %d day(s) | Status: %s\n", h.ID, h.Name, h.Streak, status)
	}
	fmt.Println()
	return nil
}

func (m *HabitManager) CheckIn(id int) error {
	data, err := m.storage.Load()
	if err != nil {
		return err
	}

	now := time.Now()
	found := false

	for i, h := range data.Habits {
		if h.ID == id {
			found = true

			if !h.LastCheck.IsZero() {
				y1, m1, d1 := h.LastCheck.Date()
				y2, m2, d2 := now.Date()
				if y1 == y2 && m1 == m2 && d1 == d2 {
					return fmt.Errorf("habit #%d already completed today", id)
				}

				// Check if last check was yesterday to increment streak
				yesterday := now.AddDate(0, 0, -1)
				y3, m3, d3 := yesterday.Date()
				if y1 == y3 && m1 == m3 && d1 == d3 {
					data.Habits[i].Streak++
				} else {
					// Reset streak if more than a day passed
					data.Habits[i].Streak = 1
				}
			} else {
				data.Habits[i].Streak = 1
			}

			data.Habits[i].LastCheck = now
			break
		}
	}

	if !found {
		return fmt.Errorf("habit #%d not found", id)
	}

	if err := m.storage.Save(data); err != nil {
		return err
	}

	fmt.Printf("Checked in for habit #%d! Keep up the momentum!\n", id)
	return nil
}
