package usecase

import (
	"context"
	"github.com/Ikhlashmulya/golang-restful-api/entity"
	"github.com/Ikhlashmulya/golang-restful-api/exception"
	"github.com/Ikhlashmulya/golang-restful-api/model"
	"github.com/Ikhlashmulya/golang-restful-api/repository"

	"github.com/google/uuid"
)

type TodoUsecaseImpl struct {
	TodoRepository repository.TodoRepository
}

func NewTodoUsecase(todoRepository repository.TodoRepository) *TodoUsecaseImpl {
	return &TodoUsecaseImpl{TodoRepository: todoRepository}
}

func (todousecaseimpl *TodoUsecaseImpl) Create(ctx context.Context, request model.CreateTodoRequest) (response model.TodoResponse) {
	todo := entity.Todo{
		Id:   uuid.NewString(),
		Name: request.Name,
	}

	todousecaseimpl.TodoRepository.Create(ctx, todo)

	return toTodoResponse(&todo)
}

func (todousecaseimpl *TodoUsecaseImpl) Delete(ctx context.Context, todoId string) {
	todo, err := todousecaseimpl.TodoRepository.GetById(ctx, todoId)
	if err != nil {
		panic(exception.NewErrorNotFound(err.Error()))
	}

	todousecaseimpl.TodoRepository.Delete(ctx, todo.Id)
}

func (todousecaseimpl *TodoUsecaseImpl) GetById(ctx context.Context, todoId string) (response model.TodoResponse) {
	todo, err := todousecaseimpl.TodoRepository.GetById(ctx, todoId)
	if err != nil {
		panic(exception.NewErrorNotFound(err.Error()))
	}

	return toTodoResponse(&todo)
}

func (todousecaseimpl *TodoUsecaseImpl) GetAll(ctx context.Context) (responses []model.TodoResponse) {
	todos := todousecaseimpl.TodoRepository.GetAll(ctx)

	for _, todo := range todos {
		responses = append(responses, toTodoResponse(&todo))
	}

	return responses
}
