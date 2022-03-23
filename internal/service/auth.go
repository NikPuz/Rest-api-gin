package service

import (
	"errors"
	"fmt"
	"restful-api-gin/internal/entity"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	"go.uber.org/zap"
)

type authService struct {
	userRepo entity.IUserRepository
	logger *zap.Logger
}

func NewAuthService(userRepo entity.IUserRepository, logger *zap.Logger) entity.IAuthService {
	authService := new(authService)
	authService.userRepo = userRepo
	authService.logger = logger
	return authService
}

func (s authService) SigninService(signinData entity.LoginData) error {
	// ПРОВЕРКА ИМЕНИ

	isTakenName, err := s.userRepo.CheckName(signinData.Name)
	if isTakenName {
		return errors.New("name is taken")
	} else if err != nil {
		return err
	}
	// ХЕШИРОВАНИЕ ПАРОЛЯ

	if err := hashPassword(&signinData.Password); err != nil {
		return err
	}
	// СОЗДАНИЕ ПОЛЬЗОВАТЕЛЯ

	s.userRepo.Add(signinData)
	return nil
}

func (s authService) LoginService(loginData entity.LoginData) (string, string, error) {
	// ПРОВЕРКА ДАННЫХ

	id, passwordHash, err := s.userRepo.GetPasswordAndIdByName(loginData.Name)
	if err != nil && err.Error() == "sql: Rows are closed" {
		return "", "", errors.New("incorrect password or name")
	} else if err != nil {
		return "", "", err
	}

	if !checkPasswordHash(loginData.Password, passwordHash) {
		return "", "", errors.New("incorrect password or name")
	}
	// СОЗДАНИЕ ТОКЕНОВ

	jwToken, err := createJWToken(id)
	if err != nil {
		return "", "", err
	}

	rJWToken, err := createRJWToken(id)
	if err != nil {
		return "", "", err
	}
	// СОХРАНЕНИЕ ТОКЕНА

	if err := s.userRepo.SaveRToken(id, rJWToken); err != nil {
		return "", "", err
	}

	return jwToken, rJWToken, err
}

func (s authService) RefreshLoginService(refreshToken string) (string, string, error) {
	// ПРОВЕРКА ДАННЫХ
	if ok, err := s.userRepo.CheckRToken(refreshToken); !ok {
		return "", "", errors.New("token is empty")
	} else if err != nil {
		return "", "", err
	}

	id, err := parseRJWToken(refreshToken)
	if err != nil {
		return "", "", errors.New("token is empty")
	}
	// СОЗДАНИЕ ТОКЕНОВ

	jwToken, err := createJWToken(id)
	if err != nil {
		return "", "", err
	}

	rJWToken, err := createRJWToken(id)
	if err != nil {
		return "", "", err
	}

	if err := s.userRepo.SaveRToken(id, rJWToken); err != nil {
		return "", "", err
	}

	return jwToken, rJWToken, err
}


func hashPassword(password *string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(*password), 10)
	*password = string(bytes)
	return err
}

func checkPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}

func createJWToken(id string) (string, error) {
	atClaim := jwt.MapClaims{}
	atClaim["authorized"] = true
	atClaim["user_id"] = id
	atClaim["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaim)
	token, err := at.SignedString([]byte(viper.GetString("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, err
}

func createRJWToken(id string) (string, error) {
	atClaim := jwt.MapClaims{}
	atClaim["authorized"] = true
	atClaim["user_id"] = id
	atClaim["exp"] = time.Now().Add(time.Hour * 720).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaim)
	token, err := at.SignedString([]byte(viper.GetString("REFRESH_SECRET")))
	if err != nil {
		return "", err
	}
	return token, err
}

func parseRJWToken(accessToken string) (string, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (i interface{}, err error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(viper.GetString("REFRESH_SECRET")), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("error get user claims from token")
	}

	return claims["user_id"].(string), nil
}

// func createRToken(id string) (string, error) {
// 	b := make([]byte, 32)

// 	s := rand.NewSource(time.Now().Unix())
// 	r := rand.New(s)

// 	_, err := r.Read(b)

// 	return fmt.Sprintf("%x%s", b, id), err
// }
