package models

import (
	"database/sql"
	"log"
	"time"
)

//User contains the basic user structure
type User struct {
	ID        uint16 `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	password  string
	StatusID  int
	CreatedAt time.Time
	UpdatedUp time.Time
}

//GetUsers gets users from DB
func (u *User) GetUsers(db *sql.DB) ([]User, error) {
	q := "SELECT first_name,last_name, email FROM users"
	rows, err := db.Query(q)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var users []User
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.FirstName, &u.LastName, &u.Email); err != nil {
			log.Println(err)
			return nil, err
		}

		users = append(users, u)
	}

	return users, nil
}

//CreateUser creates users in DB
func (u *User) CreateUser(db *sql.DB, data User) error {
	q := "INSERT INTO users(first_name, last_name, email) VALUES ($1, $2, $3)"
	if _, err := db.Exec(q, data.FirstName, data.LastName, data.Email); err != nil {
		log.Println(err)
		return err
	}

	return nil
}
