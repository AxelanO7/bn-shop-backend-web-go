package handler

import (
	"github.com/AxelanO7/bn-shop-backend-web-go/database"
	"github.com/AxelanO7/bn-shop-backend-web-go/model"
	"github.com/gofiber/fiber/v2"
)

// find output by id
func findOutputById(id string, output *model.Output) error {
	db := database.DB.Db
	// find single output in the database by id
	db.Find(&output, "id = ?", id)
	// if no output found, return an error
	if output.ID == 0 {
		return fiber.ErrNotFound
	}
	return nil
}

// create a output
func CreateOutput(c *fiber.Ctx) error {
	db := database.DB.Db
	output := new(model.Output)
	// store the body in the output and return error if encountered
	if err := c.BodyParser(output); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your output", "data": err})
	}
	// create output
	if err := db.Create(output).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create output", "data": err})
	}
	// return the created output
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Output has created", "data": output})
}

// get all Outputs from db
func GetAllOutputs(c *fiber.Ctx) error {
	db := database.DB.Db
	outputs := []model.Output{}
	// find all outputs in the database
	if err := db.Find(&outputs).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not get outputs", "data": err})
	}
	// if no output found, return an error
	if len(outputs) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Outputs not found", "data": nil})
	}
	// return outputs
	return c.Status(200).JSON(fiber.Map{"status": "sucess", "message": "Outputs Found", "data": outputs})
}

// get single output from db
func GetSingleOutput(c *fiber.Ctx) error {
	output := new(model.Output)
	// get id params
	id := c.Params("id")
	// find single output in the database by id
	if err := findOutputById(id, output); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Output not found"})
	}
	// return output
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Output Found", "data": output})
}

// update a output in db
func UpdateOutput(c *fiber.Ctx) error {
	db := database.DB.Db
	output := new(model.Output)
	// get id params
	id := c.Params("id")
	// find single output in the database by id
	if err := findOutputById(id, output); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Output not found"})
	}
	// store the body in the output and return error if encountered
	if err := c.BodyParser(output); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your output", "data": err})
	}
	// update output
	if err := db.Save(output).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not update output", "data": err})
	}
	// return the updated output
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Output has updated", "data": output})
}

// delete a output in db
func DeleteOutput(c *fiber.Ctx) error {
	db := database.DB.Db
	output := new(model.Output)
	// get id params
	id := c.Params("id")
	// find single output in the database by id
	if err := findOutputById(id, output); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Output not found"})
	}
	// delete output
	if err := db.Delete(output).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not delete output", "data": err})
	}
	// return deleted output
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Output has deleted", "data": nil})
}
