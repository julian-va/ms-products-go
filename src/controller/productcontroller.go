package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"ms-products/src/entity"
	"ms-products/src/service"
	"ms-products/src/utils"
	"net/http"
	"strconv"
)

type ProductController struct {
	productService service.ProductoService
}

func NewProductController(productService service.ProductoService) *ProductController {
	return &ProductController{
		productService: productService,
	}
}
func (controller *ProductController) GetAllProducts(w http.ResponseWriter, req *http.Request) {
	products, err := controller.productService.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if len(products) > 0 {
		utils.CreateResponse(w, products, http.StatusOK)
		return
	}
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "no hay productos")
}

func (controller *ProductController) GetByIdProduct(w http.ResponseWriter, req *http.Request) {
	id := req.PathValue("id")
	idTosearch, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	product, err := controller.productService.GetById(idTosearch)
	if product.Id == 0 {
		fmt.Print(err.Error())
		http.Error(w, err.Error(), http.StatusNoContent)
		return
	}
	utils.CreateResponse(w, product, http.StatusOK)
}

func (controller *ProductController) CreateProduct(w http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var productTocreate entity.Product
	err = json.Unmarshal(body, &productTocreate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	product, _ := controller.productService.Create(productTocreate)

	utils.CreateResponse(w, product, http.StatusCreated)
}
