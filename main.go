package main

import (
	"github.com/gin-gonic/gin"
	"knapsack-problem/controllers"
)

func main() {
  router := gin.Default()
  api := router.Group("/api")
  {
	api.POST("/solve", controllers.SolveKnapsackProblemController)
  }
  router.Run()
}