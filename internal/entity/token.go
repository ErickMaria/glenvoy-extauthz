package entity

import types "github/erickmaria/glooe-envoy-extauthz/internal/types"

type Token struct {
	Base
	Code   string       `gorm:"SIZE:100;UNIQUE;NOT NULL"`
	Status types.Status `gorm:"SIZE:10;NOT NULL"`
	AppID  uint         `gorm:"NOT NULL"`
}

// TableName ...
func (Token) TableName() string {
	return "tokens"
}
