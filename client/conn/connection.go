package conn

import (
	"fmt"
	"toritorkari-bazar/config"
	"toritorkari-bazar/internal/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func connect() {
	dbConfig := config.LocalConfig

	// certPool := x509.NewCertPool()
	// if !certPool.AppendCertsFromPEM([]byte(dbConfig.CERT)) {
	// 	log.Fatalf("Failed to append CA certificate")
	// }

	// tlsConfig := &tls.Config{
	// 	RootCAs: certPool,
	// }

	// err := sq.RegisterTLSConfig("custom", tlsConfig)
	// if err != nil {
	// 	log.Fatalf("Failed to register custom TLS config: %v", err)
	// }

	dsn := fmt.Sprintf("%s:%s@%s/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig.DBUser, dbConfig.DBPass, dbConfig.DBIp, dbConfig.DBName)

	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})

	if err != nil {
		fmt.Println("error connecting to DB")
		panic(err)
	}

	fmt.Println("Database Connected")
	db = d
}

func migrate() {
	db.Migrator().AutoMigrate(&models.Book{})
	db.Migrator().AutoMigrate(&models.Category{})
	db.Migrator().AutoMigrate(&models.SubCategory{})
	db.Migrator().AutoMigrate(&models.Product{})
	db.Migrator().AutoMigrate(&models.User{})
}

func GetDB() *gorm.DB {
	if db == nil {
		connect()
	}
	migrate()
	return db
}
