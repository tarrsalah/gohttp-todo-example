package db

import (
	"fmt"
	"github.com/jmoiron/modl"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var (
	db  *sqlx.DB
	Map *modl.DbMap
)

type Task struct {
	ID   int    `json:"id"`
	Todo string `json:"todo"`
	Done bool   `json:"done"`
}

func NewTask(todo string) *Task {
	return &Task{
		Todo: todo,
		Done: false,
	}
}

func NewTaskWithDone(todo string, done bool) *Task {
	return &Task{
		Todo: todo,
		Done: done,
	}
}

func (task *Task) String() string {
	return fmt.Sprintf("[id=%v, todo=\"%s\", done=%v]",
		task.ID,
		task.Todo,
		task.Done)
}

// BootstrapDB bring some data to the database for testing purposes
func BootStrap() {
	Map.Dbx.MustExec("drop table if exists task")
	Map.AddTable(Task{}, "task").SetKeys(true, "id")
	err := Map.CreateTablesIfNotExists()
	if err != nil {
		log.Fatal("Database not creatable: ", err)
	}

	for _, task := range []*Task{
		NewTask("do"),
		NewTask("do do"),
		NewTask("do do do "),
		NewTask("do do do do"),
	} {
		Map.Insert(task)
	}
}

func init() {
	db_file := "./todo.db"
	db = sqlx.MustConnect("sqlite3", db_file)
	// Run DB.Close() in the main function
	Map = modl.NewDbMap(db.DB, modl.SqliteDialect{})
	// Add Tables
	Map.AddTable(Task{}, "task").SetKeys(true, "id")
	// Migrate
	err := Map.CreateTablesIfNotExists()
	if err != nil {
		log.Fatal("Database not creatable: ", err)
	}
}
