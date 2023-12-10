package random

import (
	"crypto/rand"
	"github.com/shopspring/decimal"
	"math/big"
)

type Generator interface {
	GetRandom() (string, error)
}

type DefaultGenerator struct {
	rngBracket *big.Int
	precision  int32
}

func NewDefaultGenerator(
	precision int32,
) *DefaultGenerator {
	bracket := maxBracket(precision)
	g := &DefaultGenerator{
		rngBracket: big.NewInt(bracket),
		precision:  precision,
	}
	return g
}

func (g *DefaultGenerator) GetRandom() (string, error) {
	value, err := rand.Int(rand.Reader, g.rngBracket)
	if err != nil {
		return "", err
	}
	result := decimal.NewFromBigInt(value, -g.precision)
	return result.String(), nil
}
