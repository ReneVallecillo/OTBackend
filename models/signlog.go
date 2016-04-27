package models

import "time"

//SignLog contains the sign structure
type SignLog struct {
	ID           uint16    `json:"id"`
	CheckinTime  time.Time `json:"checkin"`
	CheckoutTime time.Time `json:"checkout"`
}
