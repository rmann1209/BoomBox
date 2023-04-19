package main

import (
	"encoding/json"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

// tests sign up database
func TestSignUpHandler(t *testing.T) {

	// Connect to the test database

	testDB, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	testDB.Exec("DELETE FROM users")
	if err != nil {
		t.Fatal(err)
	}

	// Create the User table in the test database
	err = testDB.AutoMigrate(&User{})
	if err != nil {
		t.Fatal(err)
	}

	// Create a new HTTP request
	result := testDB.Create(&User{
		Username: "testuser",
		Password: "testpass",
		//FirstName: "Test",
		//LastName:  "User",
	})

	// Check that the user was created in the test database
	var user User
	result = testDB.Where("username = ?", "testuser").First(&user)
	if result.Error != nil {
		t.Errorf("failed to find user in test database: %v", result.Error)
	}

	// Check that the user's details are correct
	if user.Password != "testpass" {
		t.Errorf("user has incorrect password: got %v want %v", user.Password, "testpass")
	}

}

// login tests
func TestLoginHandler1(t *testing.T) {
	testDB, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}

	// Create the User table in the test database
	err = testDB.AutoMigrate(&User{})
	if err != nil {
		t.Fatal(err)
	}

	// Create a new HTTP request
	result := testDB.Create(&User{
		Username: "testuser",
		Password: "testpass",
		//FirstName: "Test",
		//LastName:  "User",
	})

	// Check that the user was created in the test database
	var user User
	result = testDB.Where("username = ?", "testuser").First(&user)
	if result.Error != nil {
		t.Errorf("user has not been created to login yet: %v", result.Error)
	}

	// Check that the user's details are correct
	if user.Password != "testpass" {
		t.Errorf("user has incorrect password: got %v want %v", user.Password, "testpass")
	}
}

func TestLoginHandler2(t *testing.T) {
	testDB, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}

	// Create the User table in the test database
	err = testDB.AutoMigrate(&User{})
	if err != nil {
		t.Fatal(err)
	}

	// Create a new HTTP request
	result := testDB.Create(&User{
		Username: "testuser2",
		Password: "testpass2",
		//FirstName: "Test2",
		//LastName:  "User2",
	})

	// Check that the user was created in the test database
	var user User
	result = testDB.Where("username = ?", "testuser2").First(&user)
	if result.Error != nil {
		t.Errorf("user has not been created to login yet: %v", result.Error)
	}

	// Check that the user's details are correct
	if user.Password != "testpass2" {
		t.Errorf("user has incorrect password: got %v want %v", user.Password, "testpass")
	}
}

func TestSignup(t *testing.T) {
	go main()
	time.Sleep(1 * time.Second)

	// create request body
	requestBody := strings.NewReader(`{"username": "Ben", "password": "Pass"}`)

	// create request
	url := "http://localhost:8080/signup"
	req, err := http.NewRequest("POST", url, requestBody)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	// create a new response recorder to record the response
	rr := httptest.NewRecorder()

	// call the SignUpHandler function with the request and response recorder
	SignUpHandler(rr, req)

	// check if the account was created successfully
	var actualAccount User
	if err := db.First(&actualAccount, "username = ?", "Ben").Error; err != nil {
		t.Errorf("failed to retrieve account from database: %v", err)
	}

	expectedAccount := User{
		Username: "Ben",
		Password: "Pass",
	}

	if actualAccount.Username != expectedAccount.Username {
		t.Errorf("unexpected username: got %v want %v", actualAccount.Username, expectedAccount.Username)
	}
	if actualAccount.Password != expectedAccount.Password {
		t.Errorf("unexpected password: got %v want %v", actualAccount.Password, expectedAccount.Password)
	}
}

