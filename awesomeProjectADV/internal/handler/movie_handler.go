package handler

import (
	"awesomeProjectADV/internal/domain/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type MovieHandler struct {
	Usecase interface {
		CreateMovie(*model.Movie) error
		GetAllMovies() ([]*model.Movie, error)
		GetMovieByID(string) (*model.Movie, error)
		Delete(string) error // ✅ добавляем сюда
	}
}

func NewMovieHandler(uc interface {
	CreateMovie(*model.Movie) error
	GetAllMovies() ([]*model.Movie, error)
	GetMovieByID(string) (*model.Movie, error)
	Delete(string) error // ✅ добавляем сюда!
}) *MovieHandler {
	return &MovieHandler{Usecase: uc}
}

func (h *MovieHandler) RegisterRoutes(r *gin.Engine) {
	r.POST("/movies", h.CreateMovie)
	r.GET("/movies", h.GetAll)
	r.GET("/movies/:id", h.GetByID)
	r.DELETE("/movies/:id", h.DeleteMovie)
}

func (h *MovieHandler) CreateMovie(c *gin.Context) {
	var movie model.Movie
	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Usecase.CreateMovie(&movie); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Movie created"})
}

func (h *MovieHandler) GetAll(c *gin.Context) {
	movies, err := h.Usecase.GetAllMovies()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, movies)
}

func (h *MovieHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	movie, err := h.Usecase.GetMovieByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
		return
	}

	c.JSON(http.StatusOK, movie)
}
func (h *MovieHandler) DeleteMovie(c *gin.Context) {
	id := c.Param("id")

	if err := h.Usecase.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при удалении"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Фильм удалён"})
}
