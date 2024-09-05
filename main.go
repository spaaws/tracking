package main

import (
    "github.com/gin-gonic/gin"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
	"spaaws.com/tracking/handlers"
)

var db *gorm.DB

func main() {
    dsn := "root:root_password@tcp(127.0.0.1:3306)/tracking?charset=utf8mb4&parseTime=True&loc=Local"
    var err error
    db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("failed to connect to database")
    }

    // Auto migrate
    db.AutoMigrate(&handlers.User{}, &handlers.Order{})

    r := gin.Default()

    // Rotas de usuários
    r.POST("/user", func(c *gin.Context) { handlers.CreateUser(c, db) })
    r.GET("/user/:id", func(c *gin.Context) { handlers.GetUser(c, db) })
    r.PUT("/user/:id", func(c *gin.Context) { handlers.UpdateUser(c, db) })
    r.DELETE("/user/:id", func(c *gin.Context) { handlers.DeleteUser(c, db) })

    // Rotas de encomendas
    r.POST("/order", func(c *gin.Context) { handlers.CreateOrder(c, db) })
    r.GET("/order/:id", func(c *gin.Context) { handlers.GetOrder(c, db) })
    r.PUT("/order/:id", func(c *gin.Context) { handlers.UpdateOrder(c, db) })
    r.DELETE("/order/:id", func(c *gin.Context) { handlers.DeleteOrder(c, db) })
	// r.GET("/user/:id/order", func(c *gin.Context) { handlers.GetOrdersByUserID(c, db) })

    r.Run(":8080")
}
