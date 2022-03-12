package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/aveiga/basic-golang-staticfile-server/pkg/models"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
)

type GuitarRepo struct {
	db *sql.DB
}

func NewGuitarRepo(db *sql.DB) *GuitarRepo {
	return &GuitarRepo{
		db: db,
	}
}

func (r *GuitarRepo) FindAll() (*[]models.Guitar, error) {
	ctx := context.Background()
	db := bun.NewDB(r.db, pgdialect.New())

	guitars := make([]models.Guitar, 0)
	err := db.NewSelect().
		Model(&guitars).
		Scan(ctx)
	fmt.Printf("from repository: all guitars: %v\n\n", guitars)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return &guitars, nil
}

func (r *GuitarRepo) Save(guitar *models.Guitar) error {
	return nil
}
