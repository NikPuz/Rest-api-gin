package handler

import (
	"RESTful_API_Gin/interfaces"
	"RESTful_API_Gin/models"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AlbumHandler struct {
	interfaces.IAlbumService
}

func (h *AlbumHandler) GetAlbums(c *gin.Context) {

	page := c.Param("page")

	albums, err := h.GetAlbumsPageService(page)
	if err != nil {
		panic(err.Error())
	}

	c.IndentedJSON(http.StatusOK, albums)
}

func (h *AlbumHandler) GetAlbumByIdHandler(c *gin.Context) {

	id := c.Param("id")

	album, err := h.GetAlbumByIdService(id)

	if err == sql.ErrNoRows {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
		return
	} else if err != nil {
		panic(err.Error())
	}

	c.IndentedJSON(http.StatusOK, album)
}

func (h *AlbumHandler) AddAlbum(c *gin.Context) {

	var newAlbum models.Album

	err := c.BindJSON(&newAlbum)
	if err != nil {
		c.Writer.Write([]byte(err.Error()))
		return
	}

	h.AddAlbumService(newAlbum)

	c.IndentedJSON(http.StatusCreated, gin.H{"message": "command executed"})
}

func (h *AlbumHandler) DeleteAlbum(c *gin.Context) {
	id := c.Param("id")

	h.DeleteAlbumService(id)

	c.IndentedJSON(http.StatusOK, gin.H{"message": "command executed"})
}

func (h *AlbumHandler) UpdateAlbum(c *gin.Context) {

	var newAlbum models.Album

	if err := c.BindJSON(&newAlbum); err != nil {
		c.Writer.Write([]byte(err.Error()))
		return
	}

	h.UpdateAlbumService(newAlbum.Title, newAlbum.Artist, fmt.Sprint(newAlbum.Price), fmt.Sprint(newAlbum.ID))

	c.IndentedJSON(http.StatusOK, gin.H{"message": "command executed"})
}
