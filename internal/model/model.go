
package model

import (
	"sync"

	"gorm.io/gorm"
)

var once sync.Once

func AutoMigrate(db *gorm.DB, key string) {
	switch key {

	case "CloudConfig":
		once.Do(func() {
			db.AutoMigrate(CloudConfig{})
		})

	case "User":
		once.Do(func() {
			db.AutoMigrate(User{})
		})
	}
}