package interfaces

import m "RESTful_API_Gin/models"

type IAuthService interface {
	LoginService(loginData m.LoginData) (string, string, error)
	SigninService(signinData m.LoginData) error
}
