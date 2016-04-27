package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

type errResponse struct {
	Message string `json:"message"`
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
