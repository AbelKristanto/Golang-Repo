package book

import "time"

// Buat struct untuk menangkap data
type BooksInput struct {
	Title       string `json:"title" binding:"required"`
	Price       int64  `json:"price" binding:"required"`
	Description string `json:"description" binding:"required"`
	Rating      int64  `json:"rating" binding:"required"`
	Discount    int64  `json:"discount" binding:"required"`
	//Author string      `json:"authors" binding:"required"`
	UpdateAt time.Time `json:"update_at"`
}
