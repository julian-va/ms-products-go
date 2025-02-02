package impl

import (
	"ms-products/src/entity"
	"ms-products/src/repository"
	"ms-products/src/service"
)

type ProductService struct {
	repository repository.ProductoRepository
}

// Create implements service.ProductoService.
func (service *ProductService) Create(product entity.Product) (entity.Product, error) {
	return service.repository.Create(product)
}

// Delete implements service.ProductoService.
func (service *ProductService) Delete(id int) error {
	return service.repository.Delete(id)
}

// GetAll implements service.ProductoService.
func (service *ProductService) GetAll() ([]entity.Product, error) {
	return service.repository.GetAll()
}

// GetById implements service.ProductoService.
func (service *ProductService) GetById(id int) (entity.Product, error) {
	return service.repository.GetById(id)
}

// Update implements service.ProductoService.
func (service *ProductService) Update(id int, product entity.Product) (entity.Product, error) {
	return service.repository.Update(id, product)
}

func New(repository repository.ProductoRepository) service.ProductoService {
	return &ProductService{
		repository: repository,
	}
}
