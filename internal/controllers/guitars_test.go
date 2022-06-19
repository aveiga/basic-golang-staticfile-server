package controllers

import (
	"testing"

	"github.com/aveiga/basic-golang-staticfile-server/pkg/models"
)

type GuitarRepositoryMock struct{}

func (grm *GuitarRepositoryMock) FindAll() (*[]models.Guitar, error) {
	guitars := make([]models.Guitar, 1)

	guitars[0] = models.Guitar{
		Id:    1,
		Brand: "Fender",
		Model: "Strat",
	}

	return &guitars, nil
}

func (grm *GuitarRepositoryMock) Save(guitar *models.Guitar) error {
	return nil
}

type GuitarServiceMock struct{}

func (gsm *GuitarServiceMock) CreateGuitar(guitar *models.Guitar) (*models.Guitar, error) {
	return &models.Guitar{
		Id:    1,
		Brand: "Fender",
		Model: "Strat",
	}, nil
}

func (gsm *GuitarServiceMock) GetGuitars() (*[]models.Guitar, error) {
	guitars := make([]models.Guitar, 1)

	guitars[0] = models.Guitar{
		Id:    1,
		Brand: "Fender",
		Model: "Strat",
	}

	return &guitars, nil
}

// type GuitarService interface {
// 	CreateGuitar(guitar *Guitar) (*Guitar, error)
// 	GetGuitars() (*[]Guitar, error)
// }

func TestNewGuitarController(t *testing.T) {
	gsm := &GuitarServiceMock{}

	NewGuitarController(gsm)

	// if(controller.guitarService != nil && controller.logger != nil) {
	// 	t.
	// }
}
