package handler

import (
	"github.com/AxelanO7/bn-shop-backend-web-go/database"
	"github.com/AxelanO7/bn-shop-backend-web-go/model"
	"github.com/gofiber/fiber/v2"
)

// find login by id
func findLoginById(id string, login *model.Login) error {
	db := database.DB.Db
	// find single login in the database by id
	db.Find(&login, "id = ?", id)
	// if no login found, return an error
	if login.ID == 0 {
		return fiber.ErrNotFound
	}
	return nil
}

// create a login
func CreateLogin(c *fiber.Ctx) error {
	db := database.DB.Db
	login := new(model.Login)
	// store the body in the login and return error if encountered
	if err := c.BodyParser(login); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}
	// create login
	if err := db.Create(login).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create login", "data": err})
	}
	// return the created login
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Login has created", "data": login})
}

// get all Logins from db
func GetAllLogins(c *fiber.Ctx) error {
	db := database.DB.Db
	logins := []model.Login{}
	if err := db.Find(&logins).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not get logins", "data": err})
	}
	// if no login found, return an error
	if len(logins) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Logins not found", "data": nil})
	}
	// return logins
	return c.Status(200).JSON(fiber.Map{"status": "sucess", "message": "Logins Found", "data": logins})
}

// GetSingleLogin from db
func GetSingleLogin(c *fiber.Ctx) error {
	login := new(model.Login)
	// get id params
	id := c.Params("id")
	// find single login in the database by id
	if err := findLoginById(id, login); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Login not found"})
	}
	// return login
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Login Found", "data": login})
}

// update a login in db
func UpdateLogin(c *fiber.Ctx) error {
	db := database.DB.Db
	login := new(model.Login)
	// get id params
	id := c.Params("id")
	// find single login in the database by id
	if err := findLoginById(id, login); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Login not found"})
	}
	// store the body in the login and return error if encountered
	if err := c.BodyParser(login); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}
	// update login
	if err := db.Save(login).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not update login", "data": err})
	}
	// return the updated login
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "logins Found", "data": login})
}

// delete a login in db
func DeleteLogin(c *fiber.Ctx) error {
	db := database.DB.Db
	login := new(model.Login)
	// get id params
	id := c.Params("id")
	// find single login in the database by id
	if err := findLoginById(id, login); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Login not found"})
	}
	// delete login
	if err := db.Delete(login, "id = ?", id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not delete login", "data": err})
	}
	// return success message
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Login deleted"})
}

// find login by username and password
func findLoginByUsernameAndPassword(username string, password string, login *model.Login) error {
	db := database.DB.Db
	// find single login in the database by id
	db.Find(&login, "username = ? AND password = ?", username, password)
	// if no login found, return an error
	if login.ID == 0 {
		return fiber.ErrNotFound
	}
	return nil
}

// login a login in db
func Login(c *fiber.Ctx) error {
	login := new(model.Login)
	// store the body in the login and return error if encountered
	if err := c.BodyParser(login); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}
	// find single login in the database by ids
	if err := findLoginByUsernameAndPassword(login.Username, login.Password, login); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Username or Password is wrong"})
	}
	// return the login
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Login Success"})
}
