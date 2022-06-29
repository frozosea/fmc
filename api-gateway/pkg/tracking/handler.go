package tracking

import (
	"fmc-with-git/internal/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HttpHandler struct {
	client *Client
	utils.HttpUtils
}

// @Summary      Track by bill number
// @Description  tracking by container number
// @accept json
// @Produce      json
// @Param input body Track true "info"
// @Tags         Tracking
// @Success      200  {object}  ContainerNumberResponse
// @Success      204
// @Failure      400
// @Router       /tracking/trackByContainerNumber [post]
func (h *HttpHandler) TrackByContainerNumber(c *gin.Context) {
	var schema *Track
	if validateErr := c.ShouldBindJSON(&schema); validateErr != nil {
		h.ValidateSchemaError(c, http.StatusBadRequest, "invalid input body")
		return
	}
	fmt.Println(schema)
	response, err := h.client.TrackByBillNumber(c.Request.Context(), schema)
	if err != nil {
		c.JSON(204, gin.H{"message": "container not found"})
		return
	}
	c.JSON(200, response)
	return
}

// @Summary      Track by bill number
// @Description  tracking by bill number, if eta not found will be 0
// @Tags         Tracking
// @Param input body Track true "info"
// @Success      200  {object}  BillNumberResponse
// @Success      204
// @Failure      400
// @Router       /tracking/trackByBillNumber [post]
func (h *HttpHandler) TrackByBillNumber(c *gin.Context) {
	var schema *Track
	if validateErr := c.ShouldBindJSON(&schema); validateErr != nil {
		h.ValidateSchemaError(c, http.StatusBadRequest, "invalid input body")
		return
	}
	fmt.Println(schema)
	response, err := h.client.TrackByBillNumber(c.Request.Context(), schema)
	if err != nil {
		c.JSON(204, gin.H{})
		return
	}
	c.JSON(200, response)
	return
}

func NewHttpHandler(client *Client) *HttpHandler {
	return &HttpHandler{client: client}
}
