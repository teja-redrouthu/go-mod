package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/go-sql-driver/mysql"
)

type Response struct {
	CurrentTime string `json:"current_time"`
}

var db *sql.DB

func connectDB() {
	var err error
	password := os.Getenv("db_pass")
	dsn := fmt.Sprintf("root:%s@tcp(localhost:3306)/toronto_time_db", password)
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatalf("Error pinging database: %v", err)
	}
	fmt.Println("Connected to MySQL database!")
}

func getCurrentTimeHandler(w http.ResponseWriter, r *http.Request) {
	loc, err := time.LoadLocation("America/Toronto")
	if err != nil {
		http.Error(w, "Error loading timezone", http.StatusInternalServerError)
		return
	}

	currentTime := time.Now().In(loc)
	CurrentTime_db:= currentTime.Format("2006-01-02 15:04:05")
	response := Response{
		CurrentTime: currentTime.Format("2006-01-02 15:04:05"),
	}

	// Log time to database
	_, err = db.Exec("INSERT INTO time_log (timestamp) VALUES (?)", CurrentTime_db)
	if err != nil {
		http.Error(w, "Error logging time to database", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	connectDB()
	defer db.Close()

	r := mux.NewRouter()
	r.HandleFunc("/current-time", getCurrentTimeHandler).Methods("GET")

	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
