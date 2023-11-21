package main

import "fmt"

func isValidCard(number_card int64) bool {
	var sum int
	var numbers_of_card [16]int
	for i := 15; i >= 0; i-- {
		numbers_of_card[i] = (int)(number_card % 10)
		number_card = number_card / 10
	}
	for i := 14; i >= 0; i -= 2 {
		numbers_of_card[i] *= 2
		if numbers_of_card[i] >= 10 {
			sum += numbers_of_card[i] % 10
			numbers_of_card[i] /= 10
			sum += numbers_of_card[i]
		} else if numbers_of_card[i] < 10 {
			sum += numbers_of_card[i]
		}
	}
	for i := 15; i >= 0; i -= 2 {
		sum += numbers_of_card[i]
	}
	return sum%10 == 0
}

func main() {
	var number_card int64
	var choice int
	fmt.Println("Card Validator. \n",
		"Choose the operation: \n",
		"Check for the validity of the card (1)?\n",
		"Check the origin of the card (2)?\n",
		"Both (3)?")
	fmt.Scanf("%d", &choice)
	fmt.Println("Enter the card number: ")
	fmt.Scanf("%d", &number_card)
	if isValidCard(number_card) {
		fmt.Println("Card is valid.")
	} else {
		fmt.Println("Card is invalid.")
	}

}
