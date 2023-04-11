package biz

import "github.com/shopspring/decimal"

type PriceRepo interface {
	SetSymbolPrice(symbol string, price decimal.Decimal)
	GetSymbolPrice(symbol string) decimal.Decimal
	DelSymbolPrice(symbol string)
}

type PriceUseCase struct {
	repo PriceRepo
}

func NewPriceUseCase(repo PriceRepo) *PriceUseCase {
	return &PriceUseCase{repo: repo}
}

func (p *PriceUseCase) SetSymbolPrice(symbol string, price decimal.Decimal) {
	p.repo.SetSymbolPrice(symbol, price)
}

func (p *PriceUseCase) GetSymbolPrice(symbol string) decimal.Decimal {
	return p.repo.GetSymbolPrice(symbol)
}

func (p *PriceUseCase) DelSymbolPrice(symbol string) {
	p.repo.DelSymbolPrice(symbol)
}
