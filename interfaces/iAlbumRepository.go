package interfaces

import (
	m "RESTful_API_Gin/models"
)

type IAlbumRepository interface {
	GetAlbumById(id string) (m.Album, error)
	GetAlbumsPage(page int, records int) ([]m.Album, error)
	AddAlbum(album m.Album)
	DeleteAlbum(id string)
	UpdateAlbum(title string, artist string, price string, id string)
}
