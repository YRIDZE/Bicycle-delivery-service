package services

import (
	"sync"
	"time"

	"github.com/YRIDZE/Bicycle-delivery-service/cmd/parser"
	"github.com/YRIDZE/Bicycle-delivery-service/internal"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models/db_repository"
)

var wg sync.WaitGroup

type SupplierProductsParser struct {
	supplierRepo db_repository.SupplierRepository
	productsRepo db_repository.ProductRepository
}

func NewParser(supplierRepo *db_repository.SupplierRepository, productsRepo *db_repository.ProductRepository) *SupplierProductsParser {
	return &SupplierProductsParser{
		supplierRepo: *supplierRepo,
		productsRepo: *productsRepo,
	}
}

type MenuParserI interface {
	Parse()
	Save(suppliersList *[]models.Supplier)
}

func (mp *SupplierProductsParser) Save(suppliersList *[]models.Supplier) {
	for _, s := range *suppliersList {
		oldSupplierID, _ := mp.supplierRepo.GetByName(s.Name)

		if oldSupplierID != 0 {
			err := mp.supplierRepo.Delete(oldSupplierID)
			if err != nil {
				internal.Log.Errorf("supplier %d and supplier-menu didn't removed: %s", oldSupplierID, err.Error())
				return
			}
			internal.Log.Debugf("supplier %d and supplier-menu removed", oldSupplierID)
		}
		supplierID, err := mp.supplierRepo.Create(&s)
		if err != nil {
			internal.Log.Errorf("supplier %d didn't created: %s", supplierID, err.Error())
			return
		}
		internal.Log.Debugf("supplier %d created", supplierID)

		for _, m := range s.Menu {
			m.SupplierID = supplierID

			oldProductID, _ := mp.productsRepo.GetByName(m.Name)
			if oldProductID != 0 {
				err := mp.productsRepo.Delete(int(oldProductID))
				if err != nil {
					return
				}
			}
			_, err = mp.productsRepo.Create(&m)
			if err != nil {
				internal.Log.Errorf("supplier %d menu didn't created: %s", supplierID, err.Error())
				return
			}
			internal.Log.Debugf("supplier %d menu created", supplierID)
		}
	}
	return
}

func (mp *SupplierProductsParser) Parse() {
	for {
		internal.Log.Debug("New parsing iteration...")
		suppliersList, err := parser.GetSuppliers()
		if err != nil {
			return
		}
		for i := range suppliersList {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()

				supplierMenu, err2 := parser.GetSupplierProductsByID(i + 1)
				if err2 != nil {
					return
				}
				suppliersList[i].Menu = supplierMenu
			}(i)
		}
		wg.Wait()

		mp.Save(&suppliersList)
		time.Sleep(time.Minute)
	}
}
