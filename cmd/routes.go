package server

import (
	"go-bank/middleware"
	"go-bank/repo"
	"go-bank/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	Repository repo.RepositoryInterface
	Service    service.ServiceInterface
)

func AddRoutes(r *gin.Engine) {
	r.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "server is healthy and running")
	})

	public := r.Group("/api")
	public.GET("/accounts", GetAccounts)

	//Jwt authentication
	secured := r.Group("/api/v1")
	secured.Use(middleware.JwtAuthMiddleware())
	secured.POST("/transfer", Transfer)
}

func buildRoutes() {
	Repository = repo.NewRepository()
	Service = service.NewService(Repository)
}
