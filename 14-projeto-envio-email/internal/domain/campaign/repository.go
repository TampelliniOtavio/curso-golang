package campaign

type Repository interface{
    Save(campaign *Campaign) error
    Get() (*[]Campaign, error)
    GetById(id string) (*Campaign, error)
}
