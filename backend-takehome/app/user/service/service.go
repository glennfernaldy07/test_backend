package service

import (
	"app/common/snap/snapauth"
	"app/common/utils"
	"app/config"
	"app/user/domain"
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"
)

type service struct {
	cfg       config.Config
	repo      domain.Repository
	cacheRepo domain.CacheRepository
	f         utils.LogFormatter
}

func NewService(cfg config.Config, repo domain.Repository, cacheRepo domain.CacheRepository) domain.Service {
	f := utils.NewLogFormatter("user.service")
	return &service{cfg: cfg, repo: repo, cacheRepo: cacheRepo, f: f}
}

func (s *service) Login(ctx context.Context, req domain.LoginRequest) (domain.GeneralResponse, error) {
	// check is valid email
	if !s.isValidEmail(ctx, req.Email) {
		errMsg := fmt.Errorf("email not found, email=%s", req.Email)
		return domain.GeneralResponse{}, errMsg
	}
	//find email
	userModel, err := s.repo.FindUserByEmail(req.Email)
	if err != nil {
		errMsg := fmt.Errorf("email not found, email=%s, error=%s", req.Email, err.Error())
		return domain.GeneralResponse{}, errMsg
	}
	// compare password
	if !utils.VerifyPassword(req.Password, userModel.Password) {
		return domain.GeneralResponse{}, errors.New("invalid password")
	}
	// generate token
	tokenExpTime := s.cfg.GetInt(config.TokenExpTime)
	token, errCreateToken := utils.CreateToken(userModel.ID, req.Email, tokenExpTime)
	if errCreateToken != nil {
		return domain.GeneralResponse{}, errCreateToken
	}

	// store token
	dt := snapauth.AccessTokenResponse{
		AccessToken:    token,
		TokenType:      "bearer",
		ExpiresIn:      strconv.FormatInt(tokenExpTime, 10),
		AdditionalInfo: nil,
	}
	if err := s.cacheRepo.StoreAccessToken(ctx, string(rune(userModel.ID)), dt); err != nil {
		return domain.GeneralResponse{}, err
	}
	// return token
	return domain.GeneralResponse{Status: "200", Message: "success", Data: token}, nil
}

func (s *service) Register(ctx context.Context, req domain.RegisterRequest) (domain.GeneralResponse, error) {
	if s.isValidEmail(ctx, req.Email) {
		errMsg := fmt.Errorf("email has been used, email=%s", req.Email)
		return domain.GeneralResponse{}, errMsg
	}
	hashPass, err := utils.HashPassword(req.Password)
	if err != nil {
		errMsg := fmt.Errorf("err while hashing password, email=%s", req.Email)
		return domain.GeneralResponse{}, errMsg
	}
	model := domain.User{
		Name:      req.Name,
		Email:     req.Email,
		Password:  hashPass,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err := s.storeUser(ctx, model); err != nil {
		errMsg := fmt.Errorf("error while store new user, err=%s", err.Error())
		return domain.GeneralResponse{}, errMsg
	}

	resp := domain.GeneralResponse{Status: "200", Message: "success"}
	return resp, nil
}

func (s *service) isValidEmail(ctx context.Context, email string) bool {
	var err error
	user, err := s.repo.FindUserByEmail(email)
	if err != nil {
		return false
	}
	if user.ID <= 0 {
		return false
	}
	return true
}

func (s *service) storeUser(ctx context.Context, model domain.User) error {
	if err := s.repo.StoreUser(model); err != nil {
		return err
	}
	return nil
}

func (s *service) GetUserByUsernameAndPassword(ctx context.Context, username, password string) (*domain.GeneralResponse, error) {
	resp := domain.GeneralResponse{}
	userModel, err := s.repo.FindUserByEmail(username)
	if err != nil {
		return nil, err
	}
	if userModel.ID > 0 {
		resp.Message = "Success"
		resp.Status = "00"
	}

	return &resp, nil
}
