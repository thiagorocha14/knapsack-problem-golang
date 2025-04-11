package dtos

type CombinationDTO struct {
	Price int `json:"price"`
	Weight int `json:"weight"`
	Products []int `json:"products"`
}
