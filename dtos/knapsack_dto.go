package dtos

type KnapsackDTO struct {
	Capacity int `json:"capacity"`
	NumberOfCombinations int `json:"number_of_combinations"`
	Combination CombinationDTO `json:"combination"`
}

