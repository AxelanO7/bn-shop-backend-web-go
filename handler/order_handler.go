package handler

import (
	"fmt"

	"github.com/AxelanO7/bn-shop-backend-web-go/database"
	"github.com/AxelanO7/bn-shop-backend-web-go/model"
	"github.com/gofiber/fiber/v2"
)

// find order by id
func findOrderById(id string, order *model.Order) error {
	db := database.DB.Db
	// find single order in the database by id
	db.Find(&order, "id = ?", id)
	// if no order found, return an error
	if order.ID == 0 {
		return fiber.ErrNotFound
	}
	return nil
}

// create a order
func CreateOrder(c *fiber.Ctx) error {
	db := database.DB.Db
	order := new(model.Order)
	supplier := new(model.Supplier)
	// store the body in the order and return error if encountereds
	if err := c.BodyParser(order); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}
	// find supplier in the database by id
	if err := findSupplierById(fmt.Sprint(order.IdSupplier), supplier); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Supplier not found"})
	}
	// assign supplier to order
	order.Supplier = *supplier
	// create order
	if err := db.Create(order).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create order", "data": err})
	}
	// return the created order
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Order has created", "data": order})
}

// get all orders from db
func GetAllOrders(c *fiber.Ctx) error {
	db := database.DB.Db
	orders := []model.Order{}
	// find all orders in the database
	if err := db.Find(&orders).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not get orders", "data": err})
	}
	// if no order found, return an error
	if len(orders) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Orders not found", "data": nil})
	}
	responseOrders := []model.Order{}
	for _, order := range orders {
		supplier := new(model.Supplier)
		// find supplier in the database by id
		if err := findSupplierById(fmt.Sprint(order.IdSupplier), supplier); err != nil {
			return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Supplier not found"})
		}
		// assign supplier to order
		order.Supplier = *supplier
		responseOrders = append(responseOrders, order)
	}
	// return orders
	return c.Status(200).JSON(fiber.Map{"status": "sucess", "message": "Orders Found", "data": responseOrders})
}

// get single order from db
func GetSingleOrder(c *fiber.Ctx) error {
	order := new(model.Order)
	supplier := new(model.Supplier)
	// get id params
	id := c.Params("id")
	// find single order in the database by id
	if err := findOrderById(id, order); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Order not found"})
	}
	// find supplier in the database by id
	if err := findSupplierById(fmt.Sprint(order.IdSupplier), supplier); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Supplier not found"})
	}
	// assign supplier to order
	order.Supplier = *supplier
	// return order
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Order Found", "data": order})
}

// update a order in db
func UpdateOrder(c *fiber.Ctx) error {
	db := database.DB.Db
	order := new(model.Order)
	supplier := new(model.Supplier)
	// get id params
	id := c.Params("id")
	// find single order in the database by id
	if err := findOrderById(id, order); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Order not found"})
	}
	// store the body in the order and return error if encountereds
	if err := c.BodyParser(order); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}
	// find supplier in the database by id
	if err := findSupplierById(fmt.Sprint(order.IdSupplier), supplier); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Supplier not found"})
	}
	// assign supplier to order
	order.Supplier = *supplier
	// update order
	if err := db.Save(order).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not update order", "data": err})
	}
	// return the updated order
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Order has updated", "data": order})
}

// delete a order in db
func DeleteOrder(c *fiber.Ctx) error {
	db := database.DB.Db
	order := new(model.Order)
	// get id params
	id := c.Params("id")
	// find single order in the database by id
	if err := findOrderById(id, order); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Order not found"})
	}
	// delete order
	if err := db.Delete(order).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not delete order", "data": err})
	}
	// return deleted order
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Order deleted"})
}

// get all order by start date and end date from db
func GetOrderByDate(c *fiber.Ctx) error {
	db := database.DB.Db
	detailOrders := []model.DetailOrder{}
	dateStart := c.Query("date-start")
	dateEnd := c.Query("date-end")

	// find all detail orders in the database by date
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
		// find  order in the database by date
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
