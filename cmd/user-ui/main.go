package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

var db *sql.DB

// Структура для данных, которые будут вводиться в форму
type User struct {
	Username string
	Email    string
	Password string
}

func init() {
	var err error
	// Настроить подключение к базе данных
	connStr := "user=user password=yourpassword dbname=go_pr_2 sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/submit", submitForm)

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// homePage отображает форму для ввода данных
func homePage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("form").Parse(`
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<title>User Form</title>
		</head>
		<body>
			<h1>Enter your details</h1>
			<form action="/submit" method="POST">
				<label for="username">Username:</label>
				<input type="text" id="username" name="username" required><br><br>
				
				<label for="email">Email:</label>
				<input type="email" id="email" name="email" required><br><br>
				
				<label for="password">Password:</label>
				<input type="password" id="password" name="password" required><br><br>

				<label for="latitude">Latitude:</label>
    <input type="number" step="any" id="latitude" name="latitude" required><br>

    <label for="longitude">Longitude:</label>
    <input type="number" step="any" id="longitude" name="longitude" required><br>

				<button type="submit">Submit</button>
			</form>
		</body>
		</html>
	`)
	if err != nil {
		log.Fatal(err)
	}

	tmpl.Execute(w, nil)
}

// submitForm обрабатывает данные формы и сохраняет их в базе данных
func submitForm(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		username := r.FormValue("username")
		email := r.FormValue("email")
		password := r.FormValue("password")

		// Логирование полученных данных
		log.Printf("Received data: username=%s, email=%s, password=%s", username, email, password)

		// Вставляем данные в базу
		_, err := db.Exec("INSERT INTO users (username, email, password) VALUES ($1, $2, $3)", username, email, password)
		if err != nil {
			log.Println("Error inserting data:", err) // Логируем ошибку
			http.Error(w, "Failed to insert data into the database", http.StatusInternalServerError)
			return
		}

		// Успешное добавление
		fmt.Fprintf(w, "Data submitted successfully! Username: %s, Email: %s", username, email)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
