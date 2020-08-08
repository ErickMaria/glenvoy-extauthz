package entity

import types "github/erickmaria/glooe-envoy-extauthz/internal/types"

type Token struct {
	Base
	Code   string       `gorm:"COLUMN:CODE;SIZE:100;UNIQUE;NOT NULL"`
	Status types.Status `gorm:"COLUMN:STATUS;SIZE:10;UNIQUE;NOT NULL"`
	AppID  int          `gorm:"COLUMN:APP_ID"`
}

// TableName ...
func (Token) TableName() string {
	return "TOKENS"
}
