package routes

import (
	"jual-beli-barang-bekas/controllers"
	"jual-beli-barang-bekas/middlewares"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	//Auth
	r.POST("/register", func(c *gin.Context) {
		controllers.Register(c, db)
	})
	r.POST("/login", func(c *gin.Context) {
		controllers.Login(c, db)
	})

	//item
	r.GET("/items", middlewares.AuthMiddleware, func(c *gin.Context) {
		controllers.GetItems(c, db)
	})
	r.POST("/items", middlewares.AuthMiddleware, func(c *gin.Context) {
		controllers.CreateItem(c, db)
	})
	r.PUT("/items/:id", middlewares.AuthMiddleware, func(c *gin.Context) {
		controllers.UpdateItem(c, db)
	})
	r.DELETE("/items/:id", middlewares.AuthMiddleware, func(c *gin.Context) {
		controllers.DeleteItem(c, db)
	})

	//transaction
	r.GET("/transactions", middlewares.AuthMiddleware, func(c *gin.Context) {
		controllers.GetTransactions(c, db)
	})
	r.POST("/transactions", middlewares.AuthMiddleware, func(c *gin.Context) {
		controllers.CreateTransaction(c, db)
	})
}
