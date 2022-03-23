package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"restful-api-gin/internal/handler"
	"restful-api-gin/internal/repository"
	"restful-api-gin/internal/service"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
		}
	}()
	fmt.Println("Поехало!")
	logger := initZapLogger()

	initViperConfigger(logger)

	db := dbConnConfig(logger)

	albumRepository := repository.NewAlbumRepository(db, logger)
	userRepository := repository.NewUserRepository(db, logger)

	albumService := service.NewAlbumService(albumRepository, logger)
	authService := service.NewAuthService(userRepository, logger)

	router := gin.Default()
	v1 := router.Group("/v1")
	handler.RegisterAlbumHandlers(v1, albumService, logger)
	handler.RegisterAuthHandlers(v1, authService, logger)

	serverConfig(router).ListenAndServe()
}

func dbConnConfig(logger *zap.Logger) *sql.DB {
	db, err := sql.Open(
		viper.GetString("DB_DRIVER"),
		viper.GetString("DB_USER")+":"+viper.GetString("DB_PASSWORD")+"@"+viper.GetString("DB_SOURCE")+"/"+viper.GetString("DB_DATABASE"))
	if err != nil {
		logger.Error("failed connect to db", zap.Error(err))
		os.Exit(1)
	}
	return db
}

func serverConfig(router http.Handler) *http.Server {
	return &http.Server{
		Addr:           "8000",
		Handler:        router,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   5 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}

func initZapLogger() *zap.Logger {
	logger := zap.NewExample()
	return logger
}

func initViperConfigger(logger *zap.Logger) {
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AddConfigPath("./config/")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		logger.Error("failed read in config", zap.Error(err))
		return
	}
}

//dbDriver := "mysql"
//dbUser := "root"
//dbPass := "vbif15987532587"
//dbName := "albums_db"

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
