package router

import (
	"log"

	"github.com/LeMinh0706/music-go/internal/controller"
	"github.com/gin-gonic/gin"
)

func NewUserRouter(router *gin.RouterGroup) {
	uc, err := controller.NewUserController()
	if err != nil {
		log.Fatal(err)
	}
	userGroup := router.Group("/accounts")
	{
		userGroup.POST("register", uc.Register)
		userGroup.POST("login", uc.Login)
	}
}
