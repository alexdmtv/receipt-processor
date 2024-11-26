package receipt

type Repository interface {
	Create(receipt *Receipt) (*Receipt, error)
	Get(id string) (*Receipt, error)
	Update(id string, receipt *Receipt) (*Receipt, error)
	Delete(id string) error
}
