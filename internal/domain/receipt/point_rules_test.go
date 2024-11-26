package receipt

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"testing"
	"time"
)

func TestRetailerCharacterBonusRule_Calculate(t *testing.T) {
	type args struct {
		receipt *Receipt
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "empty",
			args: args{
				receipt: &Receipt{
					Retailer:         "",
					PurchaseDateTime: time.Now(),
					Items:            []Item{},
					Total:            decimal.Zero,
				},
			},
			want: 0,
		},
		{
			name: "5 chars",
			args: args{
				receipt: &Receipt{
					Retailer:         "abcde",
					PurchaseDateTime: time.Now(),
					Items:            []Item{},
					Total:            decimal.Zero,
				},
			},
			want: 5,
		},
		{
			name: "Non alphanumeric chars",
			args: args{
				receipt: &Receipt{
					Retailer:         "123!@#$%^&*()",
					PurchaseDateTime: time.Now(),
					Items:            []Item{},
					Total:            decimal.Zero,
				},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &RetailerCharacterBonusRule{}
			if got := r.Calculate(tt.args.receipt); got != tt.want {
				t.Errorf("Calculate() = %v, want %v", got, tt.want)
			}
			t.Log(r.Description())
		})
	}
}

func TestWholeNumberTotalBonusRule_Calculate(t *testing.T) {
	type args struct {
		receipt *Receipt
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "round total",
			args: args{
				receipt: &Receipt{
					Id:               uuid.New(),
					Retailer:         "Test Retailer",
					PurchaseDateTime: time.Now(),
					Items:            nil,
					Total:            decimal.NewFromFloat(10.00),
					Points:           0,
				},
			},
			want: 50,
		},
		{
			name: "non round total",
			args: args{
				receipt: &Receipt{
					Id:               uuid.New(),
					Retailer:         "Test Retailer",
					PurchaseDateTime: time.Now(),
					Items:            nil,
					Total:            decimal.NewFromFloat(10.01),
					Points:           0,
				},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &WholeNumberTotalBonusRule{}
			if got := r.Calculate(tt.args.receipt); got != tt.want {
				t.Errorf("Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuarterDollarBonusRule_Calculate(t *testing.T) {
	type args struct {
		receipt *Receipt
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "multiple of 0.25",
			args: args{
				receipt: &Receipt{
					Id:               uuid.New(),
					Retailer:         "Test Retailer",
					PurchaseDateTime: time.Now(),
					Items:            nil,
					Total:            decimal.NewFromFloat(1),
					Points:           0,
				},
			},
			want: 25,
		},
		{
			name: "not multiple of 0.25",
			args: args{
				receipt: &Receipt{
					Id:               uuid.New(),
					Retailer:         "Test Retailer",
					PurchaseDateTime: time.Now(),
					Items:            nil,
					Total:            decimal.NewFromFloat(0.26),
					Points:           0,
				},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &QuarterDollarBonusRule{}
			if got := r.Calculate(tt.args.receipt); got != tt.want {
				t.Errorf("Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestItemPairBonusRule_Calculate(t *testing.T) {
	type args struct {
		receipt *Receipt
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "4 items",
			args: args{
				receipt: &Receipt{
					Id:               uuid.New(),
					Retailer:         "Test Retailer",
					PurchaseDateTime: time.Now(),
					Items: []Item{
						{
							Id:               uuid.New(),
							ShortDescription: "Test Item 1",
							Price:            decimal.NewFromFloat(10.00),
						},
						{
							Id:               uuid.New(),
							ShortDescription: "Test Item 2",
							Price:            decimal.NewFromFloat(20.00),
						},
						{
							Id:               uuid.New(),
							ShortDescription: "Test Item 3",
							Price:            decimal.NewFromFloat(30.00),
						},
						{
							Id:               uuid.New(),
							ShortDescription: "Test Item 4",
							Price:            decimal.NewFromFloat(40.00),
						},
					},
					Total:  decimal.NewFromFloat(100.00),
					Points: 0,
				},
			},
			want: 10,
		},
		{
			name: "3 items",
			args: args{
				receipt: &Receipt{
					Id:               uuid.New(),
					Retailer:         "Test Retailer",
					PurchaseDateTime: time.Now(),
					Items: []Item{
						{
							Id:               uuid.New(),
							ShortDescription: "Test Item 1",
							Price:            decimal.NewFromFloat(10.00),
						},
						{
							Id:               uuid.New(),
							ShortDescription: "Test Item 2",
							Price:            decimal.NewFromFloat(20.00),
						},
						{
							Id:               uuid.New(),
							ShortDescription: "Test Item 3",
							Price:            decimal.NewFromFloat(30.00),
						},
					},
					Total:  decimal.NewFromFloat(60.00),
					Points: 0,
				},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ItemPairBonusRule{}
			if got := r.Calculate(tt.args.receipt); got != tt.want {
				t.Errorf("Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDescriptionLengthPriceBonusRule_Calculate(t *testing.T) {
	type args struct {
		receipt *Receipt
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "multiple of 3",
			args: args{
				receipt: &Receipt{
					Id:               uuid.New(),
					Retailer:         "Test Retailer",
					PurchaseDateTime: time.Now(),
					Items: []Item{
						{
							Id:               uuid.New(),
							ShortDescription: "123 56789",
							Price:            decimal.NewFromFloat(10.55),
						},
						{
							Id:               uuid.New(),
							ShortDescription: "123 56789",
							Price:            decimal.NewFromFloat(4.75),
						},
					},
					Total:  decimal.NewFromFloat(15.3),
					Points: 0,
				},
			},
			want: 3,
		},
		{
			name: "one multiple of 3, one not",
			args: args{
				receipt: &Receipt{
					Id:               uuid.New(),
					Retailer:         "Test Retailer",
					PurchaseDateTime: time.Now(),
					Items: []Item{
						{
							Id:               uuid.New(),
							ShortDescription: "0123 56789",
							Price:            decimal.NewFromFloat(10.55),
						},
						{
							Id:               uuid.New(),
							ShortDescription: "123 56789",
							Price:            decimal.NewFromFloat(4.75),
						},
					},
					Total:  decimal.NewFromFloat(15.25),
					Points: 0,
				},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &DescriptionLengthPriceBonusRule{}
			if got := r.Calculate(tt.args.receipt); got != tt.want {
				t.Errorf("Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOddDayBonusRule_Calculate(t *testing.T) {
	type args struct {
		receipt *Receipt
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "odd day",
			args: args{
				receipt: &Receipt{
					Id:               uuid.New(),
					Retailer:         "Test Retailer",
					PurchaseDateTime: time.Date(2023, time.January, 3, 13, 0, 0, 0, time.Local),
					Items:            nil,
					Total:            decimal.Zero,
					Points:           0,
				},
			},
			want: 6,
		},
		{
			name: "even day",
			args: args{
				receipt: &Receipt{
					Id:               uuid.New(),
					Retailer:         "Test Retailer",
					PurchaseDateTime: time.Date(2023, time.January, 4, 13, 0, 0, 0, time.Local),
					Items:            nil,
					Total:            decimal.Zero,
					Points:           0,
				},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &OddDayBonusRule{}
			if got := r.Calculate(tt.args.receipt); got != tt.want {
				t.Errorf("Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAfternoonBonusRule_Calculate(t *testing.T) {
	type args struct {
		receipt *Receipt
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "3pm purchase time",
			args: args{
				receipt: &Receipt{
					Id:               uuid.New(),
					Retailer:         "Test Retailer",
					PurchaseDateTime: time.Date(2023, time.January, 1, 15, 0, 0, 0, time.Local),
					Items:            nil,
					Total:            decimal.Zero,
					Points:           0,
				},
			},
			want: 10,
		},
		{
			name: "6am purchase time",
			args: args{
				receipt: &Receipt{
					Id:               uuid.New(),
					Retailer:         "Test Retailer",
					PurchaseDateTime: time.Date(2023, time.January, 1, 6, 0, 0, 0, time.Local),
					Items:            nil,
					Total:            decimal.Zero,
					Points:           0,
				},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &AfternoonBonusRule{}
			if got := r.Calculate(tt.args.receipt); got != tt.want {
				t.Errorf("Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}
