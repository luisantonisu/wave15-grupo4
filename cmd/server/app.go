package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/luisantonisu/wave15-grupo4/internal/repository"
	"honnef.co/go/tools/go/loader"
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
	ld := loader.Load(a.loaderFilePath)
	db, err := ld.Load()
	if err != nil {
		return
	}

	// - repository
	buyerRp := repository.NewBuyerRepository(db)
	employeeRp := repository.NewEmployeeRepository(db)
	productRp := repository.NewProductRepository(db)
	sectionRp := repository.NewSectionRepository(db)
	sellerRp := repository.NewSellerRepository(db)
	warehouseRp := repository.NewWarehouseRepository(db)

	// - service
	buyerSv := service.NewBuyerService(buyerRp)
	employeeSv := service.NewEmployeeService(employeeRp)
	productSv := service.NewProductService(productRp)
	sectionSv := service.NewSectionService(sectionRp)
	sellerSv := service.NewSellerService(sellerRp)
	warehouseSv := service.NewWarehouseService(warehouseRp)

	// - handler
	_ = handler.NewBuyerHandler(buyerSv)         // buyerHd
	_ = handler.NewEmployeeHandler(employeeSv)   // employeeHd
	_ = handler.NewProductHandler(productSv)     // productHd
	_ = handler.NewSectionHandler(sectionSv)     // sectionHd
	_ = handler.NewSellerHandler(sellerSv)       // sellerHd
	_ = handler.NewWarehouseHandler(warehouseSv) // warehouseHd

	// router
	rt := chi.NewRouter()

	// - middlewares
	rt.Use(middleware.Logger)
	rt.Use(middleware.Recoverer)

	// - endpoints
	rt.Route("/", func(rt chi.Router) {
		// - GET /
		// rt.Get("/", hd.GetAll())
	})

	// run server
	err = http.ListenAndServe(a.serverAddress, rt)
	return
}
