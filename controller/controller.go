package controller

import (
	"github.com/gin-gonic/gin"
	"web-service-gin/database"
	"web-service-gin/helper"
)

func AddRecordController(context *gin.Context) {
	helper.AddRecord(context, database.Database)
}

func GetRecordController(context *gin.Context) {
	helper.GetRecord(context, database.Database)
}

func GetRecordByIdController(context *gin.Context) {
	helper.GetRecordById(context, database.Database)
}

func UpdateRecordController(context *gin.Context) {
	helper.UpdateRecord(context, database.Database)
}

func DeleteRecordController(context *gin.Context) {
	helper.DeleteRecord(context, database.Database)
}
