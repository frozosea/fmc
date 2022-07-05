package schedule_tracking

import (
	"fmc-with-git/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HttpHandler struct {
	client *Client
	utils  *utils.HttpUtils
}

func NewHttpHandler(client *Client, utils *utils.HttpUtils) *HttpHandler {
	return &HttpHandler{client: client, utils: utils}
}

// AddContainersOnTrack
// @Summary      add containers on track
// @Security ApiKeyAuth
// @Description  add containers on track. Every day in your selected time track container and send email with result about it.
// @accept json
// @Produce      json
// @Param input body AddOnTrackRequest true "info"
// @Tags         Schedule Tracking
// @Success      200  {object}  AddOnTrackResponse
// @Failure      400
// @Failure      500  {object}  BaseResponse
// @Router       /schedule/addContainer [post]
func (h *HttpHandler) AddContainersOnTrack(c *gin.Context) {
	var s AddOnTrackRequest
	if err := h.utils.Validate(c, &s); err != nil {
		h.utils.ValidateSchemaError(c, http.StatusBadRequest, "invalid input body")
		return
	}
	userId, err := h.utils.DecodeToken(c)
	if err != nil {
		c.JSON(401, gin.H{"message": "cannot decode token"})
		return
	}
	resp, err := h.client.AddContainersOnTrack(c.Request.Context(), int(userId), &s)
	if err != nil {
		c.JSON(500, gin.H{"success": false, "error": err.Error()})
		return
	}
	c.JSON(200, resp)
	return
}

// AddBillNumbersOnTrack
// @Summary      add bill numbers on track
// @Security ApiKeyAuth
// @Description  add bill numbers on track. Every day in your selected time track bill numbers and send email with result about it.
// @accept json
// @Produce      json
// @Param input body AddOnTrackRequest true "info"
// @Tags         Schedule Tracking
// @Success      200  {object}  AddOnTrackResponse
// @Failure      400
// @Failure      500  {object}  BaseResponse
// @Router       /schedule/addBillNo [post]
func (h *HttpHandler) AddBillNumbersOnTrack(c *gin.Context) {
	var s AddOnTrackRequest
	if err := h.utils.Validate(c, &s); err != nil {
		h.utils.ValidateSchemaError(c, http.StatusBadRequest, "invalid input body")
		return
	}
	userId, err := h.utils.DecodeToken(c)
	if err != nil {
		c.JSON(401, gin.H{"message": "cannot decode token"})
		return
	}
	resp, err := h.client.AddBillNosOnTrack(c.Request.Context(), int(userId), &s)
	if err != nil {
		c.JSON(500, gin.H{"success": false, "error": err.Error()})
		return
	}
	c.JSON(200, resp)
	return
}

// UpdateTrackingTime
// @Summary      update time of tracking
// @Security ApiKeyAuth
// @Description  update time of tracking, bill or container doesn't matter
// @accept json
// @Produce      json
// @Param input body UpdateTrackingTimeRequest true "info"
// @Tags         Schedule Tracking
// @Success      200  {object}  []BaseAddOnTrackResponse
// @Failure      400
// @Failure 	 500  {object} BaseResponse
// @Router       /schedule/updateTime [put]
func (h *HttpHandler) UpdateTrackingTime(c *gin.Context) {
	var s UpdateTrackingTimeRequest
	if err := h.utils.Validate(c, &s); err != nil {
		h.utils.ValidateSchemaError(c, http.StatusBadRequest, "invalid input body")
		return
	}
	res, err := h.client.UpdateTrackingTime(c.Request.Context(), s)
	if err != nil {
		c.JSON(500, gin.H{"success": false, "error": err.Error()})
		return
	}
	c.JSON(200, res)
	return
}

