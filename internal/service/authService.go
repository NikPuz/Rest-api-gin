package service

import (
	"RESTful_API_Gin/interfaces"
	m "RESTful_API_Gin/models"
	"RESTful_API_Gin/pkg"
	"errors"
	"fmt"
	"time"
)

type AuthService struct {
	interfaces.IOtherRepository
}

func (s AuthService) LoginService(loginData m.LoginData) (string, string, error) {
	// ПРОВЕРКА ДАННЫХ

	id, _, err := s.GetPasswordAndIdByName(loginData.Name)

	if err != nil {
		return "", "", err
	}
	// СОЗДАНИЕ ТОКЕНОВ

	jwToken, err := pkg.CreateJWToken(id)
	if err != nil {
		return "", "", err
	}

	rToken, err := pkg.CreateRToken(id)
	if err != nil {
		return "", "", err
	}
	// СОХРАНЕНИЕ ТОКЕНА

	ExpiresATToken := time.Now().Add(15 * time.Minute)
	if err := s.SaveRToken(id, rToken, ExpiresATToken); err != nil {
		return "", "", err
	}

	return jwToken, rToken, nil
}

func (s AuthService) SigninService(signinData m.LoginData) error {
	// ПРОВЕРКА ИМЕНИ
	isTakenName, err := s.CheckName(signinData.Name)
	fmt.Println(isTakenName)
	if isTakenName {
		return errors.New("this name is taken")
	} else if err != nil {
		return err
	}

	// ХЕШИРОВАНИЕ ПАРОЛЯ

	if err := pkg.HashPassword(&signinData.Password); err != nil {
		return err
	}
	// СОЗДАНИЕ ПОЛЬЗОВАТЕЛЯ

	s.AddUser(signinData)
	return nil
}
