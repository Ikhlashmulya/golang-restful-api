package app

import (
	"net/http"

	handler "github.com/Ikhlashmulya/golang-restful-api/delivery/http"
	"github.com/Ikhlashmulya/golang-restful-api/exception"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(todoHandler *handler.TodoHandler) http.Handler {
	router := httprouter.New()

	router.GET("/api/todos", todoHandler.GetAll)
	router.DELETE("/api/todos/:id", todoHandler.Delete)
	router.POST("/api/todos", todoHandler.Create)

	router.PanicHandler = exception.PanicHandler

	return router
}
