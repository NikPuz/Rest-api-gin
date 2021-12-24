package handler

import (
	"RESTful_API_Gin/interfaces"
	"RESTful_API_Gin/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthHandlers struct {
	interfaces.IAuthService
}

func (h AuthHandlers) Login(c *gin.Context) {
	// ПОЛУЧЕНИЕ ДАННЫХ

	var loginData models.LoginData
	if err := c.BindJSON(&loginData); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "Invalid json provided"})
		return
	}
	// СЕРВИС АВТОРИЗАЦИИ

	jwToken, rToken, err := h.LoginService(loginData)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}
	// ОТПРАВКА ТОКЕНОВ

	c.JSON(http.StatusOK, gin.H{
		"accessToken":  jwToken,
		"refreshToken": rToken,
	})
}

func (h AuthHandlers) Signin(c *gin.Context) {

	var signinData models.LoginData
	if err := c.BindJSON(&signinData); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "Invalid json provided"})
		return
	}

	err := h.SigninService(signinData)
	if err != nil && err.Error() == "this name is taken" {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "this name is taken"})
		return
	} else if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "err"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user created"})
}
