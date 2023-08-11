package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Ikhlashmulya/golang-restful-api/exception"
	"github.com/Ikhlashmulya/golang-restful-api/model"
	"github.com/Ikhlashmulya/golang-restful-api/usecase"

	"github.com/julienschmidt/httprouter"
)

type TodoHandler struct {
	TodoUsecase usecase.TodoUsecase
}

func NewTodoHandler(todoUsecase usecase.TodoUsecase) *TodoHandler {
	return &TodoHandler{TodoUsecase: todoUsecase}
}

func (todoHandler *TodoHandler) Create(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	var createTodoRequest model.CreateTodoRequest
	err := json.NewDecoder(request.Body).Decode(&createTodoRequest)
	exception.PanicIfError(err)
	defer request.Body.Close()

	response := todoHandler.TodoUsecase.Create(request.Context(), createTodoRequest)

	toResponseJSON(writer, http.StatusCreated, model.WebResponse{
		Code:   http.StatusCreated,
		Status: "CREATED",
		Data:   response,
	})

}

func (todoHandler *TodoHandler) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	todoId := params.ByName("id")

	todoHandler.TodoUsecase.Delete(request.Context(), todoId)

	toResponseJSON(writer, http.StatusOK, model.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
	})

}

func (todoHandler *TodoHandler) GetAll(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	response := todoHandler.TodoUsecase.GetAll(request.Context())

	toResponseJSON(writer, http.StatusOK, model.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   response,
	})

}

func toResponseJSON(writer http.ResponseWriter, code int, data any) {
	writer.Header().Add("content-type", "application/json")
	writer.WriteHeader(code)

	err := json.NewEncoder(writer).Encode(data)
	exception.PanicIfError(err)
}
