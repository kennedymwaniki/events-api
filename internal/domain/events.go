package domain

import "time"

type Event struct {
	Id int `json:"id,omitempty"`
	Title string `json:"title"`
	Description string `json:"description"`
	Location string `json:"location"`
	Date string `json:"date"`
	Time string `json:"time"`
	Capacity int `json:"capacity"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}