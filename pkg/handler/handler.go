package handler

import (
	"github.com/gin-gonic/gin"
	"oooGlebusApi/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api")
	{
		client := api.Group("/client")
		{
			client.GET("/", h.getAllClient)
			client.GET("/:id", h.getClientById)
			client.PUT("/:id", h.updateClient)
			client.DELETE("/:id", h.deleteClient)
		}

		album := api.Group("/album")
		{
			album.POST("/", h.createAlbum)
			album.GET("/", h.getAllAlbum)
			album.GET("/:id", h.getAlbumById)
			album.PUT("/:id", h.updateAlbum)
			album.DELETE("/:id", h.deleteAlbum)
		}

		music := api.Group("/music")
		{
			music.POST("/", h.createMusic)
			music.GET("/", h.getAllMusic)
			music.GET("/:id", h.getMusicById)
			music.PUT("/:id", h.updateMusic)
			music.DELETE("/:id", h.deleteMusic)
		}
	}
	return router
}
