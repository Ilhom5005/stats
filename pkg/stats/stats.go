package stats

import (
	"github.com/Ilhom5005/bank/v2/pkg/types"
	)

func Avg(payments []types.Payment) types.Money {
	averagesum := types.Money(0)
	sum := types.Money(0)
	for _, payment := range payments {
		if payment.Status == types.StatusFail {
			continue
		}
		sum += payment.Amount
		averagesum = sum / (types.Money(len(payments)))
	}
	
	return averagesum
}

func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	sum := 0
	for _, payment := range payments {
		if category == payment.Category {
		if payment.Status == types.StatusFail {
			sum += int(payment.Amount)
		}
	}
}
	
	return types.Money(sum)
}
