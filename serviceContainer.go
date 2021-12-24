package main

import (
	"RESTful_API_Gin/internal/handler"
	"RESTful_API_Gin/internal/infrastructures"
	"RESTful_API_Gin/internal/repository"
	"RESTful_API_Gin/internal/service"
	"database/sql"
)

func InjectHandlers() (handler.AlbumHandler, handler.AuthHandlers) {
	mySQLHandler := &infrastructures.MySQLHandler{}
	mySQLHandler.Conn = dbConn()

	albumRepository := &repository.AlbumRepository{IDbHandler: mySQLHandler}
	albumService := service.AlbumService{IAlbumRepository: albumRepository} //Нужно ли тут &?
	albumHandler := handler.AlbumHandler{IAlbumService: albumService}

	otherRepository := &repository.OtherRepository{IDbHandler: mySQLHandler}
	authService := service.AuthService{IOtherRepository: otherRepository}
	authHandler := handler.AuthHandlers{IAuthService: authService}

	return albumHandler, authHandler
}

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "vbif15987532587"
	dbName := "albums_db"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}
