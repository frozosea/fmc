package user

import (
	"fmc-with-git/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HttpHandler struct {
	client *Client
	*utils.HttpUtils
}

func NewHttpHandler(client *Client, httpUtils *utils.HttpUtils) *HttpHandler {
	return &HttpHandler{client: client, HttpUtils: httpUtils}
}
func (h *HttpHandler) addBillOrContainer(c *gin.Context, isContainer bool) {
	var s AddContainers
	if err := h.Validate(c, &s); err != nil {
		h.ValidateSchemaError(c, http.StatusBadRequest, "invalid input body")
		return
	}
	userId, err := h.DecodeToken(c)
	if err != nil {
		c.JSON(401, gin.H{"message": "cannot decode token"})
		return
	}
	if isContainer {
		if err := h.client.AddContainerToAccount(c.Request.Context(), int64(userId), &s); err != nil {
			c.JSON(500, gin.H{"success": false, "error": err.Error()})
			return
		}
	} else {
		if err := h.client.AddBillNumbersToAccount(c.Request.Context(), int64(userId), &s); err != nil {
			c.JSON(500, gin.H{"success": false, "error": err.Error()})
			return
		}
	}
	c.JSON(200, gin.H{"success": true, "error": nil})
	return
}

// AddContainersToAccount
// @Summary Add containers to account
// @Security ApiKeyAuth
// @Description Add containers to account
// @accept json
// @Param input body AddContainers true "info"
// @Tags         User
// @Success 200 {object} BaseResponse
// @Failure 500 {object} BaseResponse
// @Router /user/addContainers [post]
func (h *HttpHandler) AddContainersToAccount(c *gin.Context) {
	h.addBillOrContainer(c, true)
}

// AddBillNumbersToAccount
// @Summary Add bill numbers to account
// @Security ApiKeyAuth
// @Description Add bill numbers to account
// @accept json
// @Param input body AddContainers true "info"
// @Tags         User
// @Success 200 {object} BaseResponse
// @Failure 500 {object} BaseResponse
// @Router /user/addBillNumbers [post]
func (h *HttpHandler) AddBillNumbersToAccount(c *gin.Context) {
	h.addBillOrContainer(c, false)
}

func (h *HttpHandler) deleteContainersOrBillNumbers(c *gin.Context, isContainer bool) {
	var s DeleteNumbers
	if err := h.Validate(c, &s); err != nil {
		h.ValidateSchemaError(c, http.StatusBadRequest, "invalid input body")
		return
	}
	userId, err := h.DecodeToken(c)
	if err != nil {
		c.JSON(401, gin.H{"message": "cannot decode token"})
		return
	}
	if isContainer {
		if err := h.client.DeleteContainersFromAccount(c.Request.Context(), int64(userId), &s); err != nil {
			c.JSON(500, gin.H{"success": false, "error": err.Error()})
			return
		}
	} else {
		if err := h.client.DeleteBillNumbersFromAccount(c.Request.Context(), int64(userId), &s); err != nil {
			c.JSON(500, gin.H{"success": false, "error": err.Error()})
			return
		}
	}
	c.JSON(200, gin.H{"success": true, "error": nil})
	return
}

// DeleteContainersFromAccount
// @Summary Delete containers from account
// @Security OAuth2Application
// @Description delete containers from account
// @accept json
// @Param input body DeleteNumbers true "info"
// @Tags         User
// @Success 200 {object} BaseResponse
// @Failure 500 {object} BaseResponse
// @Router /user/deleteContainers [delete]
func (h *HttpHandler) DeleteContainersFromAccount(c *gin.Context) {
	h.deleteContainersOrBillNumbers(c, true)
}

// DeleteBillNumbersFromAccount
// @Summary Delete bill numbers from account
// @Security OAuth2Application
// @Description delete bill numbers from account
// @accept json
// @Param input body DeleteNumbers true "info"
// @Tags         User
// @Success 200 {object} BaseResponse
// @Failure 500 {object} BaseResponse
// @Router /user/deleteBillNumbers [delete]
func (h *HttpHandler) DeleteBillNumbersFromAccount(c *gin.Context) {
	h.deleteContainersOrBillNumbers(c, false)
}

// GetAll
// @Summary Get all bill numbers and containers from account
// @Security OAuth2Application
// @Description Get all bill numbers and containers from account
// @accept json
// @Tags         User
// @Success 200 {object} AllContainersAndBillNumbers
// @Failure 500 {object} BaseResponse
// @Router /user/getAllBillsContainers [get]
func (h *HttpHandler) GetAll(c *gin.Context) {
	userId, err := h.DecodeToken(c)
	if err != nil {
		c.JSON(401, gin.H{"message": "cannot decode token"})
		return
	}
	response, err := h.client.GetAll(c.Request.Context(), int64(userId))
	if err != nil {
		c.JSON(500, gin.H{"success": false, "error": err.Error()})
		return
	}
	c.JSON(200, response)
	return
}