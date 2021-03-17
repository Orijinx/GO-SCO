package main

import (
	"./controllers"
	Lodr "./utils/AppLoader"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"net/http"
	"os"
	M"./app/middlewares"
)

func main() {

	e := godotenv.Load() //Загрузить файл .env
	if e != nil {
		fmt.Print(e)
	}

	router := mux.NewRouter()
	router.Use(M.AuthMiddlewear) // добавляем middleware проверки JWT-токена



	port := os.Getenv("PORT") //Получить порт из файла .env; мы не указали порт, поэтому при локальном тестировании должна возвращаться пустая строка
	if port == "" {
		port = "8000" //localhost
	}

	router.HandleFunc("/",
		controllers.MainView).Methods("GET")

	router.HandleFunc("/add", controllers.UploadFile).Methods("POST")

	router.HandleFunc("/test", controllers.Test).Methods("get")

	router.HandleFunc("/login",controllers.Login).Methods("Post")

	if Lodr.Load() { //Проверка на загрузку необходимых компонентов
		fmt.Println("~Server info~ :" + "localhost:" + port)

		err := http.ListenAndServe(":"+port, router) //Запустите приложение, посетите localhost:8000/api

		if err != nil {
			fmt.Print(err)
		}
	}
}
