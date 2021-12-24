package interfaces

import (
	m "RESTful_API_Gin/models"
	"time"
)

type IOtherRepository interface {
	GetPasswordAndIdByName(name string) (string, string, error)
	SaveRToken(id string, rToken string, ExpiresATToken time.Time) error
	CheckName(name string) (bool, error)
	AddUser(userData m.LoginData)
}
