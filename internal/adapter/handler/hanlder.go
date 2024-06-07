package handler

import (
	"demo-hex-go/internal/adapter/handler/rest"
	"demo-hex-go/internal/core/port"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(productService port.ProductService) *fiber.App {
	router := fiber.New()
	rest := rest.NewProductRest(productService)

	router.Get("/api2/", welcomeHandler2)
	router.Get("/api/products", rest.ReadProducts)
	router.Get("/api/product/:id", rest.ReadProduct)
	router.Post("/api/product", rest.SaveProduct)
	router.Delete("/api/product/:id", rest.DeleteProduct)
	return router
}

func welcomeHandler2(c *fiber.Ctx) error {
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"message": "Test API!",
	})
}
