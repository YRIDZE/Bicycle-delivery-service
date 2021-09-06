package parser

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/YRIDZE/Bicycle-delivery-service/internal"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models/db_repository"
	"github.com/spf13/viper"
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
		suppliersList, err := GetSuppliers()
		if err != nil {
			return
		}
		for i := range suppliersList {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()

				supplierMenu, err2 := GetSupplierProductsByID(i + 1)
				if err2 != nil {
					return
				}
				suppliersList[i].Menu = supplierMenu
			}(i)
		}
		wg.Wait()

		mp.Save(&suppliersList)
		time.Sleep(time.Duration(viper.GetInt("parser.delay")) * time.Minute)
	}
}

func GetSuppliers() ([]models.Supplier, error) {
	resp, err := http.Get(fmt.Sprintf("%s", viper.GetString("parser.url")))
	if err != nil {
		return nil, err
	}

	jsonBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		internal.Log.Errorf("error close body", err)
	}

	var suppliersList models.SuppliersResponse
	err = json.Unmarshal(jsonBytes, &suppliersList)
	if err != nil {
		return nil, err
	}

	return suppliersList.Suppliers, nil
}

func GetSupplierProductsByID(id int) ([]models.Product, error) {
	response, err := http.Get(fmt.Sprintf("%s/%d/%s", viper.GetString("parser.url"), id, "menu"))
	if err != nil {
		internal.Log.Error(err.Error())
		return nil, err
	}

	jsonBytes, err := io.ReadAll(response.Body)
	err = response.Body.Close()
	if err != nil {
		internal.Log.Errorf("error close body", err)
	}

	supplierProducts := new(models.ProductsResponse)
	err = json.Unmarshal(jsonBytes, &supplierProducts)
	if err != nil {
		internal.Log.Error(err)
		return nil, err
	}

	return supplierProducts.Products, nil
}
