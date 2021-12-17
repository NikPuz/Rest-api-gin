package handler

import (
	"RESTful_API_Gin/models"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type AlbumHandlers struct {
	repo models.Repository
}

func RegisterAlbumHandlers(group *gin.RouterGroup, repo models.Repository) {
	h := AlbumHandlers{repo: repo}

	group.GET("/albums/", h.GetAlbums)
	group.GET("/albums/:page", h.GetAlbums)

	group.GET("/album/:id", UserIdentity, h.GetAlbumById)
	group.POST("/album/add", UserIdentity, h.AddAlbum)
	group.POST("/album/update", UserIdentity, h.UpdateAlbum)
	group.DELETE("/album/:id", UserIdentity, h.DeleteAlbum)
}

func (h *AlbumHandlers) GetAlbums(c *gin.Context) {
	records := 3
	page, err := strconv.ParseUint(c.Param("page"), 0, 32)
	if err != nil || page == 0 {
		page = 1
	}

	albums, err := h.repo.GetAlbumsPage(int(page), records)
	if err != nil {
		panic(err.Error())
	}
	c.IndentedJSON(http.StatusOK, albums)
}

func (h *AlbumHandlers) GetAlbumById(c *gin.Context) {
	id := c.Param("id")
	album, err := h.repo.GetAlbumByID(id)

	if err == sql.ErrNoRows {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
		return
	} else if err != nil {
		panic(err.Error())
	}

	c.IndentedJSON(http.StatusOK, album)
}

func (h *AlbumHandlers) AddAlbum(c *gin.Context) {
	var newAlbums models.Album

	err := c.BindJSON(&newAlbums)
	if err != nil {
		c.Writer.Write([]byte(err.Error()))
		return
	}

	newAlbums, err = h.repo.AddAlbum(newAlbums)
	if err != nil {
		c.Writer.Write([]byte(err.Error()))
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"message": "command executed"})
}

func (h *AlbumHandlers) DeleteAlbum(c *gin.Context) {
	id := c.Param("id")

	err := h.repo.DeleteAlbum(id)
	if err != nil {
		c.Writer.Write([]byte(err.Error()))
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "command executed"})
}

func (h *AlbumHandlers) UpdateAlbum(c *gin.Context) {
	var newAlbum models.Album

	if err := c.BindJSON(&newAlbum); err != nil {
		c.Writer.Write([]byte(err.Error()))
		return
	}

	err := h.repo.UpdateAlbum(newAlbum.Title, newAlbum.Artist, fmt.Sprint(newAlbum.Price), fmt.Sprint(newAlbum.ID))
	if err != nil {
		c.Writer.Write([]byte(err.Error()))
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "command executed"})
}
