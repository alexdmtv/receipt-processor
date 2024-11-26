package receipt

import (
	"context"
)

type Service struct {
	receiptRepository Repository
}

func NewService(receiptRepository Repository) *Service {
	return &Service{
		receiptRepository: receiptRepository,
	}
}

func (s *Service) Create(ctx context.Context, receiptDTO CreateReceiptDTO) (*Receipt, error) {
	receipt, err := receiptDTO.ToReceipt()
	if err != nil {
		return nil, err
	}

	receipt.Points = s.calculatePoints(receipt)
	return s.receiptRepository.Create(receipt)
}

func (s *Service) GetReceiptPoints(ctx context.Context, id string) (int64, error) {
	receipt, err := s.receiptRepository.Get(id)
	if err != nil {
		return 0, err
	}

	return receipt.Points, nil
}

func (s *Service) calculatePoints(receipt *Receipt) int64 {
	pointsCalculator := NewPointCalculator(
		&RetailerCharacterBonusRule{},
		&WholeNumberTotalBonusRule{},
		&QuarterDollarBonusRule{},
		&ItemPairBonusRule{},
		&DescriptionLengthPriceBonusRule{},
		&OddDayBonusRule{},
		&AfternoonBonusRule{},
	)

	points := pointsCalculator.Calculate(receipt)
	return points
}
