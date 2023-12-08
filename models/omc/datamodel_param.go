package omc

import (
	"strings"

	"github.com/heypkg/store/jsontype"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type DataModelParameter struct {
	ID      uint              `json:"Id" gorm:"primarykey"`
	Updated jsontype.JSONTime `json:"Updated" gorm:"autoUpdateTime"`
	Created jsontype.JSONTime `json:"Created" gorm:"autoCreateTime"`
	Deleted gorm.DeletedAt    `json:"Deleted" gorm:"index"`

	DataModelID uint       `json:"DataModelId" gorm:"uniqueIndex:idx_datamodel_param_unique"`
	DataModel   *DataModel `json:"-"`
	Name        string     `json:"Name" gorm:"uniqueIndex:idx_datamodel_param_unique"`
	Default     bool       `json:"Default" gorm:"index"`

	Type         string `json:"Type" gorm:"index"`
	Writable     bool   `json:"Writable" gorm:"index"`
	Description  string `json:"Description"`
	DefaultValue string `json:"DefaultValue"`

	MetaDataRaw jsontype.JSONType[*jsontype.Tags] `json:"-" gorm:"column:meta_data"`
	MetaData    *jsontype.Tags                    `json:"MetaData" gorm:"-"`
}

func (m *DataModelParameter) SaveData() {
	m.MetaDataRaw = jsontype.NewJSONType(m.MetaData)
}

func (m *DataModelParameter) LoadData() {
	m.MetaData = m.MetaDataRaw.Data
}

func (m *DataModelParameter) BeforeSave(tx *gorm.DB) (err error) {
	if m.Type == "object" && !strings.HasSuffix(m.Name, ".") {
		return errors.Errorf("invalid object name %v", m.Name)
	}
	m.SaveData()
	return nil
}

func (m *DataModelParameter) AfterFind(tx *gorm.DB) (err error) {
	m.LoadData()
	return nil
}

func GetDataModelParameterByID(db *gorm.DB, dm *DataModel, id uint) *DataModelParameter {
	var obj DataModelParameter
	if result := db.Where(
		"data_model_id = ? AND id = ?",
		dm.ID,
		id,
	).First(&obj); result.Error != nil {
		return nil
	}
	return &obj
}

func GetDataModelParameterByName(db *gorm.DB, dm *DataModel, name string) *DataModelParameter {
	var obj DataModelParameter
	if result := db.Where(
		"data_model_id = ? AND name = ?",
		dm.ID,
		name,
	).First(&obj); result.Error != nil {
		return nil
	}
	return &obj
}

func UpsertDataModelParameter(db *gorm.DB, dm *DataModel, name string,
	typ *string, writable *bool, description *string, defaultValue *string) error {
	if dm == nil || name == "" {
		return nil
	}
	if typ == nil && writable == nil && description == nil && defaultValue == nil {
		return nil
	}
	param := GetDataModelParameterByName(db, dm, name)
	if param == nil {
		param = &DataModelParameter{
			DataModelID: dm.ID,
			Name:        name,
			Writable:    false,
		}
		if typ != nil {
			param.Type = *typ
		}
		if writable != nil {
			param.Writable = *writable
		}
		if description != nil {
			param.Description = *description
		}
		if defaultValue != nil {
			param.DefaultValue = *defaultValue
		}
		result := db.Create(param)
		if result.Error != nil {
			return errors.Wrap(result.Error, "create data model parameter")
		}
	} else {
		updateColumns := []string{"type"}
		updateData := &DataModelParameter{}
		if typ != nil {
			updateData.Type = *typ
		}
		if writable != nil {
			updateData.Writable = *writable
		}
		if description != nil {
			updateData.Description = *description
		}
		if defaultValue != nil {
			updateData.DefaultValue = *defaultValue
		}
		result := db.Model(param).Where(
			"data_model_id = ? AND name = ?",
			dm.ID, name,
		).Select(updateColumns).Updates(updateData)
		if result.Error != nil {
			return errors.Wrap(result.Error, "update data model parameter")
		}
	}
	return nil
}
