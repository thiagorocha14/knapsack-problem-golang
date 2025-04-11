package controllers

import (
	"github.com/gin-gonic/gin"
	"knapsack-problem/requests"
	"knapsack-problem/services"
	"time"
)

func SolveKnapsackProblemController(c *gin.Context) {

	var request requests.SolveKnapsackProblemRequest

	time_start := time.Now()

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	knapsack_capacity := request.KnapsackCapacity
	products := request.Products

	knapsack := services.SolveKnapsackProblem(knapsack_capacity, products)

	time_end := time.Now()

	time_elapsed := time_end.Sub(time_start)
	c.JSON(200, gin.H{
		"solution": knapsack,
		"message": "Knapsack problem solved successfully",
		"execution_time": time_elapsed.String(),
	})
}

