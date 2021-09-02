package parser

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/YRIDZE/Bicycle-delivery-service/internal"
	"github.com/YRIDZE/Bicycle-delivery-service/pkg/models"
)

var URL = "http://foodapi.true-tech.php.nixdev.co/restaurants"

func GetSuppliers() ([]models.Supplier, error) {
	response, err := http.Get(fmt.Sprintf("%s", URL))
	if err != nil {
		return nil, err
	}

	jsonBytes, err := io.ReadAll(response.Body)
	if err != nil {
		internal.Log.Errorf("error close body", err)
	}

	var suppliersList models.SuppliersResponse
	err = json.Unmarshal(jsonBytes, &suppliersList)
	if err != nil {
		return nil, err
	}
	return suppliersList.Restaurants, nil
}

func GetSupplierMenuByID(id int) ([]models.Product, error) {
	response, err := http.Get(fmt.Sprintf("%s/%d/%s", URL, id, "menu"))
	if err != nil {
		internal.Log.Error(err.Error())
		return nil, err
	}

	jsonBytes, err := io.ReadAll(response.Body)
	err = response.Body.Close()
	if err != nil {
		internal.Log.Errorf("error close body", err)
	}

	supplierMenu := new(models.MenuResponse)
	err = json.Unmarshal(jsonBytes, &supplierMenu)
	if err != nil {
		internal.Log.Error(err)
		return nil, err
	}
	return supplierMenu.Menu, nil
}
