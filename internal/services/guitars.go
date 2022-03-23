package services

import (
	"log"

	"github.com/aveiga/basic-golang-staticfile-server/pkg/models"
)

type GuitarService struct {
	guitarRepo models.GuitarRepository
}

func NewGuitarService(guitarRepo models.GuitarRepository) *GuitarService {
	return &GuitarService{
		guitarRepo: guitarRepo,
	}
}

func (h *GuitarService) CreateGuitar(guitar *models.Guitar) (*models.Guitar, error) {
	err := h.guitarRepo.Save(guitar)
	return guitar, err

}

func (h *GuitarService) GetGuitars() (*[]models.Guitar, error) {
	guitars, err := h.guitarRepo.FindAll()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return guitars, nil
}
