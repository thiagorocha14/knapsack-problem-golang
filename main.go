package main

import (
	"github.com/gin-gonic/gin"
	"knapsack-problem/controllers"
  "github.com/gin-contrib/cors"
)

func main() {
  router := gin.Default()
  config := cors.DefaultConfig()
  config.AllowOrigins = []string{"http://localhost:5173"}

  router.Use(cors.New(config))
  api := router.Group("/api")
  {
	api.POST("/solve", controllers.SolveKnapsackProblemController)
  }
  router.Run()
}