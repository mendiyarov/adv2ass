package handler

import (
	"awesomeProjectADV/internal/domain/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SessionHandler struct {
	Usecase interface {
		CreateSession(*model.Session) error
		GetAllSessions() ([]*model.Session, error)
		GetSessionByID(string) (*model.Session, error)
	}
}

func NewSessionHandler(uc interface {
	CreateSession(*model.Session) error
	GetAllSessions() ([]*model.Session, error)
	GetSessionByID(string) (*model.Session, error)
}) *SessionHandler {
	return &SessionHandler{Usecase: uc}
}

func (h *SessionHandler) RegisterRoutes(r *gin.Engine) {
	r.POST("/sessions", h.Create)
	r.GET("/sessions", h.GetAll)
	r.GET("/sessions/:id", h.GetByID)
}

func (h *SessionHandler) Create(c *gin.Context) {
	var session model.Session
	if err := c.ShouldBindJSON(&session); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Usecase.CreateSession(&session); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Session created"})
}

func (h *SessionHandler) GetAll(c *gin.Context) {
	sessions, err := h.Usecase.GetAllSessions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, sessions)
}

func (h *SessionHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	session, err := h.Usecase.GetSessionByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Session not found"})
		return
	}
	c.JSON(http.StatusOK, session)
}
