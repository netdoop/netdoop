package omc

import (
	"strconv"
	"strings"

	"github.com/netdoop/cwmp/acs"
	"github.com/netdoop/netdoop/store"

	"github.com/heypkg/store/jsontype"
	"github.com/heypkg/store/search"
	"gorm.io/gorm"
)

var SimpleDeviceInfoName = []string{
	"id", "updated", "created", "deleted",
	"schema", "oui", "product_class", "serial_number", "product_id",
	"name", "group_id",
	"enable", "online", "active_status", "last_inform_time",
	"meta_data", "properties",
}

var DeviceSearchHanleFuncs = search.SearchDataHandleFuncMap{
	"group": HanldeGroupSearch,
}

type Device struct {
	ID      uint              `json:"Id" gorm:"primarykey"`
	Updated jsontype.JSONTime `json:"Updated" gorm:"autoUpdateTime"`
	Created jsontype.JSONTime `json:"Created" gorm:"autoCreateTime"`
	Deleted gorm.DeletedAt    `json:"Deleted" gorm:"index"`

	Schema       string `json:"Schema" gorm:"uniqueIndex:idx_device_unique"`
	Oui          string `json:"Oui" gorm:"uniqueIndex:idx_device_unique"`
	ProductClass string `json:"ProductClass" gorm:"uniqueIndex:idx_device_unique"`
	SerialNumber string `json:"SerialNumber" gorm:"uniqueIndex:idx_device_unique"`
	ProductType  string `json:"ProductType" gorm:"index"`

	ActiveStatus   string            `json:"ActiveStatus" gorm:"index"`
	Enable         bool              `json:"Enable" gorm:"index"`
	Online         bool              `json:"Online" gorm:"index"`
	LastInformTime jsontype.JSONTime `json:"LastInformTime" gorm:"index"`

	ProductID uint     `json:"ProductId" gorm:"index;not null"`
	Product   *Product `json:"Product"`

	GroupID uint   `json:"GroupId" gorm:"index"`
	Group   *Group `json:"Group"`

	Name          string                            `json:"Name"`
	MetaDataRaw   jsontype.JSONType[*jsontype.Tags] `json:"-" gorm:"column:meta_data"`
	MetaData      *jsontype.Tags                    `json:"MetaData" gorm:"-"`
	PropertiesRaw jsontype.JSONType[*jsontype.Tags] `json:"-" gorm:"column:properties"`
	Properties    *jsontype.Tags                    `json:"Properties" gorm:"-"`

	MethodsRaw                jsontype.JSONType[*Methods]                `json:"-" gorm:"column:methods"`
	Methods                   *Methods                                   `json:"Methods" gorm:"-"`
	ParameterValuesRaw        jsontype.JSONType[*ParameterValues]        `json:"-" gorm:"column:parameter_values"`
	ParameterValues           *ParameterValues                           `json:"ParameterValues" gorm:"-"`
	ParameterWritablesRaw     jsontype.JSONType[*ParameterWritables]     `json:"-" gorm:"column:parameter_writables"`
	ParameterWritables        *ParameterWritables                        `json:"ParameterWritables" gorm:"-"`
	ParameterNotificationsRaw jsontype.JSONType[*ParameterNotifications] `json:"-" gorm:"column:parameter_notifications"`
	ParameterNotifications    *ParameterNotifications                    `json:"ParameterNotifications" gorm:"-"`

	// InternetGatewayDevice *igd.InternetGatewayDevice `json:"InternetGatewayDevice" gorm:"-"`
}

func (m *Device) SaveData() {
	m.MethodsRaw = jsontype.NewJSONType(m.Methods)
	m.ParameterValuesRaw = jsontype.NewJSONType(m.ParameterValues)
	m.ParameterWritablesRaw = jsontype.NewJSONType(m.ParameterWritables)
	m.ParameterNotificationsRaw = jsontype.NewJSONType(m.ParameterNotifications)
	m.MetaDataRaw = jsontype.NewJSONType(m.MetaData)
	m.PropertiesRaw = jsontype.NewJSONType(m.Properties)
}

func (m *Device) LoadData() {
	m.Methods = m.MethodsRaw.Data
	m.ParameterValues = m.ParameterValuesRaw.Data
	m.ParameterWritables = m.ParameterWritablesRaw.Data
	m.ParameterNotifications = m.ParameterNotificationsRaw.Data

	m.MetaData = m.MetaDataRaw.Data
	m.Properties = m.PropertiesRaw.Data
}

func (m *Device) BeforeSave(tx *gorm.DB) (err error) {
	m.SaveData()
	return nil
}

func (m *Device) AfterSave(tx *gorm.DB) (err error) {
	// m.UpdateIGD(true)
	return nil
}

func (m *Device) AfterFind(tx *gorm.DB) (err error) {
	m.LoadData()
	// m.UpdateIGD(false)
	return nil
}

// func (m *Device) UpdateIGD(needUpdate bool) {
// 	key := fmt.Sprintf("%v", m.ID)
// 	if m.ParameterValues != nil {
// 		cache := igd.GetDefaultCacheManager()
// 		m.InternetGatewayDevice = cache.Get(key)
// 		if needUpdate || m.InternetGatewayDevice == nil {
// 			cache.Update(key, m.ParameterValues.GetValues())
// 			m.InternetGatewayDevice = cache.Get(key)
// 		}
// 	}
// }

func (m *Device) GetProduct() acs.Product {
	if m.Product == nil {
		m.Product = GetProduct(m.Schema, m.Oui, m.ProductClass)
	}
	return m.Product
}

func (m Device) GetParameterValue(name string) *string {
	if m.ParameterValues == nil {
		return nil
	}
	if strings.HasPrefix(name, ".") {
		name = m.Product.ParameterPath + name[1:]
	}
	for k, v := range *m.ParameterValues {
		if k == name {
			return &v
		}
	}
	return nil
}

func (m Device) GetParameterInt64Value(name string) *int64 {
	str := m.GetParameterValue(name)
	if str == nil {
		return nil
	}
	v, err := strconv.ParseInt(*str, 10, 64)
	if err != nil {
		return nil
	}
	return &v
}

func GetDevice(schema string, oui string, serialNumber string) *Device {
	db := store.GetDB()
	var device Device
	if result := db.Where(
		"schema = ? AND oui = ? AND serial_number = ?",
		schema, oui, serialNumber,
	).First(&device); result.Error != nil {
		return nil
	}
	return &device
}

func GetDeviceWithProductClass(schema string, oui string, productClass string, serialNumber string) *Device {
	db := store.GetDB()

	var device Device
	if result := db.Where(
		"schema = ? AND oui = ? AND product_class = ? AND serial_number = ?",
		schema, oui, productClass, serialNumber,
	).First(&device); result.Error != nil {
		return nil
	}
	return &device
}
