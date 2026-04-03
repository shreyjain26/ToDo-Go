# ToDo-Go

A simple CLI todo manager in Go with local JSON persistence. The JSON is created when running the commands.

`It is a learning project`

## Features

- Add todo
- List todos (pretty table)
- Delete by index
- Toggle complete/uncomplete by index
- Edit title by index
- Persistent storage in todos.json

---

## Prerequisites

- Go 1.18+
- Git (optional)
- Linux/Mac/Windows shell

---

## Installation

```bash
git clone https://github.com/<your-user>/ToDo-Go.git
cd ToDo-Go
go mod tidy
go build -o todo
```

or run without build:

```bash
go run .
```

---

## Usage

### Add
```bash
./todo -add "Buy milk"
```

### List
```bash
./todo -list
```

### Delete
```bash
./todo -del 0
```

### Toggle complete/uncomplete
```bash
./todo -toggle 0
```

### Edit
Correct usage:
```bash
./todo -edit "0:Buy almond milk"
```

> Note: `-edit` uses `id:new_title`.

### Combined commands
`-list` is first-match in switch. Do one action per process or separate flags:
```bash
./todo -edit "0:Buy almond milk"
./todo -list
```

---

## Data file

- Stored in todos.json in working directory.
- `Storage.Load` reads on start, `Storage.Save` writes on exit.

---

## Code structure

- main.go
  - initializes `Todos`
  - loads storage
  - parses flags
  - executes command
  - saves storage
- todo.go
  - `Todo` struct
  - `Todos` type + methods: `add`, `delete`, `toggle`, `edit`, `print`, `validateIndex`
- command.go
  - `CmdFlags`
  - `Execute` orchestrator
- storage.go
  - `Storage[T]` for JSON load/save

---