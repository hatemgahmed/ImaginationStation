package models

import "github.com/jinzhu/gorm"

//Notification type
type Notification struct {
	gorm.Model
	ID            	 int	 `gorm:"primary_key;not null" json:"id"`
	Title            string  `gorm:"not null" json:"title"`
	Date             string  `json:"date"`
	Content			 string  `gorm:"not null" json:"title"`
}
