package handler

import (
	"fmt"

	"github.com/AxelanO7/bn-shop-backend-web-go/database"
	"github.com/AxelanO7/bn-shop-backend-web-go/model"
	"github.com/gofiber/fiber/v2"
)

// find detail output by id
func findDetailOutputById(id string, detailOutput *model.DetailOutput) error {
	db := database.DB.Db
	// find single detail output in the database by id
	db.Find(&detailOutput, "id = ?", id)
	// if no detail output found, return an error
	if detailOutput.ID == 0 {
		return fiber.ErrNotFound
	}
	return nil
}

// create a detail output
func CreateDetailOutput(c *fiber.Ctx) error {
	db := database.DB.Db
	detailOutput := new(model.DetailOutput)
	output := new(model.Output)
	// store the body in the detail output and return error if encountered
	if err := c.BodyParser(detailOutput); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your output", "data": err})
	}
	// find output in the database by id
	if err := findOutputById(fmt.Sprint(detailOutput.IdOutput), output); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Output not found"})
	}
	// assign output to detail output
	detailOutput.Output = *output
	// create detail output
	if err := db.Create(detailOutput).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create detail output", "data": err})
	}
	// return the created detail output
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Detail output has created", "data": detailOutput})
}

// get all detail outputs from db
func GetAllDetailOutputs(c *fiber.Ctx) error {
	db := database.DB.Db
	detailOutputs := []model.DetailOutput{}
	// find all detail outputs in the database
	db.Find(detailOutputs)
	// if no detail output found, return an error
	if len(detailOutputs) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Detail outputs not found", "data": nil})
	}
	responseDetailOutputs := []model.DetailOutput{}
	for _, detailOutput := range detailOutputs {
		output := new(model.Output)
		// find  output in the database by id
		if err := findOutputById(fmt.Sprint(detailOutput.IdOutput), output); err != nil {
			return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Output not found"})
		}
		// assign  output to detail output
		detailOutput.Output = *output
		responseDetailOutputs = append(responseDetailOutputs, detailOutput)
	}
	// return detail outputs
	return c.Status(200).JSON(fiber.Map{"status": "sucess", "message": "Detail outputs Found", "data": detailOutputs})
}

// get single detail output from db
func GetSingleDetailOutput(c *fiber.Ctx) error {
	detailOutput := new(model.DetailOutput)
	output := new(model.Output)
	// get id params
	id := c.Params("id")
	// find single detail output in the database by id
	if err := findDetailOutputById(id, detailOutput); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Detail output not found"})
	}
	// find  output in the database by id
	if err := findOutputById(fmt.Sprint(detailOutput.IdOutput), output); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Output not found"})
	}
	// assign  output to detail output
	detailOutput.Output = *output
	// return detail output
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Detail output Found", "data": detailOutput})
}

// update a detail output in db
func UpdateDetailOutput(c *fiber.Ctx) error {
	db := database.DB.Db
	detailOutput := new(model.DetailOutput)
	output := new(model.Output)
	// get id params
	id := c.Params("id")
	// find single detail output in the database by id
	if err := findDetailOutputById(id, detailOutput); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Detail output not found"})
	}
	// store the body in the detail output and return error if encountered
	if err := c.BodyParser(detailOutput); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your output", "data": err})
	}
	// find  output in the database by id
	if err := findOutputById(fmt.Sprint(detailOutput.IdOutput), output); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Output not found"})
	}
	// assign  output to detail output
	detailOutput.Output = *output
	// update detail output
	if err := db.Save(detailOutput).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not update detail output", "data": err})
	}
	// return the updated detail output
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Detail output has updated", "data": detailOutput})
}

// delete a detail output in db
func DeleteDetailOutput(c *fiber.Ctx) error {
	db := database.DB.Db
	detailOutput := new(model.DetailOutput)
	// get id params
	id := c.Params("id")
	// find single detail output in the database by id
	if err := findDetailOutputById(id, detailOutput); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Detail output not found"})
	}
	// delete detail output
	if err := db.Delete(detailOutput).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not delete detail output", "data": err})
	}
	// return the deleted detail output
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Detail output has deleted", "data": detailOutput})
}
