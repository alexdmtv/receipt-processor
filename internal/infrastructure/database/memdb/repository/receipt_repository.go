package repository

import (
	"receipt-processor/internal/domain/receipt"
	"receipt-processor/internal/infrastructure/database/memdb"
)

var _ receipt.Repository = (*ReceiptRepository)(nil)

type ReceiptRepository struct {
	db *memdb.DB
}

func NewReceiptRepository(db *memdb.DB) *ReceiptRepository {
	return &ReceiptRepository{
		db: db,
	}
}

func (r *ReceiptRepository) Create(receipt *receipt.Receipt) (*receipt.Receipt, error) {
	err := r.db.Set(receipt.Id.String(), receipt)
	if err != nil {
		return nil, err
	}
	return receipt, nil
}

func (r *ReceiptRepository) Get(id string) (*receipt.Receipt, error) {
	rec, err := r.db.Get(id)
	if err != nil {
		return nil, err
	}
	return rec.(*receipt.Receipt), nil
}

func (r *ReceiptRepository) Update(id string, receipt *receipt.Receipt) (*receipt.Receipt, error) {
	err := r.db.Set(id, receipt)
	if err != nil {
		return nil, err
	}
	return receipt, nil
}

func (r *ReceiptRepository) Delete(id string) error {
	err := r.db.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
