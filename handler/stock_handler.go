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

// create multiple stocks
func CreateMultipleStocks(c *fiber.Ctx) error {
	db := database.DB.Db
	stocks := new([]model.Stock)
	// store the body in the stock and return error if encountered
	if err := c.BodyParser(stocks); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}
	// create stock
	if err := db.Create(stocks).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create stocks", "data": err})
	}
	// return the created stock
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Stocks has created", "data": stocks})
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

// get all detail order where foreign key order status = 1
// func GetAllDetailOrdersByStatus(c *fiber.Ctx) error {
// 	db := database.DB.Db
// 	detailOrders := []model.DetailOrder{}
// 	orders := []model.Order{}
// 	// find all orders in the database by status
// 	if err := db.Find(&orders, "status = ?").Error; err != nil {
// 		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not get orders", "data": err})
// 	}

// 	// find all detail orders where id order = id order in orders
// 	for _, order := range orders {
// 		if err := db.Find(&detailOrders, "id_order = ?", order.ID).Error; err != nil {
// 			return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not get detail orders", "data": err})
// 		}
// 	}

// 	// if no detail order found, return an error
// 	if len(detailOrders) == 0 {
// 		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Detail orders not found", "data": nil})
// 	}
// 	responseDetailOrders := []model.DetailOrder{}
// 	for _, detailOrder := range detailOrders {
// 		order := new(model.Order)
// 		// find  order in the database by id
// 		if err := findOrderById(fmt.Sprint(detailOrder.IdOrder), order); err != nil {
// 			return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Order not found"})
// 		}
// 		// assign  order to detail order
// 		detailOrder.Order = *order
// 		responseDetailOrders = append(responseDetailOrders, detailOrder)
// 	}
// 	// return detail orders
// 	return c.Status(200).JSON(fiber.Map{"status": "sucess", "message": "Detail orders Found", "data": detailOrders})
// }

func GetAllFinished(c *fiber.Ctx) error {
	db := database.DB.Db
	stocks := []model.Stock{}
	// find all stocks in the database
	if err := db.Find(&stocks, "type_product = ?", "Barang Jadi").Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not get stocks", "data": err})
	}
	// if no stock found, return an error
	if len(stocks) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Stocks not found", "data": nil})
	}
	// return stocks
	return c.Status(200).JSON(fiber.Map{"status": "sucess", "message": "Stocks Found", "data": stocks})
}

func GetAllRaw(c *fiber.Ctx) error {
	db := database.DB.Db
	stocks := []model.Stock{}
	// find all stocks in the database
	if err := db.Find(&stocks, "type_product = ?", "Bahan Baku").Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not get stocks", "data": err})
	}
	// if no stock found, return an error
	if len(stocks) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Stocks not found", "data": nil})
	}
	// return stocks
	return c.Status(200).JSON(fiber.Map{"status": "sucess", "message": "Stocks Found", "data": stocks})
}
