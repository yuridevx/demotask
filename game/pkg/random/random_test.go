package random

import (
	"github.com/shopspring/decimal"
	"testing"
)

func TestRandom(t *testing.T) {
	g := NewDefaultGenerator(2)
	// 10000 times is just an arbitrary number, randomness is hard to test.
	// ideally we would test the distribution of the random numbers.
	// won't do that here.
	for i := 0; i < 10000; i++ {
		rnd, err := g.GetRandom()
		if err != nil {
			t.Fatal(err)
		}
		val := decimal.RequireFromString(rnd)
		if val.LessThan(decimal.NewFromInt(0)) {
			t.Fatalf("expected value >= 0, got %s", val.String())
		}
		if val.GreaterThan(decimal.NewFromInt(1)) {
			t.Fatalf("expected value <= 1, got %s", val.String())
		}
	}
}

func BenchmarkRandom(b *testing.B) {
	g := NewDefaultGenerator(2)
	for i := 0; i < b.N; i++ {
		g.GetRandom()
	}
}
