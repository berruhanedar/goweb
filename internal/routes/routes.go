package routes

import (
	"github.com/berruhanedar/goweb/internal/handler"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	e.GET("/products", handler.GetProducts)
	e.GET("/products/:id", handler.GetProduct)
	e.POST("/products", handler.CreateProduct)
	e.PUT("/products/:id", handler.UpdateProduct)
	e.DELETE("/products/:id", handler.DeleteProduct)
	e.PATCH("/products/:id", handler.PatchProduct)
}
