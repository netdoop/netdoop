package omc

import (
	"time"

	"github.com/netdoop/netdoop/store"
	"github.com/netdoop/netdoop/utils"

	"github.com/heypkg/store/jsontype"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

var DeviceAlarmStatusIndexNames = []string{"schema", "product_type"}

type DeviceAlarmStatus struct {
	Time        jsontype.JSONTime `json:"Time" gorm:"autoCreateTime;uniqueIndex:idx_device_status_unique;not null"`
	Schema      string            `json:"Schema" gorm:"uniqueIndex:idx_device_status_unique;not null"`
	ProductType string            `json:"ProductType" gorm:"uniqueIndex:idx_device_status_unique;not null"`

	Total    int64 `json:"Total"`
	Critical int64 `json:"Critical"`
	Major    int64 `json:"Major"`
	Minor    int64 `json:"Minor"`
	Warning  int64 `json:"Warning"`
}

func FetchDeviceAlarmStatus(t time.Time, schema string, productType string) error {
	db := store.GetDB()
	status := &DeviceAlarmStatus{
		Time:        jsontype.JSONTime(t),
		Schema:      schema,
		ProductType: productType,
	}
	db.Model(&DeviceAlarm{}).Where(
		"schema = ? AND product_type = ? AND alarm_cleared = ? AND perceived_severity = ?",
		schema, productType, false, "Critical").Count(&status.Critical)
	db.Model(&DeviceAlarm{}).Where(
		"schema = ? AND product_type = ? AND alarm_cleared = ? AND perceived_severity = ?",
		schema, productType, false, "Major").Count(&status.Major)
	db.Model(&DeviceAlarm{}).Where(
		"schema = ? AND product_type = ? AND alarm_cleared = ? AND perceived_severity = ?",
		schema, productType, false, "Minor").Count(&status.Minor)
	db.Model(&DeviceAlarm{}).Where(
		"schema = ? AND product_type = ? AND alarm_cleared = ? AND perceived_severity = ?",
		schema, productType, false, "Warning").Count(&status.Warning)
	status.Total = status.Warning + status.Minor + status.Major + status.Critical
	result := db.Create(status)
	if result.Error != nil {
		return errors.Wrap(result.Error, "create new device alarm status")
	}
	return nil
}

func FetchAllDeviceAlarmStatus(t time.Time) error {
	logger := utils.GetLogger()
	if err := FetchDeviceAlarmStatus(t, "", "cpe"); err != nil {
		logger.Error("fetch device status of cpe", zap.Error(err))
	}
	if err := FetchDeviceAlarmStatus(t, "", "enb"); err != nil {
		logger.Error("fetch device status of cpe", zap.Error(err))
	}
	return nil
}
