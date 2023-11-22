package handler

import (
	"fmt"

	"github.com/AxelanO7/bn-shop-backend-web-go/database"
	"github.com/AxelanO7/bn-shop-backend-web-go/model"
	"github.com/gofiber/fiber/v2"
)

// create multiple stockOpnames
func CreateMultipleStockOpnames(c *fiber.Ctx) error {
	db := database.DB.Db
	stockOpnames := new([]model.DetailOpname)
	stockopname := new(model.StockOpname)
	// store the body in the stockOpnames and return error if encountered
	if err := c.BodyParser(stockOpnames); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your stockOpnames", "data": err})
	}
	for _, stockOpname := range *stockOpnames {
		// find single stockOpname in the database by id
		if err := findStockOpnameById(fmt.Sprint(stockOpname.IdOpname), stockopname); err != nil {
			return c.Status(404).JSON(fiber.Map{"status": "error", "message": "StockOpname not found"})
		}
		// assign stockOpname to detail opname
		stockOpname.Opname = *stockopname
		// create stockOpname
		if err := db.Create(&stockOpname).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create stockOpname", "data": err})
		}
	}
	// return the created stockOpnames
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "StockOpnames has created", "data": stockOpnames})
}

// remove multiple stock from stockOpname
func RemoveMultipleStockFromStockOpname(c *fiber.Ctx) error {
	db := database.DB.Db
	detailOutputs := new([]model.DetailOutput)
	stocks := []model.Stock{}
	// store the body in the detail outputs and return error if encountered
	if err := c.BodyParser(detailOutputs); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your detail outputs", "data": err})
	}
	// if no stock found, return an error
	if err := db.Find(&stocks).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Stocks not found", "data": nil})
	}
	// if no stock found, return an error
	if len(stocks) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Stocks not found", "data": nil})
	}
	for _, detailOutput := range *detailOutputs {
		// remove total product from stock
		for _, stock := range stocks {
			if stock.CodeProduct == detailOutput.CodeProduct {
				stock.TotalProduct = detailOutput.TotalUsed
				// update stock
				if err := db.Save(&stock).Error; err != nil {
					return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not update stock", "data": err})
				}
			}
		}
	}
	// return the created detail outputs
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Detail outputs has created", "data": detailOutputs})
}
