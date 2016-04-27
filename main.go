package main

import (
	"log"
	"net/http"

	"github.com/ReneVallecillo/OTBackend/db"
	"github.com/ReneVallecillo/OTBackend/handlers"
	"github.com/julienschmidt/httprouter"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := db.NewDB()

	if err != nil {
		log.Fatalln("Could not connect to database")
	}
	r := httprouter.New()
	r.GET("/users", handlers.GetUserHandler(db))
	r.POST("/user", handlers.PostUserHandler(db))

	log.Println("Listening on :8080")
	http.ListenAndServe(":8080", r)

}
