package handler

import (
	// "fmt"
	"go_yoga_api/internal/database"
	"go_yoga_api/internal/types"
	// "log"
	"net/http"

	"github.com/gin-gonic/gin"
	// ^ Import the database package to access the service
)

func GetAllPosesHandler(c *gin.Context) {
	db := database.New() // Create a new database service instance
	rows, err := db.DB().Query("SELECT id, name, sanskrit_name, translation_name, description, pose_benefits, image_url FROM poses")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch poses",
		})
		return
	}
	// Ensure to close the rows after processing
	// This prevents resource leaks
	defer rows.Close()

	var poses []types.Pose
	for rows.Next() {
		var p types.Pose
		if err := rows.Scan(&p.ID, &p.Name, &p.SanskritName, &p.TranslationName, &p.Description, &p.PoseBenefits, &p.ImageURL); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to scan pose",
			})
			return
		}
		poses = append(poses, p)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": poses,
	})
}

func GetOnePoseHandler(c *gin.Context) {
	query := `
		SELECT id, name, sanskrit_name, translation_name, description, pose_benefits, image_url
		FROM poses
		WHERE id = $1
	`

	id := c.Param("id")
	// fmt.Println("Fetching pose with ID:", id)
	db := database.New()
	row := db.DB().QueryRow(query, id)

	var p types.Pose
	if err := row.Scan(&p.ID, &p.Name, &p.SanskritName, &p.TranslationName, &p.Description, &p.PoseBenefits, &p.ImageURL); err != nil {
		// fmt.Println("Error fetching pose with ID:", id, err)
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Pose not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": p,
		// "data": {
		// 	"id": p.ID,
		// 	"type": "pose",
		// 	"attributes": p, 
		// }
	})
}