package repository

import (
	"database/sql"
	"fmt"
	"restful-api-gin/internal/entity"

	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
)

type albumRepository struct {
	db *sql.DB
	logger *zap.Logger
}

func NewAlbumRepository(db *sql.DB, logger *zap.Logger) entity.IAlbumRepository {
	albumRepository := new(albumRepository)
	albumRepository.db = db
	albumRepository.logger = logger
	return albumRepository
}

func (r *albumRepository) GetPage(page int, records int) ([]entity.Album, error) {

	page = (page * records) - records
	rows, err := r.db.Query(fmt.Sprintf("SELECT * FROM `albums` LIMIT %d, %d", page, records))
	if err != nil {
		return []entity.Album{}, err
	}
	defer rows.Close()

	var result []entity.Album

	for rows.Next() {
		var album entity.Album
		err := rows.Scan(&album.ID, &album.Title, &album.Artist, &album.Price)
		if err != nil {
			return []entity.Album{}, err
		}
		result = append(result, album)
	}
	
	return result, err
}

func (r *albumRepository) GetById(id string) (*entity.Album, error) {

	rows, err := r.db.Query(fmt.Sprintf("SELECT id, Title, Artist, Price FROM albums WHERE id = %s", id))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var album entity.Album

	rows.Next()
	rows.Scan(&album.ID, &album.Title, &album.Artist, &album.Price)

	return &album, nil
}

func (r *albumRepository) Add(album entity.Album) {
	r.db.Exec(fmt.Sprintf("INSERT INTO albums (Title, Artist, Price) values('%s', '%s', '%f')",
		album.Title, album.Artist, album.Price))
}

func (r *albumRepository) Delete(id string) {
	r.db.Exec(fmt.Sprintf("DELETE FROM albums WHERE id = %s", id))
}

func (r *albumRepository) Update(title string, artist string, price string, id string) {
	r.db.Exec(fmt.Sprintf("UPDATE albums SET Title = %s, Artist = %s, Price = %s WHERE id = %s",
		title, artist, price, id))
}
