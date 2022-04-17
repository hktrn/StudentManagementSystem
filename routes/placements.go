package routes

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/hktrn/StudentManagementSystem/database"
	"github.com/hktrn/StudentManagementSystem/model"
)

type Placement struct {
	StudentID uint32 `json:"student_id"`
	Company   string `json:"company"`
	Title     string `json:"title"`
	Location  string `json:"location"`
	Term      string `json:"term"`
}

func createPlacement(placement model.Placement) Placement {
	return Placement{StudentID: placement.StudentID,
		Company:  placement.Company,
		Title:    placement.Title,
		Location: placement.Location,
		Term:     placement.Term}
}

func AddPlacement(c *fiber.Ctx) error {
	var placement model.Placement

	if err := c.BodyParser(&placement); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.DB.Create(&placement)
	responsePlacement := createPlacement(placement)

	return c.Status(200).JSON(responsePlacement)

}

func SeePlacements(c *fiber.Ctx) error {
	placements := []model.Placement{}

	database.Database.DB.Find(&placements)
	responsePlacements := []Placement{}
	for _, placement := range placements {
		responsePlacement := createPlacement(placement)
		responsePlacements = append(responsePlacements, responsePlacement)
	}
	return c.Status(200).JSON(responsePlacements)
}

//Finding specific placements
func findPlacement(student_id int, placement *model.Placement) error {
	database.Database.DB.Find(&placement, "student_id = ?", student_id)
	if placement.StudentID == 0 {
		return errors.New("student does not exist")
	}
	return nil
}

// seeing specific student
func SeePlacement(c *fiber.Ctx) error {
	student_id, err := c.ParamsInt("student_id")
	var placement model.Placement
	if err != nil {
		return c.Status(400).JSON("id dosent exist")
	}
	if err := findPlacement(student_id, &placement); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	responsePlacement := createPlacement(placement)

	return c.Status(200).JSON(responsePlacement)
}

//Updating all
func UpdatePlacement(c *fiber.Ctx) error {
	student_id, err := c.ParamsInt("student_id")
	var placement model.Placement
	if err != nil {
		return c.Status(400).JSON("id dosent exist")
	}
	if err := findPlacement(student_id, &placement); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	type UpdatePlacement struct {
		Company  string `json:"company"`
		Title    string `json:"title"`
		Location string `json:"location"`
		Term     string `json:"term"`
	}

	var update UpdatePlacement
	if err := c.BodyParser(&update); err != nil {
		return c.Status(500).JSON(err.Error())
	}
	placement.Company = update.Company
	placement.Title = update.Title
	placement.Location = update.Location
	placement.Term = update.Term

	database.Database.DB.Save(&placement)

	responsePlacement := createPlacement(placement)
	return c.Status(200).JSON(responsePlacement)
}

//Deleting
func DeletePlacement(c *fiber.Ctx) error {
	student_id, err := c.ParamsInt("student_id")
	var placement model.Placement
	if err != nil {
		return c.Status(400).JSON("id dosent exist")
	}

	if err := findPlacement(student_id, &placement); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err = database.Database.DB.Delete(&placement).Error; err != nil {
		return c.Status(400).JSON(err.Error())
	}
	return c.Status(200).JSON("Student deleted")
}
