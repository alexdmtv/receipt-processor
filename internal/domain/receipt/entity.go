package receipt

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"time"
)

type Receipt struct {
	Id               uuid.UUID
	Retailer         string
	PurchaseDateTime time.Time
	Items            []Item
	Total            decimal.Decimal
	Points           int64
}

type Item struct {
	Id               uuid.UUID
	ShortDescription string
	Price            decimal.Decimal
}

func (r *Receipt) TotalPrice() decimal.Decimal {
	total := decimal.Zero
	for _, item := range r.Items {
		total = total.Add(item.Price)
	}
	return total
}
