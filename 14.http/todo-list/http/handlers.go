package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/seriiserii825/go-course/14.http/todo-list/todo"
)

type HttpHandlers struct {
	todoList *todo.List
}

/*
pattern: /tasks
method: POST
info: JSON in HTTP body request

succeed:
  - status code: 201 Created
  - response body: JSON represented created task

failed:
  - status code: 400, 409, 500
  - response body: JSON with error + time
*/
func NewHttpHandlers(todoList *todo.List) *HttpHandlers {
	return &HttpHandlers{
		todoList: todoList,
	}
}

/*
pattern: /tasks
method: POST
info: JSON in HTTP body request

succeed:
  - status code: 201 Created
  - response body: JSON represented created task

failed:
  - status code: 400, 409, 500
  - response body: JSON with error + time
*/
func (h *HttpHandlers) HandleCreateTask(w http.ResponseWriter, r *http.Request) {
	var taskDto TaskDto
	if err := json.NewDecoder(r.Body).Decode(&taskDto); err != nil {
		errDto := NewErrorDto(err.Error())
		http.Error(w, errDto.ToString(), http.StatusBadRequest)
		return
	}

	if err := taskDto.ValidateForCreate(); err != nil {
		errDto := NewErrorDto(err.Error())
		http.Error(w, errDto.ToString(), http.StatusBadRequest)
		return
	}

	todoTask := todo.NewTask(taskDto.Title, taskDto.Description)
	if err := h.todoList.AddTask(todoTask); err != nil {

		errDto := NewErrorDto(err.Error())

		if errors.Is(err, todo.ErrTaskAlreadyExists) {
			http.Error(w, errDto.ToString(), http.StatusConflict)
		} else {
			http.Error(w, errDto.ToString(), http.StatusInternalServerError)
		}
		return
	}

	b, err := json.Marshal(todoTask)
	if err != nil {
		panic(err)
	}

	w.WriteHeader((http.StatusCreated))
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write response:", err)
		return
	}
}

/*
pattern: /tasks/{title}
method: GET
info: title in URL path

succeed:
  - status code: 200 OK
  - response body: JSON represented task

failed:
  - status code: 400, 404, 500
  - response body: JSON with error + time
*/
func (h *HttpHandlers) HandleGetTask(w http.ResponseWriter, r *http.Request) {
	title := mux.Vars(r)["title"]
	task, err := h.todoList.GetTask(title)
	if err != nil {
		errDto := NewErrorDto(err.Error())

		if errors.Is(err, todo.ErrTaskNotFound) {
			http.Error(w, errDto.ToString(), http.StatusNotFound)
		} else {
			http.Error(w, errDto.ToString(), http.StatusInternalServerError)
		}
		return
	}

	b, err := json.Marshal(task)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write response:", err)
		return
	}
}

/*
pattern: /tasks
method: GET

succeed:
  - status code: 200 OK
  - response body: JSON represented tasks

failed:
  - status code: 400, 500
  - response body: JSON with error + time
*/
func (h *HttpHandlers) HandleGetAllTasks(w http.ResponseWriter, r *http.Request) {
	tasks := h.todoList.ListTasks()
	b, err := json.Marshal(tasks)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write response:", err)
		return
	}
}

func (h *HttpHandlers) HandleAllUncompletedTasks(w http.ResponseWriter, r *http.Request) {

}

func (h *HttpHandlers) HandleCompleteTask(w http.ResponseWriter, r *http.Request) {

}

/*
pattern: /tasks/{title}
method: DELETE
info: pattern

succeed:
  - status code: 204 No Content
  - response body: -

failed:
  - status code: 400, 404, 409, 500
  - response body: JSON with error + time
*/
func (h *HttpHandlers) HandleDeleteTask(w http.ResponseWriter, r *http.Request) {

}
