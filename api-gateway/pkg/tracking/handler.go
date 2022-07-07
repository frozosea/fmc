package tracking

import (
	"fmc-with-git/internal/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HttpHandler struct {
	client    *Client
	validator *Validator
	*utils.HttpUtils
}

func NewHttpHandler(client *Client, httpUtils *utils.HttpUtils) *HttpHandler {
	return &HttpHandler{client: client, HttpUtils: httpUtils}
}

// TrackByContainerNumber
// @Summary      Track by bill number
// @Security ApiKeyAuth
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
	var schema Track
	if err := h.Validate(c, &schema); err != nil {
		h.ValidateSchemaError(c, http.StatusBadRequest, "invalid input body")
		return
	}
	if err := h.validator.ValidateContainer(schema.Number); err != nil {
		h.ValidateSchemaError(c, http.StatusBadRequest, "invalid input body")
		return
	}
	if err := h.validator.ValidateScac(schema.Scac); err != nil {
		h.ValidateSchemaError(c, http.StatusBadRequest, "invalid input body")
		return
	}
	response, err := h.client.TrackByContainerNumber(c.Request.Context(), schema, c.Request.Header.Get("X-REAL-IP"))
	if err != nil {
		c.JSON(204, gin.H{"message": "container not found"})
		return
	}
	c.JSON(200, response)
	return
}

// TrackByBillNumber
// @Summary      Track by bill number
// @Security ApiKeyAuth
// @Description  tracking by bill number, if eta not found will be 0
// @Tags         Tracking
// @Param input body Track true "info"
// @Success      200  {object}  BillNumberResponse
// @Success      204
// @Failure      400
// @Router       /tracking/trackByBillNumber [post]
func (h *HttpHandler) TrackByBillNumber(c *gin.Context) {
	var schema *Track
	if err := h.Validate(c, &schema); err != nil {
		h.ValidateSchemaError(c, http.StatusBadRequest, "invalid input body")
		return
	}
	if err := h.validator.ValidateBillNumber(schema.Number); err != nil {
		h.ValidateSchemaError(c, http.StatusBadRequest, "invalid input body")
		return
	}
	if err := h.validator.ValidateScac(schema.Scac); err != nil {
		h.ValidateSchemaError(c, http.StatusBadRequest, "invalid input body")
		return
	}
	response, err := h.client.TrackByBillNumber(c.Request.Context(), schema, c.Request.Header.Get("X-REAL-IP"))
	fmt.Println(err)
	if err != nil {
		c.JSON(204, gin.H{})
		return
	}
	c.JSON(200, response)
	return
}
