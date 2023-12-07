package omc

import (
	"github.com/heypkg/store/jsontype"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type TaskStatus int
type TaskExecMode int

const (
	TaskStatusInit    TaskStatus = 0
	TaskStatusWaiting TaskStatus = 1
	TaskStatusRunning TaskStatus = 2
	TaskStatusFinish  TaskStatus = 3
	TaskStatusCancel  TaskStatus = 4
)

const (
	TaskExecModeImmediately TaskExecMode = 0
	TaskExecModeSchedule    TaskExecMode = 1
)

type Task struct {
	ID      uint              `json:"Id" gorm:"primarykey"`
	Updated jsontype.JSONTime `json:"Updated" gorm:"autoUpdateTime"`
	Created jsontype.JSONTime `json:"Created" gorm:"autoCreateTime"`
	Deleted gorm.DeletedAt    `json:"Deleted" gorm:"index"`

	Schema   string `json:"Schema" gorm:"uniqueIndex:idx_task_unique"`
	TaskName string `json:"TaskName" gorm:"uniqueIndex:idx_task_unique"`
	TaskType string `json:"TaskType" gorm:"index"`
	Enable   bool   `json:"Enable" gorm:"index"`
	Default  bool   `json:"Default" gorm:"index"`

	ExecMode      TaskExecMode                      `json:"ExecMode"`
	ExecStartTime jsontype.JSONTime                 `json:"ExecStartTime"`
	ExecEndTime   jsontype.JSONTime                 `json:"ExecEndTime"`
	ExecInterval  int                               `json:"ExecInterval"`
	ExecTimes     int                               `json:"ExecTimes"`
	ExecRate      int                               `json:"ExecRate"`
	RetryInterval int                               `json:"RetryInterval"`
	RetryTimes    int                               `json:"RetryTimes"`
	RangesRaw     jsontype.JSONType[*jsontype.Tags] `json:"-" gorm:"column:ranges"`
	Ranges        *jsontype.Tags                    `json:"DeviceIds" gorm:"-"`

	Creater     string                            `json:"Creater"`
	MetaDataRaw jsontype.JSONType[*jsontype.Tags] `json:"-" gorm:"column:meta_data"`
	MetaData    *jsontype.Tags                    `json:"MetaData" gorm:"-"`

	TaskStatus   TaskStatus        `json:"TaskStatus" gorm:"index"`
	ExecProcess  int               `json:"ExecProcess"`
	ExecLastTime jsontype.JSONTime `json:"ExecLastTime"`
	ExecNextTime jsontype.JSONTime `json:"ExecNextTime"`
}

func (m *Task) BeforeSave(tx *gorm.DB) (err error) {
	if m.MetaData != nil {
		m.MetaDataRaw = jsontype.NewJSONType(m.MetaData)
	}
	if m.Ranges != nil {
		m.RangesRaw = jsontype.NewJSONType(m.Ranges)
	}
	if m.ExecInterval <= 0 {
		m.ExecInterval = 1
	}
	if m.ExecRate <= 0 {
		m.ExecRate = 1000
	}
	return nil
}

func (m *Task) AfterFind(tx *gorm.DB) (err error) {
	m.MetaData = m.MetaDataRaw.Data
	m.Ranges = m.RangesRaw.Data
	return nil
}

var TaskDeviceLogIndexNames = []string{"schema", "device_id", "task_id", "task_name", "task_type", "exec_mode"}

type TaskDeviceLog struct {
	Time     jsontype.JSONTime `json:"Time" gorm:"autoCreateTime;uniqueIndex:idx_task_log_unique;not null"`
	Schema   string            `json:"Schema" gorm:"uniqueIndex:idx_task_log_unique;not null"`
	TaskID   uint              `json:"TaskId" gorm:"uniqueIndex:idx_task_log_unique;not null"`
	Task     *Task             `json:"Task" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	DeviceID uint              `json:"DeviceId" gorm:"uniqueIndex:idx_task_log_unique;not null"`
	Device   *Device           `json:"Device" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Updated  jsontype.JSONTime `json:"Updated" gorm:"autoUpdateTime"`

	TaskName string `json:"TaskName"`
	TaskType string `json:"TaskType"`
	ExecMode int    `json:"ExecMode"`

	StartTime jsontype.JSONTime                 `json:"StartTime"`
	EndTime   jsontype.JSONTime                 `json:"EndTime"`
	Code      int                               `json:"Code"`
	Info      string                            `json:"Info"`
	RangesRaw jsontype.JSONType[*jsontype.Tags] `json:"-" gorm:"column:ranges"`
	Ranges    *jsontype.Tags                    `json:"DeviceIds" gorm:"-"`
}

func (m *TaskDeviceLog) BeforeSave(tx *gorm.DB) (err error) {
	if m.Ranges != nil {
		m.RangesRaw = jsontype.NewJSONType(m.Ranges)
	}
	return nil
}

func (m *TaskDeviceLog) AfterFind(tx *gorm.DB) (err error) {
	m.Ranges = m.RangesRaw.Data
	return nil
}

func LoadAllTasks(db *gorm.DB, schema string) ([]*Task, error) {
	all := []*Task{}
	if result := db.Model(&Task{}).Find(&all); result.Error != nil {
		return nil, result.Error
	}
	return all, nil
}

func UpdateTaskStatus(db *gorm.DB, schema string, task *Task, status TaskStatus) error {
	updateColumns := []string{"task_status"}
	updateData := &Task{
		TaskStatus: status,
	}
	result := db.Model(task).Select(updateColumns).Updates(updateData)
	if result.Error != nil {
		return errors.Wrap(result.Error, "update task status of task")
	}
	return nil
}

func CreateNBITask(db *gorm.DB,
	schema string,
	taskName string,
	protocol string, method string, source string,
	serverUrl string,
	execRate int, retryInterval int, retryTimes int,
	enable bool,
	isDefault bool,
) (*Task, error) {
	task := &Task{
		Schema:   schema,
		TaskName: taskName,
		TaskType: "NBI:HTTP:POST:DeviceAlarm",
		MetaData: &jsontype.Tags{
			"Protocol":  protocol,
			"Method":    method,
			"Source":    source,
			"ServerUrl": serverUrl,
		},
		ExecMode:      TaskExecModeImmediately,
		ExecRate:      execRate,
		RetryInterval: retryInterval,
		RetryTimes:    retryTimes,
		Enable:        enable,
		Default:       isDefault,
	}
	if err := db.Where("schema = ? AND task_name = ? ",
		schema, taskName,
	).FirstOrCreate(task).Error; err != nil {
		return nil, errors.Wrap(err, "create nbi task")
	}
	return task, nil
}
