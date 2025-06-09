package main

import (
	"encoding/json"
	"fmt"
	"os"
	"pizzaria/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	loadPizzas()
	router := gin.Default()
	router.GET("/pizzas", getPizzas)
	router.GET("/pizzas/:id", getPizzasById)
	router.POST("/pizzas", postPizzas)
	router.Run()
}

var pizzas []models.Pizza

func getPizzas(c *gin.Context) {
	c.JSON(200, gin.H{
		"pizzas": pizzas,
	})
}

func postPizzas(c *gin.Context) {
	var newPizza models.Pizza
	if err := c.ShouldBindJSON(&newPizza); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	newPizza.Id = len(pizzas) + 1 // Simple ID generation
	pizzas = append(pizzas, newPizza)
	savePizza()
	c.JSON(201, gin.H{
		"message": "Pizza created successfully",
	})
}

func getPizzasById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error()})
		return
	}
	for _, pizza := range pizzas {
		if pizza.Id == id {
			c.JSON(200, pizza)
			return
		}
	}
	c.JSON(404, gin.H{
		"error": "Pizza not found",
	})
}

func loadPizzas() {
	file, err := os.Open("/dados/pizza.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&pizzas); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
}

func savePizza() {
	file, err := os.Create("/dados/pizza.json")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	if err := encoder.Encode(pizzas); err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}
}
