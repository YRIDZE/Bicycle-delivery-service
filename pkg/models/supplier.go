package models

import "time"

type Supplier struct {
	ID          int32     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	AddressName string    `json:"addressName"`
	Type        string    `json:"type"`
	OpenTime    time.Time `json:"open_time"`
	CloseTime   time.Time `json:"close_time"`
}
