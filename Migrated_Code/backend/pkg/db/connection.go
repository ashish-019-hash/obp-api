package db

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	
	"github.com/ashish-019-hash/obp-api-backend/internal/models"
)

var DB *gorm.DB

func InitDB() error {
	var err error
	
	DB, err = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return err
	}

	err = DB.AutoMigrate(
		&models.Bank{},
		&models.BankAccount{},
		&models.Product{},
		&models.Customer{},
		&models.User{},
		&models.Transaction{},
		&models.TransactionRequest{},
		&models.Counterparty{},
		&models.Consent{},
		&models.UserCustomerLink{},
		&models.CustomerAccountLink{},
		&models.AccountRouting{},
		&models.Consumer{},
		&models.Token{},
		&models.UserCredential{},
		&models.Entitlement{},
		&models.LoginAttempt{},
		&models.Scope{},
		&models.ViewPermission{},
		&models.UserAuthContext{},
		&models.ConsentAuthContext{},
		&models.UserLock{},
		&models.AuthenticationTypeValidation{},
		&models.AuthenticationConfig{},
		&models.ConsumerRateLimit{},
		&models.SecuritySettings{},
		&models.TokenConfiguration{},
	)
	if err != nil {
		return err
	}

	log.Println("SQLite in-memory database initialized with GORM")
	log.Println("All models auto-migrated successfully")
	return nil
}

func GetDB() *gorm.DB {
	return DB
}

func CloseDB() error {
	if DB != nil {
		sqlDB, err := DB.DB()
		if err != nil {
			return err
		}
		return sqlDB.Close()
	}
	return nil
}
