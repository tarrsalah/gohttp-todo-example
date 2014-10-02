package main

import (
	"github.com/gohttp/app"
	"github.com/gohttp/response"
	"github.com/tarrsalah/DooDooDooDoo/db"
	"net/http"
)

func main() {
	db.BootStrap()
	app := app.New()
	app.Get("/api/tasks", getAllTasks)
	// app.Get("/api/tasks/:id", getTask)
	// app.Post("/api/tasks", submitNewTask)
	// app.De("/api/tasks/:id", deleteTask)
	app.Listen(":3000")
}

func getAllTasks(w http.ResponseWriter, r *http.Request) {
	var tasks []*db.Task
	err := db.Map.Select(&tasks, "select * from task")
	if err != nil {
		panic(err)
	}
	response.OK(w, tasks)
}
