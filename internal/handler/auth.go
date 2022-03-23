package handler

import (
	"net/http"
	"restful-api-gin/internal/entity"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type authHandler struct {
	authService entity.IAuthService
	logger *zap.Logger
}

func RegisterAuthHandlers(group *gin.RouterGroup, authService entity.IAuthService, logger *zap.Logger) {
	authHandler := new(authHandler)
	authHandler.authService = authService
	authHandler.logger = logger

	group.POST("/login", authHandler.Login)
	group.POST("/login/refresh", authHandler.RefreshLogin)
	group.POST("/signin", authHandler.Signin)
}

func (h authHandler) Signin(c *gin.Context) {

	var signinData entity.LoginData
	if err := c.BindJSON(&signinData); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid json provided"})
		return
	}

	err := h.authService.SigninService(signinData)
	if err != nil && err.Error() == "name is taken" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Name is taken"})
		return
	} else if err != nil {
		h.logger.Error("Failed Signin Service", zap.Error(err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Server Err"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created"})
}

func (h authHandler) Login(c *gin.Context) {
	// ПОЛУЧЕНИЕ ДАННЫХ

	var loginData entity.LoginData
	if err := c.BindJSON(&loginData); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid json provided"})
		return
	}
	// СЕРВИС АВТОРИЗАЦИИ

	jwToken, rjwToken, err := h.authService.LoginService(loginData)
	if err != nil && err.Error() == "incorrect password or name" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Incorrect password or name"})
		return
	} else if err != nil {
		h.logger.Error("Failed Login Service", zap.Error(err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Server Err"})
		return
	}
	// ОТПРАВКА ТОКЕНОВ

	c.JSON(http.StatusOK, gin.H{
		"accessToken": jwToken,
		"refreshToken": rjwToken,
	})
}

func (h authHandler) RefreshLogin(c *gin.Context) {
	// ПОЛУЧЕНИЕ ДАННЫХ

	var refreshToken entity.RJWToken // Можно лучше
	if err := c.BindJSON(&refreshToken); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid json provided"})
		return
	}
	// СЕРВИС АВТОРИЗАЦИИ

	jwToken, rjwToken, err := h.authService.RefreshLoginService(refreshToken.RJWToken)
	if err != nil && err.Error() == "token is empty" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Token is empty"})
		return
	} else if err != nil {
		h.logger.Error("Failed Login Service", zap.Error(err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Server Err"})
		return
	}
	// ОТПРАВКА ТОКЕНОВ

	c.JSON(http.StatusOK, gin.H{
		"accessToken": jwToken,
		"refreshToken": rjwToken,
	})
}