package controller

import (
	"strconv"
	"time"

	"github.com/ShankaranarayananBR/FiberApp/config"
	"github.com/ShankaranarayananBR/FiberApp/model"
	"github.com/gofiber/fiber/v2"
)

func CreateCashier(c *fiber.Ctx) error {
	var data map[string]string
	err := c.BodyParser(&data)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Invalid data",
		})
	}
	if data["name"] == "" {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Cashier name is required",
		})
	}
	if data["passcode"] == "" {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Cashier passcode is required",
		})
	}
	cashier := model.Cashier{
		Name:      data["name"],
		Passcode:  data["passcode"],
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
	config.DB.Create(&cashier)
	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Cashier creation was successful",
	})
}

func UpdateCashier(c *fiber.Ctx) error {
	cashierId := c.Params("cashierId")
	var cashier model.Cashier
	config.DB.Find(*&cashierId, "id=?", cashierId)
	if cashier.Name == "" {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "Cashier not found",
		})
	}
	var UpdateCashier model.Cashier
	err := c.BodyParser(&UpdateCashier)
	if err != nil {
		return err
	}

	if UpdateCashier.Name == "" {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "Cashier name is required",
		})
	}
	cashier.Name = UpdateCashier.Name
	config.DB.Save(&cashier)
	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "success",
		"data":    cashier,
	})
}

func EditCashier(c *fiber.Ctx) error {
	return nil
}

func DeleteCashier(c *fiber.Ctx) error {
	cashierId := c.Params("cashierId")
	var cashier model.Cashier
	config.DB.Where("id=?", cashierId).First(&cashier)
	if cashier.Id == 0 {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "cashier not found",
		})
	}

	config.DB.Where("id=?", cashierId).Delete(&cashier)
	return c.Status(404).JSON(fiber.Map{
		"success": true,
		"message": "Cashier deleted successfully",
	})

}

func GetCashierList(c *fiber.Ctx) error {
	var cashiers []model.Cashier
	limit, _ := strconv.Atoi(c.Query("limit"))
	skip, _ := strconv.Atoi(c.Query("skip"))
	var count int64
	config.DB.Select("*").Limit(limit).Offset(skip).Find(&cashiers).Count(&count)
	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Successfully listed all cashiers",
	})
}

func GetCashierDetails(c *fiber.Ctx) error {
	cashierId := c.Params("cashierId")
	var cashier model.Cashier
	config.DB.Select("id,name").Where("id=?", cashierId).First(&cashier)

	cashierData := make(map[string]interface{})
	cashierData["cashierId"] = cashier.Id
	cashierData["name"] = cashier.Name
	cashierData["createdAt"] = cashier.CreatedAt
	cashierData["updatedAt"] = cashier.UpdatedAt

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"Message": "Successfully getting details of cashiers by ID",
	})
}
