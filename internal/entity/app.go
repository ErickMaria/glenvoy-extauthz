package entity

import types "github/erickmaria/glooe-envoy-extauthz/internal/types"

// App ...
type App struct {
	Base
	Code     string       `gorm:"SIZE:100;UNIQUE;NOT NULL"`
	Status   types.Status `gorm:"SIZE:10;NOT NULL"`
	Token    []Token      `gorm:"FOREIGNKEY:ID"`
	DomainID uint         `gorm:"NOT NULL"`
}

// TableName ...
func (App) TableName() string {
	return "apps"
}
