package http

import (
	"encoding/json"
	"net/http"
	"time"

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
		errDto := ErrorDto{
			Message: err.Error(),
			Time:    time.Now(),
		}
		http.Error(w, errDto.ToString(), http.StatusBadRequest)
		return
	}

	if err := taskDto.ValidateForCreate(); err != nil {
		errDto := ErrorDto{
			Message: err.Error(),
			Time:    time.Now(),
		}
		http.Error(w, errDto.ToString(), http.StatusBadRequest)
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
