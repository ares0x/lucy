package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"lucy/app/engine/internal/biz"
	"lucy/app/engine/internal/constvar"
)

type symbolRepo struct {
	data *Data
	log  *log.Helper
}

func NewSymbolRepo(data *Data, logger log.Logger) biz.SymbolRepo {
	return &symbolRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/engine-service")),
	}
}

// SetSymbol 新增一个币对
func (s *symbolRepo) SetSymbol(symbol string) (int64, error) {
	ctx := context.TODO()
	return s.data.redisClient.SAdd(ctx, constvar.SymbolsKey, symbol).Result()
}

// GetSymbols 获取所有已经开启的币对
func (s *symbolRepo) GetSymbols() ([]string, error) {
	ctx := context.TODO()
	return s.data.redisClient.SMembers(ctx, constvar.SymbolsKey).Result()
}

// DelSymbol 删除一个币对
func (s *symbolRepo) DelSymbol(symbol string) (int64, error) {
	ctx := context.TODO()
	return s.data.redisClient.SRem(ctx, constvar.SymbolsKey, symbol).Result()
}
