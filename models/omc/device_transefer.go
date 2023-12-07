package omc

import (
	"time"

	"github.com/netdoop/netdoop/utils"

	"github.com/heypkg/s3"
	"github.com/heypkg/store/jsontype"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

const DeviceUploadBucket = "omc-upload"

const (
	TransferTypeUpload   = "upload"
	TransferTypeDownload = "download"

	UploadFileTypeConfigurationFile = "1 Vendor Configuration File"
	UploadFileTypeLogFile           = "2 Vendor Log File"

	DownloadFileTypeFirmwareUpgradeImage    = "1 Firmware Upgrade Image"
	DownloadFileTypeWebContent              = "2 Web Content"
	DownloadFileTypeVendorConfigurationFile = "3 Vendor Configuration File"
)

type DeviceTransferLog struct {
	Time     jsontype.JSONTime `json:"Time" gorm:"autoCreateTime;uniqueIndex:idx_device_log_file_unique;not null"`
	Schema   string            `json:"Schema" gorm:"uniqueIndex:idx_device_log_file_unique;not null"`
	DeviceID uint              `json:"DeviceId" gorm:"uniqueIndex:idx_device_log_file_unique;not null"`
	Device   *Device           `json:"Device" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Updated  jsontype.JSONTime `json:"Updated" gorm:"autoUpdateTime"`

	Oui          string `json:"Oui"`
	ProductClass string `json:"ProductClass"`
	SerialNumber string `json:"SerialNumber"`
	ProductType  string `json:"ProductType" gorm:"index"`

	TransferType string            `json:"TransferType" gorm:"uniqueIndex:idx_device_log_file_unique"`
	FileType     string            `json:"FileType" gorm:"uniqueIndex:idx_device_log_file_unique"`
	FileName     string            `json:"FileName" gorm:"uniqueIndex:idx_device_log_file_unique"`
	StartTime    jsontype.JSONTime `json:"StartTime"`
	CompleteTime jsontype.JSONTime `json:"CompleteTime"`
	FaultCode    int               `json:"FaultCode"`
	FaultString  string            `json:"FaultString"`
	ObjectBucket string            `json:"ObjectBucket"`
	ObjectKey    string            `json:"ObjectKey"`

	S3ObjectID *uint        `json:"S3ObjectId" gorm:"index"`
	S3Object   *s3.S3Object `json:"S3Object"`

	FirmwareID *uint     `json:"FirmwareId" gorm:"index"`
	Firmware   *Firmware `json:"Firmware"`
}

func (m *DeviceTransferLog) AfterDelete(tx *gorm.DB) (err error) {
	if m.TransferType == TransferTypeUpload && m.S3ObjectID != nil {
		s3.RemoveObjectById(m.Schema, *m.S3ObjectID)
	}
	return nil
}

func GetDeviceTransferLog(db *gorm.DB, device *Device, ts int64) *DeviceTransferLog {
	var log DeviceTransferLog
	t := time.Unix(0, ts)
	if result := db.Where("device_id = ? AND time = ?", device.ID, jsontype.JSONTime(t)).First(&log); result.Error != nil {
		return nil
	}
	return &log
}

func InsertDeviceTransferLogComplete(db *gorm.DB, device *Device, ts int64,
	bucket string,
	key string,
	fileType string,
	fileName string,
	startTime time.Time,
	completeTime time.Time,
	faultCode int,
	faultString string,
) error {
	obj, err := s3.GetObject(device.Schema, DeviceUploadBucket, key)
	if err != nil {
		return errors.Wrap(err, "get s3 object")
	} else if obj == nil {
		return errors.New("invalid s3 object")
	}

	if ts < 0 {
		ts = time.Now().UnixNano()
	}
	t := time.Unix(0, ts)

	log := DeviceTransferLog{
		Time:         jsontype.JSONTime(t),
		Oui:          device.Oui,
		ProductClass: device.ProductClass,
		SerialNumber: device.SerialNumber,
		ProductType:  device.ProductType,
		DeviceID:     device.ID,
		Schema:       device.Schema,
		TransferType: TransferTypeUpload,
		FileType:     fileType,
		FileName:     fileName,
		ObjectBucket: bucket,
		ObjectKey:    key,
		S3ObjectID:   &obj.ID,
		StartTime:    jsontype.JSONTime(startTime),
		CompleteTime: jsontype.JSONTime(completeTime),
		FaultCode:    faultCode,
		FaultString:  faultString,
	}
	result := db.Create(&log)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func UpdateDeviceTransferLogComplete(db *gorm.DB, device *Device, ts int64,
	startTime time.Time, completeTime time.Time) error {
	log := GetDeviceTransferLog(db, device, ts)
	if log != nil {
		updateColumns := []string{"start_time", "complete_time", "fault_code", "fault_string"}
		updateData := &DeviceTransferLog{}
		updateData.StartTime = jsontype.JSONTime(startTime)
		updateData.CompleteTime = jsontype.JSONTime(completeTime)
		updateData.FaultCode = 0
		updateData.FaultString = ""
		if log.TransferType == TransferTypeUpload {
			obj, err := s3.GetObject(device.Schema, log.ObjectBucket, log.ObjectKey)
			if err != nil {
				utils.GetLogger().Error("get s3 object", zap.Error(err))
			}
			if obj != nil {
				updateColumns = append(updateColumns, "s3_object_id")
				updateData.S3ObjectID = &(obj.ID)
			}
		}
		result := db.Model(log).Where("device_id = ? AND time = ?", device.ID, log.Time).Select(updateColumns).Updates(updateData)
		if result.Error != nil {
			return errors.Wrap(result.Error, "update device transfer log with complete")
		}
	}
	return nil
}

func UpdateDeviceTransferLogFault(db *gorm.DB, device *Device, ts int64,
	startTime time.Time, completeTime time.Time, faultCode int, faultString string) error {
	log := GetDeviceTransferLog(db, device, ts)
	if log != nil {
		updateColumns := []string{"start_time", "complete_time", "fault_code", "fault_string"}
		updateData := &DeviceTransferLog{}
		updateData.StartTime = jsontype.JSONTime(startTime)
		updateData.CompleteTime = jsontype.JSONTime(completeTime)
		updateData.FaultCode = faultCode
		updateData.FaultString = faultString
		result := db.Model(log).Where("device_id = ? AND time = ?", device.ID, log.Time).Select(updateColumns).Updates(updateData)
		if result.Error != nil {
			return errors.Wrap(result.Error, "update device transfer log with fault")
		}
	}
	return nil
}
