package models

import "time"

type Repository interface {
	GetAlbumsPage(page int, records int) ([]Album, error)
	GetAlbumByID(id string) (Album, error)
	AddAlbum(album Album) (Album, error)
	DeleteAlbum(id string) error
	UpdateAlbum(title string, artist string, price string, id string) error
	GetPasswordAndIdByName(name string) (string, string, error)
	SaveRToken(id string, rToken string, ExpiresATToken time.Time) error
	CheckName(name string) error
	AddUser(userData LoginData) error
}
