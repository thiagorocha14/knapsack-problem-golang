package services

import "knapsack-problem/dtos"

func CreateCombination(index int, products []dtos.ProductDTO) dtos.CombinationDTO {
	var bases []int = []int{}

	for _, product := range products {
		bases = append(bases, product.Quantity)
	}

	selected_products := selectProducts(index, bases)
	price, weight := calculatePriceAndWeight(selected_products, products)

	return dtos.CombinationDTO{
		Products: selected_products,
		Price: price,
		Weight: weight,
	}
}

func selectProducts(index int, bases []int) []int {
	selected_products := []int{}
	quotient := index
	reverse(bases)

	for _, base := range bases {
		rest := quotient % (base + 1)
		quotient = quotient / (base + 1)
		selected_products = append([]int{rest}, selected_products...)
	}
	reverse(bases)

	return selected_products
}

func calculatePriceAndWeight(selected_products []int, products []dtos.ProductDTO) (int, int) {
	price := 0
	weight := 0

	for index, quantity := range selected_products {
		product := products[index]
		price += product.Price * quantity
		weight += product.Weight * quantity
	}

	return price, weight
}

func reverse(input []int) []int {
	n := len(input)

	for i := 0; i < n/2; i++ {
		input[i], input[n-i-1] = input[n-i-1], input[i]
	}

	return input
}
