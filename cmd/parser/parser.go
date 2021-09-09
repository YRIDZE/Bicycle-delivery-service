package parser

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models/db_repository"
	yolo_log "github.com/YRIDZE/yolo-log"
	"github.com/spf13/viper"
)

var wg sync.WaitGroup

type SupplierProductsParser struct {
	logger       *yolo_log.Logger
	supplierRepo db_repository.SupplierRepository
	productsRepo db_repository.ProductRepository
}

func NewParser(logger *yolo_log.Logger, supplierRepo *db_repository.SupplierRepository, productsRepo *db_repository.ProductRepository) *SupplierProductsParser {
	return &SupplierProductsParser{
		logger:       logger,
		supplierRepo: *supplierRepo,
		productsRepo: *productsRepo,
	}
}

type MenuParserI interface {
	Parse()
	Save(suppliersList *[]models.Supplier)
	GetSuppliers() ([]models.Supplier, error)
	GetSupplierProductsByID(id int) ([]models.Product, error)
}

func (h *SupplierProductsParser) Save(suppliersList *[]models.Supplier) {
	for _, s := range *suppliersList {
		oldSupplierID, _ := h.supplierRepo.GetByName(s.Name)

		if oldSupplierID != 0 {
			err := h.supplierRepo.Delete(oldSupplierID)
			if err != nil {
				h.logger.Errorf("supplier %d and supplier-menu didn't removed: %s", oldSupplierID, err.Error())
				return
			}
			h.logger.Debugf("supplier %d and supplier-menu removed", oldSupplierID)
		}
		supplier, err := h.supplierRepo.Create(&models.SupplierResponse{ID: s.ID, Name: s.Name, Image: s.Image})
		if err != nil {
			h.logger.Errorf("supplier didn't created: %s", err.Error())
			return
		}
		h.logger.Debugf("supplier %d created", supplier.ID)

		for _, m := range s.Menu {
			m.SupplierID = supplier.ID

			oldProductID, _ := h.productsRepo.GetByName(m.Name)
			if oldProductID != 0 {
				err := h.productsRepo.Delete(int(oldProductID))
				if err != nil {
					return
				}
			}
			_, err = h.productsRepo.Create(&m)
			if err != nil {
				h.logger.Errorf("supplier %d menu didn't created: %s", supplier.ID, err.Error())
				return
			}
			h.logger.Debugf("supplier %d menu created", supplier.ID)
		}
	}
	return
}

func (h *SupplierProductsParser) Parse() {
	for {
		h.logger.Debug("New parsing iteration...")
		suppliersList, err := h.GetSuppliers()
		if err != nil {
			return
		}
		for i := range suppliersList {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()

				supplierMenu, err2 := h.GetSupplierProductsByID(i + 1)
				if err2 != nil {
					return
				}
				suppliersList[i].Menu = supplierMenu
			}(i)
		}
		wg.Wait()

		h.Save(&suppliersList)
		time.Sleep(time.Duration(viper.GetInt("parser.delay")) * time.Minute)
	}
}

func (h *SupplierProductsParser) GetSuppliers() ([]models.Supplier, error) {
	resp, err := http.Get(fmt.Sprintf("%s", viper.GetString("parser.url")))
	if err != nil {
		return nil, err
	}

	jsonBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		h.logger.Errorf("error close body", err)
	}

	var suppliersList models.SuppliersResponse
	err = json.Unmarshal(jsonBytes, &suppliersList)
	if err != nil {
		return nil, err
	}

	return suppliersList.Suppliers, nil
}

func (h *SupplierProductsParser) GetSupplierProductsByID(id int) ([]models.Product, error) {
	response, err := http.Get(fmt.Sprintf("%s/%d/%s", viper.GetString("parser.url"), id, "menu"))
	if err != nil {
		h.logger.Error(err.Error())
		return nil, err
	}

	jsonBytes, err := io.ReadAll(response.Body)
	err = response.Body.Close()
	if err != nil {
		h.logger.Errorf("error close body", err)
	}

	supplierProducts := new(models.ProductsResponse)
	err = json.Unmarshal(jsonBytes, &supplierProducts)
	if err != nil {
		h.logger.Error(err)
		return nil, err
	}

	return supplierProducts.Products, nil
}
