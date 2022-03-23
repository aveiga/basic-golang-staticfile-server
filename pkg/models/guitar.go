package models

type Guitar struct {
	Id    int64  `json:"Id" bun:"Id"`
	Brand string `json:"brand" bun:"Brand"`
	Model string `json:"model" bun:"Model"`
}

type GuitarRepository interface {
	FindAll() (*[]Guitar, error)
	Save(guitar *Guitar) error
}

type GuitarService interface {
	CreateGuitar(guitar *Guitar) (*Guitar, error)
	GetGuitars() (*[]Guitar, error)
}
