package routes

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/hktrn/StudentManagementSystem/database"
	"github.com/hktrn/StudentManagementSystem/model"
)

//Showing everything with 1 ID
type Record struct {
	StudentID uint32    `json:"student_id"`
	Student   Student   `json:"student"`
	Placement Placement `json:"placement"`
}

func createRecord(record model.Record, student Student, placement Placement) Record {
	return Record{StudentID: record.StudentID,
		Student:   student,
		Placement: placement}
}

func AddRecord(c *fiber.Ctx) error {
	var record model.Record

	if err := c.BodyParser(&record); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	var student model.Student
	if err := findStudent(int(record.StudentID), &student); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var placement model.Placement
	if err := findPlacement(int(record.StudentID), &placement); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.DB.Create(&record)

	responseStudent := createStudent(student)
	responsePlacement := createPlacement(placement)
	responseRecord := createRecord(record, responseStudent, responsePlacement)

	return c.Status(200).JSON(responseRecord)

}

func SeeRecords(c *fiber.Ctx) error {
	records := []model.Record{}
	database.Database.DB.Find(&records)
	responseRecords := []Record{}

	for _, record := range records {
		var student model.Student
		var placement model.Placement
		database.Database.DB.Find(&student, "student_id = ?", record.StudentID)
		database.Database.DB.Find(&placement, "student_id = ?", record.StudentID)
		responseRecord := createRecord(record, createStudent(student), createPlacement(placement))
		responseRecords = append(responseRecords, responseRecord)
	}

	return c.Status(200).JSON(responseRecords)
}

func findRecord(student_id int, record *model.Record) error {
	database.Database.DB.Find(&record, "student_id = ?", student_id)
	if record.StudentID == 0 {
		return errors.New("record does not exist")
	}
	return nil
}

func SeeRecord(c *fiber.Ctx) error {
	student_id, err := c.ParamsInt("student_id")
	var record model.Record

	if err != nil {
		return c.Status(400).JSON("Ensure :student_id is an integer")
	}

	if err := findRecord(student_id, &record); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var student model.Student
	var placement model.Placement
	database.Database.DB.First(&student, record.StudentID)
	database.Database.DB.First(&placement, record.StudentID)
	responseStudent := createStudent(student)
	responsePlacement := createPlacement(placement)
	responseRecord := createRecord(record, responseStudent, responsePlacement)

	return c.Status(200).JSON(responseRecord)
}

//need fixing
func DeleteRecord(c *fiber.Ctx) error {
	student_id, err := c.ParamsInt("student_int")
	var record model.Record

	if err != nil {
		return c.Status(400).JSON("id dont exist")
	}
	if err := findRecord(student_id, &record); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	if err = database.Database.DB.Delete(&record).Error; err != nil {
		return c.Status(400).JSON(err.Error())
	}
	return c.Status(200).JSON("Student deleted")
}