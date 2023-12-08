package omc

import (
	"fmt"
	"sync"

	"github.com/netdoop/netdoop/store"
	"github.com/netdoop/netdoop/utils"
)

var dataModelsMap map[string]*DataModel
var dataModelsMapLock sync.Mutex

func ReloadAllDataModels() error {
	logger := utils.GetLogger()
	logger.Debug("reload all data models")
	db := store.GetDB()
	all := []*DataModel{}
	if result := db.Model(&DataModel{}).Preload("Parameters").Preload("Templates").Find(&all); result.Error != nil {
		return result.Error
	}

	tmp := map[string]*DataModel{}
	for _, v := range all {
		key := fmt.Sprintf("%v:%v:%v", v.Schema, v.ProductType, v.Name)
		tmp[key] = v
	}
	dataModelsMapLock.Lock()
	defer dataModelsMapLock.Unlock()
	dataModelsMap = tmp
	return nil
}

func GetDataModel(schema string, productType string, name string) *DataModel {
	dataModelsMapLock.Lock()
	defer dataModelsMapLock.Unlock()
	key := fmt.Sprintf("%v:%v:%v", schema, productType, name)
	v := dataModelsMap[key]
	return v
}

func GetDataModelByProduct(product *Product) *DataModel {
	return GetDataModel(product.Schema, product.ProductType, product.ParameterPath)
}
