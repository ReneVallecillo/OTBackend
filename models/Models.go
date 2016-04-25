package models

import "time"

type User struct {
	ID        uint16 `json:"id"`
	FirstName string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	password  string
	StatusID  int
	CreatedAt time.Time
	UpdatedUp time.Time
}

type Attendance struct {
	ID           uint16    `json:"id"`
	CheckinTime  time.Time `json:"checkin`
	CheckoutTime time.Time `json:"checkout`
}
