package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"html/template"
	"log"
	"net/http"
)

type User struct {
	gorm.Model
	Username  string `gorm:"unique"` //specifies as unique, blocks repeats
	Password  string
	FirstName string
	LastName  string
}

var db *gorm.DB

func main() {
	//db -> database
	var err error
	db, err = gorm.Open(sqlite.Open("usersbase.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to open database.")
	}
	err = db.AutoMigrate(&User{})
	if err != nil {
		panic("failed to automigrate")
		return
	}

	r := mux.NewRouter()
	r.HandleFunc("/home", HomeHandler)
	r.HandleFunc("/signup", SignUpHandler)
	r.HandleFunc("/login", LoginHandler)
	http.Handle("/", r)
	//starts logger
	r.Use(logger)

	fmt.Println("Starting server on :8080")
	err = http.ListenAndServe(":8080", nil)
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

	print(r)

	username := r.FormValue("username")
	password := r.FormValue("password")
	firstName := r.FormValue("first_name")
	lastName := r.FormValue("last_name")

	result := db.Create(&User{
		Username:  username,
		Password:  password,
		FirstName: firstName,
		LastName:  lastName,
	})

	if result.Error != nil {
		http.Error(w, "Username already exists", http.StatusBadRequest)
		return
	}
	http.Redirect(w, r, "/home", http.StatusSeeOther)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		//tmpl, err := template.ParseFiles("login.html")//
		tmpl, err := template.ParseFiles("angularstuff/your-awesome-project/src/app/app.component.html")
		//tmpl, err := template.ParseFiles("angularstuff/your-awesome-project/src/index.html") //
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
	result1 := r.FormValue("action")
	if result1 == "SIGN UP" {
		http.Redirect(w, r, "/signup", http.StatusSeeOther)
	} else if result1 == "MORE INFO" {
		http.Redirect(w, r, "/home", http.StatusSeeOther)
	} else if result1 == "LOGIN" {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
	username := r.FormValue("Username")
	password := r.FormValue("Password")

	user := User{}

	result := db.Where("username = ? AND password = ?", username, password).First(&user)
	if result.Error != nil {
		http.Error(w, "Username or password is incorrect", http.StatusBadRequest)
		return
	}
	http.Redirect(w, r, "/home", http.StatusSeeOther)
}
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl, err := template.ParseFiles("homemock.html") //direct to the file
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
	result := r.FormValue("action")
	if result == "Sign up" {
		http.Redirect(w, r, "/signup", http.StatusSeeOther)
	} else if result == "Login" {
		http.Redirect(w, r, "/login", http.StatusSeeOther)

	} else {
		http.Redirect(w, r, "/home", http.StatusSeeOther)

	}

}

// begin logger - DND
func logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
