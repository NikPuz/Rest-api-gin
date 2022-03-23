package entity

type IAlbumRepository interface {
	GetById(id string) (*Album, error)
	GetPage(page int, records int) ([]Album, error)
	Add(album Album)
	Delete(id string)
	Update(title string, artist string, price string, id string)
}

type IAlbumService interface {
	GetById(string) (*Album, error)
	GetPage(page string) ([]Album, error)
	Add(album Album)
	Delete(id string)
	Update(title string, artist string, price string, id string)
}

type Album struct {
	ID uint16 `json:"id"`
	Title string `json:"title"`
	Artist string `json:"artist"`
	Price float64 `json:"price"`
}
