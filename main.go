package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/ReneVallecillo/OTBackend/db"
	"github.com/ReneVallecillo/OTBackend/models"
	"github.com/julienschmidt/httprouter"
	_ "github.com/mattn/go-sqlite3"
)

type errResponse struct {
	Message string `json:"message"`
}

func main() {
	db, err := db.NewDB()

	if err != nil {
		log.Fatalln("Could not connect to database")
	}
	r := httprouter.New()
	r.GET("/users", getUserHandler(db))
	r.POST("/user", postUserHandler(db))

	log.Println("Listening on :8080")
	http.ListenAndServe(":8080", r)

}

func getUserHandler(db *sql.DB) httprouter.Handle {
	return func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
		users, err := getUsers(db)
		if err != nil {
			respondError(rw, err)
			return
		}

		rw.Header().Set("Content-Type", "application/json")

		if err := json.NewEncoder(rw).Encode(users); err != nil {
			respondError(rw, err)
			return
		}
	}
}

func postUserHandler(db *sql.DB) httprouter.Handle {
	return func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
		var u models.User
		if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
			respondError(rw, err)
			return
		}

		if err := createUser(db, u); err != nil {
			respondError(rw, err)
			return
		}

		rw.WriteHeader(http.StatusNoContent)
	}
}

func getUsers(db *sql.DB) ([]models.User, error) {
	q := "SELECT first_name,last_name, email FROM users"
	rows, err := db.Query(q)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var users []models.User
	for rows.Next() {
		var b models.User
		if err := rows.Scan(&b.FirstName, &b.LastName, &b.Email); err != nil {
			log.Println(err)
			return nil, err
		}

		users = append(users, b)
	}

	return users, nil
}

func createUser(db *sql.DB, u models.User) error {
	q := "INSERT INTO users(first_name, last_name, email) VALUES ($1, $2, $3)"
	if _, err := db.Exec(q, u.FirstName, u.LastName, u.Email); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func respondError(rw http.ResponseWriter, err error) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusInternalServerError)

	er := errResponse{
		Message: err.Error(),
	}

	if err := json.NewEncoder(rw).Encode(er); err != nil {
		log.Println(err)
	}
}
