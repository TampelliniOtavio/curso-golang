package main

import (
	"emailn/internal/domain/campaign"
	"emailn/internal/infrastructure/database"
	"emailn/internal/infrastructure/mail"
	"log"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	println("Worker Started")
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := database.NewDB()
	repository := database.CampaignRepository{
		Db: db,
	}

	service := campaign.ServiceImp{
		Repository: &repository,
		SendMail:   mail.SendMail,
	}

	for {
		campaigns, err := repository.GetCampaignsToBeSent()

		if err != nil {
			println(err.Error())
		} else {
			println("Amount of campaigns: ", len(campaigns))
			for _, campaign := range campaigns {
				println("Campaign Sent: ", campaign.Id)
				service.SendEmailAndUpdateStatus(&campaign)
			}
		}
		time.Sleep(10 * time.Second)
	}
}
