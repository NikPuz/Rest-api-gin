package service

import (
	"fmt"
	"restful-api-gin/internal/entity"
	"strconv"

	"go.uber.org/zap"
)

type albumService struct {
	albumRepo entity.IAlbumRepository
	logger    *zap.Logger
}

func NewAlbumService(albumRepo entity.IAlbumRepository, logger *zap.Logger) entity.IAlbumService {
	albumService := new(albumService)
	albumService.albumRepo = albumRepo
	albumService.logger = logger
	return albumService
}

func (s albumService) GetById(id string) (*entity.Album, error) {
	return s.albumRepo.GetById(id)
}

func (s albumService) GetPage(pageString string) ([]entity.Album, error) {
	records := 3
	page, err := strconv.ParseUint(pageString, 0, 32)
	if err != nil || page == 0 {
		page = 1
	}
	return s.albumRepo.GetPage(int(page), records)
}

func (s albumService) Add(album entity.Album) {
	s.albumRepo.Add(album)
}

func (s albumService) Delete(id string) {
	s.albumRepo.Delete(id)
}

func (s albumService) Update(title string, artist string, price string, id string) {
	if title == "" {
		title = "Title"
	} else {
		title = fmt.Sprintf("'%s'", title)
	}
	if artist == "" {
		artist = "Artist"
	} else {
		artist = fmt.Sprintf("'%s'", artist)
	}
	if price == "0" {
		price = "Price"
	} else {
		price = fmt.Sprintf("'%s'", price)
	}

	s.albumRepo.Update(title, artist, price, id)
}
