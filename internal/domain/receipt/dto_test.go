package receipt

import (
	"testing"
	"time"
)

func TestCreateReceiptDTO_ToReceipt(t *testing.T) {
	validDTO := CreateReceiptDTO{
		Retailer:     "Test Retailer",
		PurchaseDate: "2023-01-01",
		PurchaseTime: "15:04",
		Items: []CreateItemDTO{
			{
				ShortDescription: "Test Item 1",
				Price:            "10.00",
			},
			{
				ShortDescription: "Test Item 2",
				Price:            "20.00",
			},
		},
		Total: "30.00",
	}

	gotReceipt, err := validDTO.ToReceipt()
	if err != nil {
		t.Errorf("ToReceipt() error = %v", err)
		return
	}

	if gotReceipt.Retailer != validDTO.Retailer {
		t.Errorf("ToReceipt() retailer = %v, want %v", gotReceipt.Retailer, validDTO.Retailer)
	}

	if gotReceipt.PurchaseDateTime != time.Date(2023, time.January, 1, 15, 4, 0, 0, time.Local) {
		t.Errorf("ToReceipt() purchaseDateTime = %v, want %v", gotReceipt.PurchaseDateTime, time.Date(2023, time.January, 1, 15, 4, 0, 0, time.Local))
	}

	if len(gotReceipt.Items) != 2 {
		t.Errorf("ToReceipt() items = %v, want %v", len(gotReceipt.Items), 2)
	}

	if gotReceipt.Total.StringFixed(2) != "30.00" {
		t.Errorf("ToReceipt() total = %v, want %v", gotReceipt.Total.String(), "30.00")
	}

	for i, item := range gotReceipt.Items {
		if item.ShortDescription != validDTO.Items[i].ShortDescription {
			t.Errorf("ToReceipt() item %d shortDescription = %v, want %v", i, item.ShortDescription, validDTO.Items[i].ShortDescription)
		}

		if item.Price.StringFixed(2) != validDTO.Items[i].Price {
			t.Errorf("ToReceipt() item %d price = %v, want %v", i, item.Price.StringFixed(2), validDTO.Items[i].Price)
		}
	}
}

func TestCreateReceiptDTO_Validate(t *testing.T) {
	tests := []struct {
		name    string
		fields  CreateReceiptDTO
		wantErr bool
	}{
		{
			name: "valid",
			fields: CreateReceiptDTO{
				Retailer:     "Test Retailer",
				PurchaseDate: "2023-01-01",
				PurchaseTime: "15:04",
				Items: []CreateItemDTO{
					{
						ShortDescription: "Test Item 1",
						Price:            "10.00",
					},
					{
						ShortDescription: "Test Item 2",
						Price:            "20.00",
					},
				},
				Total: "30.00",
			},
			wantErr: false,
		},
		{
			name: "missing retailer",
			fields: CreateReceiptDTO{
				PurchaseDate: "2023-01-01",
				PurchaseTime: "15:04",
				Items: []CreateItemDTO{
					{
						ShortDescription: "Test Item 1",
						Price:            "10.00",
					},
					{
						ShortDescription: "Test Item 2",
						Price:            "20.00",
					},
				},
				Total: "30.00",
			},
			wantErr: true,
		},
		{
			name: "wrong item price",
			fields: CreateReceiptDTO{
				Retailer:     "Test Retailer",
				PurchaseDate: "2023-01-01",
				PurchaseTime: "15:04",
				Items: []CreateItemDTO{
					{
						ShortDescription: "Test Item 1",
						Price:            "blah",
					},
					{
						ShortDescription: "Test Item 2",
						Price:            "20.00",
					},
				},
				Total: "30.00",
			},
			wantErr: true,
		},
		{
			name: "negative prices",
			fields: CreateReceiptDTO{
				Retailer:     "Test Retailer",
				PurchaseDate: "2023-01-01",
				PurchaseTime: "15:04",
				Items: []CreateItemDTO{
					{
						ShortDescription: "Test Item 1",
						Price:            "20.00",
					},
					{
						ShortDescription: "Test Item 2",
						Price:            "-20.00",
					},
				},
				Total: "0.00",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &CreateReceiptDTO{
				Retailer:     tt.fields.Retailer,
				PurchaseDate: tt.fields.PurchaseDate,
				PurchaseTime: tt.fields.PurchaseTime,
				Items:        tt.fields.Items,
				Total:        tt.fields.Total,
			}
			if err := r.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
