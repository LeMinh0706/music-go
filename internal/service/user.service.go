package service

import (
	"context"
	"database/sql"
	"strings"

	"github.com/LeMinh0706/music-go/internal/repo"
	"github.com/LeMinh0706/music-go/pkg/response"
	"github.com/LeMinh0706/music-go/util"
)

type UserSevice struct {
	userRepo *repo.UserRepository
}

func NewUserService() (*UserSevice, error) {
	repo, err := repo.NewUserRepo()
	if err != nil {
		return nil, err
	}
	return &UserSevice{
		userRepo: repo,
	}, nil
}

func (us *UserSevice) Register(ctx context.Context, username, password, fullname, email string, gender int32) (response.RegisterResponse, error) {
	var res response.RegisterResponse
	if strings.TrimSpace(fullname) == "" {
		fullname = username
	}
	var nullEmail sql.NullString
	if strings.TrimSpace(email) == "" {
		nullEmail = sql.NullString{Valid: false}
	} else {
		nullEmail = sql.NullString{String: email, Valid: true}
	}
	hashPassword, err := util.HashPashword(password)
	if err != nil {
		return res, err
	}

	user, err := us.userRepo.CreateUser(ctx, username, hashPassword, fullname, nullEmail, gender, 3)
	if err != nil {
		return res, err
	}
	res = response.RegisterRes(user)
	return res, nil
}

func (us *UserSevice) Login(ctx context.Context, username, password string) (response.LoginResponse, error) {
	var res response.LoginResponse
	user, err := us.userRepo.GetUser(ctx, username)
	if err == sql.ErrNoRows {
		return res, err
	}
	if err := util.CheckPassword(password, user.Password); err != nil {
		return res, err
	}
	res = response.LoginRes(user)
	return res, nil
}
