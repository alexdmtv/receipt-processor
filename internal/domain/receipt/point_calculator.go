package receipt

type PointCalculator struct {
	rules []PointRule
}

func NewPointCalculator(rules ...PointRule) *PointCalculator {
	return &PointCalculator{
		rules: rules,
	}
}

func (p *PointCalculator) AddRule(rule PointRule) {
	p.rules = append(p.rules, rule)
}

func (p *PointCalculator) Calculate(receipt *Receipt) int64 {
	points := int64(0)
	for _, rule := range p.rules {
		points += rule.Calculate(receipt)
	}
	return points
}
