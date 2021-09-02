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
	supplierRepo db_repository.SupplierDBRepository
	productsRepo db_repository.ProductDBRepository
}

func NewParser(supplierRepo *db_repository.SupplierDBRepository, productsRepo *db_repository.ProductDBRepository) *SupplierProductsParser {
	return &SupplierProductsParser{
		supplierRepo: *supplierRepo,
		productsRepo: *productsRepo,
	}
}

type MenuParserI interface {
	Parser(<-chan time.Time) error
	DBsave(suppliersList *[]models.Supplier)
}

func (mp *SupplierProductsParser) SaveParsedDataToDb(suppliersList *[]models.Supplier) {
	for _, s := range *suppliersList {
		pick, err := mp.supplierRepo.SearchByID(s.ID)
		if err != nil {
			internal.Log.Errorf("supplier %d search error: %v", s.ID, err.Error())
			return
		}
		if pick {
			err := mp.supplierRepo.Delete(s.ID)
			if err != nil {
				internal.Log.Errorf("supplier %d and supplier-menu didn't removed: %s", s.ID, err.Error())
				return
			}
			internal.Log.Debugf("supplier %d and supplier-menu removed", s.ID)
		}
		supplierID, err := mp.supplierRepo.Create(&s)
		if err != nil {
			internal.Log.Errorf("supplier %d didn't created: %s", s.ID, err.Error())
			return
		}
		internal.Log.Debugf("supplier %d created", s.ID)

		for _, m := range s.Menu {
			m.SupplierID = supplierID
			_, err := mp.productsRepo.Create(&m)
			if err != nil {
				internal.Log.Errorf("supplier %d menu didn't created: %s", s.ID, err.Error())
				return
			}
			internal.Log.Debugf("supplier %d menu created", s.ID)
		}
	}
	return
}

func (mp *SupplierProductsParser) Parser() {
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

		mp.SaveParsedDataToDb(&suppliersList)
		time.Sleep(time.Minute)
	}
}
