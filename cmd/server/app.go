package server

import (
	"net/http"
	"sync"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/luisantonisu/wave15-grupo4/infrastructure/db"
	"github.com/luisantonisu/wave15-grupo4/internal/config"
<<<<<<< HEAD
	buyerRepository "github.com/luisantonisu/wave15-grupo4/internal/repository/buyer"
	employeeRepository "github.com/luisantonisu/wave15-grupo4/internal/repository/employee"
	inboundOrderRepository "github.com/luisantonisu/wave15-grupo4/internal/repository/inbound_order"
	productRepository "github.com/luisantonisu/wave15-grupo4/internal/repository/product"
	productBatchRepository "github.com/luisantonisu/wave15-grupo4/internal/repository/product_batch"
	productRecordRepository "github.com/luisantonisu/wave15-grupo4/internal/repository/product_record"
	purchaseOrderRepository "github.com/luisantonisu/wave15-grupo4/internal/repository/purchase_order"
	sectionRepository "github.com/luisantonisu/wave15-grupo4/internal/repository/section"
	sellerRepository "github.com/luisantonisu/wave15-grupo4/internal/repository/seller"
	warehouseRepository "github.com/luisantonisu/wave15-grupo4/internal/repository/warehouse"

	buyerService "github.com/luisantonisu/wave15-grupo4/internal/service/buyer"
	employeeService "github.com/luisantonisu/wave15-grupo4/internal/service/employee"
	inboundOrderService "github.com/luisantonisu/wave15-grupo4/internal/service/inbound_order"
	productService "github.com/luisantonisu/wave15-grupo4/internal/service/product"
	productBatchService "github.com/luisantonisu/wave15-grupo4/internal/service/product_batch"
	productRecordService "github.com/luisantonisu/wave15-grupo4/internal/service/product_record"
	purchaseOrderService "github.com/luisantonisu/wave15-grupo4/internal/service/purchase_order"
	sectionService "github.com/luisantonisu/wave15-grupo4/internal/service/section"
	sellerService "github.com/luisantonisu/wave15-grupo4/internal/service/seller"
	warehouseService "github.com/luisantonisu/wave15-grupo4/internal/service/warehouse"

	buyerHandler "github.com/luisantonisu/wave15-grupo4/internal/handler/buyer"
	employeeHandler "github.com/luisantonisu/wave15-grupo4/internal/handler/employee"
	inboundOrderHandler "github.com/luisantonisu/wave15-grupo4/internal/handler/inbound_order"
	productHandler "github.com/luisantonisu/wave15-grupo4/internal/handler/product"
	productBatchHandler "github.com/luisantonisu/wave15-grupo4/internal/handler/product_batch"
	productRecordHandler "github.com/luisantonisu/wave15-grupo4/internal/handler/product_record"
	purchaseOrderHandler "github.com/luisantonisu/wave15-grupo4/internal/handler/purchase_order"
	sectionHandler "github.com/luisantonisu/wave15-grupo4/internal/handler/section"
	sellerHandler "github.com/luisantonisu/wave15-grupo4/internal/handler/seller"
	warehouseHandler "github.com/luisantonisu/wave15-grupo4/internal/handler/warehouse"
=======
>>>>>>> develop
	"github.com/luisantonisu/wave15-grupo4/internal/loader"
)

func NewServerChi(cfg *config.Config) *ServerChi {

	defaultConfig := &config.Config{
		ServerAddress: ":8080",
	}
	if cfg != nil {
		if cfg.ServerAddress != "" {
			defaultConfig.ServerAddress = cfg.ServerAddress
		}
		if cfg.LoaderFilePath != "" {
			defaultConfig.LoaderFilePath = cfg.LoaderFilePath
		}
		if cfg.DBHost != "" {
			defaultConfig.DBHost = cfg.DBHost
		}
		if cfg.DBPort != "" {
			defaultConfig.DBPort = cfg.DBPort
		}
		if cfg.DBUser != "" {
			defaultConfig.DBUser = cfg.DBUser
		}
		if cfg.DBPassword != "" {
			defaultConfig.DBPassword = cfg.DBPassword
		}
		if cfg.DBName != "" {
			defaultConfig.DBName = cfg.DBName
		}
	}

	return &ServerChi{
		serverAddress: defaultConfig.ServerAddress,
		config:        defaultConfig,
	}
}

