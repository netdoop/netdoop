package omc

import (
	"fmt"
	"strings"

	"github.com/heypkg/store/jsontype"
	"github.com/heypkg/store/search"
	"gorm.io/gorm"
)

type Group struct {
	ID      uint              `json:"Id" gorm:"primarykey"`
	Updated jsontype.JSONTime `json:"Updated" gorm:"autoUpdateTime"`
	Created jsontype.JSONTime `json:"Created" gorm:"autoCreateTime"`
	Schema  string            `json:"Schema" gorm:"uniqueIndex:idx_group_unique"`
	Name    string            `json:"Name" gorm:"uniqueIndex:idx_group_unique"`

	ParentID *uint    `json:"ParentId" gorm:"index"`
	Parent   *Group   `json:"Parent"`
	Children []*Group `json:"Children" gorm:"foreignKey:ParentID"`

	MetaDataRaw jsontype.JSONType[*jsontype.Tags] `json:"-" gorm:"column:meta_data"`
	MetaData    *jsontype.Tags                    `json:"MetaData" gorm:"-"`
}

func (m *Group) BeforeSave(tx *gorm.DB) (err error) {
	if m.MetaData != nil {
		m.MetaDataRaw = jsontype.NewJSONType(m.MetaData)
	}
	return nil
}

func (m *Group) AfterFind(tx *gorm.DB) (err error) {
	m.MetaData = m.MetaDataRaw.Data
	return nil
}
func (m Group) FindChildByID(id uint) *Group {
	for _, child := range m.Children {
		if child.ID == id {
			return child
		}
		if child2 := child.FindChildByID(id); child2 != nil {
			return child2
		}
	}
	return nil
}

func HanldeGroupSearch(values []search.SearchValue) (string, []any) {
	return handleGroupSearchWithJoin(values, "AND")
}

func handleGroupSearchWithJoin(values []search.SearchValue, join string) (string, []any) {
	args := []any{}
	if len(values) == 0 {
		return "", args
	}

	if len(values) == 1 && (values[0].Value == nil || values[0].Value == "") {
		// Handle case where the only value is empty
		if values[0].Symbol == search.SearchSymbolEq {
			return "NOT LEFT JOIN groups ON devices.group_id = groups.id", args
		} else {
			return "JOIN groups ON devices.group_id = groups.id", args
		}
	}

	conditions := []string{}
	for _, v := range values {
		symbol := v.Symbol
		// Set the default symbol to "=" if it's not one of the allowed symbols
		if symbol != search.SearchSymbolEq &&
			symbol != search.SearchSymbolGt &&
			symbol != search.SearchSymbolLt &&
			symbol != search.SearchSymbolLte {
			symbol = search.SearchSymbolEq
		}

		switch symbol {
		case search.SearchSymbolEq:
			conditions = append(conditions, "devices.group_id = ?")
			args = append(args, v.Value)
		case search.SearchSymbolLt:
			conditions = append(conditions, (`EXISTS (
					WITH RECURSIVE cte AS (
						SELECT x.* FROM groups x WHERE parent_id = ?
						UNION ALL
						SELECT y.* FROM groups y INNER JOIN cte c ON c.id = y.parent_id AND y.id != ?
					)
					SELECT * FROM cte WHERE devices.group_id = cte.id
				)`))
		case search.SearchSymbolLte:
			conditions = append(conditions, `EXISTS (
					WITH RECURSIVE cte AS (
						SELECT x.* FROM groups x WHERE id = ?
						UNION ALL
						SELECT y.* FROM groups y INNER JOIN cte c ON c.id = y.parent_id AND y.id != ?
					)
					SELECT * FROM cte WHERE devices.group_id = cte.id
				)`)
		case search.SearchSymbolGt:
			conditions = append(conditions, `NOT EXISTS (
					WITH RECURSIVE cte AS (
						SELECT x.* FROM groups x WHERE id = ?
						UNION ALL
						SELECT y.* FROM groups y INNER JOIN cte c ON c.id = y.parent_id AND y.id != ?
					)
					SELECT * FROM cte WHERE devices.group_id = cte.id
				)`)
		}
	}
	where := fmt.Sprintf("(%s)", strings.Join(conditions, join))
	return where, args
}

func GetChildrenGroups(db *gorm.DB, schema string, parentID uint) []*Group {
	var groups []*Group
	db.Preload("Children").Where("schema = ? and parent_id = ?", schema, parentID).Find(&groups)
	for i := range groups {
		groups[i].Children = GetChildrenGroups(db, schema, groups[i].ID)
	}
	return groups
}
