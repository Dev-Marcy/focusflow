# FocusFlow

A zero-dependency CLI productivity app written in Go that combines to-dos, daily habit tracking, and a real-time terminal focus timer into a single local tool.

Data is automatically persisted to a local `focusflow_data.json` file.

## Requirements

- Go 1.20 or later

## Installation & Setup

```bash
git clone [https://github.com/yourusername/focusflow.git](https://github.com/yourusername/focusflow.git)
cd focusflow
go build -o focusflow main.go storage.go todo.go habit.go timer.go
```
