package services

import "github.com/aveiga/basic-golang-staticfile-server/pkg/models"

func CreateGuitar(guitar models.Guitar) (*models.Guitar, error) {
	return &guitar, nil
}