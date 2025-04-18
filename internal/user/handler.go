package user

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) RegisterRoutes(r *gin.Engine) {
	r.POST("/users", h.createUser)
	r.GET("/users", h.getAllUsers)
	r.DELETE("/users/:id", h.deleteUser)
}

// CreateUser godoc
// @Summary Create new user
// @Description Add a new user to the system
// @Tags users
// @Accept json
// @Produce json
// @Param user body CreateUserInput true "User data"
// @Success 200 {object} ent.User
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /users [post]
func (h *Handler) createUser(c *gin.Context) {
	var input CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Invalid input"})
		return
	}

	user, err := h.service.CreateUser(context.Background(), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Could not create user"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// GetAllUsers godoc
// @Summary Get all users
// @Description Get all users from the database
// @Tags users
// @Produce json
// @Success 200 {array} ent.User
// @Failure 500 {object} ErrorResponse
// @Router /users [get]
func (h *Handler) getAllUsers(c *gin.Context) {
	users, err := h.service.GetAllUsers(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Could not retrieve users"})
		return
	}

	c.JSON(http.StatusOK, users)
}

// DeleteUser godoc
// @Summary Delete user
// @Description Delete a user by ID
// @Tags users
// @Param id path int true "User ID"
// @Produce json
// @Success 200 {object} map[string]string
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /users/{id} [delete]
func (h *Handler) deleteUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Invalid user ID"})
		return
	}

	if err := h.service.DeleteUser(context.Background(), id); err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{Message: "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}
