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
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(&taskDto); err != nil {
		writeError(w, http.StatusBadRequest, fmt.Errorf("failed to decode request body: %w", err))
		return
	}

	if err := taskDto.ValidateForCreate(); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	todoTask := todo.NewTask(taskDto.Title, taskDto.Description)
	if err := h.todoList.AddTask(todoTask); err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, todo.ErrTaskAlreadyExists) {
			status = http.StatusConflict
		}
		writeError(w, status, err)
		return
	}

	writeJSON(w, http.StatusCreated, todoTask)
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
		status := http.StatusInternalServerError

		if errors.Is(err, todo.ErrTaskNotFound) {
			status = http.StatusNotFound
		}
		writeError(w, status, err)
		return
	}

	writeJSON(w, http.StatusOK, task)
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
	writeJSON(w, http.StatusOK, tasks)
}

func (h *HttpHandlers) HandleAllUncompletedTasks(w http.ResponseWriter, r *http.Request) {
	uncompletedTasks := h.todoList.ListUncompletedTasks()
	writeJSON(w, http.StatusOK, uncompletedTasks)
}

func (h *HttpHandlers) HandleAllCompletedTasks(w http.ResponseWriter, r *http.Request) {
	completedTasks := h.todoList.ListCompletedTasks()
	writeJSON(w, http.StatusOK, completedTasks)
}

func (h *HttpHandlers) HandleCompleteTask(w http.ResponseWriter, r *http.Request) {
	var completeDto CompleteTaskDto
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(&completeDto); err != nil {
		writeError(w, http.StatusBadRequest, fmt.Errorf("failed to decode request body: %w", err))
		return
	}

	title := mux.Vars(r)["title"]
	var (
		task todo.Task
		err  error
	)

	if completeDto.Complete {
		task, err = h.todoList.CompleteTask(title)
	} else {
		task, err = h.todoList.UncompleteTask(title)
	}
	if err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, todo.ErrTaskNotFound) {
			status = http.StatusNotFound
		}
		writeError(w, status, err)
		return
	}

	writeJSON(w, http.StatusOK, task)

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
	title := mux.Vars(r)["title"]
	if err := h.todoList.DeleteTask(title); err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, todo.ErrTaskNotFound) {
			status = http.StatusNotFound
		}
		writeError(w, status, err)
		return
	}
}
