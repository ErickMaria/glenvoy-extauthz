package entity

// Base ...
type Base struct {
	ID   uint   `gorm:"PRIMARY_KEY;AUTO_INCREMENT;NOT NULL"`
	Name string `gorm:"SIZE:100;UNIQUE;NOT NULL"`
	// CreatedAt time.Timer `gorm:TYPE:datetime`
	// UpdatedAt time.Timer `gorm:TYPE:datetime`
	// DeletedAt time.Timer `gorm:type:datetime sql:INDEX`
}

// // BeforeCreate ...
// func (base Base) BeforeCreate(scope gorm.Scope) {
// 	if err := scope.SetColumn("CreatedAt", time.Now()); err != nil {
// 		// In Case of any error
// 		log.Fatal("Error during create Collumn on Object: %s", err)
// 	}
// }
