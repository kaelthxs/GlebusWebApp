package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"oooGlebusApi"
)

func (h *Handler) createAlbum(c *gin.Context) {

	var albums oooGlebusApi.Album

	if err := c.BindJSON(&albums); err != nil {
		log.Println("Error parsing input:", err)
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	query := `INSERT INTO album (name, countofplays, rating, countofmusic, status, image_uri) VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := h.db.Exec(query, albums.Name, albums.Countofplays, albums.Rating, albums.Countofmusic, albums.Status, albums.Image_uri)
	if err != nil {
		log.Println("Error creating album:", err)
		c.JSON(500, gin.H{"error": "Failed to create album"})
		return
	}

	c.JSON(200, gin.H{"message": "Album created successfully"})

	var author_albums oooGlebusApi.Author_Album

	if err := c.BindJSON(&albums); err != nil {
		log.Println("Error parsing input:", err)
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	query2 := `INSERT INTO author_album (album_id, author_id) VALUES ($1, $2)`

	_, err2 := h.db.Exec(query2, author_albums.Album_id, author_albums.Author_id)
	if err2 != nil {
		log.Println("Error creating album:", err2)
		c.JSON(500, gin.H{"error": "Failed to create album"})
		return
	}

	c.JSON(200, gin.H{"message": "Album created successfully"})

}

func (r *Handler) getAllAlbum(c *gin.Context) {
	log.Println("getAllAlbum called")

	var albums []oooGlebusApi.Album

	query := "SELECT album.id, client.username, album.name, album.countofplays, album.rating, album.countofmusic, album.status, album.image_uri FROM album JOIN author_album ON album.id = author_album.album_id JOIN client ON author_album.author_id = client.id WHERE client.role = 'AUTHOR'"

	log.Println("Executing query:", query)

	if r.db == nil {
		log.Fatalln("Database connection is nil")
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	err := r.db.Select(&albums, query)

	if err != nil {
		log.Println("Error executing query:", err)
		c.JSON(500, gin.H{"error": "Failed to fetch clients"})
		return
	}

	log.Printf("Fetched %d albums", len(albums))

	if len(albums) == 0 {
		c.JSON(200, gin.H{"message": "No clients found"})
		return
	}

	c.JSON(200, albums)
	log.Println(albums)
}

func (r *Handler) getAllAlbumsByRating(c *gin.Context) {
	log.Println("getAllAlbumByRating called")

	var albums []oooGlebusApi.Album

	query := "SELECT album.id, client.username, album.name, album.countofplays, album.rating, album.countofmusic, album.status, album.image_uri FROM album JOIN author_album ON album.id = author_album.album_id JOIN client ON author_album.author_id = client.id WHERE client.role = 'AUTHOR' ORDER BY rating DESC"
	log.Println("Executing query:", query)

	if r.db == nil {
		log.Fatalln("Database connection is nil")
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	err := r.db.Select(&albums, query)
	if err != nil {
		log.Println("Error executing query:", err)
		c.JSON(500, gin.H{"error": "Failed to fetch clients"})
		return
	}

	log.Printf("Fetched %d albums\n", len(albums))

	if len(albums) == 0 {
		c.JSON(200, gin.H{"message": "No clients found"})
		return
	}

	c.JSON(200, albums)
}

func (r *Handler) getAllAlbumsByRatingByClientId(c *gin.Context) {
	log.Println("getAllAlbumByRatingByClientId called")

	id := c.Param("id")

	var albums []oooGlebusApi.Album

	query := "SELECT album.id, client.username, album.name, album.countofplays, album.rating, album.countofmusic, album.status, album.image_uri FROM album JOIN author_album ON album.id = author_album.album_id JOIN client ON author_album.author_id = client.id WHERE client.role = 'AUTHOR' and client.id = $1 ORDER BY rating DESC"
	log.Println("Executing query:", query, id)

	if r.db == nil {
		log.Fatalln("Database connection is nil")
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	err := r.db.Select(&albums, query)
	if err != nil {
		log.Println("Error executing query:", err)
		c.JSON(500, gin.H{"error": "Failed to fetch clients"})
		return
	}

	log.Printf("Fetched %d albums\n", len(albums))

	if len(albums) == 0 {
		c.JSON(200, gin.H{"message": "No clients found"})
		return
	}

	c.JSON(200, albums)
}

func (h *Handler) getAlbumById(c *gin.Context) {
	id := c.Param("id")

	var album oooGlebusApi.Album

	query := `SELECT album.id, client.username, album.name, album.countofplays, album.rating, album.countofmusic, album.status, album.image_uri FROM album JOIN author_album ON album.id = author_album.album_id JOIN client ON author_album.author_id = client.id WHERE client.role = 'AUTHOR' AND album_id = $1`

	err := h.db.Get(&album, query, id)
	if err != nil {
		log.Println("Error fetching album by ID:", err)
		c.JSON(404, gin.H{"error": "Album not found"})
		return
	}

	c.JSON(200, album)
}

func (h *Handler) updateAlbum(c *gin.Context) {
	id := c.Param("id")

	var albums oooGlebusApi.Album
	if err := c.BindJSON(&albums); err != nil {
		log.Println("Error parsing input:", err)
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	query := `UPDATE album SET name = $1, countofplays = $2, rating = $3, countofmusic = $4, status = $5, image_uri = $6 WHERE id = $7`

	_, err := h.db.Exec(query, albums.Name, albums.Countofplays, albums.Rating, albums.Countofmusic, albums.Status, albums.Image_uri, id)
	if err != nil {
		log.Println("Error updating album:", err)
		c.JSON(500, gin.H{"error": "Failed to update client"})
		return
	}

	c.JSON(200, gin.H{"message": "Album updated successfully"})
}

func (h *Handler) deleteAlbum(c *gin.Context) {
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
