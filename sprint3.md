Issues Completed Front End:
1. Updated log in and sign up css to a sleeker design
2. Implemented page rerouting when clicking buttons in the navbar
3. Modified navbar to remain on page during rerouting
4. Reworked signup page to now send options/post requests so that new users may be added (signup funcional e2e)
5. Reworked login page to now send options/post requests (login functional e2e)
6. Implemented numerous cypress tests to veerify functionality of our buttons, rerouting, form submissions, etc

Front end tests:
1. Verify home button rerouting
2. Verify sign up button rerouting
3. Verify log in button rerouting
4. Test sign up page by filling out username and password and submitting form
5. Test log in page by filling out username and password and submitting form
6. Creates new user with sign up functionality, then reroutes to login page and logs in using user credentials.

Issues Completed Back End:
1. Added user tracking so that the currently logged in user's behavior can be tracked across the website
2. Implemented second database to house every review
3. Disabled CORS to allow BE servers to communicate with the front end and receive requests
4. Reworked signup page to now receive options/post requests so that new users may be added (signup funcional e2e)
5. Reworked login page to now receive options/post requests (login functional e2e)
6. Began adding functionality to add reviews to site
7. Removed routing
8. Reworked unit testing

Back end tests:
1. TestSignUpHandler: Tests the user database in order to ensure that when a a new account is inserted into the database, it is saved there.
2. TestLoginHandler1 and 2: Similar to TestSignUpHandler, but checks multiple other usernames and passwords within the database.
3. TestSignup: Launches the localHost and tests that when an http request is made to signup a test user, the user is saved within the database.
4. TestHomeHandler2: Creates a local host and http request to launch the homepage, and checks that the homepage contains the proper information stored in the function.
5. TestHomeHandler: Creates a local host and http requet to check that the homepage does indeed exist.
6. TestReviewHandler: Launches a local host and creates an http request to add a new review to the review database, checks that the review matches what is expected in terms of title, artist, rating, comment, and author.

Back End Documentation:

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
	
	Logger
	-func logger(next http.Handler) http.Handler
	-returns a handler
	-When called, the logger is enabled for the server. 
	-Keeps track of requests in the command line
	
	enableCors
	-func enableCors(w *http.ResponseWriter)
	-Configures settings to allow the server to receive requests.

Backend Packages Used:
	"embed"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"net/http"
	
Video Link: https://drive.google.com/file/d/1-3qxg3ebhHkZt14qFHUi0n2QPl34MTTk/view?usp=sharing
