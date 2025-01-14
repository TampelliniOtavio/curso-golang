package campaign

type CampaignResponse struct {
    Id                   string
    Name                 string
    Status               string
    AmountOfEmailsToSend int
    Content              string
    CreatedBy            string
}
