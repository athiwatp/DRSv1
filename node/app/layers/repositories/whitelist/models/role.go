package models

import (
	"gitlab.com/velo-labs/cen/node/app/constants"
	"gitlab.com/velo-labs/cen/node/app/entities"
)

type GetRole struct {
	ID   int `gorm:"primary_key"`
	Name string
	Code string
}

func (GetRole) TableName() string {
	return constants.RoleTable
}

func (m GetRole) ToEntity() (entity entities.Role) {
	return entities.Role{
		ID:   m.ID,
		Name: m.Name,
		Code: m.Code,
	}
}
