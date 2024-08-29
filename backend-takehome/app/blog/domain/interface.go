package domain

type Service interface {
	CreateBlogPost(model Blog) (Blog, error)
	GetBlogs(userID int) (*[]Blog, error)
	GetBlogByID(blogID int) (*Blog, error)
	GetBlogByTitle(title string) (*Blog, error)
	UpdateBlogByID(blog Blog) (*Blog, error)
	DeleteBlogByID(id int) error
}

type Repository interface {
	//StoreBlog
	StoreBlog(model Blog) (Blog, error)
	FindBlogByTitle(title string) (*Blog, error)
	//FindBlogs
	FindBlogs(userID int) (*[]Blog, error)
	//FindBlogByID
	FindBlogByID(blogID int) (*Blog, error)
	//UpdateBlogByID
	UpdateBlogByID(blog Blog) (*Blog, error)
	//DeleteBlogByID
	DeleteBlogByID(id int) error

	//StoreCommentByBlogID
	//GetCommentsByBlogID
}
