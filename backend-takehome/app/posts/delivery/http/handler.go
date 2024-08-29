package http

import (
	"app/common"
	"app/posts/domain"
	kitjwt "github.com/go-kit/kit/auth/jwt"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
)

func NewHandler(
	r *mux.Router,
	e domain.Endpoint,
	option ServerOption,
) {
	opt := []kithttp.ServerOption{
		kithttp.ServerErrorEncoder(option.encodeErrorResponse),
		kithttp.ServerBefore(
			kitjwt.HTTPToContext(),
		),
		kithttp.ServerAfter(
			kithttp.SetResponseHeader(common.SetResponseHeader(common.HeadXFrameOptions)),
			kithttp.SetResponseHeader(common.SetResponseHeader(common.HeadStrictTransportSecurity)),
			kithttp.SetResponseHeader(common.SetResponseHeader(common.HeadExpectCT)),
			kithttp.SetResponseHeader(common.SetResponseHeader(common.HeadContentSecurityPolicy)),
			kithttp.SetResponseHeader(common.SetResponseHeader(common.HeadXXSSProtection)),
			kithttp.SetResponseHeader(common.SetResponseHeader(common.HeadXContentTypeOptions)),
		),
	}

	// User's Blog
	r.Methods(http.MethodPost).Path("/v1/posts").Handler(
		kithttp.NewServer(e.MakePostBlogRequest(), e.DecodePostBlogRequest, option.encodeResponse, opt...),
	)
	r.Methods(http.MethodGet).Path("/v1/posts").Handler(
		kithttp.NewServer(e.MakeGetBlogRequest(), e.DecodeGetBlogRequest, option.encodeResponse, opt...),
	)
	r.Methods(http.MethodGet).Path("/v1/posts/{id}").Handler(
		kithttp.NewServer(e.MakeGetBlogRequest(), e.DecodeGetBlogRequest, option.encodeResponse, opt...),
	)
	r.Methods(http.MethodPut).Path("/v1/posts/{id}").Handler(
		kithttp.NewServer(e.MakeUpdateBlogRequest(), e.DecodeUpdateBlogRequest, option.encodeResponse, opt...),
	)
	r.Methods(http.MethodDelete).Path("/v1/posts/{id}").Handler(
		kithttp.NewServer(e.MakeDeleteBlogRequest(), e.DecodeDeleteBlogRequest, option.encodeResponse, opt...),
	)

	// Blog's Comment
	//r.Methods(http.MethodPost).Path("/v1/posts/{id}/comments").Handler(
	//	kithttp.NewServer(e.makeRegisterRequest(), e.decodeRegisterRequest, option.encodeResponse, opt...),
	//)
	//r.Methods(http.MethodGet).Path("/v1/posts/{id}/comments").Handler(
	//	kithttp.NewServer(e.makeRegisterRequest(), e.decodeRegisterRequest, option.encodeResponse, opt...),
	//)

}
