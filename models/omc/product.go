package omc

import (
	"github.com/heypkg/store/jsontype"
	"github.com/netdoop/cwmp/acs"
	"gorm.io/gorm"
)

type Product struct {
	ID      uint              `json:"Id" gorm:"primarykey"`
	Updated jsontype.JSONTime `json:"Updated" gorm:"autoUpdateTime"`
	Created jsontype.JSONTime `json:"Created" gorm:"autoCreateTime"`
	Deleted gorm.DeletedAt    `json:"Deleted" gorm:"index"`

	Schema        string `json:"Schema" gorm:"uniqueIndex:idx_product_class_unique;uniqueIndex:idx_product_model_unique"`
	ModelName     string `json:"ModelName" gorm:"uniqueIndex:idx_product_model_unique"`
	Oui           string `json:"Oui" gorm:"uniqueIndex:idx_product_class_unique"`
	ProductClass  string `json:"ProductClass" gorm:"uniqueIndex:idx_product_class_unique"`
	ProductType   string `json:"ProductType" gorm:"index"`
	Enable        bool   `json:"Enable" gorm:"index"`
	Default       bool   `json:"Default" gorm:"index"`
	ParameterPath string `json:"ParameterPath"`

	Manufacturer string                            `json:"Manufacturer"`
	MetaDataRaw  jsontype.JSONType[*jsontype.Tags] `json:"-" gorm:"column:meta_data"`
	MetaData     *jsontype.Tags                    `json:"MetaData" gorm:"-"`
	// ParameterTypesRaw          jsontype.JSONType[*ParameterTypes] `json:"-" gorm:"column:parameter_types"`
	// ParameterTypes             *ParameterTypes                 `json:"ParameterTypes" gorm:"-"`
	PerformanceValueDefinesRaw jsontype.JSONType[*jsontype.Tags] `json:"-" gorm:"column:performance_value_defines"`
	PerformanceValueDefines    *jsontype.Tags                    `json:"PerformanceValueDefines" gorm:"-"`

	SupportedAlarms []*ProductSupportedAlarm `json:"SupportedAlarms"`
	Firmwares       []*Firmware              `json:"Firmwares" gorm:"many2many:firmware_products;"`
}

func (m *Product) SaveData() {
	m.MetaDataRaw = jsontype.NewJSONType(m.MetaData)
	// m.ParameterTypesRaw = jsontype.NewJSONType(m.ParameterTypes)
	m.PerformanceValueDefinesRaw = jsontype.NewJSONType(m.PerformanceValueDefines)
}

func (m *Product) LoadData() {
	m.MetaData = m.MetaDataRaw.Data
	// m.ParameterTypes = m.ParameterTypesRaw.Data
	m.PerformanceValueDefines = m.PerformanceValueDefinesRaw.Data
}

func (m *Product) BeforeSave(tx *gorm.DB) (err error) {
	m.SaveData()
	return nil
}

func (m *Product) AfterSave(tx *gorm.DB) error {
	go ReloadAllProducts()
	return nil
}

func (m *Product) AfterFind(tx *gorm.DB) (err error) {
	m.LoadData()
	return nil
}

func (m *Product) GetDataModel() acs.DataModel {
	return GetDataModelByProduct(m)
}

// func (m Product) GetParameterType(k string) string {
// 	if m.ParameterTypes == nil {
// 		return ""
// 	}
// 	re := utils.GetRegexp(`\.[0-9]+\.`)
// 	k = re.ReplaceAllString(k, `.{i}.`)
// 	v, ok := (*m.ParameterTypes)[k]
// 	if !ok {
// 		return ""
// 	}
// 	return v
// }

// func (m *Product) UpdateParameterType(k string, v string) {
// 	if m.ParameterTypes == nil {
// 		m.ParameterTypes = &ParameterTypes{}
// 	}
// 	re := utils.GetRegexp(`\.[0-9]+\.`)
// 	k = re.ReplaceAllString(k, `.{i}.`)
// 	if parts := strings.Split(v, ":"); len(parts) > 1 {
// 		v = parts[len(parts)-1]
// 	}
// 	m.ParameterTypes.SetValue(k, v)
// }

// func (m *Product) UpdateParameterTypes(values map[string]string) {
// 	if m.ParameterTypes == nil {
// 		m.ParameterTypes = &ParameterTypes{}
// 	}
// 	m.ParameterTypes.SetValues(values)
// }

// func (m Product) EncodeParameterName(name string) string {
// 	if m.ParameterPath == "" {
// 		return name
// 	}
// 	if strings.HasPrefix(name, m.ParameterPath) {
// 		return name
// 	}
// 	if pos := strings.Index(name, "."); pos > 0 {
// 		return m.ParameterPath + name[pos+1:]
// 	}
// 	return name
// }

// func (m Product) DecodeParameterName(name string) string {
// 	if m.ParameterPath == "" {
// 		return name
// 	}
// 	if strings.HasPrefix(name, m.ParameterPath) {
// 		if pos := strings.Index(name, "."); pos > 0 {
// 			return "Device." + name[pos+1:]
// 		}
// 	}
// 	return name
// }
