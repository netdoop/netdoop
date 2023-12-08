package omc

import (
	"fmt"
	"sync"

	"github.com/netdoop/netdoop/store"
	"github.com/netdoop/netdoop/utils"

	"github.com/Knetic/govaluate"
	"github.com/heypkg/store/jsontype"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type KPIMeas struct {
	ID      uint              `json:"Id" gorm:"primarykey"`
	Updated jsontype.JSONTime `json:"Updated" gorm:"autoUpdateTime"`
	Created jsontype.JSONTime `json:"Created" gorm:"autoCreateTime"`
	Deleted gorm.DeletedAt    `json:"Deleted" gorm:"index"`

	Schema      string `json:"Schema" gorm:"uniqueIndex:idx_kpi_meas_unique"`
	ProductType string `json:"ProductType" gorm:"uniqueIndex:idx_kpi_meas_unique"`
	MeasTypeID  string `json:"MeasTypeID" gorm:"uniqueIndex:idx_kpi_meas_unique"`
	MeasTypeSet string `json:"MeasTypeSet" gorm:"index"`
	Enable      bool   `json:"Enable" gorm:"index"`
	Default     bool   `json:"Default" gorm:"index"`

	Name        string                            `json:"Name"`
	Unit        string                            `json:"Unit"`
	StatsType   string                            `json:"StatsType"`
	Formula     string                            `json:"Formula"`
	MetaDataRaw jsontype.JSONType[*jsontype.Tags] `json:"-" gorm:"column:meta_data"`
	MetaData    *jsontype.Tags                    `json:"MetaData" gorm:"-"`

	FormulaExpression *govaluate.EvaluableExpression `json:"-" gorm:"-"`
}

func (m *KPIMeas) SaveData() {
	m.MetaDataRaw = jsontype.NewJSONType(m.MetaData)
}

func (m *KPIMeas) LoadData() {
	m.MetaData = m.MetaDataRaw.Data
}

func (m *KPIMeas) BeforeSave(tx *gorm.DB) (err error) {
	m.SaveData()
	return nil
}

func (m *KPIMeas) AfterFind(tx *gorm.DB) (err error) {
	m.LoadData()
	return nil
}

func (m *KPIMeas) AfterSave(tx *gorm.DB) error {
	go ReloadAllKPIMeansure()
	return nil
}

var kpiMeasMap map[string]*KPIMeas
var kpiMeasMapLock sync.Mutex

func ReloadAllKPIMeansure() error {
	kpiMeasMapLock.Lock()
	defer kpiMeasMapLock.Unlock()

	db := store.GetDB()
	all := []*KPIMeas{}
	if result := db.Model(&KPIMeas{}).Find(&all); result.Error != nil {
		return result.Error
	}

	tmp := map[string]*KPIMeas{}
	for _, v := range all {
		key := fmt.Sprintf("%v:%v-%v", v.Schema, v.ProductType, v.MeasTypeID)
		if v.Formula != "" {
			expression, err := govaluate.NewEvaluableExpression(v.Formula)
			if err != nil {
				utils.GetLogger().Error("new formular expression", zap.String("formula", v.Formula), zap.Error(err))
			}
			v.FormulaExpression = expression
		}
		tmp[key] = v
	}

	kpiMeasMap = tmp
	return nil
}

func GetKPIMeasureIds(schema string, productType string, measTypeIds []string) []uint {
	ids := []uint{}
	for _, meas := range kpiMeasMap {
		for _, typeId := range measTypeIds {
			if meas.Schema != schema {
				continue
			}
			if meas.ProductType != productType {
				continue
			}
			if meas.MeasTypeID == typeId {
				ids = append(ids, meas.ID)
			}
		}
	}
	return ids
}

func GetKPIMeasuresBySet(schema string, productType string, measTypeSet string) []*KPIMeas {
	// kpiMeasMapLock.Lock()
	// defer kpiMeasMapLock.Unlock()
	out := []*KPIMeas{}
	for _, meas := range kpiMeasMap {
		if meas.Schema != schema {
			continue
		}
		if meas.ProductType != productType {
			continue
		}
		if measTypeSet == "" || meas.MeasTypeSet == measTypeSet {
			out = append(out, meas)
		}
	}
	return out
}

func GetKPIMeasure(schema string, productType string, measTypeID string) *KPIMeas {
	// kpiMeasMapLock.Lock()
	// defer kpiMeasMapLock.Unlock()
	key := fmt.Sprintf("%v:%v-%v", schema, productType, measTypeID)
	v := kpiMeasMap[key]
	return v
}
