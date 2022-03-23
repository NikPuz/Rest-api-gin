package repository

import (
	"database/sql"
	"fmt"
	"restful-api-gin/internal/entity"

	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
)

type userRepository struct {
	db *sql.DB
	logger *zap.Logger
}

func NewUserRepository(db *sql.DB, logger *zap.Logger) entity.IUserRepository {
	userRepository := new(userRepository)
	userRepository.db = db
	userRepository.logger = logger
	return userRepository
}

func (r *userRepository) GetPasswordAndIdByName(name string) (string, string, error) {

	rows, err := r.db.Query(fmt.Sprintf("SELECT id, password FROM user WHERE Name = '%s'", name))
	if err != nil {
		return "", "", err
	}
	defer rows.Close()

	var id string
	var password string

	rows.Next()
	err = rows.Scan(&id, &password)
	if err != nil {
		return "", "", err
	}

	return id, password, err
}

func (r *userRepository) SaveRToken(id string, token string) error {
	_, err := r.db.Query(fmt.Sprintf("UPDATE user SET RefreshToken = '%s' WHERE id = %s",
		token, id))

	return err
}

func (r *userRepository) CheckRToken(token string) (bool, error) {
	rows, err := r.db.Query(fmt.Sprintf("SELECT 1 FROM user WHERE RefreshToken = '%s'", token))
	if err != nil {
		return false, err
	}
	defer rows.Close()

	return rows.Next(), err
}

func (r *userRepository) CheckName(name string) (bool, error) {
	rows, err := r.db.Query(fmt.Sprintf("SELECT 1 FROM user WHERE Name = '%s'", name))
	if err != nil {
		return false, err
	}
	defer rows.Close()

	return rows.Next(), err
}

func (r *userRepository) Add(userData entity.LoginData) {
	r.db.Exec(fmt.Sprintf("INSERT INTO user (Name, Password) values('%s', '%s')",
		userData.Name, userData.Password))
}
