package main

import (
	"Impact/server/config"
	"Impact/server/routes/api"
	_ "database/sql"

	// "encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"

	"Impact/server/models"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"

	_ "github.com/lib/pq"
)

var db *gorm.DB //database

// postgres://zelmmuwe:kg7hf-hRvWz1Szylce4pFDvTDfBuAZDq@raja.db.elephantsql.com:5432/zelmmuwe

func enableCors(w *http.ResponseWriter) {
	fmt.Println("will be Setted")
	// (*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Origin", "*")

	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type")
	fmt.Println("Setted")

}

//wrapper functions
func getAllUsers(w http.ResponseWriter, r *http.Request) {
	api.GetAllUsers(db, w, r)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	api.CreateUser(db, w, r)

}

func getAllOrganizations(w http.ResponseWriter, r *http.Request) {
	api.GetAllOrganizations(db, w, r)
}

func createOrganization(w http.ResponseWriter, r *http.Request) {
	api.CreateOrganization(db, w, r)

}

func getAllNotifications(w http.ResponseWriter, r *http.Request) {
	api.GetAllNotifications(db, w, r)
}

func createNotification(w http.ResponseWriter, r *http.Request) {
	api.CreateNotification(db, w, r)

}

func getNotification(w http.ResponseWriter, r *http.Request) {
	api.GetNotification(db, w, r)

}

func joinNotification(w http.ResponseWriter, r *http.Request) {
	api.JoinNotification(db, w, r)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	api.GetUser(db, w, r)
}

func getOrganization(w http.ResponseWriter, r *http.Request) {
	api.GetOrganization(db, w, r)
}

func getAllVolunteers(w http.ResponseWriter, r *http.Request) {
	api.GetAllVolunteers(db, w, r)
}

func getAllNotificationVolunteers(w http.ResponseWriter, r *http.Request) {
	api.GetAllNotificationVolunteers(db, w, r)
}

func main() {

	// err := godotenv.Load(os.ExpandEnv("$GOPATH/src/Impact/server/.env"))
	err := godotenv.Load("./.env")

	if err != nil {
		log.Println("env is shitt")
		log.Fatal(err)
	}
	host := config.GetConfig().DB.Host
	portt := config.GetConfig().DB.Port
	user := config.GetConfig().DB.User
	password := config.GetConfig().DB.Password
	dbname := config.GetConfig().DB.Dbname
	//connecting to db  $$ to be imported from config file
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, portt, user, password, dbname)

	var dialect = config.GetConfig().DB.Dialect
	connection, err := gorm.Open(dialect, psqlInfo)
	if err != nil {
		panic(err)
	}
	db = connection

	db = models.DBMigrate(db)
	fmt.Println("Database connected!")

	//logging
	file, err := os.OpenFile("./logs/info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	log.SetOutput(file)
	log.Print("Logging to a file in Go!")

	//making the router
	router := mux.NewRouter()

	//routes
	router.HandleFunc("/", defaultHandler).Methods("GET")
	router.HandleFunc("/api/users", getAllUsers).Methods("GET")
	router.HandleFunc("/api/users", createUser).Methods("POST")
	router.HandleFunc("/api/users/{id}", getUser).Methods("GET")

	router.HandleFunc("/api/notifications", getAllNotifications).Methods("GET")
	router.HandleFunc("/api/notifications", createNotification).Methods("POST")
	router.HandleFunc("/api/notifications/{id}", getNotification).Methods("GET")
	router.HandleFunc("/api/notifications/join/{id}", joinNotification).Methods("POST")

	port := flag.String("p", config.GetConfig().PORT, "port to serve on")
	flag.Parse()

	log.Printf("Server up and running!!\n")
	log.Printf("Serving on HTTP port: %s\n", *port)

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	log.Fatal(http.ListenAndServe(":"+*port, handlers.CORS(originsOk, headersOk, methodsOk)(router)))

}

//default router
func defaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, Web!")
}

//GetDB returns a handle to the DB object
func GetDB() *gorm.DB {
	return db
}

//RequestLogger helper method for logging
func RequestLogger(targetMux http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		targetMux.ServeHTTP(w, r)

		// log request by who(IP address)
		requesterIP := r.RemoteAddr

		log.Printf(
			"%s\t\t%s\t\t%s\t\t%v",
			r.Method,
			r.RequestURI,
			requesterIP,
			time.Since(start),
		)
	})
}

// var books string = "welcome!"
// id := 1
// 	var col string
// 	sqlStatement := `SELECT username FROM User WHERE userId=$1`
// 	row := db.QueryRow(sqlStatement, id)
// 	err = row.Scan(&col)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			fmt.Println("Zero rows found")
// 		} else {
// 			panic(err)
// 		}
// 	}
// 	fmt.Println("Hello, Web! Users " + col)

// r := mux.NewRouter()
// 	r.HandleFunc("/api/users", getUsers).Methods("GET")
// 	r.HandleFunc("/", defaultHandler)
// 	// fmt.Println("Hello, World!")
// 	http.ListenAndServe(":3000", r)
