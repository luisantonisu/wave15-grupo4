package server

import (
	"net/http"
	"sync"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/luisantonisu/wave15-grupo4/infrastructure/db"
	"github.com/luisantonisu/wave15-grupo4/internal/config"
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

	handlers := GetHandlers(database)

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
			// - POST /api/v1/sections /
			rt.Post("/", handlers.SectionHandler.Create())
			// - PATCH /api/v1/sections /
			rt.Patch("/{id}", handlers.SectionHandler.Patch())
			// - DELETE /api/v1/sections/{id}
			rt.Delete("/{id}", handlers.SectionHandler.Delete())
			// - GET /api/v1/sections/reportProducts /
			rt.Get("/reportProducts", handlers.SectionHandler.Report())
		})
		rt.Route("/productBatches", func(rt chi.Router) {
			// - POST /api/v1/productBatches
			rt.Post("/", handlers.ProductBatchHandler.Create())
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
		rt.Route("/localities", func(rt chi.Router) {
			// - POST /api/v1/localities
			rt.Post("/", handlers.LocalityHandler.Create())
			// - GET /api/v1/localities
			rt.Get("/reportSellers", handlers.LocalityHandler.SellersReport())
			// - GET /api/v1/localities/reportCarriers?id=?
			rt.Get("/reportCarriers", handlers.LocalityHandler.CarriersReport())
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
		rt.Route("/carriers", func(rt chi.Router) {
			// - POST /api/v1/carriers
			rt.Post("/", handlers.CarryHandler.Create())
		})
	})

	// run server
	err = http.ListenAndServe(a.serverAddress, rt)
	return
}
