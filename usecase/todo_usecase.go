package usecase

import (
	"context"
	"github.com/Ikhlashmulya/golang-restful-api/model"
)

type TodoUsecase interface {
	Create(ctx context.Context, request model.CreateTodoRequest) (response model.TodoResponse)
	Delete(ctx context.Context, todoId string)
	GetById(ctx context.Context, todoId string) (response model.TodoResponse)
	GetAll(ctx context.Context) (responses []model.TodoResponse)
}
