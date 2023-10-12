package handler

import (
	"github.com/AxelanO7/bn-shop-backend-web-go/database"
	"github.com/AxelanO7/bn-shop-backend-web-go/model"
	"github.com/gofiber/fiber/v2"
)

// find stockOpname by id
func findStockOpnameById(id string, stockOpname *model.StockOpname) error {
	db := database.DB.Db
	// find single stockOpname in the database by id
	db.Find(&stockOpname, "id = ?", id)
	// if no stockOpname found, return an error
	if stockOpname.ID == 0 {
		return fiber.ErrNotFound
	}
	return nil
}

// create a stockOpname
func CreateStockOpname(c *fiber.Ctx) error {
	db := database.DB.Db
	stockOpname := new(model.StockOpname)
	// store the body in the stockOpname and return error if encountered
	if err := c.BodyParser(stockOpname); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}
	// create stockOpname
	if err := db.Create(stockOpname).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create stockOpname", "data": err})
	}
	// return the created stockOpname
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "StockOpname has created", "data": stockOpname})
}

// get all StockOpnames from db
func GetAllStockOpnames(c *fiber.Ctx) error {
	db := database.DB.Db
	stockOpnames := []model.StockOpname{}
	if err := db.Find(&stockOpnames).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not get stockOpnames", "data": err})
	}
	// if no stockOpname found, return an error
	if len(stockOpnames) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "StockOpnames not found", "data": nil})
	}
	// return stockOpnames
	return c.Status(200).JSON(fiber.Map{"status": "sucess", "message": "StockOpnames Found", "data": stockOpnames})
}

// GetSingleStockOpname from db
func GetSingleStockOpname(c *fiber.Ctx) error {
	stockOpname := new(model.StockOpname)
	// get id params
	id := c.Params("id")
	// find single stockOpname in the database by id
	if err := findStockOpnameById(id, stockOpname); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "StockOpname not found"})
	}
	// return stockOpname
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "StockOpname Found", "data": stockOpname})
}

// update a stockOpname in db
func UpdateStockOpname(c *fiber.Ctx) error {
	db := database.DB.Db
	stockOpname := new(model.StockOpname)
	// get id params
	id := c.Params("id")
	// find single stockOpname in the database by id
	if err := findStockOpnameById(id, stockOpname); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "StockOpname not found"})
	}
	// store the body in the stockOpname and return error if encountered
	if err := c.BodyParser(stockOpname); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}
	// update stockOpname
	if err := db.Save(stockOpname).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not update stockOpname", "data": err})
	}
	// return the updated stockOpname
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "stockOpnames Found", "data": stockOpname})
}

// delete a stockOpname in db
func DeleteStockOpname(c *fiber.Ctx) error {
	db := database.DB.Db
	stockOpname := new(model.StockOpname)
	// get id params
	id := c.Params("id")
	// find single stockOpname in the database by id
	if err := findStockOpnameById(id, stockOpname); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "StockOpname not found"})
	}
	// delete stockOpname
	if err := db.Delete(stockOpname, "id = ?", id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not delete stockOpname", "data": err})
	}
	// return success message
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "StockOpname deleted"})
}
