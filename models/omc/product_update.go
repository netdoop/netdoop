package omc

import (
	"github.com/heypkg/store/jsontype"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func CreateProduct(db *gorm.DB, schema string, id uint,
	productType string,
	modelName string,
	oui string,
	productClass string,
	manufacturer string,
	parameterPath string,
	isDefault bool,
) (*Product, error) {
	product := &Product{
		ID:            id,
		Schema:        schema,
		ModelName:     modelName,
		Oui:           oui,
		ProductClass:  productClass,
		ProductType:   productType,
		Manufacturer:  manufacturer,
		ParameterPath: parameterPath,
		Default:       isDefault,
		Enable:        true,
	}
	if err := db.Where("schema = ? AND model_name = ? AND oui = ? AND product_class = ?",
		schema, modelName, oui, productClass,
	).FirstOrCreate(product).Error; err != nil {
		return nil, errors.Wrap(err, "create product")
	}
	return product, nil
}

func SetProductPerformanceValueDefines(db *gorm.DB, product *Product, performanceValueDefines *jsontype.Tags) error {
	if performanceValueDefines == nil {
		return nil
	}
	updateColumns := []string{"performance_value_defines"}
	updateData := &Product{
		PerformanceValueDefines: performanceValueDefines,
	}
	updateData.SaveData()
	result := db.Model(product).Select(updateColumns).Updates(updateData)
	if result.Error != nil {
		return errors.Wrap(result.Error, "update performance value defines of product")
	}
	return nil
}
