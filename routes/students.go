package routes

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/hktrn/StudentManagementSystem/database"
	"github.com/hktrn/StudentManagementSystem/model"
)

type Student struct {
	StudentID     uint32 `json:"student_id"`
	Name          string `json:"name"`
	Program       string `json:"program"`
	Major         string `json:"major"`
	AdmissionYear uint32 `json:"admission_year"`
	Coop          string `json:"coop"`
}

func createStudent(student model.Student) Student {
	return Student{StudentID: student.StudentID,
		Name:          student.Name,
		Program:       student.Program,
		Major:         student.Major,
		AdmissionYear: student.AdmissionYear,
		Coop:          student.Coop}
}

//Adding students
func AddStudent(c *fiber.Ctx) error {
	var student model.Student
	if err := c.BodyParser(&student); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	database.Database.DB.Create(&student)
	responseStudent := createStudent(student)

	return c.Status(200).JSON(responseStudent)
}

//See all students
func SeeStudents(c *fiber.Ctx) error {
	students := []model.Student{}

	database.Database.DB.Find(&students)
	responseStudents := []Student{}
	for _, student := range students {
		responseStudent := createStudent(student)
		responseStudents = append(responseStudents, responseStudent)
	}
	return c.Status(200).JSON(responseStudents)
}

//Finding specific student
func findStudent(student_id int, student *model.Student) error {
	database.Database.DB.Find(&student, "student_id = ?", student_id)
	if student.StudentID == 0 {
		return errors.New("student does not exist")
	}
	return nil
}

// seeing specific student
func SeeStudent(c *fiber.Ctx) error {
	student_id, err := c.ParamsInt("student_id")
	var student model.Student
	if err != nil {
		return c.Status(400).JSON("id dosent exist")
	}
	if err := findStudent(student_id, &student); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	responseStudent := createStudent(student)

	return c.Status(200).JSON(responseStudent)
}

//Updating

//Updating all
func UpdateStudent(c *fiber.Ctx) error {
	student_id, err := c.ParamsInt("student_id")
	var student model.Student
	if err != nil {
		return c.Status(400).JSON("id dosent exist")
	}
	if err := findStudent(student_id, &student); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	type updateStudent struct {
		Name    string `json:"name"`
		Program string `json:"program"`
		Major   string `json:"major"`
	}

	var studies updateStudent
	if err := c.BodyParser(&studies); err != nil {
		return c.Status(500).JSON(err.Error())
	}
	student.Name = studies.Name
	student.Program = studies.Program
	student.Major = studies.Major

	database.Database.DB.Save(&student)

	responseStudent := createStudent(student)
	return c.Status(200).JSON(responseStudent)
}

//Updating Name
func UpdateName(c *fiber.Ctx) error {
	student_id, err := c.ParamsInt("student_id")
	var student model.Student
	if err != nil {
		return c.Status(400).JSON("id dosent exist")
	}
	if err := findStudent(student_id, &student); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	type updateName struct {
		Name string `json:"name"`
	}

	var name updateName
	if err := c.BodyParser(&name); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	student.Name = name.Name

	database.Database.DB.Save(&student)

	responseStudent := createStudent(student)
	return c.Status(200).JSON(responseStudent)
}

//Updating Program and Major = Studies
func UpdateStudies(c *fiber.Ctx) error {
	student_id, err := c.ParamsInt("student_id")
	var student model.Student
	if err != nil {
		return c.Status(400).JSON("id dosent exist")
	}
	if err := findStudent(student_id, &student); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	type updateStudies struct {
		Program string `json:"program"`
		Major   string `json:"major"`
	}

	var studies updateStudies
	if err := c.BodyParser(&studies); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	student.Program = studies.Program
	student.Major = studies.Major

	database.Database.DB.Save(&student)

	responseStudent := createStudent(student)
	return c.Status(200).JSON(responseStudent)
}

//Updating Coop
func UpdateCoop(c *fiber.Ctx) error {
	student_id, err := c.ParamsInt("student_id")
	var student model.Student
	if err != nil {
		return c.Status(400).JSON("id dosent exist")
	}
	if err := findStudent(student_id, &student); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	type updateCoop struct {
		Coop string `json:"coop"`
	}

	var coop updateCoop
	if err := c.BodyParser(&coop); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	student.Coop = coop.Coop

	database.Database.DB.Save(&student)

	responseStudent := createStudent(student)
	return c.Status(200).JSON(responseStudent)
}

//Deleting
func DeleteStudent(c *fiber.Ctx) error {
	student_id, err := c.ParamsInt("student_id")
	var student model.Student
	if err != nil {
		return c.Status(400).JSON("id dosent exist")
	}

	if err := findStudent(student_id, &student); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err = database.Database.DB.Delete(&student).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}
	return c.Status(200).JSON("Student deleted")
}
