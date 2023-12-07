package omc

import (
	"github.com/heypkg/s3"
	"github.com/heypkg/store/jsontype"
	"gorm.io/gorm"
)

const FirmwareBucket = "omc-firmware"

type Firmware struct {
	ID      uint              `json:"Id" gorm:"primarykey"`
	Updated jsontype.JSONTime `json:"Updated" gorm:"autoUpdateTime"`
	Created jsontype.JSONTime `json:"Created" gorm:"autoCreateTime"`
	Deleted gorm.DeletedAt    `json:"Deleted" gorm:"index"`

	Schema      string `json:"Schema" gorm:"uniqueIndex:idx_firmware_unique"`
	Name        string `json:"Name" gorm:"uniqueIndex:idx_firmware_unique"`
	Version     string `json:"Version" gorm:"uniqueIndex:idx_firmware_unique"`
	ProductType string `json:"ProductType" gorm:"index"`

	Products []*Product `gorm:"many2many:firmware_products;"`

	Uploader   string            `json:"Uploader"`
	UploadTime jsontype.JSONTime `json:"UploadTime"`
	S3ObjectID uint              `json:"S3ObjectId" gorm:"index"`
	S3Object   *s3.S3Object      `json:"S3Object"`
}

func (m *Firmware) AfterDelete(tx *gorm.DB) (err error) {
	s3.RemoveObjectById(m.Schema, m.S3ObjectID)
	return nil
}

func GetFirmware(db *gorm.DB, schema string, id uint) *Firmware {
	var firmware Firmware
	if result := db.Where(
		"schema = ? AND id = ?",
		schema, id,
	).Preload("S3Object").Preload("Products").First(&firmware); result.Error != nil {
		return nil
	}
	return &firmware
}
