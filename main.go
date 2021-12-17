package main

import (
	"RESTful_API_Gin/internal/handler"
	repo "RESTful_API_Gin/internal/repository"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/v1")
	repo := repo.NewMySQLRepository()
	handler.RegisterAlbumHandlers(v1, &repo)
	handler.RegisterAuthHandlers(v1, &repo)

	router.Run("localhost:8080")
}

//album := router.Group("/album", handler.UserIdentity)
//{

//}
//albums := router.Group("/albums")
//{
//	albums.GET("/", h.GetAlbums)
//	albums.GET("/:page", h.GetAlbums)
//}
//router.POST("/login", h.Login)
//router.POST("/signin", h.Signin)
