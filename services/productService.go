package services

import (
	PGRepository "github.com/cocoasterr/net_http/infra/db/postgres/repository/databasesql"
	"github.com/cocoasterr/net_http/models"
)

type ProductService struct {
	BaseService
}

func NewProductService(productRepo PGRepository.Repository) *ProductService {
	return &ProductService{
		BaseService{
			Repo:  productRepo,
			Model: &models.Product{},
		},
	}
}
