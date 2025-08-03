package handler

import (
	"backend/internal/model"
	"backend/internal/service"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// AuthHandler định nghĩa các handlers liên quan đến xác thực
type AuthHandler struct {
	authService service.AuthService
}

// NewAuthHandler tạo instance mới của AuthHandler
func NewAuthHandler(authService service.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

// SetupRoutes thiết lập các routes cho authentication
func (h *AuthHandler) SetupRoutes(router *gin.Engine) {
	auth := router.Group("/api/auth")
	{
		auth.POST("/signup", h.Signup)
		auth.POST("/login", h.Login)
		auth.GET("/me", h.AuthMiddleware(), h.GetMe)
	}
}

// Signup xử lý đăng ký người dùng mới
func (h *AuthHandler) Signup(c *gin.Context) {
	var userSignup model.UserSignup

	if err := c.ShouldBindJSON(&userSignup); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	response, err := h.authService.Signup(&userSignup)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, response)
}

// Login xử lý đăng nhập
func (h *AuthHandler) Login(c *gin.Context) {
	var userLogin model.UserLogin

	if err := c.ShouldBindJSON(&userLogin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	response, err := h.authService.Login(&userLogin)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

// GetMe trả về thông tin người dùng hiện tại
func (h *AuthHandler) GetMe(c *gin.Context) {
	// Lấy user từ context (được set bởi AuthMiddleware)
	userAny, exists := c.Get("currentUser")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	user, ok := userAny.(*model.User)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": model.UserResponse{
			ID:        user.ID,
			Email:     user.Email,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			CreatedAt: user.CreatedAt,
		},
	})
}

// AuthMiddleware xác thực JWT token từ Authorization header
func (h *AuthHandler) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		// Trích xuất token từ Bearer header
		tokenParts := strings.Split(authHeader, "Bearer ")
		if len(tokenParts) != 2 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization format"})
			c.Abort()
			return
		}

		tokenString := tokenParts[1]
		user, err := h.authService.GetUserFromToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// Lưu user vào context để các handler khác có thể sử dụng
		c.Set("currentUser", user)
		c.Next()
	}
}
