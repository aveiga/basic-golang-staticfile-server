package repositories

import (
	"context"
	"log"

	"github.com/aveiga/basic-golang-staticfile-server/pkg/models"
	"github.com/uptrace/bun"
)

type GuitarRepo struct {
	db *bun.DB
}

func NewGuitarRepo(db *bun.DB) *GuitarRepo {
	return &GuitarRepo{
		db: db,
	}
}

func (r *GuitarRepo) FindAll() (*[]models.Guitar, error) {
	ctx := context.Background()
	// db := bun.NewDB(r.db, pgdialect.New())

	guitars := make([]models.Guitar, 0)
	err := r.db.NewSelect().
		Model(&guitars).
		Scan(ctx)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return &guitars, nil
}

func (r *GuitarRepo) Save(guitar *models.Guitar) error {
	return nil
}
