package service

import (
	"github.com/google/wire"
	"lucy/app/engine/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(biz.NewPriceUseCase, biz.NewSymbolUseCase, NewEngineService)
