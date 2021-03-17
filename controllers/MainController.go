package controllers

import (
	u "../utils"
	L "../utils/Logger"
	"fmt"
	"io/ioutil"
	"net/http"
	DB "../DataBase"
	Usr "../Models"
)


func MainView(w http.ResponseWriter, r *http.Request)  {
	u.View(w,"index", nil)
}

func UploadFile(w http.ResponseWriter, r *http.Request) {

	err := r.ParseMultipartForm(10 << 20)
	L.ErorrLog(err)

	file, handler, err := r.FormFile("upload")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	// Create a temporary file within our temp-images directory that follows
	// a particular naming pattern
	tempFile, err := ioutil.TempFile("storage/uploads", "upload-*.png")
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()



	// read all of the contents of our uploaded file into a
	// byte array
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	// write this byte array to our temporary file
	tempFile.Write(fileBytes)
	// return that we have successfully uploaded our file!


	DB.SetValues(handler.Filename,tempFile.Name(),0)
}

func Login(w http.ResponseWriter, r *http.Request)  {

	var User Usr.User

	err:=r.ParseForm()
	L.ErorrLog(err)
	Name := r.FormValue("name")
	Password := r.FormValue("password")

	//ЖЕСТКАЯ ЗАГЛУШКА
	if Name != "Vlad"{
		return
	}else if Password !="1"{
		return
	}else {
		User.Name = Name
		User.Key = "XXX"
		User.Mandate = 1
		DB.SetAuth(&User)
	}
	/////////////////////////
	return
}

func Test(w http.ResponseWriter, r *http.Request) {

DB.GetAuthStatus()
	
}

