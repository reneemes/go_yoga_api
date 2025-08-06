package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"go_yoga_api/internal/database"
	// ^ Import the database package to access the service
)

type Pose struct {
	ID   						int    `json:"id"`
	Name 						string `json:"name"`
	SanskritName 		string `json:"sanskrit_name"`
	TranslationName string `json:"translation_name"`
	Description 		string `json:"description"`
	PoseBenefits 		string `json:"pose_benefits"`
	ImageURL 				string `json:"image_url"`
}

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

	var poses []Pose
	for rows.Next() {
		var p Pose
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