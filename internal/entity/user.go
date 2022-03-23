package entity

type IUserRepository interface {
	GetPasswordAndIdByName(name string) (string, string, error)
	SaveRToken(id string, rToken string) error
	CheckRToken(token string) (bool, error)
	CheckName(name string) (bool, error)
	Add(userData LoginData)
}

type IAuthService interface {
	SigninService(signinData LoginData) error
	LoginService(loginData LoginData) (string, string, error)
	RefreshLoginService(refreshToken string) (string, string, error)
}

type User struct {
	ID uint16 `json:"id"`
	Name string `json:"name"`
}

type LoginData struct {
	Name string `json:"name"`
	Password string `json:"password"`
}

type RJWToken struct {
	RJWToken string `json:"refresh_token"`
}

