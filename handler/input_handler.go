package handler

import (
	"fmt"

	"github.com/AxelanO7/bn-shop-backend-web-go/database"
	"github.com/AxelanO7/bn-shop-backend-web-go/model"
	"github.com/gofiber/fiber/v2"
)

// find input by id
func findInputById(id string, input *model.Input) error {
	db := database.DB.Db
	// find single input in the database by id
	db.Find(&input, "id = ?", id)
	// if no input found, return an error
	if input.ID == 0 {
		return fiber.ErrNotFound
	}
	return nil
}

// create a input
func CreateInput(c *fiber.Ctx) error {
	db := database.DB.Db
	input := new(model.Input)
	// store the body in the input and return error if encountered
	if err := c.BodyParser(input); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}
	// create input
	if err := db.Create(input).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create input", "data": err})
	}
	// return the created input
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Input has created", "data": input})
}

// get all Inputs from db
func GetAllInputs(c *fiber.Ctx) error {
	db := database.DB.Db
	inputs := []model.Input{}
	// find all inputs in the database
	if err := db.Find(&inputs).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not get inputs", "data": err})
	}
	// if no input found, return an error
	if len(inputs) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Inputs not found", "data": nil})
	}
	// return inputs
	return c.Status(200).JSON(fiber.Map{"status": "sucess", "message": "Inputs Found", "data": inputs})
}

// get single input from db
func GetSingleInput(c *fiber.Ctx) error {
	input := new(model.Input)
	// get id params
	id := c.Params("id")
	// find single input in the database by id
	if err := findInputById(id, input); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Input not found"})
	}
	// return input
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Input Found", "data": input})
}

// update a input in db
func UpdateInput(c *fiber.Ctx) error {
	db := database.DB.Db
	input := new(model.Input)
	// get id params
	id := c.Params("id")
	// find single input in the database by id
	if err := findInputById(id, input); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Input not found"})
	}
	// store the body in the input and return error if encountered
	if err := c.BodyParser(input); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}
	// update input
	if err := db.Save(input).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not update input", "data": err})
	}
	// return the updated input
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Input has updated", "data": input})
}

// delete a input in db
func DeleteInput(c *fiber.Ctx) error {
	db := database.DB.Db
	input := new(model.Input)
	// get id params
	id := c.Params("id")
	// find single input in the database by id
	if err := findInputById(id, input); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Input not found"})
	}
	// delete input
	if err := db.Delete(input).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not delete input", "data": err})
	}
	// return deleted input
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Input has deleted", "data": nil})
}

func GetInputByDate(c *fiber.Ctx) error {
	db := database.DB.Db
	detailInputs := []model.DetailInput{}
	dateStart := c.Query("date-start")
	dateEnd := c.Query("date-end")

	// find all detailInputs in the database
	if err := db.Find(&detailInputs, "created_at BETWEEN ? AND ?", dateStart, dateEnd).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not get detailInputs", "data": err})
	}
	// if no input found, return an error
	if len(detailInputs) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Inputs not found", "data": nil})
	}
	responseDetailInputs := []model.DetailInput{}
	for _, detailInput := range detailInputs {
		input := new(model.Input)
		// find single input in the database by id
		if err := findInputById(fmt.Sprint(detailInput.IdInput), input); err != nil {
			return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Input not found"})
		}
		// assign input to detailInput
		detailInput.Input = *input
		responseDetailInputs = append(responseDetailInputs, detailInput)
	}
	// return detailInputs
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Inputs Found", "data": responseDetailInputs})
}
