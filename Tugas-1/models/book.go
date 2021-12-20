//models/movie.go, digunakan untuk membuat struct/ struktur.

package models

import "time"

type (
	Book struct {
		ID           int       `json: "id"`
		Title        string    `json:"title"`
		Description  string    `json:"description"`
		Image_url    string    `json:"image_url"`
		Release_year int       `json:"release_year"`
		Price        string    `json:"price"`
		Total_page   int       `json:"total_page"`
		Thickness    string    `json:"thickness"`
		CategoryId   int       `json: "category_id"`
		CreatedAt    time.Time `json:"created_at"`
		UpdatedAt    time.Time `json:"updated_at"`
	}
)
