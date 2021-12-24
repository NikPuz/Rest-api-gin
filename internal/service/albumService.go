package service

import (
	"RESTful_API_Gin/interfaces"
	"RESTful_API_Gin/models"
	"strconv"
)

type AlbumService struct {
	interfaces.IAlbumRepository
}

func (s AlbumService) GetAlbumByIdService(id string) (models.Album, error) {
	return s.GetAlbumById(id)
}

func (s AlbumService) GetAlbumsPageService(pageString string) ([]models.Album, error) {
	records := 3
	page, err := strconv.ParseUint(pageString, 0, 32)
	if err != nil || page == 0 {
		page = 1
	}
	return s.GetAlbumsPage(int(page), records)
}

func (s AlbumService) AddAlbumService(album models.Album) {
	s.AddAlbum(album)
}

func (s AlbumService) DeleteAlbumService(id string) {
	s.DeleteAlbum(id)
}

func (s AlbumService) UpdateAlbumService(title string, artist string, price string, id string) {
	s.UpdateAlbum(title, artist, price, id)
}
