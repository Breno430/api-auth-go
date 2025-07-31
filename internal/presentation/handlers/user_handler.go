package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"api-auth-go/internal/domain/entities"
	"api-auth-go/internal/domain/usecases"
)

type UserHandler struct {
	userUseCase *usecases.UserUseCase
}

func NewUserHandler(userUseCase *usecases.UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: userUseCase,
	}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var input usecases.CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	output, err := h.userUseCase.CreateUser(c.Request.Context(), input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, output)
}

func (h *UserHandler) Login(c *gin.Context) {
	var input usecases.LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	output, err := h.userUseCase.Login(c.Request.Context(), input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, output)
}

func (h *UserHandler) GetProfile(c *gin.Context) {
	userID := c.GetString("user_id")
	userEmail := c.GetString("user_email")
	userName := c.GetString("user_name")
	userRole := c.GetString("user_role")

	c.JSON(http.StatusOK, gin.H{
		"id":      userID,
		"email":   userEmail,
		"name":    userName,
		"role":    userRole,
		"message": "Profile retrieved successfully",
	})
}

func (h *UserHandler) RequestPasswordReset(c *gin.Context) {
	var input usecases.RequestPasswordResetInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	output, err := h.userUseCase.RequestPasswordReset(c.Request.Context(), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, output)
}

func (h *UserHandler) ResetPassword(c *gin.Context) {
	var input usecases.ResetPasswordInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	output, err := h.userUseCase.ResetPassword(c.Request.Context(), input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, output)
}

func (h *UserHandler) ListUsers(c *gin.Context) {
	userID := c.GetString("user_id")

	filters := &entities.UserFilters{
		Name:      c.Query("name"),
		Email:     c.Query("email"),
		Role:      c.Query("role"),
		Page:      1,
		Limit:     10,
		SortBy:    c.Query("sort_by"),
		SortOrder: c.Query("sort_order"),
	}

	if pageStr := c.Query("page"); pageStr != "" {
		if page, err := strconv.Atoi(pageStr); err == nil && page > 0 {
			filters.Page = page
		}
	}

	if limitStr := c.Query("limit"); limitStr != "" {
		if limit, err := strconv.Atoi(limitStr); err == nil && limit > 0 {
			filters.Limit = limit
		}
	}

	output, err := h.userUseCase.ListUsers(c.Request.Context(), userID, filters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, output)
}

func (h *UserHandler) GetUserByID(c *gin.Context) {
	userID := c.Param("id")

	output, err := h.userUseCase.GetUserByID(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, output)
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	userID := c.Param("id")

	var input usecases.UpdateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	output, err := h.userUseCase.UpdateUser(c.Request.Context(), userID, input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, output)
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	userID := c.Param("id")
	output, err := h.userUseCase.DeleteUser(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, output)
}
