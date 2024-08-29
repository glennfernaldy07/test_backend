package service

import (
	_blogDomain "app/blog/domain"
	"app/common/utils"
	"app/config"
	"app/posts/domain"
	"fmt"
	"gopkg.in/jinzhu/gorm.v1"
	"time"
)

type service struct {
	cfg     config.Config
	blogSvc _blogDomain.Service
	f       utils.LogFormatter
}

func NewService(cfg config.Config, blogSvc _blogDomain.Service) domain.Service {
	f := utils.NewLogFormatter("user.service")
	return &service{cfg: cfg, blogSvc: blogSvc, f: f}
}

func (s *service) CreateBlogPost(request domain.PostsBlogRequest) (int, error) {
	//validate blog by title
	if err := s.isValidBlogTitle(request.Title); err != nil {
		//if already found return error blog title has been used
		return 0, err
	}
	//parse data to blog model
	model := _blogDomain.Blog{
		Title:     request.Title,
		Content:   request.Content,
		AuthorID:  int(request.UserID),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	//store blog model
	model, err := s.blogSvc.CreateBlogPost(model)
	if err != nil {
		return 0, err
	}
	//return blog id
	return model.ID, nil
}
func (s *service) isValidBlogTitle(title string) error {
	blog, err := s.blogSvc.GetBlogByTitle(title)
	if err == nil {
		errMsg := fmt.Errorf("blog post %s already exists, blog id=%d", title, blog.ID)
		return errMsg
	} else {
		if gorm.IsRecordNotFoundError(err) {
			return nil
		}
		errMsg := fmt.Errorf("blog post %s already exists, blog id=%d", title, blog.ID)
		return errMsg
	}

}

func (s *service) GetBlog(request domain.GetBlogRequest) ([]_blogDomain.Blog, error) {
	blog, err := s.blogSvc.GetBlogs(int(request.UserID))
	if err != nil {
		return []_blogDomain.Blog{}, err
	}
	return *blog, nil
}

func (s *service) GetBlogByID(request domain.GetBlogRequest) (_blogDomain.Blog, error) {
	blog, err := s.blogSvc.GetBlogByID(int(request.ID))
	if err != nil {
		return _blogDomain.Blog{}, err
	}
	return *blog, nil
}

func (s *service) UpdateBlogByID(req domain.UpdateBlogRequest) (_blogDomain.Blog, error) {

	model := _blogDomain.Blog{
		ID:        req.ID,
		AuthorID:  int(req.UserID),
		UpdatedAt: time.Now(),
	}
	if req.Title != "" {
		model.Title = req.Title
	}
	if req.Content != "" {
		model.Content = req.Content
	}
	blog, err := s.blogSvc.UpdateBlogByID(model)
	if err != nil {
		return _blogDomain.Blog{}, err
	}
	return *blog, nil
}

func (s *service) DeleteBlogPostByID(request domain.DeleteBlogRequest) (int, error) {
	if err := s.blogSvc.DeleteBlogByID(request.ID); err != nil {
		return 0, err
	}
	return request.ID, nil
}
