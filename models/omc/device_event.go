package omc

import (
	"time"

	"github.com/heypkg/store/jsontype"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

var DeviceEventIndexNames = []string{"schema", "product_type", "event_type"}
var DeviceEventDeviceIndexNames = []string{"schema", "product_type", "device_id", "event_type"}

type DeviceEvent struct {
	Time     jsontype.JSONTime `json:"Time" gorm:"autoCreateTime;uniqueIndex:idx_device_event_unique;not null"`
	Schema   string            `json:"Schema" gorm:"uniqueIndex:idx_device_event_unique;not null"`
	DeviceID uint              `json:"DeviceId" gorm:"uniqueIndex:idx_device_event_unique;not null"`
	Device   *Device           `json:"Device" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Updated  jsontype.JSONTime `json:"Updated" gorm:"autoUpdateTime"`

	Oui          string `json:"Oui"`
	ProductClass string `json:"ProductClass"`
	SerialNumber string `json:"SerialNumber"`
	ProductType  string `json:"ProductType" gorm:"index"`

	EventType   string            `json:"EventType" gorm:"uniqueIndex:idx_device_event_unique;not null"`
	CurrentTime jsontype.JSONTime `json:"CurrentTime"`

	MetaDataRaw jsontype.JSONType[*jsontype.Tags] `json:"-" gorm:"column:meta_data"`
	MetaData    *jsontype.Tags                    `json:"MetaData" gorm:"-"`
}

func (m *DeviceEvent) SaveData() {
	m.MetaDataRaw = jsontype.NewJSONType(m.MetaData)
}
func (m *DeviceEvent) LoadData() {
	m.MetaData = m.MetaDataRaw.Data
}

func (m *DeviceEvent) BeforeSave(tx *gorm.DB) (err error) {
	m.SaveData()
	return nil
}

func (m *DeviceEvent) AfterFind(tx *gorm.DB) (err error) {
	m.LoadData()
	return nil
}

func InsertDevieEvent(db *gorm.DB, device *Device, eventType string, currentTime time.Time, metaData *jsontype.Tags) error {
	if eventType == "" {
		return nil
	}
	event := &DeviceEvent{
		DeviceID:     device.ID,
		Oui:          device.Oui,
		ProductClass: device.ProductClass,
		SerialNumber: device.SerialNumber,
		ProductType:  device.ProductType,
		EventType:    eventType,
		CurrentTime:  jsontype.JSONTime(currentTime),
		MetaData:     metaData,
	}
	result := db.Create(event)
	if result.Error != nil {
		return errors.Wrap(result.Error, "create new device event")
	}
	return nil
}
