package tax

import (
	"errors"
)

type Repository interface {
	SaveTax(amount float64) error
}

func CalculateTaxAndSave(amount float64, repository Repository) error {
	tax := CalculateTax2(amount)
	return repository.SaveTax(tax)
}

func CalculateTax(amount float64) (float64, error) {
	if amount <= 0 {
		return 0.0, errors.New("amout must be greter then 0")
	}
	if amount >= 1000 && amount < 2000 {
		return 10, nil
	}
	if amount >= 2000 {
		return 20.0, nil
	}
	return 5.0, nil
}

func CalculateTax2(amount float64) float64 {
	if amount <= 0 {
		return 0.0
	}
	if amount >= 1000 {
		return 10
	}
	return 5.0
}
