package main

import (
	"url_shortener/internal/config"
	"url_shortener/internal/repository/redis_repo"
	"url_shortener/internal/repository/sql"
	"url_shortener/internal/service"
	"url_shortener/internal/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	db := config.ConnectDB()
	repo := sql.NewURLRepository(db)
	cache := redis_repo.NewCache()
	svc := service.NewURLService(repo, cache)
	h := handler.NewURLHandler(svc)
	r := gin.Default()

	r.POST("/shorten", h.ShortenURL)
	r.GET("/:code", h.Redirect)
	r.Run(":8080")
}
