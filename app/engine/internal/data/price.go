package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/shopspring/decimal"
	"lucy/app/engine/internal/biz"
	"lucy/app/engine/internal/constvar"
)

type priceRepo struct {
	data *Data
	log  *log.Helper
}

func NewPriceRepo(data *Data, logger log.Logger) biz.PriceRepo {
	return &priceRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/engine-service")),
	}
}

func (p *priceRepo) SetSymbolPrice(symbol string, price decimal.Decimal) {
	ctx := context.TODO()
	p.data.redisClient.Set(ctx, constvar.PricePrefix+symbol, price.String(), 0)
}

func (p *priceRepo) GetSymbolPrice(symbol string) decimal.Decimal {
	ctx := context.TODO()
	priceStr, err := p.data.redisClient.Get(ctx, constvar.PricePrefix+symbol).Result()
	if err != nil {
		return decimal.Zero
	}
	price, _ := decimal.NewFromString(priceStr)
	return price
}

func (p *priceRepo) DelSymbolPrice(symbol string) {
	ctx := context.TODO()
	p.data.redisClient.Del(ctx, constvar.PricePrefix+symbol)
}
