package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"oooGlebusApi"
)

func (h *Handler) createReview(c *gin.Context) {

	var reviews oooGlebusApi.Review
	if err := c.BindJSON(&reviews); err != nil {
		log.Println("Error parsing input:", err)
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	query := `INSERT INTO review(rating, description, client_id, music_id) VALUES ($1, $2, $3, $4)`

	_, err := h.db.Exec(query, reviews.Rating, reviews.Description, reviews.Client_id, reviews.Music_id)
	if err != nil {
		log.Println("Error creating review:", err)
		c.JSON(500, gin.H{"error": "Failed to create review"})
		return
	}

	c.JSON(200, gin.H{"message": "Review created successfully"})
}

func (r *Handler) getAllReview(c *gin.Context) {
	log.Println("getAllAlbum called")

	var reviews []oooGlebusApi.Review

	query := "SELECT id, rating, description, client_id, music_id FROM review"
	log.Println("Executing query:", query)

	if r.db == nil {
		log.Fatalln("Database connection is nil")
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	err := r.db.Select(&reviews, query)
	if err != nil {
		log.Println("Error executing query:", err)
		c.JSON(500, gin.H{"error": "Failed to fetch reviews"})
		return
	}

	log.Printf("Fetched %d reviews\n", len(reviews))

	if len(reviews) == 0 {
		c.JSON(200, gin.H{"message": "No reviews found"})
		return
	}

	c.JSON(200, reviews)
}

func (h *Handler) getReviewById(c *gin.Context) {
	id := c.Param("id")

	var reviews oooGlebusApi.Review
	query := "SELECT id, rating, description, client_id, music_id FROM review WHERE id = $1"

	err := h.db.Get(&reviews, query, id)
	if err != nil {
		log.Println("Error fetching album by ID:", err)
		c.JSON(404, gin.H{"error": "Album not found"})
		return
	}
	c.JSON(200, reviews)
}

func (h *Handler) updateReview(c *gin.Context) {
	id := c.Param("id")

	var reviews oooGlebusApi.Review
	if err := c.BindJSON(&reviews); err != nil {
		log.Println("Error parsing input:", err)
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	query := `UPDATE review SET rating = $1, description = $2, client_id = $3, music_id = $4 WHERE id = $5`

	_, err := h.db.Exec(query, reviews.Rating, reviews.Description, reviews.Client_id, reviews.Music_id, id)
	if err != nil {
		log.Println("Error updating review:", err)
		c.JSON(500, gin.H{"error": "Failed to update review"})
		return
	}

	c.JSON(200, gin.H{"message": "Review updated successfully"})
}

func (h *Handler) deleteReview(c *gin.Context) {
	id := c.Param("id")

	query := "DELETE FROM album WHERE id = $1"

	_, err := h.db.Exec(query, id)
	if err != nil {
		log.Println("Error deleting album:", err)
		c.JSON(500, gin.H{"error": "Failed to delete album"})
		return
	}

	c.JSON(200, gin.H{"message": "album deleted successfully"})
}
