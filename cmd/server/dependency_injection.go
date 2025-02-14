package server

import (
	"database/sql"

	buyerRepository "github.com/luisantonisu/wave15-grupo4/internal/repository/buyer"
	countryRepository "github.com/luisantonisu/wave15-grupo4/internal/repository/country"
	employeeRepository "github.com/luisantonisu/wave15-grupo4/internal/repository/employee"
	inboundOrderRepository "github.com/luisantonisu/wave15-grupo4/internal/repository/inbound_order"
	localityRepository "github.com/luisantonisu/wave15-grupo4/internal/repository/locality"
	productRepository "github.com/luisantonisu/wave15-grupo4/internal/repository/product"
	productRecordRepository "github.com/luisantonisu/wave15-grupo4/internal/repository/product_record"
	provinceRepository "github.com/luisantonisu/wave15-grupo4/internal/repository/province"
	purchaseOrderRepository "github.com/luisantonisu/wave15-grupo4/internal/repository/purchase_order"
	sectionRepository "github.com/luisantonisu/wave15-grupo4/internal/repository/section"
	sellerRepository "github.com/luisantonisu/wave15-grupo4/internal/repository/seller"
	warehouseRepository "github.com/luisantonisu/wave15-grupo4/internal/repository/warehouse"
	carryRepository "github.com/luisantonisu/wave15-grupo4/internal/repository/carry"

	buyerService "github.com/luisantonisu/wave15-grupo4/internal/service/buyer"
	employeeService "github.com/luisantonisu/wave15-grupo4/internal/service/employee"
	inboundOrderService "github.com/luisantonisu/wave15-grupo4/internal/service/inbound_order"
	localityService "github.com/luisantonisu/wave15-grupo4/internal/service/locality"
	productService "github.com/luisantonisu/wave15-grupo4/internal/service/product"
	productRecordService "github.com/luisantonisu/wave15-grupo4/internal/service/product_record"
	purchaseOrderService "github.com/luisantonisu/wave15-grupo4/internal/service/purchase_order"
	sectionService "github.com/luisantonisu/wave15-grupo4/internal/service/section"
	sellerService "github.com/luisantonisu/wave15-grupo4/internal/service/seller"
	warehouseService "github.com/luisantonisu/wave15-grupo4/internal/service/warehouse"
	carryService "github.com/luisantonisu/wave15-grupo4/internal/service/carry"

	buyerHandler "github.com/luisantonisu/wave15-grupo4/internal/handler/buyer"
	employeeHandler "github.com/luisantonisu/wave15-grupo4/internal/handler/employee"
	inboundOrderHandler "github.com/luisantonisu/wave15-grupo4/internal/handler/inbound_order"
	localityHandler "github.com/luisantonisu/wave15-grupo4/internal/handler/locality"
	productHandler "github.com/luisantonisu/wave15-grupo4/internal/handler/product"
	productRecordHandler "github.com/luisantonisu/wave15-grupo4/internal/handler/product_record"
	purchaseOrderHandler "github.com/luisantonisu/wave15-grupo4/internal/handler/purchase_order"
	sectionHandler "github.com/luisantonisu/wave15-grupo4/internal/handler/section"
	sellerHandler "github.com/luisantonisu/wave15-grupo4/internal/handler/seller"
	warehouseHandler "github.com/luisantonisu/wave15-grupo4/internal/handler/warehouse"
	carryHandler "github.com/luisantonisu/wave15-grupo4/internal/handler/carry"
)

type Handlers struct {
	BuyerHandler         *buyerHandler.BuyerHandler
	PurchaseOrderHandler *purchaseOrderHandler.PurchaseOrderHandler
	EmployeeHandler      *employeeHandler.EmployeeHandler
	InboundOrderHandler  *inboundOrderHandler.InboundOrderHandler
	ProductHandler       *productHandler.ProductHandler
	ProductRecordHandler *productRecordHandler.ProductRecordHandler
	SectionHandler       *sectionHandler.SectionHandler
	SellerHandler        *sellerHandler.SellerHandler
	WarehouseHandler     *warehouseHandler.WarehouseHandler
	LocalityHandler      *localityHandler.LocalityHandler
	CarryHandler         *carryHandler.CarryHandler
}

func GetHandlers(db *sql.DB) Handlers {
	buyerRp := buyerRepository.NewBuyerRepository(db)
	purchaseOrderRp := purchaseOrderRepository.NewPurchaseOrderRepository(db)
	employeeRp := employeeRepository.NewEmployeeRepository(db)
	inboundOrderRp := inboundOrderRepository.NewInboundOrderRepository(db)
	productRp := productRepository.NewProductRepository(db)
	productRecordRp := productRecordRepository.NewProductRecordRepository(db)
	sectionRp := sectionRepository.NewSectionRepository(db)
	sellerRp := sellerRepository.NewSellerRepository(db)
	warehouseRp := warehouseRepository.NewWarehouseRepository(db)
	countryRp := countryRepository.NewCountryRepository(db)
	provinceRp := provinceRepository.NewProvinceRepository(db)
	localityRp := localityRepository.NewLocalityRepository(db)
	carryRp := carryRepository.NewCarryRepository(db)

	// - service
	buyerSv := buyerService.NewBuyerService(buyerRp)
	PurchaseOrderSv := purchaseOrderService.NewPurchaseOrderService(purchaseOrderRp)
	employeeSv := employeeService.NewEmployeeService(employeeRp, warehouseRp)
	inboundOrderSv := inboundOrderService.NewInboundOrderService(inboundOrderRp, employeeRp, warehouseRp)
	productSv := productService.NewProductService(productRp)
	productRecordSv := productRecordService.NewProductRecordService(productRecordRp, productRp)
	sectionSv := sectionService.NewSectionService(sectionRp)
	sellerSv := sellerService.NewSellerService(sellerRp)
	warehouseSv := warehouseService.NewWarehouseService(warehouseRp)
	localitySv := localityService.NewLocalityService(countryRp, provinceRp, localityRp)
	carrySv := carryService.NewCarryService(carryRp)

	// - handler
	buyerHd := buyerHandler.NewBuyerHandler(buyerSv)
	purchaseOrderHd := purchaseOrderHandler.NewPurchaseOrderHandler(PurchaseOrderSv)
	employeeHd := employeeHandler.NewEmployeeHandler(employeeSv)
	inboundOrderHd := inboundOrderHandler.NewInboundOrderHandler(inboundOrderSv)
	productHd := productHandler.NewProductHandler(productSv)
	productRecordHd := productRecordHandler.NewProductRecordHandler(productRecordSv)
	sectionHd := sectionHandler.NewSectionHandler(sectionSv)
	sellerHd := sellerHandler.NewSellerHandler(sellerSv)
	warehouseHd := warehouseHandler.NewWarehouseHandler(warehouseSv)
	localityHd := localityHandler.NewLocalityHandler(localitySv)
	carryHd := carryHandler.NewCarryHandler(carrySv)

	return Handlers{
		BuyerHandler:         buyerHd,
		PurchaseOrderHandler: purchaseOrderHd,
		EmployeeHandler:      employeeHd,
		InboundOrderHandler:  inboundOrderHd,
		ProductHandler:       productHd,
		ProductRecordHandler: productRecordHd,
		SectionHandler:       sectionHd,
		SellerHandler:        sellerHd,
		WarehouseHandler:     warehouseHd,
		LocalityHandler:      localityHd,
		CarryHandler:         carryHd,
	}
}
