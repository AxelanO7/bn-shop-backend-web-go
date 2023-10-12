package handler

import (
	"github.com/AxelanO7/bn-shop-backend-web-go/database"
	"github.com/AxelanO7/bn-shop-backend-web-go/model"
	"github.com/gofiber/fiber/v2"
)

// find stock by id
func findStockById(id string, stock *model.Stock) error {
	db := database.DB.Db
	// find single stock in the database by id
	db.Find(&stock, "id = ?", id)
	// if no stock found, return an error
	if stock.ID == 0 {
		return fiber.ErrNotFound
	}
	return nil
}

// create a stock
func CreateStock(c *fiber.Ctx) error {
	db := database.DB.Db
	stock := new(model.Stock)
	// store the body in the stock and return error if encountered
	if err := c.BodyParser(stock); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}
	// create stock
	if err := db.Create(stock).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create stock", "data": err})
	}
	// return the created stock
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Stock has created", "data": stock})
}

// get all Stocks from db
func GetAllStocks(c *fiber.Ctx) error {
	db := database.DB.Db
	stocks := []model.Stock{}
	if err := db.Find(&stocks).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not get stocks", "data": err})
	}
	// if no stock found, return an error
	if len(stocks) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Stocks not found", "data": nil})
	}
	// return stocks
	return c.Status(200).JSON(fiber.Map{"status": "sucess", "message": "Stocks Found", "data": stocks})
}

// GetSingleStock from db
func GetSingleStock(c *fiber.Ctx) error {
	stock := new(model.Stock)
	// get id params
	id := c.Params("id")
	// find single stock in the database by id
	if err := findStockById(id, stock); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Stock not found"})
	}
	// return stock
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Stock Found", "data": stock})
}

// update a stock in db
func UpdateStock(c *fiber.Ctx) error {
	db := database.DB.Db
	stock := new(model.Stock)
	// get id params
	id := c.Params("id")
	// find single stock in the database by id
	if err := findStockById(id, stock); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Stock not found"})
	}
	// store the body in the stock and return error if encountered
	if err := c.BodyParser(stock); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}
	// update stock
	if err := db.Save(stock).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not update stock", "data": err})
	}
	// return the updated stock
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "stocks Found", "data": stock})
}

// delete a stock in db
func DeleteStock(c *fiber.Ctx) error {
	db := database.DB.Db
	stock := new(model.Stock)
	// get id params
	id := c.Params("id")
	// find single stock in the database by id
	if err := findStockById(id, stock); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Stock not found"})
	}
	// delete stock
	if err := db.Delete(stock, "id = ?", id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not delete stock", "data": err})
	}
	// return success message
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Stock deleted"})
}
