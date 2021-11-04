package entity

import "time"

type Album struct {
	ID        string    `form:"id"         json:"id"         binding:"-"`
	Title     string    `form:"title"      json:"title"      binding:"required"`
	Artist    string    `form:"artist"     json:"artist"     binding:"required"`
	Price     float64   `form:"price"      json:"price"      binding:"required"`
	CreatedAt time.Time `form:"created_at" json:"created_at" binding:"-"`
	UpdatedAt time.Time `form:"updated_at" json:"updated_at" binding:"-"`
}
