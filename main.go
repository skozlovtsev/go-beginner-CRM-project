package main

import (
	"fmt"
	"github.com/skozlovtsev/go-beginner-crm-project/database"
	"github.com/skozlovtsev/go-beginner-crm-project/lead"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func setupRoutes(app *fiber.App){
	//добавление обработчиков в app
	app.Get("/api/v1/lead", lead.GetLeads)
	app.Get("/api/v1/lead/:id", lead.GetLead)
	app.Post("/api/v1/lead", lead.NewLead)
	app.Delete("/api/v1/lead/:id", lead.DeleteLead)
}

func initDatabase(){
	var err error
	database.DBConn, err = gorm.Open("postgres", "host=localhost user=postgres password=1234 dbname=postgres port=5432 sslmode=disable TimeZone=Europe/Moscow")  //подключение базы данных
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection opened to database")
	database.DBConn.AutoMigrate(&lead.Lead{})
	fmt.Println("Database Migrated")
}

func main() {
	app := fiber.New()  //создание нового объекта типа fiber.App
	initDatabase()  //инициация базы данных
	setupRoutes(app)  //добавление обработчиков
	app.Listen(3000)  //запуск сервера на локальном хосте
	defer database.DBConn.Close()  //после того как выполнение функции полностью завершено, закрываем подключение к базе данных 
}