// tests that local host functions and launches
func TestHomeHandler(t *testing.T) {
	// start server
	go main()
	time.Sleep(2 * time.Second)

	// create request
	url := "http://localhost:8080/home"
	resp, err := http.Get(url)
	if err != nil {
		t.Errorf("error making request: %v", err)
		return
	}
	defer resp.Body.Close()

	// check response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("error reading response body: %v", err)
		return
	}
	if len(body) == 0 {
		t.Errorf("empty response body")
		return
	}
}

func TestReviewHandler(t *testing.T) {
	// Set the active username
	activeUsername = "evan"

	// Create a new HTTP request with method POST and empty body
	go main()
	time.Sleep(1 * time.Second)

	// create request
	url := "http://localhost:8080/review"
	resp, err := http.Get(url)
	if err != nil {
		t.Errorf("error making request: %v", err)
		return
	}
	defer resp.Body.Close()
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new response recorder to record the response
	rr := httptest.NewRecorder()

	// Call the ReviewHandler function with the request and response recorder
	AddReviewHandler(rr, req)

	// Check if the HTTP status code is 201 Created
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

	// Check if the review was properly inputted
	var actualReview SongReview
	if err := db2.Last(&actualReview).Error; err != nil {
		t.Errorf("failed to retrieve review from database: %v", err)
	}

	expectedReview := SongReview{
		SongTitle: "Sky",
		Artist:    "Playboi Carti",
		Rating:    5,
		Comment:   "Fire",
		Author:    "evan",
	}

	if actualReview.SongTitle != expectedReview.SongTitle {
		t.Errorf("unexpected Song title: got %v want %v",
			actualReview.SongTitle, expectedReview.SongTitle)
	}
	if actualReview.Artist != expectedReview.Artist {
		t.Errorf("unexpected Artist: got %v want %v",
			actualReview.Artist, expectedReview.Artist)
	}
	if actualReview.Rating != expectedReview.Rating {
		t.Errorf("unexpected Rating: got %v want %v",
			actualReview.Rating, expectedReview.Rating)
	}
	if actualReview.Comment != expectedReview.Comment {
		t.Errorf("unexpected Review: got %v want %v",
			actualReview.Comment, expectedReview.Comment)
	}
	if actualReview.Author != expectedReview.Author {
		t.Errorf("unexpected Author: got %v want %v",
			actualReview.Author, expectedReview.Author)
	}
}

func TestViewReviewHandler(t *testing.T) {

	go main()
	time.Sleep(1 * time.Second)

	// Initialize the database connection and migrate the schema
	db2, err := gorm.Open(sqlite.Open("reviews.db"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect to database: %v", err)
	}
	if err := db2.AutoMigrate(&SongReview{}); err != nil {
		t.Fatalf("failed to migrate database schema: %v", err)
	}

	activeUsername := "evan"
	activeReview := SongReview{SongTitle: "Song Title", Artist: "Artist Name", Rating: 5, Comment: "eh", Author: activeUsername}
	db.Create(&activeReview)

	// create a new HTTP request
	req, err := http.NewRequest("GET", fmt.Sprintf("/viewreview/%s", activeUsername), nil)
	if err != nil {
		t.Fatalf("failed to create HTTP request: %v", err)
	}

	// create a new HTTP response recorder to capture the response
	w := httptest.NewRecorder()

	// Call the ReviewHandler function with the request and response recorder
	viewReviewHandler(w, req)

	// Check that the HTTP status code is 200 OK
	if w.Code != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, w.Code)
	}

	// Decode the response body into a slice of SongReview objects
	var responseReviews []SongReview
	responseBody, err := ioutil.ReadAll(w.Body)
	if err != nil {
		t.Fatalf("failed to read response body: %v", err)
	}
	if err := json.Unmarshal(responseBody, &responseReviews); err != nil {
		t.Fatalf("failed to decode response body: %v", err)
	}
	fmt.Printf("Response Reviews: %+v\n", responseReviews)

	// Check that the response contains the test review
	found := false
	for _, review := range responseReviews {
		if review.Author == activeUsername {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("expected review with author %s to be in response, but not found", activeUsername)
	}

}
