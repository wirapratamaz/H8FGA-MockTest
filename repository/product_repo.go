package repository

import "mocktest/entity"

type ProductRepository interface {
	FindById(id string) *entity.Product
}
