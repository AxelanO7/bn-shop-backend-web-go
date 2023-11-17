package handler

import (
	"fmt"

	"github.com/AxelanO7/bn-shop-backend-web-go/database"
	"github.com/AxelanO7/bn-shop-backend-web-go/model"
	"github.com/gofiber/fiber/v2"
)

// find detail order by id
func findDetailOrderById(id string, detailOrder *model.DetailOrder) error {
	db := database.DB.Db
	// find single detail order in the database by id
	db.Find(&detailOrder, "id = ?", id)
	// if no detail order found, return an error
	if detailOrder.ID == 0 {
		return fiber.ErrNotFound
	}
	return nil
}

// create a detail order
func CreateDetailOrder(c *fiber.Ctx) error {
	db := database.DB.Db
	detailOrder := new(model.DetailOrder)
	order := new(model.Order)
	// store the body in the detail order and return error if encountered
	if err := c.BodyParser(detailOrder); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}
	// find order in the database by id
	if err := findOrderById(fmt.Sprint(detailOrder.IdOrder), order); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Order not found"})
	}
	// assign order to detail order
	detailOrder.Order = *order
	// create detail order
	if err := db.Create(detailOrder).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create detail order", "data": err})
	}
	// return the created detail order
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Detail order has created", "data": detailOrder})
}

// create multiple detail orders
func CreateMultipleDetailOrders(c *fiber.Ctx) error {
	db := database.DB.Db
	detailOrders := new([]model.DetailOrder)
	// store the body in the detail orders and return error if encountered
	if err := c.BodyParser(detailOrders); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}

	for _, detailOrder := range *detailOrders {
		order := new(model.Order)
		// find order in the database by id
		if err := findOrderById(fmt.Sprint(detailOrder.IdOrder), order); err != nil {
			return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Order not found"})
		}
		// assign order to detail order
		detailOrder.Order = *order
		// create detail order
		if err := db.Create(&detailOrder).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create detail order", "data": err})
		}
	}
	// return the created detail orders
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Detail orders has created", "data": detailOrders})
}

// get all detail orders from db
func GetAllDetailOrders(c *fiber.Ctx) error {
	db := database.DB.Db
	detailOrders := []model.DetailOrder{}
	// find all detail orders in the database
	if err := db.Find(&detailOrders).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not get detail orders", "data": err})
	}
	// if no detail order found, return an error
	if len(detailOrders) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Detail orders not found", "data": nil})
	}
	responseDetailOrders := []model.DetailOrder{}
	for _, detailOrder := range detailOrders {
		order := new(model.Order)
		// find  order in the database by id
		if err := findOrderById(fmt.Sprint(detailOrder.IdOrder), order); err != nil {
			return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Order not found"})
		}
		// assign  order to detail order
		detailOrder.Order = *order
		responseDetailOrders = append(responseDetailOrders, detailOrder)
	}
	// return detail orders
	return c.Status(200).JSON(fiber.Map{"status": "sucess", "message": "Detail orders Found", "data": detailOrders})
}

// get detail order by id order from db
func GetAllDetailOrdersByOrder(c *fiber.Ctx) error {
	db := database.DB.Db
	detailOrders := []model.DetailOrder{}
	// get id params
	id := c.Params("id")
	// find all detail orders in the database by id
	if err := db.Find(&detailOrders, "id_order = ?", id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not get detail orders", "data": err})
	}
	// if no detail order found, return an error
	if len(detailOrders) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Detail orders not found", "data": nil})
	}
	responseDetailOrders := []model.DetailOrder{}
	for _, detailOrder := range detailOrders {
		order := new(model.Order)
		// find  order in the database by id
		if err := findOrderById(fmt.Sprint(detailOrder.IdOrder), order); err != nil {
			return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Order not found"})
		}
		// assign  order to detail order
		detailOrder.Order = *order
		responseDetailOrders = append(responseDetailOrders, detailOrder)
	}
	// return detail orders
	return c.Status(200).JSON(fiber.Map{"status": "sucess", "message": "Detail orders Found", "data": detailOrders})
}

