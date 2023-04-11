package biz

// SymbolRepo 币对引擎
type SymbolRepo interface {
	SetSymbol(symbol string) (int64, error)
	GetSymbols() ([]string, error)
	DelSymbol(symbol string) (int64, error)
}

type SymbolUseCase struct {
	repo SymbolRepo
}

func NewSymbolUseCase(repo SymbolRepo) *SymbolUseCase {
	return &SymbolUseCase{repo: repo}
}

func (s *SymbolUseCase) SetSymbol(symbol string) (int64, error) {
	return s.repo.SetSymbol(symbol)
}

func (s *SymbolUseCase) GetSymbols() ([]string, error) {
	return s.repo.GetSymbols()
}

func (s *SymbolUseCase) DelSymbol(symbol string) (int64, error) {
	return s.repo.DelSymbol(symbol)
}
