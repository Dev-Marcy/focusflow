package main

import (
	"fmt"
	"time"
)

type FocusTimer struct{}

func NewFocusTimer() *FocusTimer {
	return &FocusTimer{}
}

func (t *FocusTimer) Start(workMinutes, breakMinutes int) {
	fmt.Printf("\n Starting Focus Session: %d minutes\n", workMinutes)
	t.runCountdown(workMinutes, "Focus Time")

	fmt.Printf("\n Focus Session complete! Starting Break: %d minutes\n", breakMinutes)
	t.runCountdown(breakMinutes, "Break Time")

	fmt.Println("\n Full session finished! Great work.")
}

func (t *FocusTimer) runCountdown(minutes int, phase string) {
	totalSeconds := minutes * 60

	for totalSeconds > 0 {
		m := totalSeconds / 60
		s := totalSeconds % 60

		// ANSI escape sequence to rewrite line in terminal
		fmt.Printf("\r[%s] Remaining: %02d:%02d ", phase, m, s)
		time.Sleep(1 * time.Second)
		totalSeconds--
	}
	fmt.Printf("\r[%s] Remaining: 00:00 - Done!\n", phase)
}
