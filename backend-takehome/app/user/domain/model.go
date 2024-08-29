package domain

import "time"

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Name           string `json:"name"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	ReTypePassword string `json:"retype_password"`
}

type User struct {
	ID        int       `gorm:"column:id;primaryKey"`
	Name      string    `gorm:"column:name"`
	Email     string    `gorm:"column:email"`
	Password  string    `gorm:"column:password"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (User) TableName() string {
	return "user"
}

type GeneralResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type LoginDataResponse struct {
	Token string `json:"token"`
}
