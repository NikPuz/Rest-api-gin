package interfaces

import m "RESTful_API_Gin/models"

type IAlbumService interface {
	GetAlbumByIdService(string) (m.Album, error)
	GetAlbumsPageService(page string) ([]m.Album, error)
	AddAlbumService(album m.Album)
	DeleteAlbumService(id string)
	UpdateAlbumService(title string, artist string, price string, id string)
}
