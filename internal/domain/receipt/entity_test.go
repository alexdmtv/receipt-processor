package receipt

import (
	"github.com/shopspring/decimal"
	"reflect"
	"testing"
	"time"
)

func TestReceipt_TotalPrice(t *testing.T) {
	tests := []struct {
		name   string
		fields Receipt
		want   decimal.Decimal
	}{
		{
			name: "empty",
			fields: Receipt{
				Retailer:         "Test Retailer",
				PurchaseDateTime: time.Date(2023, time.January, 1, 15, 4, 0, 0, time.Local),
				Items:            []Item{},
				Total:            decimal.Zero,
			},
			want: decimal.Zero,
		},
		{
			name: "one item",
			fields: Receipt{
				Retailer:         "Test Retailer",
				PurchaseDateTime: time.Date(2023, time.January, 1, 15, 4, 0, 0, time.Local),
				Items: []Item{
					{
						ShortDescription: "Test Item 1",
						Price:            decimal.NewFromFloat(10.00),
					},
				},
				Total: decimal.NewFromFloat(10.00),
			},
			want: decimal.NewFromFloat(10.00),
		},
		{
			name: "two items",
			fields: Receipt{
				Retailer:         "Test Retailer",
				PurchaseDateTime: time.Date(2023, time.January, 1, 15, 4, 0, 0, time.Local),
				Items: []Item{
					{
						ShortDescription: "Test Item 1",
						Price:            decimal.NewFromFloat(10.00),
					},
					{
						ShortDescription: "Test Item 2",
						Price:            decimal.NewFromFloat(20.00),
					},
				},
				Total: decimal.NewFromFloat(30.00),
			},
			want: decimal.NewFromFloat(30.00),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Receipt{
				Id:               tt.fields.Id,
				Retailer:         tt.fields.Retailer,
				PurchaseDateTime: tt.fields.PurchaseDateTime,
				Items:            tt.fields.Items,
				Total:            tt.fields.Total,
			}
			if got := r.TotalPrice(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TotalPrice() = %v, want %v", got, tt.want)
			}
		})
	}
}
