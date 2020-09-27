package stats

import (
	"github.com/Ilhom5005/bank/pkg/types"
	)

func Avg(payments []types.Payment) types.Money {
	averagesum := types.Money(0)
	sum := types.Money(0)
	for _, payment := range payments {
		sum += payment.Amount
		averagesum = sum / (types.Money(len(payments)))
	}
	
	return averagesum
}


func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	sum := 0
	for _, payment := range payments {
		if category == payment.Category {
			sum += int(payment.Amount)
		}
	}
	
	return types.Money(sum)
}
