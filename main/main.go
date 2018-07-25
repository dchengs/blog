package main

import (
	"fmt"
	"html/template"
	"net/http"
	"crypto/sha256"
	"./logger"
)

//list of pages handler
func Index(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("html/index.html")
	if err != nil {
		l.ErrLogger(err)
	}
	else{
		template.Execute(w, nil)
	}
}

func Project(w http.ResponseWriter, r *http.Request) {
	Project, err := template.ParseFiles("html/currentProject.html")
	if err != nil {
		l.ErrLogger(err)
	}
	else{
		template.Execute(w, nil)
	}
}
func AboutMe(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("html/aboutMe.html")
	if err != nil {
		l.ErrLogger(err)
	}
	else{
		template.Execute(w, nil)
	}
}


func Resume(w http.ResponseWriter, r *http.Request){
	template, err := template.ParseFiles("html/resume.html")
	if err != nil{
		l.ErrLogger(err)
	}
	else{
		template.Execute(w, nil)
	}

}

func Login(w http.ResponseWriter, r *http.Request){
	//user requests the login page
	if(r.Method() == "GET"){
		template, err = := template.ParseFiles("html/login.html")
		template.execute(w, nil)
		if err != nil{
			ult.ErrLogger(err)
		}
	}
	else{
		//simple debug testing
		debugUsername := "chengs"
		debugPassword := "123123123"
		username:=r.Form["username"]
		password:= r.Form["password"]
		if (debugUsername == username && debugPassword == password){
			fmt.Println("Correct password")
		}
		else{
			fmt.Println("Incorrect password")
		}
		//encrypt password input
		sha_256 := sha256.New()
		//convert password to byte array
		sha_256.Write([]byte(password))
		encryptedPW := hex.EncodeToString(sha_256.Sum(nil))
		//check with server

	}
}

func main() {
	http.HandleFunc("/", Index)
	http.HandleFunc("/Project/", Project)
	http.HandleFunc("/AboutMe/", AboutMe)
	http.HandleFunc("/login", auth.login)
	http.HandleFunc("Resume", Resume)
	http.ListenAndServe(":5000", nil)
}