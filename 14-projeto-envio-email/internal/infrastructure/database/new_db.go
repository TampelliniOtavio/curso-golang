package database

import (
	"emailn/internal/domain/campaign"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
    dsn := "host=localhost user=user password=password dbname=emailn_dev port=5432 sslmode=disable"

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
