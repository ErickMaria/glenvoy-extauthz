package entity

// Domain ...
type Domain struct {
	Base
	Host   string `gorm:"SIZE:100;NOT NULL"`
	Prefix string `gorm:"SIZE:100;NOT NULL"`
	App    []App  `gorm:"FOREIGNKEY:ID"`
}

func (Domain) TableName() string {
	return "domains"
}
