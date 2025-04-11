package requests

import "knapsack-problem/dtos"

type SolveKnapsackProblemRequest struct {
	KnapsackCapacity int `json:"knapsack_capacity"`
	Products []dtos.ProductDTO `json:"products"`
}




