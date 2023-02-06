package futures

import (
	"github.com/nntaoli-project/goex/v2/model"
	"github.com/nntaoli-project/goex/v2/okx/common"
	"github.com/nntaoli-project/goex/v2/options"
)

type Swap struct {
	*common.OKxV5
	currencyPairM map[string]model.CurrencyPair
}

func NewSwap() *Swap {
	var currencyPairM = make(map[string]model.CurrencyPair, 64)
	return &Swap{
		OKxV5:         common.New(),
		currencyPairM: currencyPairM}
}

func (f *Swap) GetExchangeInfo() (map[string]model.CurrencyPair, []byte, error) {
	m, b, er := f.OKxV5.GetExchangeInfo("SWAP")
	f.currencyPairM = m
	return m, b, er
}

func (f *Swap) NewCurrencyPair(baseSym, quoteSym string) model.CurrencyPair {
	return f.currencyPairM[baseSym+quoteSym]
}

func (f *Swap) NewPrvApi(apiOpts ...options.ApiOption) *PrvApi {
	return NewPrvApi(f.OKxV5, apiOpts...)
}
