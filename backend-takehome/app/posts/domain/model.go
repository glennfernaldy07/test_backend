package domain

import "time"

type GeneralResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type PostsBlogRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	UserID  int64  `json:"-"`
}

type PostsBlogResponse struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GetBlogResponse struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GetBlogRequest struct {
	ID     int   `schema:"id"`
	UserID int64 `json:"-"`
}

type DeleteBlogRequest struct {
	ID     int   `schema:"id"`
	UserID int64 `json:"-"`
}

type UpdateBlogRequest struct {
	ID      int    `schema:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	UserID  int64  `json:"-"`
}

type UpdateBlogResponse struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
