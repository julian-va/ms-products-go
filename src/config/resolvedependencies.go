package config

import (
	"ms-products/src/controller"
	"ms-products/src/entity"
	"ms-products/src/repository"
	repositoryProduct "ms-products/src/repository/impl"
	"ms-products/src/router"
	"ms-products/src/service"
	productService "ms-products/src/service/impl"
	"ms-products/src/utils"
	"os"

	"gopkg.in/yaml.v3"
)

type ResolveDependencies struct {
	ApplicationConfigs ApplicationConfigs
	productRepository  repository.ProductoRepository
	productService     service.ProductoService
	productController  controller.ProductController
	ProductsRouter     router.ProductsRouter
	pingController     controller.PingController
	PingRouter         router.PingRouter
}

func NewResolveDependencies() *ResolveDependencies {
	dependencies := &ResolveDependencies{}
	dependencies.ApplicationConfigs = createApplicationConfigs()
	dependencies.productRepository = createProductRepository()
	dependencies.productService = createproductService(dependencies.productRepository)
	dependencies.productController = createproductController(dependencies.productService)
	dependencies.ProductsRouter = createProductsRouter(dependencies.productController)
	dependencies.pingController = createPingController()
	dependencies.PingRouter = createPingRouter(dependencies.pingController)
	return dependencies
}

func createApplicationConfigs() ApplicationConfigs {
	var applicationConfigs ApplicationConfigs

	yamlFile, err := os.ReadFile(utils.YML_APPLICATION_PATH)
	utils.PanicError(err)

	err = yaml.Unmarshal(yamlFile, &applicationConfigs)
	utils.PanicError(err)
	return applicationConfigs
}

func createProductRepository() repository.ProductoRepository {
	return repositoryProduct.New(make(map[int]entity.Product))
}

func createproductService(productRepository repository.ProductoRepository) service.ProductoService {
	return productService.New(productRepository)
}

func createproductController(productService service.ProductoService) controller.ProductController {
	return *controller.NewProductController(productService)
}

func createProductsRouter(productController controller.ProductController) router.ProductsRouter {
	return *router.NewProductsRouter(productController)
}

func createPingController() controller.PingController {
	return *controller.NewPingController()
}

func createPingRouter(pingController controller.PingController) router.PingRouter {
	return *router.NewPingRouter(pingController)
}
