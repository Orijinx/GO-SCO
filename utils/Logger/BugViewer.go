package Logger

import (
	"html/template"
	"net/http"
	"path/filepath"

)

func debugView(w http.ResponseWriter, data interface{}){
	pages := []string{
		filepath.Base("templates")+"/Debug.page.html",
	}
	ts ,err := template.ParseFiles(pages...)
	if err != nil {
		panic(err)
	}
	// Затем мы используем метод Execute() для записи содержимого
	// шаблона в тело HTTP ответа. Последний параметр в Execute() предоставляет
	// возможность отправки динамических данных в шаблон.
	err = ts.Execute(w, data)
	if err != nil{
		panic(err)
	}
}
