package handler

import (
	"fmt"

	"github.com/AxelanO7/bn-shop-backend-web-go/database"
	"github.com/AxelanO7/bn-shop-backend-web-go/model"
	"github.com/gofiber/fiber/v2"
)

// find detail input by id
func findDetailInputById(id string, detailInput *model.DetailInput) error {
	db := database.DB.Db
	// find single detail input in the database by id
	db.Find(&detailInput, "id = ?", id)
	// if no detail input found, return an error
	if detailInput.ID == 0 {
		return fiber.ErrNotFound
	}
	return nil
}

// create a detail input
func CreateDetailInput(c *fiber.Ctx) error {
	db := database.DB.Db
	detailInput := new(model.DetailInput)
	input := new(model.Input)
	// store the body in the detail input and return error if encountered
	if err := c.BodyParser(detailInput); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}
	// find input in the database by id
	if err := findInputById(fmt.Sprint(detailInput.IdInput), input); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Input not found"})
	}
	// assign input to detail input
	detailInput.Input = *input
	// create detail input
	if err := db.Create(detailInput).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create detail input", "data": err})
	}
	// return the created detail input
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Detail input has created", "data": detailInput})
}

// get all detail inputs from db
func GetAllDetailInputs(c *fiber.Ctx) error {
	db := database.DB.Db
	detailInputs := []model.DetailInput{}
	// find all detail inputs in the database
	db.Find(detailInputs)
	// if no detail input found, return an error
	if len(detailInputs) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Detail inputs not found", "data": nil})
	}
	responseDetailInputs := []model.DetailInput{}
	for _, detailInput := range detailInputs {
		input := new(model.Input)
		// find  input in the database by id
		if err := findInputById(fmt.Sprint(detailInput.IdInput), input); err != nil {
			return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Input not found"})
		}
		// assign  input to detail input
		detailInput.Input = *input
		responseDetailInputs = append(responseDetailInputs, detailInput)
	}
	// return detail inputs
	return c.Status(200).JSON(fiber.Map{"status": "sucess", "message": "Detail inputs Found", "data": detailInputs})
}

// get single detail input from db
func GetSingleDetailInput(c *fiber.Ctx) error {
	detailInput := new(model.DetailInput)
	input := new(model.Input)
	// get id params
	id := c.Params("id")
	// find single detail input in the database by id
	if err := findDetailInputById(id, detailInput); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Detail input not found"})
	}
	// find  input in the database by id
	if err := findInputById(fmt.Sprint(detailInput.IdInput), input); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Input not found"})
	}
	// assign  input to detail input
	detailInput.Input = *input
	// return detail input
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Detail input Found", "data": detailInput})
}

// update a detail input in db
func UpdateDetailInput(c *fiber.Ctx) error {
	db := database.DB.Db
	detailInput := new(model.DetailInput)
	input := new(model.Input)
	// get id params
	id := c.Params("id")
	// find single detail input in the database by id
	if err := findDetailInputById(id, detailInput); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Detail input not found"})
	}
	// store the body in the detail input and return error if encountered
	if err := c.BodyParser(detailInput); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}
	// find  input in the database by id
	if err := findInputById(fmt.Sprint(detailInput.IdInput), input); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Input not found"})
	}
	// assign  input to detail input
	detailInput.Input = *input
	// update detail input
	if err := db.Save(detailInput).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not update detail input", "data": err})
	}
	// return the updated detail input
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Detail input has updated", "data": detailInput})
}

// delete a detail input in db
func DeleteDetailInput(c *fiber.Ctx) error {
	db := database.DB.Db
	detailInput := new(model.DetailInput)
	// get id params
	id := c.Params("id")
	// find single detail input in the database by id
	if err := findDetailInputById(id, detailInput); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Detail input not found"})
	}
	// delete detail input
	if err := db.Delete(detailInput).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not delete detail input", "data": err})
	}
	// return the deleted detail input
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Detail input has deleted", "data": detailInput})
}
