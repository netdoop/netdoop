package omc

import (
	"time"

	"github.com/netdoop/netdoop/store"
	"github.com/netdoop/netdoop/utils"

	"github.com/heypkg/store/jsontype"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

var DeviceStatusIndexNames = []string{"schema", "product_type"}

type DeviceStatus struct {
	Time        jsontype.JSONTime `json:"Time" gorm:"autoCreateTime;uniqueIndex:idx_device_status_unique;not null"`
	Schema      string            `json:"Schema" gorm:"uniqueIndex:idx_device_status_unique;not null"`
	ProductType string            `json:"ProductType" gorm:"uniqueIndex:idx_device_status_unique;not null"`

	Total   int64 `json:"Total"`
	Online  int64 `json:"Online"`
	Active  int64 `json:"Active"`
	Offline int64 `json:"Offline"`
}

func FetchDeviceStatus(t time.Time, schema string, productType string) error {
	db := store.GetDB()
	status := &DeviceStatus{
		Time:        jsontype.JSONTime(t),
		Schema:      schema,
		ProductType: productType,
	}
	db.Model(&Device{}).Where("schema = ? AND product_type = ? AND online = ?", schema, productType, true).Count(&status.Online)
	db.Model(&Device{}).Where("schema = ? AND product_type = ? AND online = ?", schema, productType, false).Count(&status.Offline)
	db.Model(&Device{}).Where("schema = ? AND product_type = ? AND online = ? AND active_status = ? ", schema, productType, true, "active").Count(&status.Active)

	status.Total = status.Online + status.Offline
	result := db.Create(status)
	if result.Error != nil {
		return errors.Wrap(result.Error, "create new device status")
	}
	return nil
}

func FetchAllDeviceStatus(t time.Time) error {
	logger := utils.GetLogger()
	if err := FetchDeviceStatus(t, "", "cpe"); err != nil {
		logger.Error("fetch device status of cpe", zap.Error(err))
	}
	if err := FetchDeviceStatus(t, "", "enb"); err != nil {
		logger.Error("fetch device status of cpe", zap.Error(err))
	}

	if err := FetchDeviceAlarmStatus(t, "", "cpe"); err != nil {
		logger.Error("fetch device alarm status of cpe", zap.Error(err))
	}
	if err := FetchDeviceAlarmStatus(t, "", "enb"); err != nil {
		logger.Error("fetch device alarm status of cpe", zap.Error(err))
	}
	return nil
}
