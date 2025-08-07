package handler

import (
	"fmt"
	"go_yoga_api/internal/database"
	"net/http"
	"go_yoga_api/internal/types"

	"github.com/gin-gonic/gin"
)

func GetAllRoutinesHandler(c *gin.Context) {
	db := database.New()

	// Get all routines
	routineRows, err := db.DB().Query("SELECT id, name, description, difficulty FROM routines")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch routines",
		})
		return
	}
	defer routineRows.Close()

	var routines []types.Routine
	for routineRows.Next() {
		var r types.Routine
		if err := routineRows.Scan(&r.ID, &r.Name, &r.Description, &r.Difficulty); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to scan routines",
			})
			return
		}
		// Get routine poses for this routine
		poseQuery := `
			SELECT p.id, p.name, p.sanskrit_name, p.translation_name, p.description, p.image_url
			FROM routine_poses rp
			JOIN poses p ON p.id = rp.pose_id
			WHERE rp.routine_id = $1
		`

		poseRows, err := db.DB().Query(poseQuery, r.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch poses for routine"})
			return
		}

		var poses []types.Pose
		for poseRows.Next() {
			var p types.Pose
			if err := poseRows.Scan(&p.ID, &p.Name, &p.SanskritName, &p.TranslationName, &p.Description, &p.ImageURL); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan pose"})
				return
			}
			poses = append(poses, p)
		}
		poseRows.Close()

		r.RoutinePoses = poses
		routines = append(routines, r)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": routines,
	})
}

func GetOneRoutineHandler(c *gin.Context) {
	query := `
		SELECT r.id, r.name, r.description, r.difficulty,
			p.id, p.name, p.sanskrit_name, p.translation_name,
			p.description, p.image_url
		FROM routine_poses rp
		JOIN routines r ON r.id = rp.routine_id
		JOIN poses p ON p.id = rp.pose_id
		WHERE r.id = $1
	`
	id := c.Param("id")
	db := database.New()
	row, err := db.DB().Query(query, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch poses",
		})
		return
	}
	defer row.Close()

	var r types.Routine
	var poses []types.Pose

	for row.Next() {
		var pose types.Pose
		if err := row.Scan(
			&r.ID, &r.Name, &r.Description, &r.Difficulty,
			&pose.ID, &pose.Name, &pose.SanskritName, &pose.TranslationName,
			&pose.Description, &pose.ImageURL,
		); err != nil {
			fmt.Println("Error fetching pose with ID:", id, err)
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Routine not found",
			})
			return
		}
		poses = append(poses, pose)
	}

	r.RoutinePoses = poses
	c.JSON(http.StatusOK, gin.H{
		"data": r,
	})
}