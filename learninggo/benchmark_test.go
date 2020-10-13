package learninggo

/*
go test -run=Bench -bench=.
go test -run=Sqrt -bench=. -benchtime=20s
*/

import (
	"math"
	"testing"
	"time"
)

var (
	numbersToCheck = []float64{1, 1001, 1000001, 10000001.100000001}
)

func BenchmarkMathSqrt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		math.Sqrt(numbersToCheck[0])
		math.Sqrt(numbersToCheck[1])
		math.Sqrt(numbersToCheck[2])
		math.Sqrt(numbersToCheck[3])
	}
}

func BenchmarkMySqrt64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MySqrt64(numbersToCheck[0])
		MySqrt64(numbersToCheck[1])
		MySqrt64(numbersToCheck[2])
		MySqrt64(numbersToCheck[3])
	}
}

func MySqrt64(number float64) float64 {
	time.Sleep(time.Millisecond * 10)
	return math.Sqrt(number)
}