type ServerChi struct {
	serverAddress string
	config        *config.Config
}

func (a *ServerChi) Run(cfg config.Config) (err error) {
	database := db.ConnectDB(&cfg)
	defer database.Close()

	// Optional, if deploy, delete this
	var o sync.Once
	o.Do(func() {
		err = loader.Load(database)
		if err != nil {
			return
		}
	},
	)

<<<<<<< HEAD
	// - repository
	buyerRp := buyerRepository.NewBuyerRepository(database)
	purchaseOrderRp := purchaseOrderRepository.NewPurchaseOrderRepository(database)
	employeeRp := employeeRepository.NewEmployeeRepository(database)
	inboundOrderRp := inboundOrderRepository.NewInboundOrderRepository(database)
	productRp := productRepository.NewProductRepository(database)
	productRecordRp := productRecordRepository.NewProductRecordRepository(database)
	sectionRp := sectionRepository.NewSectionRepository(database)
	productBatchRp := productBatchRepository.NewProductBatchRepository(database)
	sellerRp := sellerRepository.NewSellerRepository(database)
	warehouseRp := warehouseRepository.NewWarehouseRepository(db.Warehouses)

	// - service
	buyerSv := buyerService.NewBuyerService(buyerRp)
	PurchaseOrderSv := purchaseOrderService.NewPurchaseOrderService(purchaseOrderRp)
	employeeSv := employeeService.NewEmployeeService(employeeRp)
	inboundOrderSv := inboundOrderService.NewInboundOrderService(inboundOrderRp)
	productSv := productService.NewProductService(productRp)
	productRecordSv := productRecordService.NewProductRecordService(productRecordRp)
	sectionSv := sectionService.NewSectionService(sectionRp)
	productBatchSv := productBatchService.NewProductBatchService(productBatchRp)
	sellerSv := sellerService.NewSellerService(sellerRp)
	warehouseSv := warehouseService.NewWarehouseService(warehouseRp)

	// - handler
	buyerHd := buyerHandler.NewBuyerHandler(buyerSv)                                 // buyerHd
	purchaseOrderHd := purchaseOrderHandler.NewPurchaseOrderHandler(PurchaseOrderSv) // purchaseOrderHd
	employeeHd := employeeHandler.NewEmployeeHandler(employeeSv)                     // employeeHd
	inboundOrderHd := inboundOrderHandler.NewInboundOrderHandler(inboundOrderSv)     // inboundOrderHd
	productHd := productHandler.NewProductHandler(productSv)                         // productHd
	productRecordHd := productRecordHandler.NewProductRecordHandler(productRecordSv) // productRecordHd
	sectionHd := sectionHandler.NewSectionHandler(sectionSv)                         // sectionHd
	productBatchHd := productBatchHandler.NewProductBatchHandler(productBatchSv)     // productBatchHd
	sellerHd := sellerHandler.NewSellerHandler(sellerSv)                             // sellerHd
	warehouseHd := warehouseHandler.NewWarehouseHandler(warehouseSv)                 // warehouseHd
=======
	handlers := GetHandlers(database)
>>>>>>> develop

	// router
	rt := chi.NewRouter()

	// - middlewares
	rt.Use(middleware.Logger)
	rt.Use(middleware.Recoverer)

	// - endpoints
	rt.Route("/api/v1", func(rt chi.Router) {
		rt.Route("/buyers", func(rt chi.Router) {
			// - GET /api/v1/buyers
			rt.Get("/", handlers.BuyerHandler.GetAll())
			rt.Get("/{id}", handlers.BuyerHandler.GetByID())
			// - POST /api/v1/buyers
			rt.Post("/", handlers.BuyerHandler.Create())
			// - PUT /api/v1/buyers
			rt.Patch("/{id}", handlers.BuyerHandler.Update())
			// - DELETE /api/v1/buyers
			rt.Delete("/{id}", handlers.BuyerHandler.Delete())
			// - GET /api/v1/buyers/reportPurchaseOrders
			rt.Get("/reportPurchaseOrders", handlers.BuyerHandler.Report())
		})
		rt.Route("/purchaseOrders", func(rt chi.Router) {
			// - POST /api/v1/purchaseOrders
			rt.Post("/", handlers.PurchaseOrderHandler.Create())
		})
		rt.Route("/employees", func(rt chi.Router) {
			// - GET /api/v1/employees
			rt.Get("/", handlers.EmployeeHandler.GetAll())
			rt.Get("/{id}", handlers.EmployeeHandler.GetByID())
			// - POST /api/v1/employees
			rt.Post("/", handlers.EmployeeHandler.Create())
			// - PUT /api/v1/employees/{id}
			rt.Patch("/{id}", handlers.EmployeeHandler.Update())
			// - DELETE /api/v1/employees/{id}
			rt.Delete("/{id}", handlers.EmployeeHandler.Delete())
			// - GET /api/v1/employees/reportInboundOrders?id=?
			rt.Get("/reportInboundOrders", handlers.EmployeeHandler.Report())
		})
		rt.Route("/inboundOrders", func(rt chi.Router) {
			// - POST /api/v1/inboundOrders
			rt.Post("/", handlers.InboundOrderHandler.Create())
		})
		rt.Route("/products", func(rt chi.Router) {
			// - GET /api/v1/products /
			rt.Get("/", handlers.ProductHandler.GetAll())
			rt.Get("/{id}", handlers.ProductHandler.GetByID())
			// - GET /api/v1/products/reportRecords /
			rt.Get("/reportRecords", handlers.ProductHandler.GetRecord())
			// - POST /api/v1/products /
			rt.Post("/", handlers.ProductHandler.Create())
			// - DELETE /api/v1/products /
			rt.Delete("/{id}", handlers.ProductHandler.Delete())
			// - PATCH /api/v1/products /
			rt.Patch("/{id}", handlers.ProductHandler.Update())
		})
		rt.Route("/productRecords", func(rt chi.Router) {
			// - POST /api/v1/products /
			rt.Post("/", handlers.ProductRecordHandler.Create())
		})
		rt.Route("/sections", func(rt chi.Router) {
			// - GET /api/v1/sections
			rt.Get("/", handlers.SectionHandler.GetAll())
			rt.Get("/{id}", handlers.SectionHandler.GetByID())
			// - POST /api/v1/products /
			rt.Post("/", handlers.SectionHandler.Create())
			// - PATCH /api/v1/products /
			rt.Patch("/{id}", handlers.SectionHandler.Patch())
			// - DELETE /api/v1/employees/{id}
			rt.Delete("/{id}", handlers.SectionHandler.Delete())
		})
		rt.Route("/productBatches", func(rt chi.Router) {
			// - POST /api/v1/productBatches
			rt.Post("/", productBatchHd.Create())
		})
		rt.Route("/sellers", func(rt chi.Router) {
			// - GET /api/v1/sellers
			rt.Get("/", handlers.SellerHandler.GetAll())
			// - GET /api/v1/sellers/{id}
			rt.Get("/{id}", handlers.SellerHandler.GetByID())
			// - POST /api/v1/sellers
			rt.Post("/", handlers.SellerHandler.Create())
			// - PATCH /api/v1/sellers/{id}
			rt.Patch("/{id}", handlers.SellerHandler.Update())
			// - DELETE /api/v1/sellers/{id}
			rt.Delete("/{id}", handlers.SellerHandler.Delete())
		})
		rt.Route("/warehouses", func(rt chi.Router) {
			// - GET /api/v1/warehouses
			rt.Get("/", handlers.WarehouseHandler.GetAll())
			// - GET /api/v1/warehouses/{id}
			rt.Get("/{id}", handlers.WarehouseHandler.GetByID())
			// - POST /api/v1/warehouses
			rt.Post("/", handlers.WarehouseHandler.Create())
			// - PATCH /api/v1/warehouses/{id}
			rt.Patch("/{id}", handlers.WarehouseHandler.Update())
			// - DELETE /api/v1/warehouses/{id}
			rt.Delete("/{id}", handlers.WarehouseHandler.Delete())
		})
		rt.Route("/localities", func(rt chi.Router){
			rt.Post("/", handlers.LocalityHandler.Create())
		})
		rt.Route("/carriers", func(rt chi.Router){
			// - POST /api/v1/carriers
			rt.Post("/", handlers.CarryHandler.Create())
		})
	})

	// run server
	err = http.ListenAndServe(a.serverAddress, rt)
	return
}
