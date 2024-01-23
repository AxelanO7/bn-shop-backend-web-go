package handler

import (
	"github.com/AxelanO7/bn-shop-backend-web-go/database"
	"github.com/AxelanO7/bn-shop-backend-web-go/model"
	"github.com/gofiber/fiber/v2"
)

// find user by id
func findUserById(id string, user *model.User) error {
	db := database.DB.Db
	// find single user in the database by id
	db.Find(&user, "id = ?", id)
	// if no user found, return an error
	if user.ID == 0 {
		return fiber.ErrNotFound
	}
	return nil
}

// create a user
func CreateUser(c *fiber.Ctx) error {
	db := database.DB.Db
	user := new(model.User)
	// store the body in the user and return error if encountered
	if err := c.BodyParser(user); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}
	// create user
	if err := db.Create(user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create user", "data": err})
	}
	// return the created user
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "User has created", "data": user})
}

// get all Users from db
func GetAllUsers(c *fiber.Ctx) error {
	db := database.DB.Db
	users := []model.User{}
	if err := db.Find(&users).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not get users", "data": err})
	}
	// if no user found, return an error
	if len(users) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Users not found", "data": nil})
	}
	// return users
	return c.Status(200).JSON(fiber.Map{"status": "sucess", "message": "Users Found", "data": users})
}

// GetSingleUser from db
func GetSingleUser(c *fiber.Ctx) error {
	user := new(model.User)
	// get id params
	id := c.Params("id")
	// find single user in the database by id
	if err := findUserById(id, user); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "User not found"})
	}
	// return user
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "User Found", "data": user})
}

// update a user in db
func UpdateUser(c *fiber.Ctx) error {
	db := database.DB.Db
	user := new(model.User)
	// get id params
	id := c.Params("id")
	// find single user in the database by id
	if err := findUserById(id, user); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "User not found"})
	}
	// store the body in the user and return error if encountered
	if err := c.BodyParser(user); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}
	// update user
	if err := db.Save(user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not update user", "data": err})
	}
	// return the updated user
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "users Found", "data": user})
}

// delete a user in db
func DeleteUser(c *fiber.Ctx) error {
	db := database.DB.Db
	user := new(model.User)
	// get id params
	id := c.Params("id")
	// find single user in the database by id
	if err := findUserById(id, user); err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "User not found"})
	}
	// delete user
	if err := db.Delete(user, "id = ?", id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not delete user", "data": err})
	}
	// return success message
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "User deleted"})
}

func GetLogedUser(c *fiber.Ctx) error {
	db := database.DB.Db
	users := []model.User{}
	db.Find(&users)
	if len(users) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Users not found", "data": nil})
	}
	if len(users) > 1 {
		for i, user := range users {
			if i > 0 {
				user.Status = 0
			}
		}
		db.Save(users)
		return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Users Found", "data": users[0]})
	} else {
		for _, user := range users {
			if user.Status == 1 {
				return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Users Found", "data": user})
			}
		}
	}
	return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Users not found", "data": nil})
}
