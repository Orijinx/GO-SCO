package middlewares

import (
	DB "../../DataBase"
	U "../../utils"
	"context"
	"fmt"
	"net/http"
)

var AuthMiddlewear = func(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		notAuth := []string{"/api/user/new", "/login"} //Список эндпоинтов, для которых не требуется авторизация
		requestPath := r.URL.Path                               //текущий путь запроса

		//проверяем, не требует ли запрос аутентификации, обслуживаем запрос, если он не нужен
		for _, value := range notAuth {

			if value == requestPath {
				next.ServeHTTP(w, r)
				return
			}
		}

		User := DB.GetAuthStatus()

		//response := make(map[string]interface{})
		//tokenHeader := r.Header.Get("Authorization") //Получение токена

		if User.Status {
			//Всё прошло хорошо, продолжаем выполнение запроса
			fmt.Sprintf("User %", User.Name) //Полезно для мониторинга
			ctxU := context.WithValue(r.Context(), "user", User)
			r = r.WithContext(ctxU)
			next.ServeHTTP(w, r) //передать управление следующему обработчику!
		}else {
			U.View(w,"login.page",nil)
			return
		}
	})
}


