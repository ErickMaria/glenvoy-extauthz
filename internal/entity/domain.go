package entity

// Domain ...
type Domain struct {
	Base
	Url string `gorm:"SIZE:100;UNIQUE_INDEX;NOT NULL"`
	App []App  `gorm:"FOREIGNKEY:ID"`
}

func (Domain) TableName() string {
	return "domains"
}
