package handler

import (
	"go_yoga_api/internal/database"
	"go_yoga_api/internal/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllPosesHandler(c *gin.Context) {
	dbService := database.New()
	db := dbService.DB()

	var poses []types.Pose
	result := db.Find(&poses)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch poses",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": poses,
		})
	}
}

func GetOnePoseHandler(c *gin.Context) {
	dbService := database.New()
	db := dbService.DB()

	var pose []types.Pose
	result := db.Find(&pose, c.Param("id"))

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch poses",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": pose,
		})
	}
}
