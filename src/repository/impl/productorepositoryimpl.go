package impl

import (
	"fmt"
	"ms-products/src/entity"
	"ms-products/src/repository"
	"sort"

	"github.com/google/uuid"
)

type ProductoRepositoryMemory struct {
	productsMap map[int]entity.Product
}

// GetAll implements repository.ProductoRepository.
func (productoRepositoryMemory *ProductoRepositoryMemory) GetAll() ([]entity.Product, error) {
	var products []entity.Product
	for _, product := range productoRepositoryMemory.productsMap {
		products = append(products, product)
	}
	return products, nil
}

// Create implements repository.ProductoRepository.
func (productoRepositoryMemory *ProductoRepositoryMemory) Create(prduct entity.Product) (entity.Product, error) {
	prduct.Id = lastIdCreated(productoRepositoryMemory.productsMap)
	prduct.IdSku = uuid.New().String()
	productoRepositoryMemory.productsMap[prduct.Id] = prduct
	return prduct, nil
}

// Delete implements repository.ProductoRepository.
func (productoRepositoryMemory *ProductoRepositoryMemory) Delete(id int) error {
	_, err := productoRepositoryMemory.GetById(id)
	if err != nil {
		return err
	}
	delete(productoRepositoryMemory.productsMap, id)
	return nil
}

// GetById implements repository.ProductoRepository.
func (productoRepositoryMemory *ProductoRepositoryMemory) GetById(id int) (entity.Product, error) {
	if product, exists := productoRepositoryMemory.productsMap[id]; exists {
		return product, nil
	}
	return entity.Product{}, fmt.Errorf("producto con ID %d no encontrado", id)
}

// Update implements repository.ProductoRepository.
func (productoRepositoryMemory *ProductoRepositoryMemory) Update(id int, product entity.Product) (entity.Product, error) {
	productToUpdate, err := productoRepositoryMemory.GetById(id)
	if err != nil {
		return entity.Product{}, err
	}
	productToUpdate.Descripcion = validateValueStr(product.Descripcion, productToUpdate.Descripcion)
	productToUpdate.Nombre = validateValueStr(product.Nombre, productToUpdate.Nombre)
	productToUpdate.Precio = validateValueint(product.Precio, productToUpdate.Precio)
	productToUpdate.Stock = validateValueint(product.Stock, productToUpdate.Stock)
	productoRepositoryMemory.productsMap[id] = productToUpdate
	return productToUpdate, nil
}

func New(products map[int]entity.Product) repository.ProductoRepository {
	return &ProductoRepositoryMemory{
		productsMap: initialData(products),
	}
}

func initialData(products map[int]entity.Product) map[int]entity.Product {
	product1 := entity.Product{Id: 1, Nombre: "test product_1", Descripcion: "test Descripcion_1", Precio: 15, Stock: 20, IdSku: uuid.New().String()}
	product2 := entity.Product{Id: 2, Nombre: "test product_2", Descripcion: "test Descripcion_2", Precio: 150, Stock: 5, IdSku: uuid.New().String()}
	product3 := entity.Product{Id: 3, Nombre: "test product_3", Descripcion: "test Descripcion_3", Precio: 1000, Stock: 200, IdSku: uuid.New().String()}
	product4 := entity.Product{Id: 4, Nombre: "test product_4", Descripcion: "test Descripcion_4", Precio: 5, Stock: 1000, IdSku: uuid.New().String()}
	product5 := entity.Product{Id: 5, Nombre: "test product_5", Descripcion: "test Descripcion_5", Precio: 3000, Stock: 20, IdSku: uuid.New().String()}
	products[product1.Id] = product1
	products[product2.Id] = product2
	products[product3.Id] = product3
	products[product4.Id] = product4
	products[product5.Id] = product5
	return products
}

func lastIdCreated(products map[int]entity.Product) int {
	keys := make([]int, 0, len(products))

	for k := range products {
		keys = append(keys, k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))
	id := keys[0]
	if id != 0 {
		id++
		return id
	}
	id++
	return id
}

func validateValueStr(newValue string, oldValue string) string {
	if newValue == "" {
		return oldValue
	} else {
		return newValue
	}
}

func validateValueint(newValue int64, oldValue int64) int64 {
	if newValue == 0 {
		return oldValue
	} else {
		return newValue
	}
}
