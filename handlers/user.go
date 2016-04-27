package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/ReneVallecillo/OTBackend/models"
	"github.com/julienschmidt/httprouter"
)

//GetUserHandler defines handlker for Getting users
func GetUserHandler(db *sql.DB) httprouter.Handle {
	return func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
		user := models.User{}
		users, err := user.GetUsers(db)
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

//PostUserHandler defines handler for Posting users
func PostUserHandler(db *sql.DB) httprouter.Handle {
	return func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
		var u models.User
		if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
			respondError(rw, err)
			return
		}

		if err := u.CreateUser(db, u); err != nil {
			respondError(rw, err)
			return
		}

		rw.WriteHeader(http.StatusNoContent)
	}
}
