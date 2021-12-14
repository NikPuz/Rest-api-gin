package handler

import (
	repo "RESTful_API_Gin/internal/repository"
	"RESTful_API_Gin/models"
	"RESTful_API_Gin/pkg"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func GetAlbums(c *gin.Context) {
	records := 3
	page, err := strconv.ParseUint(c.Param("page"), 0, 32)
	if err != nil || page == 0 {
		page = 1
	}

	albums, err := repo.GetAlbumsPage(int(page), records)
	if err != nil {
		panic(err.Error())
	}
	c.IndentedJSON(http.StatusOK, albums)
}

func GetAlbumById(c *gin.Context) {
	id := c.Param("id")
	album, err := repo.GetAlbumByID(id)

	if err == sql.ErrNoRows {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
		return
	} else if err != nil {
		panic(err.Error())
	}

	c.IndentedJSON(http.StatusOK, album)
}

func AddAlbum(c *gin.Context) {
	var newAlbums models.Album

	err := c.BindJSON(&newAlbums)
	if err != nil {
		c.Writer.Write([]byte(err.Error()))
		return
	}

	newAlbums, err = repo.AddAlbum(newAlbums)
	if err != nil {
		c.Writer.Write([]byte(err.Error()))
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"message": "command executed"})
}

func DeleteAlbum(c *gin.Context) {
	id := c.Param("id")

	err := repo.DeleteAlbum(id)
	if err != nil {
		c.Writer.Write([]byte(err.Error()))
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "command executed"})
}

func UpdateAlbum(c *gin.Context) {
	var newAlbum models.Album

	if err := c.BindJSON(&newAlbum); err != nil {
		c.Writer.Write([]byte(err.Error()))
		return
	}

	err := repo.UpdateAlbum(newAlbum.Title, newAlbum.Artist, fmt.Sprint(newAlbum.Price), fmt.Sprint(newAlbum.ID))
	if err != nil {
		c.Writer.Write([]byte(err.Error()))
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "command executed"})
}

func Login(c *gin.Context) {
	// ПОЛУЧЕНИЕ ДАННЫХ

	var loginData models.LoginData
	if err := c.BindJSON(&loginData); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "Invalid json provided"})
		return
	}
	// ПРОВЕРКА ДАННЫХ

	id, password, err := repo.GetPasswordAndIdByName(loginData.Name)

	if err == sql.ErrNoRows || !pkg.CheckPasswordHash(loginData.Password, password) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "username or password is incorrect"})
		return
	} else if err != nil {
		c.Writer.Write([]byte(err.Error()))
		return
	}
	// СОЗДАНИЕ ТОКЕНОВ

	jwToken, err := pkg.CreateJWToken(id)
	if err != nil {
		c.Writer.Write([]byte(err.Error()))
		return
	}

	rToken, err := pkg.CreateRToken(id)
	if err != nil {
		c.Writer.Write([]byte(err.Error()))
		return
	}
	// СОХРАНЕНИЕ ТОКЕНА

	ExpiresATToken := time.Now().Add(15 * time.Minute)
	if err := repo.SaveRToken(id, rToken, ExpiresATToken); err != nil {
		c.Writer.Write([]byte(err.Error()))
		return
	}
	// ОТПРАВКА ТОКЕНОВ

	c.JSON(http.StatusOK, gin.H{
		"accessToken":  jwToken,
		"refreshToken": rToken,
	})
}

func Signin(c *gin.Context) {
	// ПОЛУЧЕНИЕ ДАННЫХ

	var signinData models.LoginData
	if err := c.BindJSON(&signinData); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "Invalid json provided"})
		return
	}
	fmt.Println(signinData)
	// ПРОВЕРКА ИМЕНИ

	if err := repo.CheckName(signinData.Name); err == nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "this name is taken"})
		return
	} else if err != sql.ErrNoRows {
		c.Writer.Write([]byte(err.Error()))
		return
	}
	// ХЕШИРОВАНИЕ ПАРОЛЯ

	if err := pkg.HashPassword(&signinData.Password); err != nil {
		c.Writer.Write([]byte(err.Error()))
		return
	}
	// СОЗДАНИЕ ПОЛЬЗОВАТЕЛЯ

	err := repo.AddUser(signinData)
	if err != nil {
		c.Writer.Write([]byte(err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "user created"})
}
