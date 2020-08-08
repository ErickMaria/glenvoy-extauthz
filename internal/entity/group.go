package entity

// Domain ...
type Group struct {
	Base
	DomainID int   `gorm:"COLUMN:DOMAIN_ID"`
	App      []App `gorm:"FOREIGNKEY:ID"`
}

// TableName ...
func (Group) TableName() string {
	return "GROUPS"
}
