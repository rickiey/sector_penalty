package pkg

import (
	"github.com/filecoin-project/go-state-types/abi"
	"math/big"
)

const FilecoinPrecision = uint64(1_000_000_000_000_000_000)

func ToFloat64(f abi.TokenAmount) float64 {
	var zero float64 = 0

	if f.Int == nil {
		return zero
	}

	fp := big.NewInt(int64(FilecoinPrecision))
	fil, _ := new(big.Rat).SetFrac(f.Int, fp).Float64()

	return fil
}
