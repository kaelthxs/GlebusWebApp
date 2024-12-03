package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"oooGlebusApi"
)

func (h *Handler) createMusic(c *gin.Context) {

	var musics oooGlebusApi.Music
	if err := c.BindJSON(&musics); err != nil {
		log.Println("Error parsing input:", err)
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	query := `INSERT INTO music (name, rating, countofplays, album_id, image_uri, audio_uri) VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := h.db.Exec(query, musics.Name, musics.Rating, musics.Countofplays, musics.Album_id, musics.Image_uri, musics.Audio_uri)
	if err != nil {
		log.Println("Error creating album:", err)
		c.JSON(500, gin.H{"error": "Failed to create album"})
		return
	}

	c.JSON(200, gin.H{"message": "Album created successfully"})

	var author_music oooGlebusApi.Author_Music

	if err := c.BindJSON(&musics); err != nil {
		log.Println("Error parsing input:", err)
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	query2 := `INSERT INTO author_music (music_id, author_id) VALUES ($1, $2)`

	_, err2 := h.db.Exec(query2, author_music.Music_id, author_music.Author_id)
	if err2 != nil {
		log.Println("Error creating album:", err2)
		c.JSON(500, gin.H{"error": "Failed to create album"})
		return
	}

	c.JSON(200, gin.H{"message": "Album created successfully"})
}

func (r *Handler) getAllMusic(c *gin.Context) {
	log.Println("getAllMusic called")

	var musics []oooGlebusApi.Music

	query := "SELECT music.id, client.username, music.name, music.rating, music.countofplays, music.album_id, music.image_uri, music.audio_uri FROM music JOIN author_music ON music.id = author_music.music_id JOIN client ON author_music.author_id = client.id WHERE client.role = 'AUTHOR'"
	log.Println("Executing query:", query)

	if r.db == nil {
		log.Fatalln("Database connection is nil")
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	err := r.db.Select(&musics, query)
	if err != nil {
		log.Println("Error executing query:", err)
		c.JSON(500, gin.H{"error": "Failed to fetch clients"})
		return
	}

	log.Printf("Fetched %d musics\n", len(musics))

	if len(musics) == 0 {
		c.JSON(200, gin.H{"message": "No musics found"})
		return
	}

	c.JSON(200, musics)
}

func (r *Handler) getAllMusicByAlbumId(c *gin.Context) {
	log.Println("getAllMusic called")

	id := c.Param("id")

	var musics []oooGlebusApi.Music

	query := "SELECT music.id, client.username, music.name, music.rating, music.countofplays, music.album_id, music.image_uri, music.audio_uri FROM music JOIN author_music ON music.id = author_music.music_id JOIN client ON author_music.author_id = client.id WHERE client.role = 'AUTHOR' AND album_id = $1"
	log.Println("Executing query:", query, id)

	if r.db == nil {
		log.Fatalln("Database connection is nil")
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	err := r.db.Select(&musics, query, id)
	if err != nil {
		log.Println("Error executing query:", err)
		c.JSON(500, gin.H{"error": "Failed to fetch clients"})
		return
	}

	log.Printf("Fetched %d musics\n", len(musics))

	if len(musics) == 0 {
		c.JSON(200, gin.H{"message": "No musics found"})
		return
	}

	c.JSON(200, musics)
}

func (h *Handler) getMusicById(c *gin.Context) {
	id := c.Param("id")

	var musics oooGlebusApi.Music
	query := "SELECT music.id, music.name, music.rating, music.countofplays, music.album_id, music.image_uri, music.audio_uri FROM music JOIN author_music ON music.id = author_music.music_id JOIN client ON author_music.author_id = client.id WHERE client.role = 'AUTHOR' AND id = $1"

	err := h.db.Get(&musics, query, id)
	if err != nil {
		log.Println("Error fetching music by ID:", err)
		c.JSON(404, gin.H{"error": "Music not found"})
		return
	}

	c.JSON(200, musics)
}

func (h *Handler) updateMusic(c *gin.Context) {
	id := c.Param("id")

	var musics oooGlebusApi.Music
	if err := c.BindJSON(&musics); err != nil {
		log.Println("Error parsing input:", err)
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	query := `UPDATE music SET name = $1, rating = $2, countofplays = $3, album_id = $4, image_uri = $5, audio_uri = $6 WHERE id = $7`

	_, err := h.db.Exec(query, musics.Name, musics.Rating, musics.Countofplays, musics.Album_id, musics.Image_uri, musics.Audio_uri, id)
	if err != nil {
		log.Println("Error updating music:", err)
		c.JSON(500, gin.H{"error": "Failed to update music"})
		return
	}

	c.JSON(200, gin.H{"message": "Music updated successfully"})
}

func (h *Handler) deleteMusic(c *gin.Context) {
	id := c.Param("id")

	query := "DELETE FROM music WHERE id = $1"

	_, err := h.db.Exec(query, id)
	if err != nil {
		log.Println("Error deleting music:", err)
		c.JSON(500, gin.H{"error": "Failed to delete music"})
		return
	}

	c.JSON(200, gin.H{"message": "music deleted successfully"})
}
