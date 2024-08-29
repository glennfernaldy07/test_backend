package service

import (
	"app/blog/domain"
	"app/common/utils"
	"app/config"
)

type service struct {
	cfg  config.Config
	repo domain.Repository
	f    utils.LogFormatter
}

func (s *service) GetBlogs(userID int) (*[]domain.Blog, error) {
	blogs, err := s.repo.FindBlogs(userID)
	if err != nil {
		return nil, err
	}
	return blogs, nil
}

func (s *service) GetBlogByID(blogID int) (*domain.Blog, error) {
	blog, err := s.repo.FindBlogByID(blogID)
	if err != nil {
		return nil, err
	}
	return blog, nil
}

func (s *service) UpdateBlogByID(model domain.Blog) (*domain.Blog, error) {
	blog, err := s.repo.UpdateBlogByID(model)
	if err != nil {
		return nil, err
	}
	return blog, nil
}

func NewService(cfg config.Config, repo domain.Repository) domain.Service {
	f := utils.NewLogFormatter("user.service")
	return &service{cfg: cfg, repo: repo, f: f}
}

func (s *service) CreateBlogPost(model domain.Blog) (domain.Blog, error) {
	blog, err := s.repo.StoreBlog(model)
	if err != nil {
		return domain.Blog{}, err
	}
	return blog, nil
}

func (s *service) GetBlogByTitle(title string) (*domain.Blog, error) {
	blog, err := s.repo.FindBlogByTitle(title)
	if err != nil {
		return nil, err
	}
	return blog, nil
}

func (s *service) DeleteBlogByID(id int) error {
	if err := s.repo.DeleteBlogByID(id); err != nil {
		return err
	}
	return nil
}