// get all order by start date and end date from db
func GetOrderByDate(c *fiber.Ctx) error {
	// return c.Status(200).JSON(fiber.Map{"status": "sucess", "message": "Detail orders Found"})
	db := database.DB.Db
	detailOrders := []model.DetailOrder{}
	// get start date params
	dateStart := c.Params("date-start")
	// get end date params
	dateEnd := c.Params("date-end")
	// find all detail orders in the database by id
	if err := db.Find(&detailOrders, "created_at BETWEEN ? AND ?", dateStart, dateEnd).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not get detail orders", "data": err})
	}
	// if no detail order found, return an error
	if len(detailOrders) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Detail orders not found", "data": nil})
	}
	responseDetailOrders := []model.DetailOrder{}
	for _, detailOrder := range detailOrders {
		order := new(model.Order)
		// find  order in the database by id
		if err := findOrderById(fmt.Sprint(detailOrder.IdOrder), order); err != nil {
			return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Order not found"})
		}
		// assign  order to detail order
		detailOrder.Order = *order
		responseDetailOrders = append(responseDetailOrders, detailOrder)
	}
	// return detail orders
	return c.Status(200).JSON(fiber.Map{"status": "sucess", "message": "Detail orders Found", "data": detailOrders})
}

// get single detail order from db
func GetSingleDetailOrder(c *fiber.Ctx) error {
	detailOrder := new(model.DetailOrder)
	order := new(model.Order)
	// get id params
	id := c.Params("id")
	// find single detail order in the database by id
	if err := findDetailOrderById(id, detailOrder); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Detail order not found"})
	}
	// find  order in the database by id
	if err := findOrderById(fmt.Sprint(detailOrder.IdOrder), order); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Order not found"})
	}
	// assign  order to detail order
	detailOrder.Order = *order
	// return detail order
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Detail order Found", "data": detailOrder})
}

// update a detail order in db
func UpdateDetailOrder(c *fiber.Ctx) error {
	db := database.DB.Db
	detailOrder := new(model.DetailOrder)
	order := new(model.Order)
	// get id params
	id := c.Params("id")
	// find single detail order in the database by id
	if err := findDetailOrderById(id, detailOrder); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Detail order not found"})
	}
	// store the body in the detail order and return error if encountered
	if err := c.BodyParser(detailOrder); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}
	// find  order in the database by id
	if err := findOrderById(fmt.Sprint(detailOrder.IdOrder), order); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Order not found"})
	}
	// assign  order to detail order
	detailOrder.Order = *order
	// update detail order
	if err := db.Save(detailOrder).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not update detail order", "data": err})
	}
	// return the updated detail order
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Detail order has updated", "data": detailOrder})
}

// delete a detail order in db
func DeleteDetailOrder(c *fiber.Ctx) error {
	db := database.DB.Db
	detailOrder := new(model.DetailOrder)
	// get id params
	id := c.Params("id")
	// find single detail order in the database by id
	if err := findDetailOrderById(id, detailOrder); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Detail order not found"})
	}
	// delete detail order
	if err := db.Delete(detailOrder).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not delete detail order", "data": err})
	}
	// return the deleted detail order
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Detail order has deleted", "data": detailOrder})
}

// update multiple detail orders
func UpdateMultipleDetailOrders(c *fiber.Ctx) error {
	db := database.DB.Db
	detailOrders := new([]model.DetailOrder)
	// get id params
	id := c.Params("id")
	// store the body in the detail orders and return error if encountered
	if err := c.BodyParser(detailOrders); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}

	for _, detailOrder := range *detailOrders {
		order := new(model.Order)
		// find  order in the database by id
		if err := findOrderById(fmt.Sprint(id), order); err != nil {
			return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Order not found"})
		}
		// assign  order to detail order
		detailOrder.Order = *order
		// update detail order
		if err := db.Save(&detailOrder).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not update detail order", "data": err})
		}
	}
	// return the updated detail orders
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Detail orders has updated", "data": detailOrders})
}
