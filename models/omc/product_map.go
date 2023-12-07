package omc

import (
	"fmt"
	"sync"

	"github.com/netdoop/netdoop/store"
	"github.com/netdoop/netdoop/utils"

	"gorm.io/gorm"
)

var productsMap map[string]*Product
var productsMapLock sync.Mutex

func ReloadAllProducts() error {
	logger := utils.GetLogger()
	logger.Debug("reload all products")
	db := store.GetDB()
	all := []*Product{}
	if result := db.Model(&Product{}).Preload("SupportedAlarms").Find(&all); result.Error != nil {
		return result.Error
	}

	tmp := map[string]*Product{}
	for _, v := range all {
		key := fmt.Sprintf("%v:%v-%v", v.Schema, v.Oui, v.ProductClass)
		tmp[key] = v
	}
	productsMapLock.Lock()
	defer productsMapLock.Unlock()
	productsMap = tmp
	return nil
}

func GetProductByModelName(db *gorm.DB, schema string, modelName string) *Product {
	var product Product
	if result := db.Where("schema = ? AND model_name = ?", schema, modelName).First(&product); result.Error != nil {
		return nil
	}
	return &product
}

func GetProduct(schema string, oui string, productClass string) *Product {
	productsMapLock.Lock()
	defer productsMapLock.Unlock()
	key := fmt.Sprintf("%v:%v-%v", schema, oui, productClass)
	v := productsMap[key]
	return v
}
