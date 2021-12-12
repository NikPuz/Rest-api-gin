package repository

import (
	m "RESTful_API_Gin/models"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

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

func GetAlbumsPage(page int, records int) ([]m.Album, error) {
	page = (page * records) - 3
	conn := dbConn()
	defer conn.Close()
	rovs, err := conn.Query(fmt.Sprintf("SELECT * FROM `albums` LIMIT %d, %d", page, records))
	if err != nil {
		return []m.Album{}, err
	}

	var result []m.Album
	for rovs.Next() {
		var album m.Album
		err := rovs.Scan(&album.ID, &album.Title, &album.Artist, &album.Price)
		if err != nil {
			return []m.Album{}, err
		}
		result = append(result, album)
	}
	return result, err
}

func GetAlbumByID(id string) (m.Album, error) {
	conn := dbConn()
	defer conn.Close()
	rov := conn.QueryRow(fmt.Sprintf("SELECT * FROM `albums` WHERE id = %s", id))

	if rov.Err() != nil  {
		return m.Album{}, rov.Err()
	}

	var album m.Album
	err := rov.Scan(&album.ID, &album.Title, &album.Artist, &album.Price)

	if err == sql.ErrNoRows {
		return m.Album{}, sql.ErrNoRows
	} else if err != nil {
		return m.Album{}, err
	}

	return album, nil
}

func AddAlbum (album m.Album) (m.Album, error) {
	conn := dbConn()
	defer conn.Close()
	_, err := conn.Exec(fmt.Sprintf("INSERT INTO albums (Title, Artist, Price) values('%s', '%s', '%f')",
		album.Title, album.Artist, album.Price))
	if err != nil {
		return m.Album{}, err
	}

	return album, err
}

func DeleteAlbum(id string) error {
	conn := dbConn()
	defer conn.Close()
	_, err := conn.Exec(fmt.Sprintf("DELETE FROM albums WHERE id = %s", id))

	return err
}

func UpdateAlbum (title string, artist string, price string, id string) error {
	if title == "" { title = "Title"} else {
		title = fmt.Sprintf("'%s'", title)
	}
	if artist == "" { artist = "Artist"} else {
		artist = fmt.Sprintf("'%s'", artist)
	}
	if price == "0" { price = "Price"} else {
		price = fmt.Sprintf("'%s'", price)
	}

	conn := dbConn()
	defer conn.Close()
	_, err := conn.Exec(fmt.Sprintf("UPDATE albums SET Title = %s, Artist = %s, Price = %s WHERE id = %s",
		title, artist, price, id))

	return err
}

func GetPasswordAndIdByName(name string) (string, string, error) {
	conn := dbConn()
	defer conn.Close()
	rov := conn.QueryRow(fmt.Sprintf("SELECT id, password FROM user WHERE Name = '%s'", name))

	if rov.Err() != nil {
		return "", "", rov.Err()
	}

	var id string
	var password string
	err := rov.Scan(&id, &password)

	if err == sql.ErrNoRows {
		return "", "", sql.ErrNoRows
	} else if err != nil {
		return "", "", err
	}

	return id, password, nil
}

func SaveRToken(id string, rToken string, ExpiresATToken time.Time) error {
	conn := dbConn()
	defer conn.Close()
	_, err := conn.Query(fmt.Sprintf("UPDATE user SET RefreshToken = '%s', ExpiresATToken = '%s' WHERE id = %s",
		rToken, ExpiresATToken.Format("2006-01-02 15:04:05"), id))

	return err
}

func CheckName(name string) error {
	conn := dbConn()
	defer conn.Close()
	rov := conn.QueryRow(fmt.Sprintf("SELECT id FROM user WHERE Name = '%s'", name))
	fmt.Println(rov.Scan())

	return nil
}