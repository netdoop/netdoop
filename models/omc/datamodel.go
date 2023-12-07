package omc

import (
	"github.com/heypkg/store/jsontype"
	"gorm.io/gorm"
)

type DataModel struct {
	ID      uint              `json:"Id" gorm:"primarykey"`
	Updated jsontype.JSONTime `json:"Updated" gorm:"autoUpdateTime"`
	Created jsontype.JSONTime `json:"Created" gorm:"autoCreateTime"`
	Deleted gorm.DeletedAt    `json:"Deleted" gorm:"index"`
	Default bool              `json:"Default" gorm:"index"`

	Schema        string                            `json:"Schema" gorm:"uniqueIndex:idx_datamodel_unique"`
	ProductType   string                            `json:"ProductType" gorm:"uniqueIndex:idx_datamodel_unique"`
	Name          string                            `json:"Name" gorm:"uniqueIndex:idx_datamodel_unique"`
	ParameterPath string                            `json:"ParameterPath" gorm:"index"`
	MetaDataRaw   jsontype.JSONType[*jsontype.Tags] `json:"-" gorm:"column:meta_data"`
	MetaData      *jsontype.Tags                    `json:"MetaData" gorm:"-"`

	Templates  []*DataModelTemplate  `json:"Templates"`
	Parameters []*DataModelParameter `json:"Parameters"`
}

func (m *DataModel) SaveData() {
	m.MetaDataRaw = jsontype.NewJSONType(m.MetaData)
}

func (m *DataModel) LoadData() {
	m.MetaData = m.MetaDataRaw.Data
}

func (m *DataModel) BeforeSave(tx *gorm.DB) (err error) {
	m.SaveData()
	return nil
}

func (m *DataModel) AfterFind(tx *gorm.DB) (err error) {
	m.LoadData()
	return nil
}

func (m DataModel) GetParameter(name string) *DataModelParameter {
	for _, v := range m.Parameters {
		if v.Name == name {
			return v
		}
	}
	return nil
}

func (m DataModel) GetParameterType(name string) string {
	if param := m.GetParameter(name); param != nil {
		return param.Type
	}
	return ""
}
