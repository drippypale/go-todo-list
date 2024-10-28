package model

import (
	"log"
	"strconv"
	"time"
)

type Priority int

const (
	HighPriority Priority = iota // Default Priority is set High if not set
	MidPriority
	LowPriority
)

type Todo struct {
	Id       int
	Task     string
	DueTime  time.Time
	Priority Priority
}

func (todo Todo) FromRecord(record []string) Todo {
	id, _ := strconv.Atoi(record[0])
	task := record[1]
	dueTime, err := time.Parse("2006-01-02 15:04", record[2])
	if err != nil {
		log.Printf("failed to parse dueTime %v: %s", dueTime, err)
	}
	priority, _ := strconv.Atoi(record[3])

	todo.Id = id
	todo.Task = task
	todo.DueTime = dueTime
	todo.Priority = Priority(priority)

	return todo

}

func (todo *Todo) ToRecord() []string {
	record := []string{"", "", "", ""}

	dueTime := ""
	if !todo.DueTime.IsZero() {
		dueTime = todo.DueTime.Format("2006-01-02 15:04")
	}

	record[0] = strconv.Itoa(todo.Id)
	record[1] = todo.Task
	record[2] = dueTime
	record[3] = strconv.Itoa(int(todo.Priority))

	return record
}

func TodoFromRecords(records [][]string) []Todo {
	todoList := make([]Todo, len(records))

	for i, record := range records {
		todoList[i] = Todo{}.FromRecord(record)
	}

	return todoList
}

func TodoToRecords(todoList []Todo) [][]string {
	records := make([][]string, len(todoList))

	for i, todo := range todoList {
		records[i] = todo.ToRecord()
	}

	return records
}
