package handler

import (
	"github.com/AxelanO7/bn-shop-backend-web-go/database"
	"github.com/AxelanO7/bn-shop-backend-web-go/model"
	"github.com/gofiber/fiber/v2"
)

// find supplier by id
func findSupplierById(id string, supplier *model.Supplier) error {
	db := database.DB.Db
	// find single supplier in the database by id
	db.Find(&supplier, "id = ?", id)
	// if no supplier found, return an error
	if supplier.ID == 0 {
		return fiber.ErrNotFound
	}
	return nil
}

// create a supplier
func CreateSupplier(c *fiber.Ctx) error {
	db := database.DB.Db
	supplier := new(model.Supplier)
	// store the body in the supplier and return error if encountered
	if err := c.BodyParser(supplier); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}
	// create supplier
	if err := db.Create(supplier).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create supplier", "data": err})
	}
	// return the created supplier
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Supplier has created", "data": supplier})
}

// get all Suppliers from db
func GetAllSuppliers(c *fiber.Ctx) error {
	db := database.DB.Db
	suppliers := []model.Supplier{}
	if err := db.Find(suppliers).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not get suppliers", "data": err})
	}
	// if no supplier found, return an error
	if len(suppliers) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Suppliers not found", "data": nil})
	}
	// return suppliers
	return c.Status(200).JSON(fiber.Map{"status": "sucess", "message": "Suppliers Found", "data": suppliers})
}

// GetSingleSupplier from db
func GetSingleSupplier(c *fiber.Ctx) error {
	supplier := new(model.Supplier)
	// get id params
	id := c.Params("id")
	// find single supplier in the database by id
	if err := findSupplierById(id, supplier); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Supplier not found"})
	}
	// return supplier
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Supplier Found", "data": supplier})
}

// update a supplier in db
func UpdateSupplier(c *fiber.Ctx) error {
	db := database.DB.Db
	supplier := new(model.Supplier)
	// get id params
	id := c.Params("id")
	// find single supplier in the database by id
	if err := findSupplierById(id, supplier); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Supplier not found"})
	}
	// store the body in the supplier and return error if encountered
	if err := c.BodyParser(supplier); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}
	// update supplier
	if err := db.Save(supplier).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not update supplier", "data": err})
	}
	// return the updated supplier
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "suppliers Found", "data": supplier})
}

// delete a supplier in db
func DeleteSupplier(c *fiber.Ctx) error {
	db := database.DB.Db
	supplier := new(model.Supplier)
	// get id params
	id := c.Params("id")
	// find single supplier in the database by id
	if err := findSupplierById(id, supplier); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Supplier not found"})
	}
	// delete supplier
	if err := db.Delete(supplier, "id = ?", id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not delete supplier", "data": err})
	}
	// return success message
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Supplier deleted"})
}
