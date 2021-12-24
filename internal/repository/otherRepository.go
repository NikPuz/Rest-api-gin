package repository

import (
	"RESTful_API_Gin/interfaces"
	m "RESTful_API_Gin/models"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type OtherRepository struct {
	interfaces.IDbHandler
}

func (r *OtherRepository) GetPasswordAndIdByName(name string) (string, string, error) {

	rov, err := r.Query(fmt.Sprintf("SELECT id, password FROM user WHERE Name = '%s'", name))
	if err != nil {
		return "", "", err
	}

	var id string
	var password string

	rov.Next()
	err = rov.Scan(&id, &password)
	if err != nil {
		return "", "", err
	}

	return id, password, nil
}

func (r *OtherRepository) SaveRToken(id string, rToken string, ExpiresATToken time.Time) error {
	_, err := r.Query(fmt.Sprintf("UPDATE user SET RefreshToken = '%s', ExpiresATToken = '%s' WHERE id = %s",
		rToken, ExpiresATToken.Format("2006-01-02 15:04:05"), id))

	return err
}

func (r *OtherRepository) CheckName(name string) (bool, error) {
	rov, err := r.Query(fmt.Sprintf("SELECT 1 FROM user WHERE Name = '%s'", name))
	if err != nil {
		return false, err
	}

	return rov.Next(), err
}

func (r *OtherRepository) AddUser(userData m.LoginData) {
	r.Execute(fmt.Sprintf("INSERT INTO user (Name, Password) values('%s', '%s')",
		userData.Name, userData.Password))
}
