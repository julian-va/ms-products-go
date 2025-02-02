package service

import "ms-products/src/entity"

type ProductoService interface {
	Create(p entity.Product) (entity.Product, error)
	GetById(id int) (entity.Product, error)
	Update(id int, p entity.Product) (entity.Product, error)
	Delete(id int) error
	GetAll() ([]entity.Product, error)
}
