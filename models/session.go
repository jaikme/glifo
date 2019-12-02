package models

import "github.com/jinzhu/gorm"

// Session type details
type Session struct {
	gorm.Model
	Name           string  `gorm:"size:255"`
	Email          string  `gorm:"type:varchar(100);unique_index"`
	SerialNumber   *string `gorm:"unique;not null"` // set member number to unique and not null
	ImageLocalPath string  `gorm:"index:addr"`      // create index with name `addr` for address
}
