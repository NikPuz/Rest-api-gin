package repository

import (
	"RESTful_API_Gin/interfaces"
	m "RESTful_API_Gin/models"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type AlbumRepository struct {
	interfaces.IDbHandler
}

func (r *AlbumRepository) GetAlbumsPage(page int, records int) ([]m.Album, error) {

	page = (page * records) - records
	row, err := r.Query(fmt.Sprintf("SELECT * FROM `albums` LIMIT %d, %d", page, records))
	if err != nil {
		return []m.Album{}, err
	}

	var result []m.Album

	for row.Next() {
		var album m.Album
		err := row.Scan(&album.ID, &album.Title, &album.Artist, &album.Price)
		if err != nil {
			return []m.Album{}, err
		}
		result = append(result, album)
	}

	return result, err
}

func (r *AlbumRepository) GetAlbumById(id string) (m.Album, error) {

	row, err := r.Query(fmt.Sprintf("SELECT * FROM `albums` WHERE id = %s", id))
	if err != nil {
		return m.Album{}, err
	}

	var album m.Album

	row.Next()
	err = row.Scan(&album.ID, &album.Title, &album.Artist, &album.Price)

	return album, nil
}

func (r *AlbumRepository) AddAlbum(album m.Album) {
	r.Execute(fmt.Sprintf("INSERT INTO albums (Title, Artist, Price) values('%s', '%s', '%f')",
		album.Title, album.Artist, album.Price))
}

func (r *AlbumRepository) DeleteAlbum(id string) {
	r.Execute(fmt.Sprintf("DELETE FROM albums WHERE id = %s", id))
}

func (r *AlbumRepository) UpdateAlbum(title string, artist string, price string, id string) {

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

	r.Execute(fmt.Sprintf("UPDATE albums SET Title = %s, Artist = %s, Price = %s WHERE id = %s",
		title, artist, price, id))
}
