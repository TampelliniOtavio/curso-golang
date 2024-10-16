package database

import (
	"emailn/internal/domain/campaign"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
    dsn := os.Getenv("DATABASE")

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

    if err != nil {
        panic("Failed to Connect to Database")
    }

    db.AutoMigrate(
        &campaign.Campaign{},
        &campaign.Contact{},
    )

    return db
}
