package lead

import (
	"github.com/skozlovtsev/go-beginner-crm-project/database"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Lead struct {  //Лидер?
	gorm.Model
	Name    string `json:"name"`
	Company string `json:"company"`
	Email   string `json:"email"`
	Phone   int    `json:"phone"`
}

func GetLeads(c *fiber.Ctx){
	db := database.DBConn
	var leads []Lead
	db.Find(&leads)  //получаем все записи в leads
	c.JSON(leads)  //переводим leads в формат json и посылаем в качестве ответа
}

func GetLead(c *fiber.Ctx){
	id := c.Params("id")
	db := database.DBConn
	var lead Lead
	db.Find(&lead, id)  //получаем все записи с указаным id в lead
	c.JSON(lead)  //переводим lead в формат json и посылаем в качестве ответа
}

func NewLead(c *fiber.Ctx){
	db := database.DBConn
	lead := new(Lead)
	if err := c.BodyParser(lead); err != nil {
		c.Status(503).Send(err)  //посылаем ошибку с статусом 503
		return
	}
	db.Create(&lead)  //добавляем в бд новую запись
	c.JSON(lead)  //переводим lead в формат json и посылаем в качестве ответа
}

func DeleteLead(c *fiber.Ctx){
	id := c.Params("id")
	db := database.DBConn

	var lead Lead
	db.First(&lead, id)  //получаем первую запись с указаным id в lead
	if lead.Name == "" {  //В случае отсутствия 
		c.Status(500).Send("No lead found with ID")  //посылаем сообщение "No lead found with ID" с статусом 500
		return
	}
	db.Delete(&lead)  //удаляем запись lead из базы данных
	c.Send("Lead successfully Deleted")
}