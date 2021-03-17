package utils

import (
	L "../utils/Logger"
	"html/template"
	"net/http"
	"path/filepath"
)


func View(w http.ResponseWriter,PageName string, data interface{})  {

	pages := []string{
		filepath.Base("templates")+"/"+PageName+".html",
		filepath.Base("templates")+"/base.layout.html",

	}
	ts ,err := template.ParseFiles(pages...)
	L.ErorrLog(err)
	// Затем мы используем метод Execute() для записи содержимого
	// шаблона в тело HTTP ответа. Последний параметр в Execute() предоставляет
	// возможность отправки динамических данных в шаблон.
	err = ts.Execute(w, data)
	L.ErorrLog(err)
}

