package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Order struct
type Order struct {
	ID        uint    `json:"id" gorm:"primaryKey"`
	Symbol    string  `json:"symbol"`
	Price     float64 `json:"price"`
	Quantity  int     `json:"quantity"`
	OrderType string  `json:"order_type"`
}

var db *gorm.DB
var err error

func init() {
	dsn := "host=host.docker.internal user=postgres password=Cse@40668 dbname=orders_db port=5432 sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	db.AutoMigrate(&Order{})

}

func createOrder(c *gin.Context) {
	var order Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Create(&order)
	c.JSON(http.StatusOK, order)
}

func getOrders(c *gin.Context) {
	var orders []Order
	db.Find(&orders)
	c.JSON(http.StatusOK, orders)
}

func main() {
	r := gin.Default()
	r.POST("/orders", createOrder)
	r.GET("/orders", getOrders)

	fmt.Println("Server running on port 8000...")
	r.Run("0.0.0.0:8000")
}
