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
	result := db.Preload("RoutinePoses").Find(&routines)
	// .Preload("RoutinePoses") loads the associated poses via the many2many relationship
	// .Find(&routines) retrieves all routines

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch poses",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": routines,
		})
	}
}

func GetOneRoutineHandler(c *gin.Context) {
	dbService := database.New()
	db := dbService.DB()

	var routine types.Routine
	result := db.Preload("RoutinePoses").First(&routine, c.Param("id"))
	// .First(&routine, c.Param("id")) retrieves the routine by ID

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch poses",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": routine,
		})
	}
}