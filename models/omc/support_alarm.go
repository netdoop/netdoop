package omc

import (
	"fmt"

	"github.com/netdoop/netdoop/models/omc/define/igd"

	"github.com/heypkg/store/jsontype"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type ProductSupportedAlarm struct {
	ID      uint              `json:"Id" gorm:"primarykey"`
	Updated jsontype.JSONTime `json:"Updated" gorm:"autoUpdateTime"`
	Created jsontype.JSONTime `json:"Created" gorm:"autoCreateTime"`

	Schema    string   `json:"Schema" gorm:"uniqueIndex:idx_product_support_alarm_unique;not null"`
	ProductID uint     `json:"ProductId" gorm:"uniqueIndex:idx_product_support_alarm_unique;not null"`
	Product   *Product `json:"Product"`

	Oui          string `json:"Oui"`
	ProductClass string `json:"ProductClass"`
	SerialNumber string `json:"SerialNumber"`
	ProductType  string `json:"ProductType" gorm:"index"`

	AlarmIdentifier   string `json:"AlarmIdentifier" gorm:"uniqueIndex:idx_product_support_alarm_unique;not null"`
	EventType         string `json:"EventType"`
	ProbableCause     string `json:"ProbableCause"`
	SpecificProblem   string `json:"SpecificProblem"`
	PerceivedSeverity string `json:"PerceivedSeverity"`
}

func getSupportedAlarm(db *gorm.DB, product *Product, alarmIdentifier string) *ProductSupportedAlarm {
	var obj ProductSupportedAlarm
	if result := db.Where(
		"product_id = ? AND alarm_identifier = ?",
		product.ID,
		alarmIdentifier,
	).First(&obj); result.Error != nil {
		return nil
	}
	return &obj
}

func UpsertSupportedAlarm(db *gorm.DB, product *Product, alarmIdentifier int, data *igd.SupportedAlarm) error {
	if data == nil || alarmIdentifier < 1 {
		return nil
	}
	alarm := getSupportedAlarm(db, product, *data.EventType)
	if alarm == nil {
		alarm = &ProductSupportedAlarm{
			ProductID:    product.ID,
			Oui:          product.Oui,
			ProductClass: product.ProductClass,
			ProductType:  product.ProductType,
		}
		alarm.AlarmIdentifier = fmt.Sprintf("%d", alarmIdentifier)
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
		result := db.Create(alarm)
		if result.Error != nil {
			return errors.Wrap(result.Error, "create new supported alarm")
		}
	} else {
		updateColumns := []string{"event_type", "probable_cause", "specific_problem", "perceived_severity"}
		updateData := &DeviceAlarm{}
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
		result := db.Model(alarm).Where(
			"product_id = ? AND alarm_identifier = ?",
			product.ID,
			alarm.AlarmIdentifier,
		).Select(updateColumns).Updates(updateData)
		if result.Error != nil {
			return errors.Wrap(result.Error, "update supported alarm")
		}
	}
	return nil
}
