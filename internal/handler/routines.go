package handler

import (
	"net/http"
	"go_yoga_api/internal/database"
	"go_yoga_api/internal/types"

	"github.com/gin-gonic/gin"
)

func GetAllRoutinesHandler(c *gin.Context) {
	dbService := database.New()
	// ^ new database connection
	db := dbService.DB()
	// ^ gorm DB instance

	var routines []types.Routine
	if err := db.Preload("RoutinePoses").Find(&routines).Error; err != nil {
	// .Preload("RoutinePoses") loads the associated poses via the many2many relationship
	// .Find(&routines) retrieves all routines
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch routines",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": routines,
	})
}

func GetOneRoutineHandler(c *gin.Context) {
	dbService := database.New()
	db := dbService.DB()
	id := c.Param("id")

	var routine types.Routine
	if err := db.Preload("RoutinePoses").First(&routine, id).Error; err != nil {
	// .First(&routine, id) retrieves the routine by ID
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Routine not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": routine,
	})
}