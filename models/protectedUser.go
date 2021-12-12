package models

type ProtectedUser struct {
	ID uint16 `json:"id"`
	Password string `json:"password"`
	JWT string `json:"jwt"`
}
