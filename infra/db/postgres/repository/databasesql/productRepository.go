package PGRepository

import (
	"database/sql"

	"github.com/cocoasterr/net_http/models"
)

type ProductRepository struct {
	Repository
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{
		Repository{
			Db:    db,
			Model: &models.Product{},
		},
	}
}
