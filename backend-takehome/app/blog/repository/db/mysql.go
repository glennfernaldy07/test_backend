package db

import (
	"app/blog/domain"
	"app/common/utils"
	"app/config"
	"gopkg.in/jinzhu/gorm.v1"
)

type repository struct {
	cfg config.Config
	f   utils.LogFormatter
	*gorm.DB
}

func NewRepository(cfg config.Config, DB *gorm.DB) domain.Repository {
	f := utils.NewLogFormatter("user.repository")
	return &repository{cfg: cfg, DB: DB, f: f}
}

func (r *repository) StoreBlog(model domain.Blog) (domain.Blog, error) {
	db := r.DB
	row := db.Create(&model)
	if row.Error != nil {
		return domain.Blog{}, row.Error
	}

	return model, nil
}

func (r *repository) FindBlogByTitle(title string) (*domain.Blog, error) {
	db := r.DB
	var model domain.Blog
	row := db.Where("title = ?", title).Find(&model)
	if row.Error != nil {
		return nil, row.Error
	}
	return &model, nil
}

func (r *repository) FindBlogs(userID int) (*[]domain.Blog, error) {
	db := r.DB
	var models []domain.Blog
	row := db.Where("author_id = ?", userID).Find(&models)
	if row.Error != nil {
		return nil, row.Error
	}
	return &models, nil
}

func (r *repository) FindBlogByID(blogID int) (*domain.Blog, error) {
	db := r.DB
	var model domain.Blog
	row := db.Where("id = ?", blogID).Find(&model)
	if row.Error != nil {
		return nil, row.Error
	}
	return &model, nil
}

func (r *repository) UpdateBlogByID(blog domain.Blog) (*domain.Blog, error) {
	db := r.DB
	model := domain.Blog{
		Title:     blog.Title,
		Content:   blog.Content,
		UpdatedAt: blog.UpdatedAt,
	}
	row := db.Model(&domain.Blog{}).
		Where("id = ?", blog.ID).
		Where("author_id = ?", blog.AuthorID).
		Updates(model)
	if row.Error != nil {
		return nil, row.Error
	}

	updatedModel := domain.Blog{}
	db.Where("id = ? and author_id = ?", blog.ID, blog.AuthorID).Find(&updatedModel)
	return &updatedModel, nil
}

func (r *repository) DeleteBlogByID(id int) error {
	db := r.DB
	if err := db.Model(&domain.Blog{}).Delete(&domain.Blog{}, id).Error; err != nil {
		return err
	}
	return nil
}
