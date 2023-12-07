package omc

import (
	"fmt"

	"github.com/heypkg/store/jsontype"
	"gorm.io/gorm"
)

type KPITemplate struct {
	ID      uint              `json:"Id" gorm:"primarykey"`
	Updated jsontype.JSONTime `json:"Updated" gorm:"autoUpdateTime"`
	Created jsontype.JSONTime `json:"Created" gorm:"autoCreateTime"`
	Deleted gorm.DeletedAt    `json:"Deleted" gorm:"index"`

	Schema      string `json:"Schema" gorm:"uniqueIndex:idx_kpi_template_unique"`
	ProductType string `json:"ProductType" gorm:"uniqueIndex:idx_kpi_template_unique"`
	Name        string `json:"Name" gorm:"uniqueIndex:idx_kpi_template_unique"`
	Default     bool   `json:"Default" gorm:"index"`

	PeriodicInterval int64                      `json:"PeriodicInterval"`
	SelectType       string                     `json:"SelectType"`
	SelectIds        []uint                     `json:"SelectIds" gorm:"-"`
	SelectIdsRaw     jsontype.JSONSlice[uint]   `json:"-" gorm:"column:select_ids"`
	MeasTypeIds      []string                   `json:"MeasTypeIds" gorm:"-"`
	MeasTypeIdsRaw   jsontype.JSONSlice[string] `json:"-" gorm:"column:meas_type_ids"`

	MetaDataRaw jsontype.JSONType[*jsontype.Tags] `json:"-" gorm:"column:meta_data"`
	MetaData    *jsontype.Tags                    `json:"MetaData" gorm:"-"`
}

func (m KPITemplate) ViewTableName() string {
	return fmt.Sprintf("device_kpi_template_%v_view", m.Name)
}

func (m *KPITemplate) SaveData() {
	m.SelectIdsRaw = jsontype.NewJSONSlice(m.SelectIds)
	m.MeasTypeIdsRaw = jsontype.NewJSONSlice(m.MeasTypeIds)
	m.MetaDataRaw = jsontype.NewJSONType(m.MetaData)
}

func (m *KPITemplate) LoadData() {
	m.MetaData = m.MetaDataRaw.Data
	m.SelectIds = m.SelectIdsRaw
	m.MeasTypeIds = m.MeasTypeIdsRaw
}

func (m *KPITemplate) BeforeSave(tx *gorm.DB) (err error) {
	m.SaveData()
	return nil
}

func (m *KPITemplate) AfterFind(tx *gorm.DB) (err error) {
	m.LoadData()
	return nil
}
