package omc

import (
	"time"

	"github.com/netdoop/netdoop/models/omc/define/igd"

	"github.com/heypkg/store/jsontype"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

var DeviceAlarmDeviceIndexNames = []string{"schema", "product_type", "device_id", "alarm_identifier"}
var DeviceAlarmTypeIndexNames = []string{"schema", "product_type", "alarm_identifier", "event_type"}
var DeviceAlarmPerceivedSeverityIndexNames = []string{"schema", "perceived_severity"}

type DeviceAlarm struct {
	Time     jsontype.JSONTime `json:"Time" gorm:"autoCreateTime;uniqueIndex:idx_device_alarm_unique;not null"`
	Schema   string            `json:"Schema" gorm:"uniqueIndex:idx_device_alarm_unique;not null"`
	DeviceID uint              `json:"DeviceId" gorm:"uniqueIndex:idx_device_alarm_unique;not null"`
	Device   *Device           `json:"Device" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Updated  jsontype.JSONTime `json:"Updated" gorm:"autoUpdateTime"`

	Oui          string `json:"Oui"`
	ProductClass string `json:"ProductClass"`
	SerialNumber string `json:"SerialNumber"`
	ProductType  string `json:"ProductType" gorm:"index"`

	AlarmIdentifier       string            `json:"AlarmIdentifier" gorm:"uniqueIndex:idx_device_alarm_unique;not null"`
	AlarmRaisedTime       jsontype.JSONTime `json:"AlarmRaisedTime"`
	AlarmChangedTime      jsontype.JSONTime `json:"AlarmChangedTime"`
	AlarmClearedTime      jsontype.JSONTime `json:"AlarmClearedTime"`
	ManagedObjectInstance string            `json:"ManagedObjectInstance"`
	EventType             string            `json:"EventType"`
	ProbableCause         string            `json:"ProbableCause"`
	SpecificProblem       string            `json:"SpecificProblem"`
	PerceivedSeverity     string            `json:"PerceivedSeverity"`
	AdditionalText        string            `json:"AdditionalText"`
	AdditionalInformation string            `json:"AdditionalInformation"`

	AlarmCleared       bool              `json:"AlarmCleared"`
	AlarmConfirmed     bool              `json:"AlarmConfirmed"`
	AlarmConfirmedTime jsontype.JSONTime `json:"AlarmConfirmedTime"`
}

func getAlarm(db *gorm.DB, device *Device, alarmIdentifier string) *DeviceAlarm {
	var obj DeviceAlarm
	if result := db.Where(
		"device_id = ? AND alarm_identifier = ?",
		device.ID,
		alarmIdentifier,
	).First(&obj); result.Error != nil {
		// utils.GetLogger().Error("get active alarm", zap.Error(result.Error), zap.String("alarm_identifier", alarmIdentifier))
		return nil
	}
	return &obj
}

func getActiveAlarms(db *gorm.DB, device *Device) ([]*DeviceAlarm, error) {
	objs := []*DeviceAlarm{}
	if result := db.Where(
		"device_id = ? AND alarm_cleared = ?",
		device.ID,
		false,
	).Find(&objs); result.Error != nil {
		return objs, result.Error
	}
	return objs, nil
}

func postAlarm(db *gorm.DB, device *Device, alarmIdentifier string) {
	alarm := getAlarm(db, device, alarmIdentifier)
	if alarm != nil {
		tsrv := GetTaskServer()
		tsrv.PushDeviceAlarm(alarm)
	}
}

func DebugPostAlarm(db *gorm.DB, cleared bool, alarmIdentifier string) {
	device := GetDevice("", "20AC9C", "CTBU022D200052-NR")
	if device != nil {
		objs := []*DeviceAlarm{}
		db.Where("device_id = ?", device.ID).Find(&objs)
		if len(objs) > 0 {
			obj := objs[0]
			alarm := getAlarm(db, device, obj.AlarmIdentifier)
			if alarm != nil {
				now := jsontype.JSONTime(time.Now())
				if cleared {
					alarm.AlarmIdentifier = alarmIdentifier
					alarm.AlarmCleared = true
					alarm.AlarmChangedTime = now
					alarm.AlarmClearedTime = now
				} else {
					alarm.AlarmIdentifier = alarmIdentifier
					alarm.AlarmCleared = false
					alarm.AlarmRaisedTime = now
					alarm.AlarmChangedTime = now
				}
				tsrv := GetTaskServer()
				tsrv.PushDeviceAlarm(alarm)
			}
		}
		// for _, obj := range objs {
		// 	postAlarm(db, device, obj.AlarmIdentifier)
		// }
	}
}

func ClearCurrentAlarm(db *gorm.DB, device *Device, alarmIdentifier string) error {
	alarm := getAlarm(db, device, alarmIdentifier)
	if alarm != nil && !alarm.AlarmCleared {
		updateColumns := []string{"alarm_cleared_time", "alarm_cleared"}
		updateData := &DeviceAlarm{}
		updateData.AlarmClearedTime = jsontype.JSONTime(time.Now())
		updateData.AlarmCleared = true
		result := db.Model(alarm).Where(
			"device_id = ? AND alarm_identifier = ? AND alarm_cleared = ?",
			device.ID,
			alarmIdentifier,
			false,
		).Select(updateColumns).Updates(updateData)
		if result.Error != nil {
			return errors.Wrap(result.Error, "clear alarm")
		}
		postAlarm(db, device, alarm.AlarmIdentifier)
	}
	return nil
}

func UpsertCurrentAlarm(db *gorm.DB, device *Device, data *igd.CurrentAlarm) error {
	if data == nil || data.AlarmIdentifier == nil {
		return nil
	}
	alarmIdentifier := *data.AlarmIdentifier
	alarm := getAlarm(db, device, alarmIdentifier)
	if alarm == nil {
		alarm = &DeviceAlarm{
			DeviceID:     device.ID,
			Oui:          device.Oui,
			ProductClass: device.ProductClass,
			SerialNumber: device.SerialNumber,
			ProductType:  device.ProductType,
		}
		alarm.AlarmIdentifier = alarmIdentifier
		if data.AlarmRaisedTime != nil {
			alarm.AlarmRaisedTime = jsontype.JSONTime(ParseTime(*data.AlarmRaisedTime))
		}
		if data.AlarmChangedTime != nil {
			alarm.AlarmChangedTime = jsontype.JSONTime(ParseTime(*data.AlarmChangedTime))
		}
		if data.ManagedObjectInstance != nil {
			alarm.ManagedObjectInstance = *data.ManagedObjectInstance
		}
		if data.EventType != nil {
			alarm.EventType = *data.EventType
		}
		if data.ProbableCause != nil {
			alarm.ProbableCause = *data.ProbableCause
		}
		if data.SpecificProblem != nil {
			alarm.SpecificProblem = *data.SpecificProblem
		}
		if data.PerceivedSeverity != nil {
			alarm.PerceivedSeverity = *data.PerceivedSeverity
		}
		if data.AdditionalText != nil {
			alarm.AdditionalText = *data.AdditionalText
		}
		if data.AdditionalInformation != nil {
			alarm.AdditionalInformation = *data.AdditionalInformation
		}
		result := db.Create(alarm)
		if result.Error != nil {
			return errors.Wrap(result.Error, "create new alarm")
		}
	} else {
		updateColumns := []string{"alarm_changed_time"}
		updateData := &DeviceAlarm{}
		if data.AlarmChangedTime != nil {
			updateData.AlarmChangedTime = jsontype.JSONTime(ParseTime(*data.AlarmChangedTime))
		}
		result := db.Model(alarm).Where(
			"device_id = ? AND alarm_identifier = ? AND alarm_cleared = ?",
			device.ID,
			alarm.AlarmIdentifier,
			false,
		).Select(updateColumns).Updates(updateData)
		if result.Error != nil {
			return errors.Wrap(result.Error, "update alarm")
		}
	}
	postAlarm(db, device, alarmIdentifier)
	return nil
}

func UpsertExpeditedEvent(db *gorm.DB, device *Device, data *igd.ExpeditedEvent) error {
	if data == nil || data.AlarmIdentifier == nil {
		return nil
	}
	alarmIdentifier := *data.AlarmIdentifier
	alarm := getAlarm(db, device, alarmIdentifier)
	if alarm == nil {
		alarm = &DeviceAlarm{
			DeviceID:     device.ID,
			Oui:          device.Oui,
			ProductClass: device.ProductClass,
			SerialNumber: device.SerialNumber,
			ProductType:  device.ProductType,
		}
		alarm.AlarmIdentifier = alarmIdentifier
		if data.NotificationType != nil {
			switch *data.NotificationType {
			case "NewAlarm":
				if data.EventTime != nil {
					alarm.AlarmRaisedTime = jsontype.JSONTime(ParseTime(*data.EventTime))
					alarm.AlarmChangedTime = jsontype.JSONTime(ParseTime(*data.EventTime))
				}
			case "ChangedAlarm":
				if data.EventTime != nil {
					alarm.AlarmRaisedTime = jsontype.JSONTime(ParseTime(*data.EventTime))
					alarm.AlarmChangedTime = jsontype.JSONTime(ParseTime(*data.EventTime))
				}
			case "ClearedAlarm":
				if data.EventTime != nil {
					alarm.AlarmRaisedTime = jsontype.JSONTime(ParseTime(*data.EventTime))
					alarm.AlarmChangedTime = jsontype.JSONTime(ParseTime(*data.EventTime))
					alarm.AlarmClearedTime = jsontype.JSONTime(ParseTime(*data.EventTime))
				}
				alarm.AlarmCleared = true
			}
		}

		if data.ManagedObjectInstance != nil {
			alarm.ManagedObjectInstance = *data.ManagedObjectInstance
		}
		if data.EventType != nil {
			alarm.EventType = *data.EventType
		}
		if data.ProbableCause != nil {
			alarm.ProbableCause = *data.ProbableCause
		}
		if data.SpecificProblem != nil {
			alarm.SpecificProblem = *data.SpecificProblem
		}
		if data.PerceivedSeverity != nil {
			alarm.PerceivedSeverity = *data.PerceivedSeverity
		}
		if data.AdditionalText != nil {
			alarm.AdditionalText = *data.AdditionalText
		}
		if data.AdditionalInformation != nil {
			alarm.AdditionalInformation = *data.AdditionalInformation
		}
		result := db.Create(alarm)
		if result.Error != nil {
			return errors.Wrap(result.Error, "create new alarm")
		}
	} else if !alarm.AlarmCleared {
		updateColumns := []string{"alarm_cleared"}
		updateData := &DeviceAlarm{}
		if data.NotificationType != nil {
			switch *data.NotificationType {
			case "NewAlarm":
				if data.EventTime != nil {
					updateData.AlarmRaisedTime = jsontype.JSONTime(ParseTime(*data.EventTime))
					updateColumns = append(updateColumns, "alarm_raised_time")
				}
				updateData.AlarmCleared = false
			case "ChangedAlarm":
				if data.EventTime != nil {
					updateData.AlarmChangedTime = jsontype.JSONTime(ParseTime(*data.EventTime))
					updateColumns = append(updateColumns, "alarm_changed_time")
				}
				updateData.AlarmCleared = false
			case "ClearedAlarm":
				if data.EventTime != nil {
					updateData.AlarmClearedTime = jsontype.JSONTime(ParseTime(*data.EventTime))
					updateColumns = append(updateColumns, "alarm_cleared_time")
				}
				updateData.AlarmCleared = true
			}
		}
		if data.PerceivedSeverity != nil {
			if *(data.PerceivedSeverity) == "Cleared" {
				updateData.AlarmCleared = true
			}
		}
		result := db.Model(alarm).Where(
			"device_id = ? AND alarm_identifier = ? AND alarm_cleared = ?",
			device.ID,
			alarm.AlarmIdentifier,
			false,
		).Select(updateColumns).Updates(updateData)
		if result.Error != nil {
			return errors.Wrap(result.Error, "update alarm")
		}
	}
	postAlarm(db, device, alarmIdentifier)
	return nil
}
