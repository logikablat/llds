package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", "postgres://username:password@ngrok_url:ngrok_port/dbname?sslmode=disable")
	if err != nil {
		logrus.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		logrus.Fatal(err)
	}
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/register", RegisterHandler).Methods("POST")
	router.HandleFunc("/login", LoginHandler).Methods("POST")
	router.HandleFunc("/logout", LogoutHandler).Methods("POST")

	initLogging()

	if err := http.ListenAndServe(":8080", router); err != nil {
		logrus.Fatal(err)
	}
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var credentials map[string]string
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	username, password := credentials["username"], credentials["password"]

	_, err := db.Exec("INSERT INTO users (username, password) VALUES ($1, $2)", username, password)
	if err != nil {
		logrus.Error("Failed to insert user into the database:", err)
		http.Error(w, "Failed to register user", http.StatusInternalServerError)
		return
	}

	logrus.Infof("User registered: %s", username)

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Registration successful"))
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var credentials map[string]string
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	username, password := credentials["username"], credentials["password"]

	row := db.QueryRow("SELECT password FROM users WHERE username = $1", username)
	var storedPassword string
	err := row.Scan(&storedPassword)
	if err != nil {
		logrus.Error("Failed to fetch user from the database:", err)
		http.Error(w, "Failed to log in", http.StatusUnauthorized)
		return
	}

	if storedPassword == password {
		logrus.Infof("User logged in: %s", username)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Login successful"))
	} else {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
	}
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	logrus.Infof("User logged out")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Logout successful"))
}

func initLogging() {
	logrus.SetFormatter(&logrus.TextFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.InfoLevel)
}
