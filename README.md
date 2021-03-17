# Gorijinx
Gorijinx - это микрофреймворк реализованный на языке Golang.
## Поставки
- Под капотом находится сервер Mux ([github.com/gorilla/mux](http://github.com/gorilla/mux))
- Гибкая настройка с Godotenv ([github.com/joho/godotenv](http://github.com/joho/godotenv))
- Поставляется совместно с CLI-Gori
## Getting started
В _main.go_ находится запуск сервера и рутинг нашего приложения.
```go
router.HandleFunc("/", controllers.MainView).Methods("GET")//Настройка рута
//Вместо контроллера можно использовать обычную функцию
router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
	//
	...
}).Methods("GET")
```
Пакет _controllers_ отвечает за реализацию запросов
```go
func MainView(w http.ResponseWriter, r *http.Request)  {
	//Функция Рендерит шаблон с данными из пакета __utils__
	u.View(w,"index", nil)
}
```

Все страницы находятся в _$APP_PATH/templates_ 

### Middlewear

Добавление middlewear в _main.go_

```go
	router.Use(M.Middlewear)
```

