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
	if err := db.Find(&poses).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch poses",
		})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"data": poses,
	})
}

func GetOnePoseHandler(c *gin.Context) {
	dbService := database.New()
	db := dbService.DB()

	var pose types.Pose
	if err := db.First(&pose, c.Param("id")).Error; err != nil {
		// First(&pose, c.Param("id")) finds the pose by ID
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Pose not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": pose,
	})
}