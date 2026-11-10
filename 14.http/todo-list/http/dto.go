package http

import (
	"encoding/json"
	"errors"
	"time"
)

type TaskDto struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (t TaskDto) ValidateForCreate() error {
	if t.Title == "" {
		return errors.New("title is empty")
	}
	if t.Description == "" {
		return errors.New("description is empty")
	}
	return nil
}

type ErrorDto struct {
	Message string    `json:"message"`
	Time    time.Time `json:"time"`
}

func (e ErrorDto) ToString() string {
	b, err := json.Marshal(e)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func NewErrorDto(message string) ErrorDto {
	return ErrorDto{
		Message: message,
		Time:    time.Now(),
	}
}
