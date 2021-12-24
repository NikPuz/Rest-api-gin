package main

import (
	h "RESTful_API_Gin/internal/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	albumHandler, authHandlers := InjectHandlers()

	router.GET("/albums/", albumHandler.GetAlbums)
	router.GET("/albums/:page", albumHandler.GetAlbums)

	router.GET("/album/:id", albumHandler.GetAlbumByIdHandler)
	router.POST("/album/add", h.UserIdentity, albumHandler.AddAlbum)
	router.POST("/album/update", h.UserIdentity, albumHandler.UpdateAlbum)
	router.DELETE("/album/:id", h.UserIdentity, albumHandler.DeleteAlbum)

	router.POST("/login", authHandlers.Login)
	router.POST("/signin", authHandlers.Signin)

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

//func RegisterAlbumHandlers(group *gin.RouterGroup, repo interfaces.IAlbumRepository) {
//	h := AlbumHandlers{repo: repo}
//
//	//group.GET("/albums/", h.GetAlbums)
//	//group.GET("/albums/:page", h.GetAlbums)
//
//	group.GET("/album/:id", UserIdentity, h.GetAlbumById)
//	//group.POST("/album/add", UserIdentity, h.AddAlbum)
//	//group.POST("/album/update", UserIdentity, h.UpdateAlbum)
//	//group.DELETE("/album/:id", UserIdentity, h.DeleteAlbum)
//}