// AddEmailsOnTracking
// @Summary      add new email to tracking
// @Security ApiKeyAuth
// @Description  add new email to tracking, bill or container doesn't matter
// @accept json
// @Produce      json
// @Param input body UpdateTrackingTimeRequest true "info"
// @Tags         Schedule Tracking
// @Success      200 {object} BaseResponse
// @Failure      400
// @Failure 	 500  {object} BaseResponse
// @Router       /schedule/addEmail [put]
func (h *HttpHandler) AddEmailsOnTracking(c *gin.Context) {
	var s AddEmailRequest
	if err := h.utils.Validate(c, &s); err != nil {
		h.utils.ValidateSchemaError(c, http.StatusBadRequest, "invalid input body")
		return
	}
	if err := h.client.AddEmailsOnTracking(c.Request.Context(), s); err != nil {
		c.JSON(500, gin.H{"success": false, "error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"success": true, "error": nil})
	return
}

// DeleteEmailFromTrack
// @Summary      delete email from tracking
// @Security ApiKeyAuth
// @Description  delete email from tracking, bill or container doesn't matter
// @accept json
// @Produce      json
// @Param input body DeleteEmailFromTrackRequest true "info"
// @Tags         Schedule Tracking
// @Success      200 {object} BaseResponse
// @Failure      400
// @Failure 	 500  {object} BaseResponse
// @Router       /schedule/deleteEmail [delete]
func (h *HttpHandler) DeleteEmailFromTrack(c *gin.Context) {
	var s DeleteEmailFromTrackRequest
	if err := h.utils.Validate(c, &s); err != nil {
		h.utils.ValidateSchemaError(c, http.StatusBadRequest, "invalid input body")
		return
	}
	if err := h.client.DeleteEmailFromTrack(c.Request.Context(), s); err != nil {
		c.JSON(500, gin.H{"success": false, "error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"success": true, "error": nil})
	return
}

// DeleteContainersFromTrack
// @Summary      delete containers from tracking
// @Security ApiKeyAuth
// @Description  delete containers from tracking
// @accept json
// @Produce      json
// @Param input body DeleteFromTrackRequest true "info"
// @Tags         Schedule Tracking
// @Success      200 {object} BaseResponse
// @Failure      400
// @Failure 	 500  {object} BaseResponse
// @Router       /schedule/deleteContainers [delete]
func (h *HttpHandler) DeleteContainersFromTrack(c *gin.Context) {
	var s DeleteFromTrackRequest
	if err := h.utils.Validate(c, &s); err != nil {
		h.utils.ValidateSchemaError(c, http.StatusBadRequest, "invalid input body")
		return
	}
	userId, err := h.utils.DecodeToken(c)
	if err != nil {
		c.JSON(401, gin.H{"message": "cannot decode token"})
		return
	}
	if err := h.client.DeleteFromTracking(c.Request.Context(), true, int64(userId), s); err != nil {
		c.JSON(500, gin.H{"success": false, "error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"success": true, "error": nil})
	return
}

// DeleteBillNumbersFromTrack
// @Summary      delete bill numbers from tracking
// @Security ApiKeyAuth
// @Description  delete numbers from tracking
// @accept json
// @Produce      json
// @Param input body DeleteFromTrackRequest true "info"
// @Tags         Schedule Tracking
// @Success      200 {object} BaseResponse
// @Failure      400
// @Failure 	 500  {object} BaseResponse
// @Router       /schedule/deleteBillNumbers [delete]
func (h *HttpHandler) DeleteBillNumbersFromTrack(c *gin.Context) {
	var s DeleteFromTrackRequest
	if err := h.utils.Validate(c, &s); err != nil {
		h.utils.ValidateSchemaError(c, http.StatusBadRequest, "invalid input body")
		return
	}
	userId, err := h.utils.DecodeToken(c)
	if err != nil {
		c.JSON(401, gin.H{"message": "cannot decode token"})
		return
	}
	if err := h.client.DeleteFromTracking(c.Request.Context(), false, int64(userId), s); err != nil {
		c.JSON(500, gin.H{"success": false, "error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"success": true, "error": nil})
	return
}

// GetInfoAboutTracking
// @Summary      get info about number on tracking
// @Security ApiKeyAuth
// @Description  get info about number on tracking
// @accept json
// @Produce      json
// @Param input body GetInfoAboutTrackRequest true "info"
// @Tags         Schedule Tracking
// @Success      200 {object} GetInfoAboutTrackResponse
// @Failure      400
// @Failure 	 500  {object} BaseResponse
// @Router       /schedule/getInfo [get]
func (h *HttpHandler) GetInfoAboutTracking(c *gin.Context) {
	var s GetInfoAboutTrackRequest
	if err := h.utils.Validate(c, &s); err != nil {
		h.utils.ValidateSchemaError(c, http.StatusBadRequest, "invalid input body")
		return
	}
	res, err := h.client.GetInfoAboutTrack(c.Request.Context(), s)
	if err != nil {
		c.JSON(500, gin.H{"success": false, "error": err.Error()})
		return
	}
	c.JSON(200, res)
	return
}
