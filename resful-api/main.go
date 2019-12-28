package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"resful-api/controllers"
	"resful-api/databases"

	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
	"gopkg.in/natefinch/lumberjack.v2"
)

var db *sql.DB

func init() {
	gotenv.Load()
	logfile := os.Getenv("LOG_FILE_LOCATION")
	if logfile != "" {
		log.SetOutput(&lumberjack.Logger{
			Filename:   logfile,
			MaxSize:    500, // megabytes
			MaxBackups: 3,
			MaxAge:     28,   //days
			Compress:   true, // disabled by default
		})
	}
}

func main() {

	log.Println("Starting Server")
	db = databases.ConnectDB()
	fmt.Println("Kết nối database thành công")
	router := mux.NewRouter()
	// Configure Logging

	router.HandleFunc("/students", controllers.GetAllStudent(db)).Methods("GET")
	router.HandleFunc("/students/{id}", controllers.GetStudentByID(db)).Methods("GET")
	router.HandleFunc("/students", controllers.AddStudent(db)).Methods("POST")
	fmt.Println("Server is running at port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
