package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"oooGlebusApi/pkg/service"
)

type Handler struct {
	services *service.Service
	db       *sqlx.DB
}

func NewHandler(services *service.Service, db *sqlx.DB) *Handler {
	return &Handler{
		services: services,
		db:       db,
	}
}
func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.clientIdentity)
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
			album.GET("/byrating/", h.getAllAlbumsByRating)
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

		review := api.Group("/review")
		{
			review.POST("/", h.createReview)
			review.GET("/", h.getAllReview)
			review.GET("/:id", h.getReviewById)
			review.PUT("/:id", h.updateReview)
			review.DELETE("/:id", h.deleteReview)
		}
	}
	return router
}
