package omc

import (
	"github.com/heypkg/store/jsontype"
	"gorm.io/gorm"
)

func SetupDefaultDataModel(db *gorm.DB, schema string) error {
	all := []DataModel{
		{Default: true, Schema: schema, ProductType: "enb", Name: "InternetGatewayDevice", ParameterPath: "InternetGatewayDevice."},
		{Default: true, Schema: schema, ProductType: "enb", Name: "Device", ParameterPath: "Device."},
		{Default: true, Schema: schema, ProductType: "cpe", Name: "InternetGatewayDevice", ParameterPath: "InternetGatewayDevice."},
		{Default: true, Schema: schema, ProductType: "cpe", Name: "Device", ParameterPath: "Device."},
	}
	if result := db.CreateInBatches(all, 100); result.Error != nil {
		return result.Error
	}

	return nil
}

func SetupDefaultDataModelTemplate(db *gorm.DB, dm *DataModel) error {
	all := []DataModelTemplate{}

	all = append(all, DataModelTemplate{DataModelID: dm.ID, Name: "", ParameterNames: &jsontype.Tags{
		"Device.Services.FAPService.{i}.CellConfig.": "Cell Config",
		"Device.Services.FAPService.{i}.FAPControl.": "FAP Control",
	}})

	if result := db.CreateInBatches(all, 100); result.Error != nil {
		return result.Error
	}
	return nil
}
