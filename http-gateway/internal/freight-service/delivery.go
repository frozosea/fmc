package freight_service

import (
	_ "fmc-gateway/internal/schedule-tracking"
	"github.com/gin-gonic/gin"
)

type Http struct {
	client IClient
}

func NewHttp(cli IClient) *Http {
	return &Http{client: cli}
}

// GetFreights
// @Summary get all freights
// @accept json
// @Tags         Freights
// @Param input body GetFreight true "body"
// @Success 200 {object} []Freight
// @Failure 500 {object} BaseResponse
// @Router /freights [get]
func (h *Http) GetFreights(c *gin.Context) {
	var r *GetFreight
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(400, gin.H{"success": false, "error": "cannot validate your request"})
		return
	}
	result, err := h.client.GetFreights(c.Request.Context(), r)
	if err != nil {
		c.JSON(500, gin.H{"success": false, "error": err.Error()})
		return
	}
	c.JSON(200, result)
}

// GetAllCities
// @Summary get all cities
// @Description get all cities
// @accept json
// @Tags         Freights
// @Success 200 {object} []City
// @Failure 500 {object} BaseResponse
// @Router /cities [get]
func (h *Http) GetAllCities(c *gin.Context) {
	ctx := c.Request.Context()
	result, err := h.client.GetCities(ctx)
	if err != nil {
		c.JSON(500, gin.H{"success": false, "error": err.Error()})
		return
	}
	c.JSON(200, result)
}

// GetAllCompanies
// @Summary get all contacts
// @Tags         Freights
// @Success 200 {object} []Company
// @Failure 500 {object} BaseResponse
// @Router /companies [get]
func (h *Http) GetAllCompanies(c *gin.Context) {
	ctx := c.Request.Context()
	result, err := h.client.GetCompanies(ctx)
	if err != nil {
		c.JSON(500, gin.H{"success": false, "error": err.Error()})
		return
	}
	c.JSON(200, result)
}

// GetAllContainers
// @Summary get all containers
// @Tags         Freights
// @Success 200 {object} []Container
// @Failure 500 {object} BaseResponse
// @Router /containers [get]
func (h *Http) GetAllContainers(c *gin.Context) {
	ctx := c.Request.Context()
	r, err := h.client.GetContainers(ctx)
	if err != nil {
		c.JSON(500, gin.H{"success": false, "error": err.Error()})
		return
	}
	c.JSON(200, r)
}
