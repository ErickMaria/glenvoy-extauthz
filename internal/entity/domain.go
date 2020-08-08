package entity

// Domain ...
type Domain struct {
	Base
	Url   string  `gorm:"COLUMN:URL;SIZE:100;UNIQUE_INDEX;NOT NULL"`
	Group []Group `gorm:"FOREIGNKEY:ID"`
}

func (Domain) TableName() string {
	return "DOMAINS"
}
