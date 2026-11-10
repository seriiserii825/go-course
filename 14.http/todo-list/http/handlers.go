package http

import (
	"net/http"

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
func (h *HttpHandlers) HandleNewTask(w http.ResponseWriter, r *http.Request) {

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
