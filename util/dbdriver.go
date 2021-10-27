package util

import (
	"altastore/config"
	"altastore/models"

	"github.com/labstack/gommon/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func MysqlDatabaseConnection(config *config.AppConfig) *gorm.DB {
	// uri := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true",
	// 	config.Database.Username,
	// 	config.Database.Password,
	// 	config.Database.Address,
	// 	config.Database.Port,
	// 	config.Database.Name)

	db, err := gorm.Open(mysql.Open(config.Database.Connection), &gorm.Config{})

	if err != nil {
		log.Info("failed to connect database: ", err)
		panic(err)
	}
	// Uncommand For Migration
	DatabaseMigration(db)

	return db
}

// Create Migration Here
func DatabaseMigration(db *gorm.DB) {
	db.AutoMigrate(models.Customer{})
	db.AutoMigrate(models.Product{})
	db.AutoMigrate(models.Category{})
	db.AutoMigrate(models.Carts{})
	db.AutoMigrate(models.CartDetails{})
	db.AutoMigrate(models.Checkout{})

}
