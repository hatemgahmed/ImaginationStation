package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//User type
type User struct {
	gorm.Model
	Username string `gorm:"unique;not null" json:"username"`
	Password string `gorm:"not null" json:"password"`
	Admin	 bool	`gorm:"not null" json:"admin"`
}


// DBMigrate will create and migrate the tables
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&User{}, &Notification{})
	//db.Model(&Event{}).AddForeignKey("organization_id", "organizations(id)", "CASCADE", "CASCADE")

	// var e Event
	// var v Volunteer

	// db.First(&e, 1)
	// db.Model(&e).Related(&v)
	return db
}
