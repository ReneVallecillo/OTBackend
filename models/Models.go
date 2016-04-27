package models

import "time"

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

//SignLog contains the sign structure
type SignLog struct {
	ID           uint16    `json:"id"`
	CheckinTime  time.Time `json:"checkin"`
	CheckoutTime time.Time `json:"checkout"`
}
