package main

import (
	"emailn/internal/domain/campaign"

	"github.com/go-playground/validator/v10"
)

func main() {
    contacts := []campaign.Contact{{Email: "emai"}, {Email: ""}}
    campaign := campaign.Campaign{
        // Id: "123",
        Name: "213123123123asdasdasdasdadasdasdasdasdasdad",
        // Content: "Content",
        Contacts: contacts,
        // CreatedOn: time.Now(),
    }

    validate := validator.New()

    err := validate.Struct(campaign)

    if err == nil {
        println("Nenhum Erro")
    } else {
        validationErrors := err.(validator.ValidationErrors)


        for _, v := range validationErrors {
            switch v.Tag() {
            case "required":
                println(v.StructField() + " is required")
            case "min":
                println(v.StructField() + " is required with min " + v.Param())
            case "max":
                println(v.StructField() + " is required with max " + v.Param())
            case "email":
                println(v.StructField() + " is invalid")
            }
        }
    }
}
