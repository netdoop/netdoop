package omc

import (
	"github.com/heypkg/store/jsontype"
	"gorm.io/gorm"
)

type DataModelTemplate struct {
	ID      uint              `json:"Id" gorm:"primarykey"`
	Updated jsontype.JSONTime `json:"Updated" gorm:"autoUpdateTime"`
	Created jsontype.JSONTime `json:"Created" gorm:"autoCreateTime"`
	Deleted gorm.DeletedAt    `json:"Deleted" gorm:"index"`
	Default bool              `json:"Default" gorm:"index"`

	DataModelID       uint                              `json:"DataModelId" gorm:"uniqueIndex:idx_datamodel_template_unique"`
	DataModel         *DataModel                        `json:"-"`
	Name              string                            `json:"Name" gorm:"uniqueIndex:idx_datamodel_template_unique"`
	ParameterNames    *jsontype.Tags                    `json:"ParameterNames" gorm:"-"`
	ParameterNamesRaw jsontype.JSONType[*jsontype.Tags] `json:"-" gorm:"column:parameter_names"`

	MetaDataRaw jsontype.JSONType[*jsontype.Tags] `json:"-" gorm:"column:meta_data"`
	MetaData    *jsontype.Tags                    `json:"MetaData" gorm:"-"`
}

func (m *DataModelTemplate) SaveData() {
	m.MetaDataRaw = jsontype.NewJSONType(m.MetaData)
	m.ParameterNamesRaw = jsontype.NewJSONType(m.ParameterNames)
}

func (m *DataModelTemplate) LoadData() {
	m.MetaData = m.MetaDataRaw.Data
	m.ParameterNames = m.ParameterNamesRaw.Data
}

func (m *DataModelTemplate) BeforeSave(tx *gorm.DB) (err error) {
	m.SaveData()
	return nil
}

func (m *DataModelTemplate) AfterFind(tx *gorm.DB) (err error) {
	m.LoadData()
	return nil
}

func GetDataModelTemplate(db *gorm.DB, dm *DataModel, id uint) *DataModelTemplate {
	var obj DataModelTemplate
	if result := db.Where(
		"data_model_id = ? AND id = ?",
		dm.ID,
		id,
	).First(&obj); result.Error != nil {
		return nil
	}
	return &obj
}
