package rest

import (
	"demo-hex-go/internal/core/domain/entity"

	"demo-hex-go/internal/core/port"

	"github.com/gofiber/fiber/v2"
)

type ProductRest struct {
	productService port.ProductService
}

func NewProductRest(productSeervice port.ProductService) *ProductRest {
	return &ProductRest{
		productService: productSeervice,
	}
}

func (rest *ProductRest) SaveProduct(c *fiber.Ctx) error {
	product := new(entity.Product)
	err := c.BodyParser(product)
	if err != nil {
		return c.Status(fiber.StatusServiceUnavailable).JSON(err.Error())
	}

	if product.Id == "" {
		rest.productService.SaveProduct(*product)
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Success Created new Product",
		})
	} else {
		rest.productService.UpdateProduct(*product)
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Success Updated new Product",
		})
	}

}

func (rest *ProductRest) ReadProducts(c *fiber.Ctx) error {
	products, err := rest.productService.ReadProducts()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message":    "Product not found",
			"errMessage": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success",
		"data":    products,
	})
}

func (rest *ProductRest) ReadProduct(c *fiber.Ctx) error {

	id := c.Params("id")
	product, err := rest.productService.ReadProduct(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message":    "Product not found",
			"errMessage": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success",
		"data":    product,
	})
}

func (rest *ProductRest) DeleteProduct(c *fiber.Ctx) error {

	id := c.Params("id")
	err := rest.productService.DeleteProduct(id)
	if err != nil {
		return c.Status(fiber.StatusServiceUnavailable).JSON(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success Deleted Data",
	})

}
