package omc

import (
	"fmt"
	"time"

	"github.com/netdoop/netdoop/store"
	"github.com/netdoop/netdoop/utils"
	"go.uber.org/zap"
)

func StartAllActiveDeviceKeepaliveTimers() {
	db := store.GetDB()
	all := []*Device{}
	db.Model(&Device{}).Where("online = ?", true).Find(&all)
	for _, device := range all {
		device.Product = GetProduct(device.Schema, device.Oui, device.ProductClass)
		StartDeviceKeepaliveTimer(device)
	}
}

func StopDeviceKeepaliveTimer(device *Device) {
	key := fmt.Sprintf("deivce:%v:alive", device.ID)
	utils.GetTimingServer().StopTimer(key)
}

func StartDeviceKeepaliveTimer(device *Device) {
	duration := time.Hour
	if v := device.GetParameterInt64Value(".ManagementServer.PeriodicInformInterval"); v != nil {
		duration = time.Duration(*v) * time.Second * 3
	}
	now := time.Now()
	last := time.Time(device.LastInformTime)
	if last.Unix() < 0 {
		last = time.Unix(0, 0)
	}
	if diff := now.Sub(last); diff > 0 && diff < duration {
		duration = duration - diff
	}
	key := fmt.Sprintf("deivce:%v:alive", device.ID)
	utils.GetTimingServer().StartTimer(key, duration, func() {
		HandleDeviceTimeout(device, time.Now())
	})
}

func HandleDeviceAlive(device *Device, t time.Time, lastOnlineStatus bool) {
	db := store.GetDB()
	logger := utils.GetLogger()

	StartDeviceKeepaliveTimer(device)
	if err := ClearCurrentAlarm(db, device, "-1"); err != nil {
		logger.Error("clear device lost alarm", zap.Error(err))
	}
	if !lastOnlineStatus {
		if err := InsertDevieEvent(db, device, "Online", t, nil); err != nil {
			logger.Error("insert device online event", zap.Error(err))
		}
	}
}

func HandleDeviceTimeout(device *Device, t time.Time) {
	db := store.GetDB()
	logger := utils.GetLogger()

	StopDeviceKeepaliveTimer(device)

	result := db.Model(device).Select("online", "active_status").Updates(Device{Online: false, ActiveStatus: ""})
	if result.Error != nil {
		logger.Error("update device online status", zap.Error(result.Error))
	}
	if err := InsertDevieEvent(db, device, "Offline", t, nil); err != nil {
		logger.Error("insert device offline event", zap.Error(err))
	}
	// alarmIdentifier := "-1"
	// alarmRaiseTime := FormateTime(t)
	// eventType := "LostConnection"
	// probableCause := "InformTimeout"
	// perceivedSeverity := "Critical"
	// alarm := &igd.CurrentAlarm{
	// 	AlarmIdentifier:   &alarmIdentifier,
	// 	PerceivedSeverity: &perceivedSeverity,
	// 	AlarmRaisedTime:   &alarmRaiseTime,
	// 	AlarmChangedTime:  &alarmRaiseTime,
	// 	EventType:         &eventType,
	// 	ProbableCause:     &probableCause,
	// }
	// if err := UpsertCurrentAlarm(db, device, alarm); err != nil {
	// 	logger.Error("upsert device lost alarm", zap.Error(err))
	// }
}
