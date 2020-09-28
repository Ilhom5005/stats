package stats

import (
	"github.com/Ilhom5005/bank/v2/pkg/types"
	"fmt"
)

func ExampleAvg() {
	payments := []types.Payment {
{
	ID: 1,
	Amount: 10_000,
	Category: "jacket",
	Status: types.StatusFAIL,
},
{
	ID: 2,
	Amount: 10_000,
	Category: "jeans",
	Status: types.StatusOK,
},
{
	ID: 3,
	Amount: 10_000,
	Category: "T-shirt",
	Status: types.StatusOK,
},
	}
fmt.Println(Avg(payments))

//Output:
//20000

}

func TotalInCategory() {
	payments := []types.Payment {
		{
			ID: 1,
			Amount: 10_000,
			Category: "Cafe",
			Status: types.StatusOK,
		},
		{
			ID: 2,
			Amount: 10_000,
			Category: "Cafe",
			Status: types.StatusOK,
		},
		{
			ID: 3,
			Amount: 10_000,
			Category: "Auto",
			Status: types.StatusFAIL,
		},
		{
			ID: 4,
			Amount: 10_000,
			Category: "Cafe",
			Status: types.StatusOK,
		},
	}
	fmt.Println(TotalInCategory(payments, "Cafe"))
	//Output:
	//30000
}