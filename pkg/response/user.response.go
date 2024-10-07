package response

import (
	"time"

	"github.com/LeMinh0706/music-go/db"
)

type RegisterResquest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Gender   int32  `json:"gender" binding:"required,min=0,max=1"`
}

type RegisterResponse struct {
	ID       int64     `json:"id"`
	Username string    `json:"username"`
	Fullname string    `json:"fullname"`
	Email    string    `json:"email"`
	Gender   int32     `json:"gender"`
	RoleID   int32     `json:"role_id"`
	Avt      string    `json:"avt"`
	CreateAt time.Time `json:"create_at"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	ID       int64     `json:"id"`
	Fullname string    `json:"fullname"`
	Email    string    `json:"email"`
	Gender   int32     `json:"gender"`
	RoleID   int32     `json:"role_id"`
	Avt      string    `json:"avt"`
	CreateAt time.Time `json:"create_at"`
}

func RegisterRes(user db.User) RegisterResponse {
	return RegisterResponse{
		ID:       user.ID,
		Username: user.Username,
		Fullname: user.Fullname,
		Email:    user.Email.String,
		Gender:   user.Gender,
		RoleID:   user.RoleID,
		Avt:      user.Avt,
		CreateAt: user.CreateAt,
	}
}

func LoginRes(user db.User) LoginResponse {
	return LoginResponse{
		ID:       user.ID,
		Fullname: user.Fullname,
		Email:    user.Email.String,
		Gender:   user.Gender,
		RoleID:   user.RoleID,
		Avt:      user.Avt,
		CreateAt: user.CreateAt,
	}
}
