package http

import (
	"app/common/utils"
	"app/posts/domain"
	"context"
	"encoding/json"
	"errors"
	kitendpoint "github.com/go-kit/kit/endpoint"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"time"
)

const tokenHeader = "Authorization"

type endpoint struct {
	ps domain.Service
	f  utils.LogFormatter
}

func (e endpoint) MakePostBlogRequest() kitendpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(domain.PostsBlogRequest)
		if !ok {
			return nil, errors.New("format tidak sesuai")
		}

		id, err := e.ps.CreateBlogPost(req)
		if err != nil {
			return nil, err
		}

		resp := domain.PostsBlogResponse{
			ID:        id,
			Title:     req.Title,
			Content:   req.Content,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		return Response{HTTPCode: http.StatusOK, Data: resp}, nil
	}
}

func (e endpoint) DecodePostBlogRequest(ctx context.Context, request *http.Request) (interface{}, error) {
	req := domain.PostsBlogRequest{}

	err := json.NewDecoder(request.Body).Decode(&req)
	if err != nil {
		return nil, err
	}

	token := request.Header.Get(tokenHeader)
	tokenString := token[len("Bearer "):]
	userID, err := e.getUserIdByAuthToken(tokenString)
	if err != nil {
		return nil, err
	}

	//ctx = context.WithValue(ctx, "userId", userID)
	req.UserID, _ = strconv.ParseInt(userID, 10, 64)

	return req, nil
}
func (e endpoint) getUserIdByAuthToken(token string) (string, error) {
	//GET AUTHORIZATION HEADER
	userID, err := utils.VerifyToken(token)
	if err != nil {
		return "", err
	}
	return userID, nil
}

func (e endpoint) MakeGetBlogRequest() kitendpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(domain.GetBlogRequest)
		if !ok {
			return nil, errors.New("format tidak sesuai")
		}

		if req.ID > 0 {
			blog, err := e.ps.GetBlogByID(req)
			if err != nil {
				return nil, err
			}
			resp := domain.GetBlogResponse{
				ID:        blog.ID,
				Title:     blog.Title,
				Content:   blog.Content,
				CreatedAt: blog.CreatedAt,
				UpdatedAt: blog.UpdatedAt,
			}
			return Response{HTTPCode: http.StatusOK, Data: resp}, nil
		}
		blogs, err := e.ps.GetBlog(req)
		if err != nil {
			return nil, err
		}

		resps := []domain.GetBlogResponse{}
		for _, blog := range blogs {
			resp := domain.GetBlogResponse{
				ID:        blog.ID,
				Title:     blog.Title,
				Content:   blog.Content,
				CreatedAt: blog.CreatedAt,
				UpdatedAt: blog.UpdatedAt,
			}
			resps = append(resps, resp)
		}

		return Response{HTTPCode: http.StatusOK, Data: resps}, nil
	}
}

func (e endpoint) DecodeGetBlogRequest(ctx context.Context, request *http.Request) (interface{}, error) {
	req := domain.GetBlogRequest{}

	vars := mux.Vars(request)
	id := vars["id"]
	if id != "" {
		req.ID, _ = strconv.Atoi(id)
	} else {
		req.ID = 0
	}

	token := request.Header.Get(tokenHeader)
	tokenString := token[len("Bearer "):]
	userID, err := e.getUserIdByAuthToken(tokenString)
	if err != nil {
		return nil, err
	}

	req.UserID, _ = strconv.ParseInt(userID, 10, 64)

	return req, nil
}

func (e endpoint) MakeUpdateBlogRequest() kitendpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(domain.UpdateBlogRequest)
		if !ok {
			return nil, errors.New("format tidak sesuai")
		}

		blog, err := e.ps.UpdateBlogByID(req)
		if err != nil {
			return nil, err
		}

		resp := domain.UpdateBlogResponse{
			ID:        blog.ID,
			Title:     blog.Title,
			Content:   blog.Content,
			UpdatedAt: blog.UpdatedAt,
		}

		return Response{HTTPCode: http.StatusOK, Data: resp}, nil
	}
}

func (e endpoint) DecodeUpdateBlogRequest(ctx context.Context, request *http.Request) (interface{}, error) {
	req := domain.UpdateBlogRequest{}

	err := json.NewDecoder(request.Body).Decode(&req)
	if err != nil {
		return nil, err
	}

	vars := mux.Vars(request)
	id := vars["id"]
	req.ID, _ = strconv.Atoi(id)

	token := request.Header.Get(tokenHeader)
	tokenString := token[len("Bearer "):]
	userID, err := e.getUserIdByAuthToken(tokenString)
	if err != nil {
		return nil, err
	}

	req.UserID, _ = strconv.ParseInt(userID, 10, 64)

	return req, nil
}

func (e endpoint) MakeDeleteBlogRequest() kitendpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(domain.DeleteBlogRequest)
		if !ok {
			return nil, errors.New("format tidak sesuai")
		}

		id, err := e.ps.DeleteBlogPostByID(req)
		if err != nil {
			return nil, err
		}

		resp := domain.PostsBlogResponse{
			ID: id,
		}

		return Response{HTTPCode: http.StatusOK, Data: resp}, nil
	}
}

func (e endpoint) DecodeDeleteBlogRequest(ctx context.Context, request *http.Request) (interface{}, error) {
	req := domain.DeleteBlogRequest{}

	vars := mux.Vars(request)
	id := vars["id"]
	req.ID, _ = strconv.Atoi(id)

	token := request.Header.Get(tokenHeader)
	tokenString := token[len("Bearer "):]
	userID, err := e.getUserIdByAuthToken(tokenString)
	if err != nil {
		return nil, err
	}

	req.UserID, _ = strconv.ParseInt(userID, 10, 64)

	return req, nil
}

func NewEndpoint(us domain.Service) domain.Endpoint {
	f := utils.NewLogFormatter("posts.delivery.endpoint")
	return &endpoint{us, f}
}
