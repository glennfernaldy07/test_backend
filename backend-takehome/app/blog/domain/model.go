package domain

import "time"

type Blog struct {
	ID        int       `gorm:"column:id;primaryKey"`
	Title     string    `gorm:"column:title"`
	Content   string    `gorm:"column:content"`
	AuthorID  int       `gorm:"column:author_id"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (Blog) TableName() string {
	return "Blog"
}
