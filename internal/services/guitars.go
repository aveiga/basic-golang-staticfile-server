package services

import (
	"github.com/aveiga/basic-golang-staticfile-server/pkg/models"
	"github.com/aveiga/basic-golang-staticfile-server/pkg/utils/customamqp"
	"github.com/aveiga/basic-golang-staticfile-server/pkg/utils/customlogger"
	"go.uber.org/zap"
)

type GuitarService struct {
	guitarRepo models.GuitarRepository
	messaging  *customamqp.MessagingClient
	logger     *zap.SugaredLogger
}

func NewGuitarService(guitarRepo models.GuitarRepository, messaging *customamqp.MessagingClient) *GuitarService {
	return &GuitarService{
		guitarRepo: guitarRepo,
		messaging:  messaging,
		logger:     customlogger.NewCustomLogger(),
	}
}

func (s *GuitarService) CreateGuitar(guitar *models.Guitar) (*models.Guitar, error) {
	err := s.guitarRepo.Save(guitar)
	// serializedGuitar, _ := serialization.Serialize(guitar)
	// s.messaging.Publish(serializedGuitar, "guitars", "topic")
	// for i := 0; i < 10000; i++ {
	// 	fmt.Printf(`%d`, i)
	// }
	return guitar, err
}

func (s *GuitarService) GetGuitars() (*[]models.Guitar, error) {
	guitars, err := s.guitarRepo.FindAll()
	if err != nil {
		s.logger.Fatal(err)
		return nil, err
	}

	return guitars, nil
}
