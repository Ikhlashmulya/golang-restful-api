package handler

import (
	"encoding/json"
	"github.com/Ikhlashmulya/golang-restful-api/model"
	"github.com/Ikhlashmulya/golang-restful-api/usecase"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCreate(t *testing.T) {

	ctrl := gomock.NewController(t)
	todoUsecase := usecase.NewMockTodoUsecase(ctrl)

	todoUsecase.EXPECT().Create(gomock.Any(), model.CreateTodoRequest{Name: "Testing"}).
		Return(model.TodoResponse{Id: "1", Name: "Testing"})

	requestBody := strings.NewReader(`{"name":"Testing"}`)

	request := httptest.NewRequest(http.MethodPost, "/api/todos", requestBody)
	recorder := httptest.NewRecorder()

	todoHandler := NewTodoHandler(todoUsecase)
	todoHandler.Create(recorder, request, httprouter.Params{})

	result := recorder.Result()
	responseBody, err := io.ReadAll(result.Body)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusCreated, result.StatusCode)

	var webResponse model.WebResponse
	json.Unmarshal(responseBody, &webResponse)

	assert.Equal(t, http.StatusCreated, webResponse.Code)
	assert.Equal(t, "CREATED", webResponse.Status)

	data := webResponse.Data.(map[string]any)
	assert.Equal(t, "1", data["id"])
	assert.Equal(t, "Testing", data["name"])

	assert.True(t, ctrl.Satisfied())

}

func TestDelete(t *testing.T) {
	ctrl := gomock.NewController(t)
	todoUsecase := usecase.NewMockTodoUsecase(ctrl)

	todoUsecase.EXPECT().Delete(gomock.Any(), "1").Times(1)

	todoId := "1"

	request := httptest.NewRequest(http.MethodPost, "/api/todos"+todoId, nil)
	recorder := httptest.NewRecorder()

	todoHandler := NewTodoHandler(todoUsecase)
	todoHandler.Delete(recorder, request, httprouter.Params{httprouter.Param{Key: "id", Value: todoId}})

	result := recorder.Result()

	assert.Equal(t, http.StatusOK, result.StatusCode)

	responseBody, err := io.ReadAll(result.Body)
	assert.NoError(t, err)

	var webResponse model.WebResponse
	json.Unmarshal(responseBody, &webResponse)

	assert.Equal(t, http.StatusOK, webResponse.Code)
	assert.Equal(t, "OK", webResponse.Status)

	assert.True(t, ctrl.Satisfied())

}

func TestGetAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	todoUsecase := usecase.NewMockTodoUsecase(ctrl)

	todos := []model.TodoResponse{
		{
			Id:   "1",
			Name: "testing 1",
		},
		{
			Id:   "2",
			Name: "testing 2",
		},
	}

	todoUsecase.EXPECT().GetAll(gomock.Any()).Return(todos)

	request := httptest.NewRequest(http.MethodPost, "/api/todos", nil)
	recorder := httptest.NewRecorder()

	todoHandler := NewTodoHandler(todoUsecase)
	todoHandler.GetAll(recorder, request, httprouter.Params{})

	result := recorder.Result()

	assert.Equal(t, http.StatusOK, result.StatusCode)

	requestBody, err := io.ReadAll(result.Body)
	assert.NoError(t, err)

	var webResponse model.WebResponse
	json.Unmarshal(requestBody, &webResponse)

	assert.Equal(t, http.StatusOK, webResponse.Code)
	assert.Equal(t, "OK", webResponse.Status)

	data := webResponse.Data.([]any)

	for i := 0; i < 2; i++ {
		assert.Equal(t, todos[i].Id, data[i].(map[string]any)["id"].(string))
		assert.Equal(t, todos[i].Name, data[i].(map[string]any)["name"].(string))
	}

	assert.True(t, ctrl.Satisfied())

}
