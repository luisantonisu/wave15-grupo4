package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/luisantonisu/wave15-grupo4/infrastructure/db"
	"github.com/luisantonisu/wave15-grupo4/internal/config"
	buyerRepository "github.com/luisantonisu/wave15-grupo4/internal/repository/buyer"
	employeeRepository "github.com/luisantonisu/wave15-grupo4/internal/repository/employee"
	inboundOrderRepository "github.com/luisantonisu/wave15-grupo4/internal/repository/inbound_order"
	productRepository "github.com/luisantonisu/wave15-grupo4/internal/repository/product"
	productRecordRepository "github.com/luisantonisu/wave15-grupo4/internal/repository/product_record"
	purchaseOrderRepository "github.com/luisantonisu/wave15-grupo4/internal/repository/purchase_order"
	sectionRepository "github.com/luisantonisu/wave15-grupo4/internal/repository/section"
	sellerRepository "github.com/luisantonisu/wave15-grupo4/internal/repository/seller"
	warehouseRepository "github.com/luisantonisu/wave15-grupo4/internal/repository/warehouse"

	buyerService "github.com/luisantonisu/wave15-grupo4/internal/service/buyer"
	employeeService "github.com/luisantonisu/wave15-grupo4/internal/service/employee"
	inboundOrderService "github.com/luisantonisu/wave15-grupo4/internal/service/inbound_order"
	productService "github.com/luisantonisu/wave15-grupo4/internal/service/product"
	productRecordService "github.com/luisantonisu/wave15-grupo4/internal/service/product_record"
	purchaseOrderService "github.com/luisantonisu/wave15-grupo4/internal/service/purchase_order"
	sectionService "github.com/luisantonisu/wave15-grupo4/internal/service/section"
	sellerService "github.com/luisantonisu/wave15-grupo4/internal/service/seller"
	warehouseService "github.com/luisantonisu/wave15-grupo4/internal/service/warehouse"

	buyerHandler "github.com/luisantonisu/wave15-grupo4/internal/handler/buyer"
	employeeHandler "github.com/luisantonisu/wave15-grupo4/internal/handler/employee"
	inboundOrderHandler "github.com/luisantonisu/wave15-grupo4/internal/handler/inbound_order"
	productHandler "github.com/luisantonisu/wave15-grupo4/internal/handler/product"
	productRecordHandler "github.com/luisantonisu/wave15-grupo4/internal/handler/product_record"
	purchaseOrderHandler "github.com/luisantonisu/wave15-grupo4/internal/handler/purchase_order"
	sectionHandler "github.com/luisantonisu/wave15-grupo4/internal/handler/section"
	sellerHandler "github.com/luisantonisu/wave15-grupo4/internal/handler/seller"
	warehouseHandler "github.com/luisantonisu/wave15-grupo4/internal/handler/warehouse"
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
	//TODO
	database := db.ConnectDB(&cfg)
	defer database.Close()

	db, err := loader.Load()
	if err != nil {
		return
	}

	// - repository
	buyerRp := buyerRepository.NewBuyerRepository(database)
	purchaseOrderRp := purchaseOrderRepository.NewPurchaseOrderRepository(database)
	employeeRp := employeeRepository.NewEmployeeRepository(database)
	inboundOrderRp := inboundOrderRepository.NewInboundOrderRepository(database)
	productRp := productRepository.NewProductRepository(database)
	productRecordRp := productRecordRepository.NewProductRecordRepository(database)
	sectionRp := sectionRepository.NewSectionRepository(db.Sections)
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
	sellerHd := sellerHandler.NewSellerHandler(sellerSv)                             // sellerHd
	warehouseHd := warehouseHandler.NewWarehouseHandler(warehouseSv)                 // warehouseHd

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
			// - POST /api/v1/buyers
			rt.Post("/", buyerHd.Create())
			// - PUT /api/v1/buyers
			rt.Patch("/{id}", buyerHd.Update())
			// - DELETE /api/v1/buyers
			rt.Delete("/{id}", buyerHd.Delete())
		})
		rt.Route("/purchaseOrders", func(rt chi.Router) {
			// - POST /api/v1/buyers
			rt.Post("/", purchaseOrderHd.Create())
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
			// - GET /api/v1/employees/reportInboundOrders?id=?
			rt.Get("/reportInboundOrders", employeeHd.Report())
		})
		rt.Route("/inboundOrders", func(rt chi.Router) {
			// - POST /api/v1/inboundOrders
			rt.Post("/", inboundOrderHd.Create())
		})
		rt.Route("/products", func(rt chi.Router) {
			// - GET /api/v1/products /
			rt.Get("/", productHd.GetAll())
			rt.Get("/{id}", productHd.GetByID())
			// - GET /api/v1/products/reportRecords /
			rt.Get("/reportRecords", productHd.GetRecord())
			// - POST /api/v1/products /
			rt.Post("/", productHd.Create())
			// - DELETE /api/v1/products /
			rt.Delete("/{id}", productHd.Delete())
			// - PATCH /api/v1/products /
			rt.Patch("/{id}", productHd.Update())
		})
		rt.Route("/productRecords", func(rt chi.Router) {
			// - POST /api/v1/products /
			rt.Post("/", productRecordHd.Create())
		})
		rt.Route("/sections", func(rt chi.Router) {
			// - GET /api/v1/sections
			rt.Get("/", sectionHd.GetAll())
			rt.Get("/{id}", sectionHd.GetByID())
			// - POST /api/v1/products /
			rt.Post("/", sectionHd.Create())
			// - PATCH /api/v1/products /
			rt.Patch("/{id}", sectionHd.Patch())
			// - DELETE /api/v1/employees/{id}
			rt.Delete("/{id}", sectionHd.Delete())
		})
		rt.Route("/sellers", func(rt chi.Router) {
			// - GET /api/v1/sellers
			rt.Get("/", sellerHd.GetAll())
			// - GET /api/v1/sellers/{id}
			rt.Get("/{id}", sellerHd.GetByID())
			// - POST /api/v1/sellers
			rt.Post("/", sellerHd.Create())
			// - PATCH /api/v1/sellers/{id}
			rt.Patch("/{id}", sellerHd.Update())
			// - DELETE /api/v1/sellers/{id}
			rt.Delete("/{id}", sellerHd.Delete())
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
			// - DELETE /api/v1/warehouses/{id}
			rt.Delete("/{id}", warehouseHd.Delete())
		})
	})

	// run server
	err = http.ListenAndServe(a.serverAddress, rt)
	return
}
