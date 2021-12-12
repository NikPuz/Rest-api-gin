package main

import (
	h "RESTful_API_Gin/internal/handler"
	"RESTful_API_Gin/pkg"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main()  {
	router := gin.Default()
	fmt.Println(pkg.HashPassword("Mishasha"))
	album := router.Group("/album", h.UserIdentity)
	{
		album.POST("/:id", h.GetAlbumById)
		album.GET("/delete/:id", h.DeleteAlbum)
		album.POST("/add", h.AddAlbum)
		album.POST("/update", h.UpdateAlbum)
	}
	albums := router.Group("/albums")
	{
		albums.GET("/", h.GetAlbums)
		albums.GET("/:page", h.GetAlbums)
	}
	router.POST("/login", h.Login)
	router.POST("/signin", h.Signin)
	router.Run("localhost:8080")

}