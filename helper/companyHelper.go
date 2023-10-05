package helper

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"web-service-gin/model"
)

func GetRecord(c *gin.Context, db *gorm.DB) {
	var rec []model.Company

	user, _ := CurrentUser(c)
	db.Find(&rec)
	c.JSON(200, user.Companies)

}

func GetRecordById(c *gin.Context, db *gorm.DB) {
	var data model.Company
	todoID := c.Param("id")

	result := db.First(&data, todoID)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Record not found"})
		return
	}

	c.JSON(200, data)

}

func AddRecord(c *gin.Context, db *gorm.DB) {
	var data model.Company
	userId, err := CurrentUser(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err = c.ShouldBindJSON(&data); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON data"})
		return
	}
	data.UserID = userId.ID
	db.Create(&data)

	c.JSON(200, data)
}

func UpdateRecord(c *gin.Context, db *gorm.DB) {
	var data model.Company
	todoID := c.Param("id")

	result := db.First(&data, todoID)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Todo not found"})
		return
	}

	var updatedTodo model.Company
	if err := c.ShouldBindJSON(&updatedTodo); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON data"})
		return
	}

	data.Title = updatedTodo.Title
	data.Description = updatedTodo.Description
	db.Save(&data)

	c.JSON(200, data)
}

func DeleteRecord(c *gin.Context, db *gorm.DB) {
	var data model.Company
	compId := c.Param("id")

	result := db.First(&data, compId)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Todo not found"})
		return
	}

	db.Delete(&data)

	c.JSON(200, gin.H{"message": fmt.Sprintf("Todo with ID %s deleted", compId)})
}
