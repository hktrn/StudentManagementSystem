package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/hktrn/StudentManagementSystem/database"
	"github.com/hktrn/StudentManagementSystem/routes"
)

func SetupRoutes(app *fiber.App) {
	//Students
	app.Post("/student/add", routes.AddStudent)
	app.Get("/student/see/all", routes.SeeStudents)
	app.Get("/student/see/:student_id", routes.SeeStudent)
	app.Put("/student/update/all/:student_id", routes.UpdateStudent)
	app.Put("/student/update/name/:student_id", routes.UpdateName)
	app.Put("/student/update/studies/:student_id", routes.UpdateStudies)
	app.Put("/student/update/coop/:student_id", routes.UpdateCoop)
	app.Delete("/student/delete/:student_id", routes.DeleteStudent)

	//Placements
	app.Post("/placement/add", routes.AddPlacement)
	app.Get("/placement/see/all", routes.SeePlacements)
	app.Get("/placement/see/:student_id", routes.SeePlacement)
	app.Put("/placement/update/all/:student_id", routes.UpdatePlacement)
	app.Delete("/placement/delete/:student_id", routes.DeletePlacement)

	//Records
	app.Post("/record/add", routes.AddRecord)
	app.Get("/record/see/all", routes.SeeRecords)
	app.Get("/record/see/:student_id", routes.SeeRecord)
	app.Delete("/record/delete/:student_id", routes.DeleteRecord)

}

func Connect() {
	database.ConnectToDatabase()
	app := fiber.New()
	SetupRoutes(app)
	log.Fatal(app.Listen(":5000"))
}
