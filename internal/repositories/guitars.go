package repositories

import (
	"context"
	"log"

	"github.com/aveiga/basic-golang-staticfile-server/pkg/models"
	"github.com/uptrace/bun"
)

type GuitarRepo struct {
	db  *bun.DB
	ctx context.Context
}

func NewGuitarRepo(db *bun.DB, ctx context.Context) *GuitarRepo {
	db.NewCreateTable().Model((*models.Guitar)(nil)).IfNotExists().Exec(ctx)
	return &GuitarRepo{
		db:  db,
		ctx: ctx,
	}
}

func (r *GuitarRepo) FindAll() (*[]models.Guitar, error) {
	guitars := make([]models.Guitar, 0)
	err := r.db.NewSelect().
		Model(&guitars).
		Scan(r.ctx)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return &guitars, nil
}

func (r *GuitarRepo) Save(guitar *models.Guitar) error {
	_, err := r.db.NewInsert().
		Model(guitar).
		Exec(r.ctx)

	return err
}
