package handler

import (
	"fmt"
	"net/http"
	"restful-api-gin/internal/entity"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type albumHandler struct {
	albumService entity.IAlbumService
	logger *zap.Logger
}

func RegisterAlbumHandlers(g *gin.RouterGroup, albumService entity.IAlbumService, logger *zap.Logger) {
	albumHandler := new(albumHandler)
	albumHandler.albumService = albumService
	albumHandler.logger = logger

	g.GET("/albums/", albumHandler.GetAlbums)
	g.GET("/albums/:page", albumHandler.GetAlbums)

	g.GET("/album/:id", UserIdentity, albumHandler.GetById)
	g.POST("/album/add", UserIdentity, albumHandler.Add)
	g.POST("/album/update", UserIdentity, albumHandler.Update)
	g.DELETE("/album/:id", UserIdentity, albumHandler.Delete)
}

func (h albumHandler) GetAlbums(c *gin.Context) {
	
	page := c.Param("page")

	albums, err := h.albumService.GetPage(page)
	if err != nil {
		h.logger.Error("Failed fetch albums page", zap.Error(err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Server Err"})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"albums": albums})
}

func (h albumHandler) GetById(c *gin.Context) {

	id := c.Param("id")

	album, err := h.albumService.GetById(id)
	
	if err != nil {
		h.logger.Error("Failed fetch album by id", zap.Error(err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Server Err"})
		return
	} else if *album == *new(entity.Album) {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "Album not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"album": album})
}

func (h albumHandler) Add(c *gin.Context) {

	var newAlbum entity.Album

	err := c.BindJSON(&newAlbum)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Incorrect album"})
		return
	}

	h.albumService.Add(newAlbum)

	c.JSON(http.StatusOK, gin.H{"message": "Command executed"})
}

func (h albumHandler) Delete(c *gin.Context) {
	id := c.Param("id")

	h.albumService.Delete(id)

	c.JSON(http.StatusOK, gin.H{"message": "Command executed"})
}

func (h albumHandler) Update(c *gin.Context) {

	var newAlbum entity.Album

	if err := c.BindJSON(&newAlbum); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Incorrect album"})
		return
	}
	h.albumService.Update(newAlbum.Title, newAlbum.Artist, fmt.Sprint(newAlbum.Price), fmt.Sprint(newAlbum.ID))

	c.JSON(http.StatusOK, gin.H{"message": "Command executed"})
}
