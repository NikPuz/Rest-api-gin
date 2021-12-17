package handler

import (
	"RESTful_API_Gin/models"
	"RESTful_API_Gin/pkg"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type AuthHandlers struct {
	repo models.Repository
}

func RegisterAuthHandlers(group *gin.RouterGroup, repo models.Repository) {
	h := AuthHandlers{repo: repo}
	group.POST("/login", h.Login)
	group.POST("/signin", h.Signin)
}

func (h *AuthHandlers) Login(c *gin.Context) {
	// ПОЛУЧЕНИЕ ДАННЫХ

	var loginData models.LoginData
	if err := c.BindJSON(&loginData); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "Invalid json provided"})
		return
	}
	// ПРОВЕРКА ДАННЫХ

	id, password, err := h.repo.GetPasswordAndIdByName(loginData.Name)

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
	if err := h.repo.SaveRToken(id, rToken, ExpiresATToken); err != nil {
		c.Writer.Write([]byte(err.Error()))
		return
	}
	// ОТПРАВКА ТОКЕНОВ

	c.JSON(http.StatusOK, gin.H{
		"accessToken":  jwToken,
		"refreshToken": rToken,
	})
}

func (h *AuthHandlers) Signin(c *gin.Context) {
	// ПОЛУЧЕНИЕ ДАННЫХ

	var signinData models.LoginData
	if err := c.BindJSON(&signinData); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "Invalid json provided"})
		return
	}
	fmt.Println(signinData)
	// ПРОВЕРКА ИМЕНИ

	if err := h.repo.CheckName(signinData.Name); err == nil {
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

	err := h.repo.AddUser(signinData)
	if err != nil {
		c.Writer.Write([]byte(err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "user created"})
}
