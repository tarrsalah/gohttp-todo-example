package main

import (
	"encoding/json"
	"fmt"
	"github.com/gohttp/app"
	"github.com/gohttp/logger"
	"github.com/gohttp/response"
	"github.com/tarrsalah/gohttp-todo-example/db"
	"net/http"
	"strconv"
)

func main() {
	db.BootStrap()
	a := app.New()
	a.Use(logger.New())

	// serve static files
	fs := http.FileServer(http.Dir("./static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	a.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/index.html")
	})

	a.Get("/api/tasks", func(w http.ResponseWriter, r *http.Request) {
		var tasks []*db.Task
		err := db.Map.Select(&tasks, "select * from task")
		if err != nil {
			response.InternalServerError(w)
			return
		}
		response.OK(w, tasks)
	})

	a.Get("/api/tasks/:id", func(w http.ResponseWriter, r *http.Request) {
		task := db.NewTask("")
		err := db.Map.Get(task, r.URL.Query().Get(":id"))
		if err != nil {
			response.NotFound(w)
			return
		}
		response.OK(w, task)
	})

	a.Post("/api/tasks", func(w http.ResponseWriter, r *http.Request) {
		task := db.NewTask("")
		dec := json.NewDecoder(r.Body)
		err := dec.Decode(task)
		if err != nil {
			response.InternalServerError(w)
			return
		}
		err = db.Map.Insert(task)
		if err != nil {
			response.InternalServerError(w)
			return
		}
		response.Created(w)
	})

	a.Del("/api/tasks/:id", func(w http.ResponseWriter, r *http.Request) {
		var err error
		task := db.NewTask("")
		task.ID, err = strconv.Atoi(r.URL.Query().Get(":id"))
		if err != nil {
			fmt.Println(err)
			response.InternalServerError(w)
			return
		}
		_, err = db.Map.Delete(task)
		if err != nil {
			fmt.Println(err)
			response.InternalServerError(w)
			return
		}
		response.NoContent(w)

	})
	a.Listen(":3000")
}
