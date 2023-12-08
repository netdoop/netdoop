package omc

import (
	"time"

	"github.com/netdoop/netdoop/store"
	"github.com/netdoop/netdoop/utils"

	"github.com/heypkg/s3"
	"github.com/heypkg/store/tsdb"
)

func Setup() {
	env := utils.GetEnv()
	duration := time.Duration(time.Second * env.GetDuration("data_retention_period"))
	secret := utils.GetEnv().GetString("secret")
	db := store.GetDB()
	mdb := store.GetMongoDatabase()

	s3.SetupMongoStorage(db, mdb, secret)

	db.AutoMigrate(&DataModel{})
	db.AutoMigrate(&DataModelParameter{})
	db.AutoMigrate(&DataModelTemplate{})

	db.AutoMigrate(&Group{})
	db.AutoMigrate(&Product{})
	db.AutoMigrate(&ProductSupportedAlarm{})
	db.AutoMigrate(&KPIMeas{})
	SetupDefaultKPIMeasures(db, "")
	db.AutoMigrate(&KPITemplate{})
	SetupDefaultKPITemplates(db, "")

	db.AutoMigrate(&Firmware{})
	db.AutoMigrate(&Device{})

	db.AutoMigrate(&DeviceMethodCall{})
	tsdb.CreateHyperTable(db, "device_method_calls", duration)
	db.AutoMigrate(&DeviceTransferLog{})
	tsdb.CreateHyperTable(db, "device_transfer_logs", duration)

	db.AutoMigrate(&DeviceStatus{})
	tsdb.CreateHyperTable(db, "device_statuses", duration)
	tsdb.CreateHyperTableCountView(db, "device_statuses", "device_statuses_hourly_view", "1h", DeviceStatusIndexNames)
	tsdb.CreateHyperTableCountView(db, "device_statuses", "device_statuses_daily_view", "1d", DeviceStatusIndexNames)

	db.AutoMigrate(&DeviceAlarm{})
	tsdb.CreateHyperTable(db, "device_alarms", duration)
	tsdb.CreateHyperTableCountView(db, "device_alarms", "device_alarms_device_hourly_view", "1h", DeviceAlarmDeviceIndexNames)
	tsdb.CreateHyperTableCountView(db, "device_alarms", "device_alarms_device_daily_view", "1d", DeviceAlarmDeviceIndexNames)
	tsdb.CreateHyperTableCountView(db, "device_alarms", "device_alarms_type_hourly_view", "1h", DeviceAlarmTypeIndexNames)
	tsdb.CreateHyperTableCountView(db, "device_alarms", "device_alarms_type_daily_view", "1d", DeviceAlarmTypeIndexNames)
	tsdb.CreateHyperTableCountView(db, "device_alarms", "device_alarms_severity_hourly_view", "1h", DeviceAlarmPerceivedSeverityIndexNames)
	tsdb.CreateHyperTableCountView(db, "device_alarms", "device_alarms_severity_daily_view", "1d", DeviceAlarmPerceivedSeverityIndexNames)

	db.AutoMigrate(&DeviceAlarmStatus{})
	tsdb.CreateHyperTable(db, "device_alarm_statuses", duration)
	tsdb.CreateHyperTableCountView(db, "device_alarm_statuses", "device_alarm_statuses_hourly_view", "1h", DeviceAlarmStatusIndexNames)
	tsdb.CreateHyperTableCountView(db, "device_alarm_statuses", "device_alarm_statuses_daily_view", "1d", DeviceAlarmStatusIndexNames)

	db.AutoMigrate(&DeviceEvent{})
	tsdb.CreateHyperTable(db, "device_events", duration)
	tsdb.CreateHyperTableCountView(db, "device_events", "device_events_hourly_view", "1h", DeviceEventIndexNames)
	tsdb.CreateHyperTableCountView(db, "device_events", "device_events_daily_view", "1d", DeviceEventIndexNames)
	tsdb.CreateHyperTableCountView(db, "device_events", "device_events_device_hourly_view", "1h", DeviceEventDeviceIndexNames)
	tsdb.CreateHyperTableCountView(db, "device_events", "device_events_device_daily_view", "1d", DeviceEventDeviceIndexNames)

	db.AutoMigrate(&DevicePerformanceValue{})
	tsdb.CreateHyperTable(db, "device_performance_values", duration)
	tsdb.CreateHyperTableCountView(db, "device_performance_values", "device_performance_values_hourly_view", "1h", DevicePerformanceValueIndexNames)
	tsdb.CreateHyperTableCountView(db, "device_performance_values", "device_performance_values_daily_view", "1d", DevicePerformanceValueIndexNames)
	tsdb.CreateHyperTableCountView(db, "device_performance_values", "device_performance_values_device_hourly_view", "1h", DevicePerformanceValueDeviceIndexNames)
	tsdb.CreateHyperTableCountView(db, "device_performance_values", "device_performance_values_device_daily_view", "1d", DevicePerformanceValueDeviceIndexNames)

	db.AutoMigrate(&Task{})
	db.AutoMigrate(&TaskDeviceLog{})
	tsdb.CreateHyperTable(db, "task_device_logs", duration)

}
