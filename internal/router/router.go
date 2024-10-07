package router

import (
	"github.com/gin-gonic/gin"
)

func NewRouter(r *gin.Engine) {
	v1 := r.Group("api/v1")
	{
		NewUserRouter(v1)
	}
}
