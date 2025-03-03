package assignment

import "fmt"

func OrderFood(input string) {
	fmt.Print("What would you like to order? ")
	fmt.Scan(&input)
	fmt.Println("You ordered: ", input)
}

func main() {
	var OrderPrice int = 0
	var OrderItem string

	Menu := map[string]int{
		"pizza":  10,
		"burger": 5,
		"pasta":  8,
		"salad":  3,
	}

	OrderFood(OrderItem)

	for key, value := range Menu {
		if OrderItem == key {
			OrderPrice += value
		}
	}

}
