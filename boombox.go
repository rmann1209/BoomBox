/*
	func main() {
		http.HandleFunc("/homepage", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("Boombox homepage (i want to kms)")) })
		log.Println("Server starting...")
		mux := http.NewServeMux()
		//this be the main redirect code
		mux.Handle("/", http.RedirectHandler("https://local", http.StatusSeeOther))
		http.ListenAndServe(":8088", mux)
	}
*/
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type User struct {
	Username  string
	Password  string
	FirstName string
	LastName  string
}

var users = []User{}

func main() {
	http.HandleFunc("/signup", SignUpHandler)
	http.HandleFunc("/login", LoginHandler)
	
	fmt.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl, err := template.ParseFiles("signup.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, nil)
		return
	}

	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")
	firstName := r.FormValue("first_name")
	lastName := r.FormValue("last_name")

	for _, u := range users {
		if u.Username == username {
			http.Error(w, "Username already exists", http.StatusBadRequest)
			return
		}
	}

	user := User{
		Username:  username,
		Password:  password,
		FirstName: firstName,
		LastName:  lastName,
	}

	users = append(users, user)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl, err := template.ParseFiles("login.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, nil)
		return
	}

	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	for _, u := range users {
		if u.Username == username && u.Password == password {
			w.Write([]byte("Login successful"))
			return
		}
	}

	http.Redirect(w, r, "/signup", http.StatusSeeOther)
}
