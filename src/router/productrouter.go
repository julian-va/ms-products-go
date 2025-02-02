package router

import (
	"fmt"
	"ms-products/src/controller"
	"net/http"
)

type ProductsRouter struct {
	productController controller.ProductController
}

func NewProductsRouter(productController controller.ProductController) *ProductsRouter {
	return &ProductsRouter{productController: productController}
}

func (router *ProductsRouter) RegisterRouters(routers *http.ServeMux) *http.ServeMux {
	routers.HandleFunc(fmt.Sprintf("%s /products", http.MethodGet), router.productController.GetAllProducts)
	routers.HandleFunc(fmt.Sprintf("%s /products/{id}", http.MethodGet), router.productController.GetByIdProduct)
	routers.HandleFunc(fmt.Sprintf("%s /products", http.MethodPost), router.productController.CreateProduct)
	return routers
}
