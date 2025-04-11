package services

import (
	"knapsack-problem/dtos"
)

var knapsack dtos.KnapsackDTO = dtos.KnapsackDTO{
	Capacity: 0,
	NumberOfCombinations: 0,
	Combination: dtos.CombinationDTO{
		Products: []int{},
		Price: 0,
		Weight: 0,
	},
}

func SolveKnapsackProblem(knapsack_capacity int, products []dtos.ProductDTO) dtos.KnapsackDTO {
	knapsack.Capacity = knapsack_capacity
	knapsack.NumberOfCombinations = calculateNumberOfCombinations(products)

	begin := 0
	end := knapsack.NumberOfCombinations

	for i := begin; i < end; i++ {
		combination := CreateCombination(i, products)

		if !canStoreProduct(combination, knapsack_capacity) {
			continue
		}

		if isBetterCombination(combination, knapsack) {
			knapsack.Combination = combination
		}
	}

	return knapsack
}

func calculateNumberOfCombinations(products []dtos.ProductDTO) int {
	var quantities []int = []int{}

	for _, product := range products {
		quantities = append(quantities, (product.Quantity + 1))
	}

	var number_of_combinations int = 1

	for _, quantity := range quantities {
		number_of_combinations *= quantity
	}

	return number_of_combinations
}

func canStoreProduct(combination dtos.CombinationDTO, knapsack_capacity int) bool {
	return combination.Weight <= knapsack_capacity
}

func isBetterCombination(combination dtos.CombinationDTO, knapsack dtos.KnapsackDTO) bool {
	if combination.Price > knapsack.Combination.Price {
		return true
	}

	if combination.Price == knapsack.Combination.Price {
		if combination.Weight < knapsack.Combination.Weight {
			return true
		}
	}

	return false
}
