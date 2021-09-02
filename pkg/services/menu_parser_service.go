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

type MenuParser struct {
	supplierRepo db_repository.SupplierDBRepository
	productsRepo db_repository.ProductDBRepository
}

func NewParser(supplierRepo *db_repository.SupplierDBRepository, productsRepo *db_repository.ProductDBRepository) *MenuParser {
	return &MenuParser{
		supplierRepo: *supplierRepo,
		productsRepo: *productsRepo,
	}
}

type MenuParserI interface {
	Parser(<-chan time.Time) error
	DBsave(suppliersList *[]models.Supplier)
}

func (mp *MenuParser) DBsave(suppliersList *[]models.Supplier) {
	for _, s := range *suppliersList {
		dbSupplier, err := mp.supplierRepo.SearchByID(s.ID)
		if err != nil {
			internal.Log.Errorf("supplier %d search error: %v", s.ID, err.Error())
			return
		}
		if dbSupplier {
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
		for _, menu := range s.Menu {
			menu.SupplierID = supplierID
			_, err := mp.productsRepo.Create(&menu)
			if err != nil {
				internal.Log.Errorf("supplier %d menu didn't created: %s", s.ID, err.Error())
				return
			}
			internal.Log.Debugf("supplier %d menu created", s.ID)
		}
	}

	return
}

func (mp *MenuParser) Parser() {
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

				supplierMenu, err2 := parser.GetSupplierMenuByID(i + 1)
				if err2 != nil {
					return
				}
				suppliersList[i].Menu = supplierMenu
			}(i)
		}
		wg.Wait()

		mp.DBsave(&suppliersList)

		time.Sleep(time.Minute)
	}
}
