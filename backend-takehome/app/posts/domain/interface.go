package domain

import (
	_blogDomain "app/blog/domain"
	"context"
	kitendpoint "github.com/go-kit/kit/endpoint"
	"net/http"
)

type Endpoint interface {
	// POST BLOG ENDPOINT
	MakePostBlogRequest() kitendpoint.Endpoint
	DecodePostBlogRequest(context.Context, *http.Request) (interface{}, error)
	// GET BLOG ENDPOINT
	MakeGetBlogRequest() kitendpoint.Endpoint
	DecodeGetBlogRequest(context.Context, *http.Request) (interface{}, error)
	// UPDATE BLOG REQUEST
	MakeUpdateBlogRequest() kitendpoint.Endpoint
	DecodeUpdateBlogRequest(context.Context, *http.Request) (interface{}, error)
	// DELETE BLOG REQUEST
	MakeDeleteBlogRequest() kitendpoint.Endpoint
	DecodeDeleteBlogRequest(context.Context, *http.Request) (interface{}, error)
}

type Service interface {
	CreateBlogPost(request PostsBlogRequest) (int, error)
	GetBlog(request GetBlogRequest) ([]_blogDomain.Blog, error)
	GetBlogByID(request GetBlogRequest) (_blogDomain.Blog, error)
	UpdateBlogByID(request UpdateBlogRequest) (_blogDomain.Blog, error)
	DeleteBlogPostByID(request DeleteBlogRequest) (int, error)
}
