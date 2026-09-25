package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	storage := NewStorage()
	todoMgr := NewTodoManager(storage)
	habitMgr := NewHabitManager(storage)
	timer := NewFocusTimer()

	subcommand := os.Args[1]

	switch subcommand {
	case "todo":
		handleTodoCommands(todoMgr, os.Args[2:])
	case "habit":
		handleHabitCommands(habitMgr, os.Args[2:])
	case "timer":
		handleTimerCommands(timer, os.Args[2:])
	default:
		fmt.Printf("Unknown subcommand: %s\n", subcommand)
		printUsage()
		os.Exit(1)
	}
}

func handleTodoCommands(mgr *TodoManager, args []string) {
	if len(args) < 1 {
		fmt.Println("Usage: focusflow todo [add|list|done|delete]")
		return
	}

	switch args[0] {
	case "add":
		if len(args) < 2 {
			fmt.Println("Error: Task description required.")
			return
		}
		if err := mgr.Add(args[1]); err != nil {
			fmt.Printf("Error adding task: %v\n", err)
		}
	case "list":
		if err := mgr.List(); err != nil {
			fmt.Printf("Error listing tasks: %v\n", err)
		}
	case "done":
		if len(args) < 2 {
			fmt.Println("Error: Task ID required.")
			return
		}
		id, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("Error: Invalid task ID.")
			return
		}
		if err := mgr.Complete(id); err != nil {
			fmt.Printf("Error completing task: %v\n", err)
		}
	case "delete":
		if len(args) < 2 {
			fmt.Println("Error: Task ID required.")
			return
		}
		id, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("Error: Invalid task ID.")
			return
		}
		if err := mgr.Delete(id); err != nil {
			fmt.Printf("Error deleting task: %v\n", err)
		}
	default:
		fmt.Println("Unknown todo command. Use: add, list, done, or delete.")
	}
}

func handleHabitCommands(mgr *HabitManager, args []string) {
	if len(args) < 1 {
		fmt.Println("Usage: focusflow habit [add|list|check]")
		return
	}

	switch args[0] {
	case "add":
		if len(args) < 2 {
			fmt.Println("Error: Habit name required.")
			return
		}
		if err := mgr.Add(args[1]); err != nil {
			fmt.Printf("Error adding habit: %v\n", err)
		}
	case "list":
		if err := mgr.List(); err != nil {
			fmt.Printf("Error listing habits: %v\n", err)
		}
	case "check":
		if len(args) < 2 {
			fmt.Println("Error: Habit ID required.")
			return
		}
		id, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("Error: Invalid habit ID.")
			return
		}
		if err := mgr.CheckIn(id); err != nil {
			fmt.Printf("Error checking in: %v\n", err)
		}
	default:
		fmt.Println("Unknown habit command. Use: add, list, or check.")
	}
}

func handleTimerCommands(timer *FocusTimer, args []string) {
	timerCmd := flag.NewFlagSet("timer", flag.ExitOnError)
	work := timerCmd.Int("work", 25, "Work duration in minutes")
	rest := timerCmd.Int("break", 5, "Break duration in minutes")

	if len(args) > 0 && args[0] == "start" {
		timerCmd.Parse(args[1:])
		timer.Start(*work, *rest)
	} else {
		fmt.Println("Usage: focusflow timer start [-work mins] [-break mins]")
	}
}

func printUsage() {
	fmt.Println(`FocusFlow - CLI Productivity Suite

Usage:
  focusflow <command> <subcommand> [arguments]

Commands:
  todo      Manage to-do tasks (add, list, done, delete)
  habit     Track daily habits (add, list, check)
  timer     Run focus countdown (start -work 25 -break 5)`)
}
