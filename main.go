package main

import (
	"log"
	"net/http"

	"github.com/Ikhlashmulya/golang-restful-api/app"
	handler "github.com/Ikhlashmulya/golang-restful-api/delivery/http"
	"github.com/Ikhlashmulya/golang-restful-api/exception"
	"github.com/Ikhlashmulya/golang-restful-api/repository"
	"github.com/Ikhlashmulya/golang-restful-api/usecase"
	"github.com/rs/cors"
)

func main() {
	db := app.NewDB()
	todoRepository := repository.NewTodoRepository(db)
	todoUsecase := usecase.NewTodoUsecase(todoRepository)
	todoHandler := handler.NewTodoHandler(todoUsecase)
	router := app.NewRouter(todoHandler)

	server := http.Server{
		Addr:    ":8080",
		Handler: cors.AllowAll().Handler(router),
	}

	log.Println("Server running on port :8080")
	err := server.ListenAndServe()
	exception.PanicIfError(err)
}
