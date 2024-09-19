package campaign

type Repository interface{
    Create(campaign *Campaign) error
    Update(campaign *Campaign) error
    Get() (*[]Campaign, error)
    GetById(id string) (*Campaign, error)
    Delete(campaign *Campaign) error
}
