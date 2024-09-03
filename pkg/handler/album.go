package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) createAlbum(c *gin.Context) {
	id, _ := c.Get(clientCTX)
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAllAlbum(c *gin.Context) {

}

func (h *Handler) getAlbumById(c *gin.Context) {

}

func (h *Handler) updateAlbum(c *gin.Context) {

}

func (h *Handler) deleteAlbum(c *gin.Context) {

}
