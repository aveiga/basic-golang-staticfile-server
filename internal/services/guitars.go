package services

import (
	"fmt"
	"log"

	"github.com/aveiga/basic-golang-staticfile-server/pkg/models"
	"github.com/aveiga/basic-golang-staticfile-server/pkg/utils/customamqp"
)

type GuitarService struct {
	guitarRepo models.GuitarRepository
	messaging  *customamqp.MessagingClient
}

func NewGuitarService(guitarRepo models.GuitarRepository, messaging *customamqp.MessagingClient) *GuitarService {
	return &GuitarService{
		guitarRepo: guitarRepo,
		messaging:  messaging,
	}
}

func (s *GuitarService) CreateGuitar(guitar *models.Guitar) (*models.Guitar, error) {
	err := s.guitarRepo.Save(guitar)
	// serializedGuitar, _ := serialization.Serialize(guitar)
	// s.messaging.Publish(serializedGuitar, "guitars", "topic")
	// time.Sleep(5 * time.Second)
	for i := 0; i < 10000; i++ {
		fmt.Printf(`%d`, i)
	}
	return guitar, err
}

func (s *GuitarService) GetGuitars() (*[]models.Guitar, error) {
	guitars, err := s.guitarRepo.FindAll()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return guitars, nil
}
