package tax

import (
	"testing"
)

func TestCalculateTax(t *testing.T) {
	amount := 500.0
	expect := 5.0

	result := CalculateTax(amount)

	if result != expect {
		t.Errorf("Expected %f but got %f", expect, result)
	}
}

func TestCalculateTaxBatch(t *testing.T) {
	type calcTax struct {
		amount, expect float64
	}

	table := []calcTax{
		{500.0, 5.0},
		{1000.0, 10.0},
		{1500.0, 10.0},
	}

	for _, item := range table {
		result := CalculateTax(item.amount)
		if result != item.expect {
			t.Errorf("Expected %f but got %f", item.expect, result)
		}
	}
}

func BenchmarkCalculateText(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax(500.0)
	}
}
func BenchmarkCalculateText2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax2(500.0)
	}
}
