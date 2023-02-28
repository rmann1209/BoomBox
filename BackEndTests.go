package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var testDB *gorm.DB

func TestSignUpHandler(t *testing.T) {
	// Initialize the test database

	var err error
	testDB, err = gorm.Open(sqlite.Open("testDB.db"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open database: %v", err)
	}

	// Clear the users table before running the test
	testDB.Exec("DELETE FROM users")

	// Create a test request
	req, err := http.NewRequest("POST", "/signup", strings.NewReader("username=testuser&password=testpassword&first_name=Test&last_name=User"))
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	// Create a test response recorder
	rr := httptest.NewRecorder()

	// Call the signup handler with the test request and response recorder
	SignUpHandler(rr, req)

	// Check if the response status code is OK
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check if the user was created in the database
	var user User
	if err := testDB.Where("username = ?", "testuser").First(&user).Error; err != nil {
		t.Errorf("failed to find user: %v", err)
	}

	// Check if the user's details were correctly saved in the database
	if user.Username != "testuser" {
		t.Errorf("username not saved correctly: got %v want %v", user.Username, "testuser")
	}
	if user.Password != "testpassword" {
		t.Errorf("password not saved correctly: got %v want %v", user.Password, "testpassword")
	}
	if user.FirstName != "Test" {
		t.Errorf("first name not saved correctly: got %v want %v", user.FirstName, "Test")
	}
	if user.LastName != "User" {
		t.Errorf("last name not saved correctly: got %v want %v", user.LastName, "User")
	}
}

func TestLoginHandler(t *testing.T) {
	// Initialize the test database
	var err error
	testDB, err = gorm.Open(sqlite.Open("testDB.db"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open database: %v", err)
	}

	// Create a test user in the database
	testUser := User{
		Username:  "testuser",
		Password:  "testpassword",
		FirstName: "Test",
		LastName:  "User",
	}
	testDB.Create(&testUser)

	// Create a test request with valid login credentials
	reqValid, err := http.NewRequest("POST", "/login", strings.NewReader("username=testuser&password=testpassword"))
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	// Create a test response recorder
	rrValid := httptest.NewRecorder()

	// Call the login handler with the valid test request and response recorder
	LoginHandler(rrValid, reqValid)

	// Check if the response status code is a redirect to the home page
	if status := rrValid.Code; status != http.StatusSeeOther {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusSeeOther)
	}

	// Create a test request with invalid login credentials
	reqInvalid, err := http.NewRequest("POST", "/login", strings.NewReader("username=testuser&password=wrongpassword"))
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	// Create a test response recorder
	rrInvalid := httptest.NewRecorder()

	// Call the login handler with the invalid test request and response recorder
	LoginHandler(rrInvalid, reqInvalid)

	// Check if the response status code is a bad request
	if status := rrInvalid.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}
}

func TestSignupAndLogin(t *testing.T) {
	// Initialize the test database

	var err error
	testDB, err = gorm.Open(sqlite.Open("testDB.db"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open database: %v", err)
	}

	// Clear the users table before running the test
	testDB.Exec("DELETE FROM users")

	// Create a test request to signup
	signupReq, err := http.NewRequest("POST", "/signup", strings.NewReader("username=testuser&password=testpassword&first_name=Test&last_name=User"))
	if err != nil {
		t.Fatalf("failed to create signup request: %v", err)
	}

	// Create a test response recorder for signup
	signupRR := httptest.NewRecorder()

	// Call the signup handler with the test request and response recorder
	SignUpHandler(signupRR, signupReq)

	// Check if the response status code is OK
	if status := signupRR.Code; status != http.StatusOK {
		t.Errorf("signup handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Create a test request to login
	loginReq, err := http.NewRequest("POST", "/login", strings.NewReader("username=testuser&password=testpassword"))
	if err != nil {
		t.Fatalf("failed to create login request: %v", err)
	}

	// Create a test response recorder for login
	loginRR := httptest.NewRecorder()

	// Call the login handler with the test request and response recorder
	LoginHandler(loginRR, loginReq)

	// Check if the response status code is OK
	if status := loginRR.Code; status != http.StatusSeeOther {
		t.Errorf("login handler returned wrong status code: got %v want %v", status, http.StatusSeeOther)
	}

	// Check if the response redirected to the correct page
	if redirectURL := loginRR.Header().Get("Location"); redirectURL != "/home" {
		t.Errorf("login handler redirected to wrong URL: got %v want %v", redirectURL, "/home")
	}
}

func TestHomeHandler_SignUpButton(t *testing.T) {
	// Create a test request with "Sign up" as the form value
	req, err := http.NewRequest("POST", "/home", strings.NewReader("action=Sign up"))
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	// Create a test response recorder
	rr := httptest.NewRecorder()

	// Call the home handler with the test request and response recorder
	HomeHandler(rr, req)

	// Check if the response status code is a redirect
	if status := rr.Code; status != http.StatusSeeOther {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusSeeOther)
	}

	// Check if the response header "Location" is "/signup"
	if location := rr.Header().Get("Location"); location != "/signup" {
		t.Errorf("handler returned wrong location header: got %v want %v", location, "/signup")
	}
}

func TestHomeHandler_LoginButton(t *testing.T) {
	// Create a test request with "Login" as the form value
	req, err := http.NewRequest("POST", "/home", strings.NewReader("action=Login"))
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	// Create a test response recorder
	rr := httptest.NewRecorder()

	// Call the home handler with the test request and response recorder
	HomeHandler(rr, req)

	// Check if the response status code is a redirect
	if status := rr.Code; status != http.StatusSeeOther {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusSeeOther)
	}

	// Check if the response header "Location" is "/login"
	if location := rr.Header().Get("Location"); location != "/login" {
		t.Errorf("handler returned wrong location header: got %v want %v", location, "/login")
	}
}
