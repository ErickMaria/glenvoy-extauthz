package entity

import types "github/erickmaria/glooe-envoy-extauthz/internal/types"

// App ...
type App struct {
	Base
	ClientID string       `gorm:"COLUMN:CLIENT_ID;SIZE:100;UNIQUE;NOT NULL"`
	Status   types.Status `gorm:"COLUMN:STATUS;SIZE:10;UNIQUE;NOT NULL"`
	Token    []Token      `gorm:"FOREIGNKEY:ID"`
	GroupID  int          `gorm:"COLUMN:GROUP_ID"`
}

// TableName ...
func (App) TableName() string {
	return "APPS"
}
