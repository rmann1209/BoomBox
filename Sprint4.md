Video Link: https://drive.google.com/file/d/1ZjcW8C-5Jrv3T8-JEOzs8InqFhdi3JO5/view?usp=sharing

Work Completed Front End:
1. Updated log in and sign up css to a sleeker design
2. Fixed bug where containers would scroll over top of the navbar
3. Set initial and unknown page reroutes to redirect back to home page
4. Updated home page design to introduce the site
5. Created a review component where users can input song, album, or artist names and leave a written review
6. Implemented a rating system for users to give a rating from 1 to 5 stars
7. Rerouted login submission to redirect to newly created review page
8. Created a profile component for users to store information related to their account on login
9. Implemented numerous cypress tests to verify functionality of our buttons, rerouting, form submissions, etc (listed below)
10. Implemented new Unit tests to verify proper functionality of all components (listed below)

FRONT END UNIT TESTS:

AppComponent
- should create the app

NavbarComponent
- should navigate to sign up when sign up button is clicked
- should create the navbar component
- should navigate to home when home button is clicked
- should navigate to login when login button is clicked

ProfileComponent
- should create the profile component

LoginComponent
- form should be valid when fields are filled out
- loginUser function should be called on button click
- should create the login component
- form should be invalid initially

SignupComponent
- should have a valid form when both fields are filled
- should create the signup component
- should add a user when form is submitted with valid data
- should have an invalid form when empty

HomeComponent
- should create the home component

ReviewComponent
- should submit a review when the user clicks the submit button
- should initialize with default values
- should show the review box when user submits a search query
- should create the review component

FRONT END CYPRESS TESTS:
- Verify home button rerouting
- Verify sign up button rerouting
- Verify log in button rerouting
- Test sign up page by filling out username and password and submitting form
- Test log in page by filling out username and password and submitting form
- Creates new user with sign up functionality, then reroutes to login page and logs in using user credentials.
- Test review functionality by logging in, entering an artist, writing a review, and leaving 5 stars.


BACKEND DOCUMENTATION:

Objects Created:
1. User struct
	-string fields for username and password.
	-Gorm model for primary database
	-username field has to be unique to be added to the database
2. Review struct
	-string fields for song title, artist, comment, and author. Integer field for rating
	-Gorm model for secondary database

Functions:
SignUpHandler
- func SignUpHandler(w http.ResponseWriter, r *http.Request)
- This function takes a request, usually from the /signup page
- If the request method is options, we know to enable cors
- Otherwise, we receive the post request, read the data, and write to the primary database to add a new user
- If the username already exists within the database, we return an error

LoginHandler
- func LoginHandler(w http.ResponseWriter, r *http.Request)
- This function takes a request, typically from the /login page
- If the request method is options, we know to enable cors
- Otherwise the method is post and we read in the data from the request.
- Query the db for matching username and password
- If the info is confirmed, that user becomes tracked by the server.
- Otherwise we return an error.

HomeHandler
-func HomeHandler(w http.ResponseWriter, r *http.Request)
-This function will take in requests from the /home page
- All this does is enable cors

ReviewHandler
-func ReviewHandler(w http.ResponseWriter, r *http.Request)
- This function takes a request, usually from the /addreview page
- If the request method is options, we know to enable cors
- Otherwise if the request is post, we read in the request and add it into the secondary database.

ViewReviewHandler
-func viewReviewHandler(w http.ResponseWriter, r *http.Request)
- This function receieves get requests from the front end
- If ther is an authenticated user, the function queries the database of reviews
- and returns an array of all reviews that are from that user. May add other queries based on needs in future

Logger
-func logger(next http.Handler) http.Handler
-returns a handler
-When called, the logger is enabled for the server. 
-Keeps track of requests in the command line

enableCors
-func enableCors(w *http.ResponseWriter)
-Configures settings to allow the server to receive requests.

BackEnd Unit Tests:
1. TestSignupHandler: Ensures that the database is functional and will log singup information, checking that the username and password added to the database exists.
2. TestLoginHandler1 and 2: Tests the database with other usernames and passwords and ensures they are added.
3. TestSignup: Launches a local host and creates an http request to sign up a username and password, looks at the most recent entry in the database to see if the info has been added.
4. TestHomeHandler: Launches a local host and ensures the home page exists and contains the proper information.
5. TestReviewHandler: Launches a local host and creates an http request to add a review, checks that the review has been logged in the database.
6. TestViewReviewHandler: Launches a local host and creates an http request to view an array of reviews from the viewReviewHandler method. Searches the array for a review added into the database.
