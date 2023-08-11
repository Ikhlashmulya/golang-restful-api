package usecase

import (
	"github.com/Ikhlashmulya/golang-restful-api/entity"
	"github.com/Ikhlashmulya/golang-restful-api/model"
)

func toTodoResponse(todo *entity.Todo) model.TodoResponse {
	return model.TodoResponse{
		Id:   todo.Id,
		Name: todo.Name,
	}
}
