package receipt

import (
	"github.com/shopspring/decimal"
	"strings"
	"unicode"
)

type PointRule interface {
	Calculate(*Receipt) int64
	Description() string
}

type RetailerCharacterBonusRule struct{}

func (r *RetailerCharacterBonusRule) Calculate(receipt *Receipt) int64 {
	count := int64(0)
	for _, r := range receipt.Retailer {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			count++
		}
	}
	return count
}
func (r *RetailerCharacterBonusRule) Description() string {
	return "One point for every alphanumeric character in the retailer name"
}

type WholeNumberTotalBonusRule struct{}

func (r *WholeNumberTotalBonusRule) Calculate(receipt *Receipt) int64 {
	if receipt.Total.Truncate(0).Equal(receipt.Total) {
		return 50
	}
	return 0
}

func (r *WholeNumberTotalBonusRule) Description() string {
	return "50 points if the total is a round dollar amount with no cents"
}

type QuarterDollarBonusRule struct{}

func (r *QuarterDollarBonusRule) Calculate(receipt *Receipt) int64 {
	quarter := decimal.NewFromFloat(0.25)
	if receipt.Total.Mod(quarter).Equal(decimal.Zero) {
		return 25
	}
	return 0
}

func (r *QuarterDollarBonusRule) Description() string {
	return "25 points if the total is a multiple of 0.25"
}

type ItemPairBonusRule struct{}

func (r *ItemPairBonusRule) Calculate(receipt *Receipt) int64 {
	return int64(len(receipt.Items) / 2 * 5)
}

func (r *ItemPairBonusRule) Description() string {
	return "5 points for every two items on the receipt"
}

type DescriptionLengthPriceBonusRule struct{}

func (r *DescriptionLengthPriceBonusRule) Calculate(receipt *Receipt) int64 {
	points := int64(0)
	for _, item := range receipt.Items {
		itemDescriptionLength := len(strings.TrimSpace(item.ShortDescription))
		if (itemDescriptionLength % 3) == 0 {
			points += int64(item.Price.Mul(decimal.NewFromFloat(0.2)).Round(0).IntPart())
		}
	}

	return points
}

func (r *DescriptionLengthPriceBonusRule) Description() string {
	return "If the trimmed length of the item description is a multiple of 3, multiply the price by 0.2 and round up to the nearest integer. The result is the number of points earned"
}

type OddDayBonusRule struct{}

func (r *OddDayBonusRule) Calculate(receipt *Receipt) int64 {
	if receipt.PurchaseDateTime.Day()%2 != 0 {
		return 6
	}
	return 0
}

func (r *OddDayBonusRule) Description() string {
	return "6 points if the day in the purchase date is odd"
}

type AfternoonBonusRule struct{}

func (r *AfternoonBonusRule) Calculate(receipt *Receipt) int64 {
	if receipt.PurchaseDateTime.Hour() > 14 && receipt.PurchaseDateTime.Hour() < 16 {
		return 10
	}
	return 0
}

func (r *AfternoonBonusRule) Description() string {
	return "10 points if the time of purchase is after 2:00pm and before 4:00pm"
}
