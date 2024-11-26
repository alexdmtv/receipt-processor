package receipt

import (
	"fmt"
	validator "github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"regexp"
	"time"
)

type CreateReceiptDTO struct {
	Retailer     string          `json:"retailer" validate:"required,regexp=^[\\w\\s\\-&]+$"`
	PurchaseDate string          `json:"purchaseDate" validate:"required,datetime=2006-01-02"`
	PurchaseTime string          `json:"purchaseTime" validate:"required,datetime=15:04"`
	Items        []CreateItemDTO `json:"items" validate:"required,min=1,dive"`
	Total        string          `json:"total" validate:"required,regexp=^\\d+\\.\\d{2}$,gt=0"`
}

type CreateItemDTO struct {
	ShortDescription string `json:"shortDescription" validate:"required,regexp=^[\\w\\s\\-]+$"`
	Price            string `json:"price" validate:"required,regexp=^\\d+\\.\\d{2}$,gt=0"`
}

func (r *CreateReceiptDTO) Validate() error {
	validate := validator.New()
	err := validate.RegisterValidation("regexp", func(fl validator.FieldLevel) bool {
		pattern := fl.Param()
		regex := regexp.MustCompile(pattern)
		return regex.MatchString(fl.Field().String())
	})
	if err != nil {
		return err
	}

	if err := validate.Struct(r); err != nil {
		return err
	}

	var sum decimal.Decimal
	total, err := decimal.NewFromString(r.Total)
	if err != nil {
		return err
	}

	for _, item := range r.Items {
		price, err := decimal.NewFromString(item.Price)
		if err != nil {
			return err
		}
		sum = sum.Add(price)
	}

	if !sum.Equal(total) {
		return fmt.Errorf("sum of items (%s) does not match total (%s)", sum, total)
	}

	return nil
}

func (r *CreateReceiptDTO) ToReceipt() (*Receipt, error) {
	// Parse purchase date
	purchaseDate, err := time.Parse("2006-01-02", r.PurchaseDate)
	if err != nil {
		return nil, err
	}

	// Parse purchase time
	purchaseTime, err := time.Parse("15:04", r.PurchaseTime)
	if err != nil {
		return nil, err
	}

	// Combine date and time
	fullPurchaseTime := time.Date(
		purchaseDate.Year(),
		purchaseDate.Month(),
		purchaseDate.Day(),
		purchaseTime.Hour(),
		purchaseTime.Minute(),
		0, 0,
		time.Local,
	)

	// Parse total
	total, err := decimal.NewFromString(r.Total)
	if err != nil {
		return nil, err
	}

	// Convert items
	items := make([]Item, len(r.Items))
	for i, item := range r.Items {
		price, err := decimal.NewFromString(item.Price)
		if err != nil {
			return nil, err
		}
		items[i] = Item{
			Id:               uuid.New(),
			ShortDescription: item.ShortDescription,
			Price:            price,
		}
	}

	return &Receipt{
		Id:               uuid.New(),
		Retailer:         r.Retailer,
		PurchaseDateTime: fullPurchaseTime,
		Items:            items,
		Total:            total,
	}, nil
}
