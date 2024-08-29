package db

import (
	"app/common/utils"
	"app/config"
	"app/user/domain"
	"fmt"
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

func (r *repository) StoreUser(model domain.User) error {
	db := r.DB
	err := db.Create(&model).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) FindUserByEmail(email string) (domain.User, error) {

	var result domain.User
	db := r.DB
	if err := db.Where(&domain.User{Email: email}).
		Find(&result).Error; err != nil {
		msg := fmt.Errorf("cannot find user: error=%s, username=%s", err, email)
		fmt.Println(msg)
		return result, err
	}

	return result, nil
}

func (r *repository) FindUserByUsernameAndPassword(email, password string) (*domain.User, error) {

	db := r.DB
	var result domain.User
	if err := db.Where(&domain.User{Email: email, Password: password}).
		Find(&result).Error; err != nil {
		msg := fmt.Errorf("cannot find user: error=%s, username=%s", err, email)
		fmt.Println(msg)
		return nil, err
	}

	return &result, nil
}
