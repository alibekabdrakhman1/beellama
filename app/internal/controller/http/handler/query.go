package handler

import (
	"github.com/alibekabdrakhman1/beellama/app/internal/model"
	"github.com/alibekabdrakhman1/beellama/app/internal/service"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

func NewQueryHandler(service *service.Service, logger *slog.Logger) *QueryHandler {
	return &QueryHandler{
		service: service,
		logger:  logger,
	}
}

type QueryHandler struct {
	service *service.Service
	logger  *slog.Logger
}

type RequestProcess struct {
	Text string `json:"text"`
}

type ResponseProcessSuccess struct {
	Response string `json:"response"`
}

type ResponseProcessFailure struct {
	Error string `json:"error"`
}

// @Summary Process user query with tinyllama
// @Description Process user query with tinyllama
// @ID process
// @Tags auth
// @Accept json
// @Produce json
// @Param request body RequestProcess true "Process request payload"
// @Success 200 {object} ResponseProcessSuccess "Successful response"
// @Failure 401 {object} ResponseProcessFailure "Unauthorized"
// @Failure 500 {object} ResponseProcessFailure "Unauthorized"
// @Router /api/process [post]
func (h *QueryHandler) Process(c *gin.Context) {
	var request struct {
		Text string `json:"text" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := h.service.Query.ProcessQuery(c.Request.Context(), request.Text)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": response})
}

type SuccessResponseHistory struct {
	History []model.Query `json:"history"`
}

// History @Summary History user queries
// @Description History user queries
// @ID history
// @Tags query
// @Produce json
// @Success 200 {object} SuccessResponseHistory "Successful response"
// @Failure 401 {object} ResponseProcessFailure "Unauthorized"
// @Failure 500 {object} ResponseProcessFailure "Unauthorized"
// @Router /api/history [get]
func (h *QueryHandler) History(c *gin.Context) {
	history, err := h.service.Query.GetHistory(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"history": history})
}
