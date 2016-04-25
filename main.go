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
	// r.POST("/users", postUserHander(db))

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

func getUsers(db *sql.DB) ([]models.User, error) {
	q := "SELECT firstname, email FROM users"
	rows, err := db.Query(q)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var users []models.User
	for rows.Next() {
		var b models.User
		if err := rows.Scan(&b.FirstName, &b.Email); err != nil {
			log.Println(err)
			return nil, err
		}

		users = append(users, b)
	}

	return users, nil
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
