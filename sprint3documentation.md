Issues Completed Front End:

Front end tests:

Issues Completed Back End:
1: Added user tracking so that the currently logged in user's behavior can be tracked across the website
2: Implemented second database to house every review
3: Disabled CORS to allow BE servers to communicate with the front end and receive requests
4: Reworked signup page to now receive options/post requests so that new users may be added (signup funcional e2e)
5: Reworkd login page to now receive options/post requests (login functional e2e)
6: Began adding functionality to add reviews to site
7: Removed routing
8: Reworked unit testing

Back end tests:

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