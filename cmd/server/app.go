package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	buyerRepository "github.com/luisantonisu/wave15-grupo4/internal/repository/buyer"
	employeeRepository "github.com/luisantonisu/wave15-grupo4/internal/repository/employee"
	productRepository "github.com/luisantonisu/wave15-grupo4/internal/repository/product"
	sectionRepository "github.com/luisantonisu/wave15-grupo4/internal/repository/section"
	sellerRepository "github.com/luisantonisu/wave15-grupo4/internal/repository/seller"
	warehouseRepository "github.com/luisantonisu/wave15-grupo4/internal/repository/warehouse"

	buyerService "github.com/luisantonisu/wave15-grupo4/internal/service/buyer"
	employeeService "github.com/luisantonisu/wave15-grupo4/internal/service/employee"
	productService "github.com/luisantonisu/wave15-grupo4/internal/service/product"
	sectionService "github.com/luisantonisu/wave15-grupo4/internal/service/section"
	sellerService "github.com/luisantonisu/wave15-grupo4/internal/service/seller"
	warehouseService "github.com/luisantonisu/wave15-grupo4/internal/service/warehouse"

	buyerHandler "github.com/luisantonisu/wave15-grupo4/internal/handler/buyer"
	employeeHandler "github.com/luisantonisu/wave15-grupo4/internal/handler/employee"
	productHandler "github.com/luisantonisu/wave15-grupo4/internal/handler/product"
	sectionHandler "github.com/luisantonisu/wave15-grupo4/internal/handler/section"
	sellerHandler "github.com/luisantonisu/wave15-grupo4/internal/handler/seller"
	warehouseHandler "github.com/luisantonisu/wave15-grupo4/internal/handler/warehouse"
	"github.com/luisantonisu/wave15-grupo4/internal/loader"
)

type ConfigServerChi struct {
	ServerAddress  string
	LoaderFilePath string
}

func NewServerChi(cfg *ConfigServerChi) *ServerChi {
	defaultConfig := &ConfigServerChi{
		ServerAddress: ":8080",
	}
	if cfg != nil {
		if cfg.ServerAddress != "" {
			defaultConfig.ServerAddress = cfg.ServerAddress
		}
		if cfg.LoaderFilePath != "" {
			defaultConfig.LoaderFilePath = cfg.LoaderFilePath
		}
	}

	return &ServerChi{
		serverAddress:  defaultConfig.ServerAddress,
		loaderFilePath: defaultConfig.LoaderFilePath,
	}
}

type ServerChi struct {
	serverAddress  string
	loaderFilePath string
}

func (a *ServerChi) Run() (err error) {
	//TODO
	db, err := loader.Load()
	if err != nil {
		return
	}

	// - repository
	buyerRp := buyerRepository.NewBuyerRepository(db.Buyers)
	employeeRp := employeeRepository.NewEmployeeRepository(db.Employees)
	productRp := productRepository.NewProductRepository(db.Products)
	sectionRp := sectionRepository.NewSectionRepository(db.Sections)
	sellerRp := sellerRepository.NewSellerRepository(db.Sellers)
	warehouseRp := warehouseRepository.NewWarehouseRepository(db.Warehouses)

	// - service
	buyerSv := buyerService.NewBuyerService(buyerRp)
	employeeSv := employeeService.NewEmployeeService(employeeRp)
	productSv := productService.NewProductService(productRp)
	sectionSv := sectionService.NewSectionService(sectionRp)
	sellerSv := sellerService.NewSellerService(sellerRp)
	warehouseSv := warehouseService.NewWarehouseService(warehouseRp)

	// - handler
	buyerHd := buyerHandler.NewBuyerHandler(buyerSv)                 // buyerHd
	employeeHd := employeeHandler.NewEmployeeHandler(employeeSv)     // employeeHd
	productHd := productHandler.NewProductHandler(productSv)         // productHd
	sectionHd := sectionHandler.NewSectionHandler(sectionSv)         // sectionHd
	sellerHd := sellerHandler.NewSellerHandler(sellerSv)             // sellerHd
	warehouseHd := warehouseHandler.NewWarehouseHandler(warehouseSv) // warehouseHd

	// router
	rt := chi.NewRouter()

	// - middlewares
	rt.Use(middleware.Logger)
	rt.Use(middleware.Recoverer)

	// - endpoints
	rt.Route("/api/v1", func(rt chi.Router) {
		rt.Route("/buyers", func(rt chi.Router) {
			// - GET /api/v1/buyers
			rt.Get("/", buyerHd.GetAll())
			rt.Get("/{id}", buyerHd.GetByID())
			rt.Post("/", buyerHd.Create())
			rt.Delete("/{id}", buyerHd.Delete())

		})
		rt.Route("/employees", func(rt chi.Router) {
			// - GET /api/v1/employees
			rt.Get("/", employeeHd.GetAll())
			rt.Get("/{id}", employeeHd.GetByID())
			// - POST /api/v1/employees
			rt.Post("/", employeeHd.Create())
			// - PUT /api/v1/employees/{id}
			rt.Patch("/{id}", employeeHd.Update())
			// - DELETE /api/v1/employees/{id}
			rt.Delete("/{id}", employeeHd.Delete())
		})
		rt.Route("/products", func(rt chi.Router) {
			// - GET /api/v1/products /
			rt.Get("/", productHd.GetAll())
			rt.Get("/{id}", productHd.GetByID())
			// - POST /api/v1/products /
			rt.Post("/", productHd.Create())
			// - DELETE /api/v1/products /
			rt.Delete("/{id}", productHd.Delete())
			// - PATCH /api/v1/products /
			rt.Patch("/{id}", productHd.Update())
		})
		rt.Route("/sections", func(rt chi.Router) {
			// - GET /api/v1/sections
			rt.Get("/", sectionHd.GetAll())
			rt.Get("/{id}", sectionHd.GetByID())
			// - POST /api/v1/products /
			rt.Post("/", sectionHd.Create())
			// - DELETE /api/v1/employees/{id}
			rt.Delete("/{id}", sectionHd.Delete())
		})
		rt.Route("/sellers", func(rt chi.Router) {
			// - GET /api/v1/sellers
			rt.Get("/", sellerHd.GetAll())
			// -GET /api/v1/sellers/{id}
			rt.Get("/{id}", sellerHd.GetByID())
			// -POST /api/v1/sellers
			rt.Post("/", sellerHd.Create())
		})
		rt.Route("/warehouses", func(rt chi.Router) {
			// - GET /api/v1/warehouses
			rt.Get("/", warehouseHd.GetAll())
			// - GET /api/v1/warehouses/{id}
			rt.Get("/{id}", warehouseHd.GetByID())
			// - POST /api/v1/warehouses
			rt.Post("/", warehouseHd.Create())
			// - PATCH /api/v1/warehouses/{id}
			rt.Patch("/{id}", warehouseHd.Update())
		})
	})

	// run server
	err = http.ListenAndServe(a.serverAddress, rt)
	return
}
