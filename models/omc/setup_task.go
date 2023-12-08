package omc

import (
	"github.com/netdoop/netdoop/utils"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

func SetupTasks(db *gorm.DB, schema string) {
	logger := utils.GetLogger()

	if _, err := CreateNBITask(db, schema,
		"Alarm Post",
		"HTTP", "POST", "DeviceAlarm",
		// "http://54.246.203.22:9999/",
		"http://43.143.43.123:9176/dev/post",
		100, 0, 0,
		true, true,
	); err != nil {
		logger.Error("create nbi task", zap.Error(err), zap.String("TaskName", "Alarm Post"))
	}

	// go func() {
	// 	time.Sleep(time.Second * 10)
	// 	DebugPostAlarm(db)
	// }()
}
