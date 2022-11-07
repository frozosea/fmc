package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HttpHandler struct {
	client IClient
}

func NewHttpHandler(client IClient) *HttpHandler {
	return &HttpHandler{client: client}
}

// Register
// @Summary Register user by username and password
// @Description register user by username and password
// @accept json
// @Param input body RegisterUser true "body"
// @Tags         Auth
// @Success 200 {object} BaseResponse
// @Failure 500 {object} BaseResponse
// @Router /auth/register [post]
func (h *HttpHandler) Register(c *gin.Context) {
	var s *RegisterUser
	if err := c.ShouldBindJSON(&s); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}
	if err := h.client.Register(c.Request.Context(), s); err != nil {
		switch err.(type) {
		case *AlreadyRegisterError:
			c.JSON(422, BaseResponse{
				Success: false,
				Error:   err.Error(),
			})
			return
		default:
			c.JSON(500, BaseResponse{
				Success: false,
				Error:   err.Error(),
			})
			return
		}
	}
	c.JSON(200, BaseResponse{
		Success: true,
		Error:   "",
	})
	return
}

// Login
// @Summary Login user by username and password
// @Description login user by username and password, tokens expires is unix timestamp
// @accept json
// @Param input body User true "info"
// @Tags         Auth
// @Success 200 {object} LoginUserResponse
// @Failure 500 {object} BaseResponse
// @Failure 422 {object} BaseResponse
// @Router /auth/login [post]
func (h *HttpHandler) Login(c *gin.Context) {
	var s *User
	if err := c.ShouldBindJSON(&s); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}
	token, err := h.client.Login(c.Request.Context(), s)
	if err != nil {
		switch err.(type) {
		case *InvalidUserError:
			c.JSON(422, gin.H{"success": false, "error": "no user with this data"})
			return
		default:
			c.JSON(500, gin.H{"success": false, "error": err.Error()})
			return
		}
	}
	c.JSON(200, token)
	return
}

// Refresh
// @Summary Refresh token
// @Security ApiKeyAuth
// @Description refresh token by refresh token
// @accept json
// @Param input body RefreshTokenRequest true "info"
// @Tags         Auth
// @Success 200 {object} LoginUserResponse
// @Failure 500 {object} BaseResponse
// @Router /auth/refresh [post]
func (h *HttpHandler) Refresh(c *gin.Context) {
	var s RefreshTokenRequest
	if err := c.ShouldBindJSON(&s); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}
	token, err := h.client.RefreshToken(c.Request.Context(), s.RefreshToken)
	if err != nil {
		c.JSON(500, gin.H{"success": false, "error": err.Error()})
		return
	}
	c.Request.Header.Set("Authorization", "Bearer "+token.Token)
	c.JSON(200, token)
	return
}
