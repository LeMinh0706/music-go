package controller

import (
	"github.com/LeMinh0706/music-go/internal/service"
	"github.com/LeMinh0706/music-go/pkg/response"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserSevice
}

func NewUserController() (*UserController, error) {
	service, err := service.NewUserService()
	if err != nil {
		return nil, err
	}
	return &UserController{
		userService: service,
	}, nil
}

func (uc *UserController) Register(g *gin.Context) {
	var req response.RegisterResquest

	if err := g.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(g, 404, response.ErrBadRequest)
		return
	}

	user, err := uc.userService.Register(g, req.Username, req.Password, req.Fullname, req.Email, req.Gender)
	if err != nil {
		response.ErrorResponse(g, 401, response.ErrFailed)
		return
	}
	response.SuccessResponse(g, 201, user)
}

func (uc *UserController) Login(g *gin.Context) {
	var req response.LoginRequest
	if err := g.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(g, 404, response.ErrBadRequest)
		return
	}
	user, err := uc.userService.Login(g, req.Username, req.Password)
	if err != nil {
		response.ErrorResponse(g, 401, response.ErrFailed)
		return
	}
	response.SuccessResponse(g, 200, user)
}
