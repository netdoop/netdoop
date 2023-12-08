package omc

import (
	"github.com/heypkg/store/jsontype"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

var DevicePerformanceValueIndexNames = []string{"schema", "product_type", "name", "type"}
var DevicePerformanceValueDeviceIndexNames = []string{"schema", "product_type", "device_id", "name", "type"}

type DevicePerformanceValue struct {
	Time     jsontype.JSONTime `json:"Time" gorm:"autoCreateTime;uniqueIndex:idx_device_performance_value_unique;not null"`
	Schema   string            `json:"Schema" gorm:"uniqueIndex:idx_device_performance_value_unique;not null"`
	DeviceID uint              `json:"DeviceId" gorm:"uniqueIndex:idx_device_performance_value_unique;not null"`
	Device   *Device           `json:"Device" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Updated  jsontype.JSONTime `json:"Updated" gorm:"autoUpdateTime"`

	Oui          string `json:"Oui"`
	ProductClass string `json:"ProductClass"`
	SerialNumber string `json:"SerialNumber"`
	ProductType  string `json:"ProductType" gorm:"index"`

	Name   string  `json:"Name" gorm:"uniqueIndex:idx_device_performance_value_unique;not null"`
	Source string  `json:"Source"`
	Type   string  `json:"Type"`
	Value  float64 `json:"Value"`

	MetaDataRaw jsontype.JSONType[*jsontype.Tags] `json:"-" gorm:"column:meta_data"`
	MetaData    *jsontype.Tags                    `json:"MetaData" gorm:"-"`
}

func (m *DevicePerformanceValue) SaveData() {
	m.MetaDataRaw = jsontype.NewJSONType(m.MetaData)
}
func (m *DevicePerformanceValue) LoadData() {
	m.MetaData = m.MetaDataRaw.Data
}

func (m *DevicePerformanceValue) BeforeSave(tx *gorm.DB) (err error) {
	m.SaveData()
	return nil
}

func (m *DevicePerformanceValue) AfterFind(tx *gorm.DB) (err error) {
	m.LoadData()
	return nil
}

func InsertDeviePerformanceValue(db *gorm.DB, device *Device, name string, source string, value float64, metaData *jsontype.Tags) error {
	if name == "" {
		return nil
	}
	data := &DevicePerformanceValue{
		DeviceID:     device.ID,
		Oui:          device.Oui,
		ProductClass: device.ProductClass,
		SerialNumber: device.SerialNumber,
		ProductType:  device.ProductType,
		Name:         name,
		Source:       source,
		Value:        value,
		MetaData:     metaData,
	}
	result := db.Create(data)
	if result.Error != nil {
		return errors.Wrap(result.Error, "create new device performance value")
	}
	return nil
}
