package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"oooGlebusApi"
)

func (r *Handler) getAllClient(c *gin.Context) {
	log.Println("getAllClient called")

	var clients []oooGlebusApi.Client

	query := "SELECT id, username, email, phone_number, role, password FROM client"
	log.Println("Executing query:", query)

	if r.db == nil {
		log.Fatalln("Database connection is nil")
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	err := r.db.Select(&clients, query)
	if err != nil {
		log.Println("Error executing query:", err)
		c.JSON(500, gin.H{"error": "Failed to fetch clients"})
		return
	}

	log.Printf("Fetched %d clients\n", len(clients))

	if len(clients) == 0 {
		c.JSON(200, gin.H{"message": "No clients found"})
		return
	}

	c.JSON(200, clients)
}

func (h *Handler) getClientById(c *gin.Context) {
	id := c.Param("id")

	var client oooGlebusApi.Client
	query := "SELECT id, username, email, phone_number, role, password FROM client WHERE id = $1"

	err := h.db.Get(&client, query, id)
	if err != nil {
		log.Println("Error fetching client by ID:", err)
		c.JSON(404, gin.H{"error": "Client not found"})
		return
	}

	c.JSON(200, client)
}

func (h *Handler) updateClient(c *gin.Context) {
	id := c.Param("id")

	var input oooGlebusApi.Client
	if err := c.BindJSON(&input); err != nil {
		log.Println("Error parsing input:", err)
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	query := `
		UPDATE client
		SET username = $1, email = $2, phone_number = $3, role = $4, password = $5
		WHERE id = $6
	`

	_, err := h.db.Exec(query, input.Username, input.Email, input.Phone_Number, input.Role, input.Password, id)
	if err != nil {
		log.Println("Error updating client:", err)
		c.JSON(500, gin.H{"error": "Failed to update client"})
		return
	}

	c.JSON(200, gin.H{"message": "Client updated successfully"})
}

func (h *Handler) deleteClient(c *gin.Context) {
	id := c.Param("id")

	query := "DELETE FROM client WHERE id = $1"

	_, err := h.db.Exec(query, id)
	if err != nil {
		log.Println("Error deleting client:", err)
		c.JSON(500, gin.H{"error": "Failed to delete client"})
		return
	}

	c.JSON(200, gin.H{"message": "Client deleted successfully"})
}
