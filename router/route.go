package router

import (
	"github.com/AxelanO7/bn-shop-backend-web-go/handler"
	"github.com/gofiber/fiber/v2"
)

// SetupRoutes func
func SetupRoutes(app *fiber.App) {
	// grouping
	api := app.Group("/api")

	// login
	login := api.Group("/login")
	// routes
	login.Get("/", handler.GetAllLogins)
	login.Get("/:id", handler.GetSingleLogin)
	login.Post("/", handler.CreateLogin)
	login.Put("/:id", handler.UpdateLogin)
	login.Delete("/:id", handler.DeleteLogin)

	// supplier
	supplier := api.Group("/supplier")
	// routes
	supplier.Get("/", handler.GetAllSuppliers)
	supplier.Get("/:id", handler.GetSingleSupplier)
	supplier.Post("/", handler.CreateSupplier)
	supplier.Put("/:id", handler.UpdateSupplier)
	supplier.Delete("/:id", handler.DeleteSupplier)

	// user
	user := api.Group("/user")
	// routes
	user.Get("/", handler.GetAllUsers)
	user.Get("/:id", handler.GetSingleUser)
	user.Post("/", handler.CreateUser)
	user.Put("/:id", handler.UpdateUser)
	user.Delete("/:id", handler.DeleteUser)

	// input
	input := api.Group("/input")
	// routes
	input.Get("/", handler.GetAllInputs)
	input.Get("/:id", handler.GetSingleInput)
	input.Post("/", handler.CreateInput)
	input.Put("/:id", handler.UpdateInput)
	input.Delete("/:id", handler.DeleteInput)

	// detail input
	detailInput := api.Group("/detail-input")
	// routes
	detailInput.Get("/", handler.GetAllDetailInputs)
	detailInput.Get("/:id", handler.GetSingleDetailInput)
	detailInput.Post("/", handler.CreateDetailInput)
	detailInput.Put("/:id", handler.UpdateDetailInput)
	detailInput.Delete("/:id", handler.DeleteDetailInput)

	// output
	output := api.Group("/output")
	// routes
	output.Get("/", handler.GetAllOutputs)
	output.Get("/:id", handler.GetSingleOutput)
	output.Post("/", handler.CreateOutput)
	output.Put("/:id", handler.UpdateOutput)
	output.Delete("/:id", handler.DeleteOutput)

	// detail output
	detailOutput := api.Group("/detail-output")
	// routes
	detailOutput.Get("/", handler.GetAllDetailOutputs)
	detailOutput.Get("/:id", handler.GetSingleDetailOutput)
	detailOutput.Post("/", handler.CreateDetailOutput)
	detailOutput.Put("/:id", handler.UpdateDetailOutput)
	detailOutput.Delete("/:id", handler.DeleteDetailOutput)

	// order
	order := api.Group("/order")
	// routes
	order.Get("/", handler.GetAllOrders)
	order.Get("/:id", handler.GetSingleOrder)
	order.Post("/", handler.CreateOrder)
	order.Put("/:id", handler.UpdateOrder)
	order.Delete("/:id", handler.DeleteOrder)

	// detail order
	detailOrder := api.Group("/detail-order")
	// routes
	detailOrder.Get("/", handler.GetAllDetailOrders)
	detailOrder.Get("/:id", handler.GetSingleDetailOrder)
	detailOrder.Post("/", handler.CreateDetailOrder)
	detailOrder.Put("/:id", handler.UpdateDetailOrder)
	detailOrder.Delete("/:id", handler.DeleteDetailOrder)

	// stock
	stock := api.Group("/stock")
	// routes
	stock.Get("/", handler.GetAllStocks)
	stock.Get("/:id", handler.GetSingleStock)
	stock.Post("/", handler.CreateStock)
	stock.Put("/:id", handler.UpdateStock)
	stock.Delete("/:id", handler.DeleteStock)

	// stock opname
	stockOpname := api.Group("/stock-opname")
	// routes
	stockOpname.Get("/", handler.GetAllStockOpnames)
	stockOpname.Get("/:id", handler.GetSingleStockOpname)
	stockOpname.Post("/", handler.CreateStockOpname)
	stockOpname.Put("/:id", handler.UpdateStockOpname)
	stockOpname.Delete("/:id", handler.DeleteStockOpname)
}
