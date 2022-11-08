package schedule_tracking

import (
	"fmc-gateway/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HttpHandler struct {
	client    IClient
	validator *Validator
	utils     *utils.HttpUtils
}

func NewHttpHandler(client IClient, utils *utils.HttpUtils) *HttpHandler {
	return &HttpHandler{client: client, utils: utils}
}

// AddContainersOnTrack
// @Summary      add containers on track
// @Security ApiKeyAuth
// @Description  add containers on track. Every day in your selected time track container and send email with result about it. You can add on track only if container/bill already in your account.
// @accept json
// @Produce      json
// @Param input body AddOnTrackRequest true "info"
// @Tags         Schedule Tracking
// @Success      200  {object}  AddOnTrackResponse
// @Failure      400
// @Failure      500  {object}  BaseResponse
// @Router       /schedule/containers [post]
func (h *HttpHandler) AddContainersOnTrack(c *gin.Context) {
	var s *AddOnTrackRequest
	if err := c.ShouldBindJSON(&s); err != nil {
		h.utils.ValidateSchemaError(c, http.StatusBadRequest, "invalid input body")
		return
	}
	for _, b := range s.Numbers {
		if err := h.validator.ValidateContainer(b); err != nil {
			h.utils.ValidateSchemaError(c, http.StatusBadRequest, "invalid input body")
			return
		}
	}
	if err := h.validator.ValidateEmails(s.Emails); err != nil {
		h.utils.ValidateSchemaError(c, http.StatusBadRequest, "invalid input body")
		return
	}
	if err := h.validator.ValidateTime(s.Time); err != nil {
		h.utils.ValidateSchemaError(c, http.StatusBadRequest, "invalid input body")
		return
	}
	userId, err := h.utils.DecodeToken(c)
	if err != nil {
		c.JSON(401, gin.H{"message": "cannot decode token"})
		return
	}
	resp, err := h.client.AddContainersOnTrack(c.Request.Context(), int(userId), s)
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
// @Description  add bill numbers on track. Every day in your selected time track bill numbers and send email with result about it. You can add on track only if container/bill already in your account.
// @accept json
// @Produce      json
// @Param input body AddOnTrackRequest true "info"
// @Tags         Schedule Tracking
// @Success      200  {object}  AddOnTrackResponse
// @Failure      400
// @Failure      500  {object}  BaseResponse
// @Router       /schedule/bills [post]
func (h *HttpHandler) AddBillNumbersOnTrack(c *gin.Context) {
	var s *AddOnTrackRequest
	if err := c.ShouldBindJSON(&s); err != nil {
		h.utils.ValidateSchemaError(c, http.StatusBadRequest, "invalid input body")
		return
	}
	for _, b := range s.Numbers {
		if err := h.validator.ValidateBill(b); err != nil {
			h.utils.ValidateSchemaError(c, http.StatusBadRequest, "invalid input body")
			return
		}
	}
	if err := h.validator.ValidateEmails(s.Emails); err != nil {
		h.utils.ValidateSchemaError(c, http.StatusBadRequest, "invalid input body")
		return
	}
	if err := h.validator.ValidateTime(s.Time); err != nil {
		h.utils.ValidateSchemaError(c, http.StatusBadRequest, "invalid input body")
		return
	}
	userId, err := h.utils.DecodeToken(c)
	if err != nil {
		c.JSON(401, gin.H{"message": "cannot decode token"})
		return
	}
	resp, err := h.client.AddBillNosOnTrack(c.Request.Context(), int(userId), s)
	if err != nil {
		c.JSON(500, gin.H{"success": false, "error": err.Error()})
		return
	}
	c.JSON(200, resp)
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
// @Router       /schedule/containers [delete]
func (h *HttpHandler) DeleteContainersFromTrack(c *gin.Context) {
	var s *DeleteFromTrackRequest
	userId, err := h.utils.DecodeToken(c)
	if err != nil {
		c.AbortWithStatus(401)
		return
	}
	s.userId = int64(userId)
	if err := h.utils.Validate(c, &s); err != nil {
		h.utils.ValidateSchemaError(c, http.StatusBadRequest, "invalid input body")
		return
	}
	if err := h.validator.ValidateContainers(s.Numbers); err != nil {
		h.utils.ValidateSchemaError(c, http.StatusBadRequest, "invalid input body")
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
// @Router       /schedule/billNumbers [delete]
func (h *HttpHandler) DeleteBillNumbersFromTrack(c *gin.Context) {
	var s *DeleteFromTrackRequest
	userId, err := h.utils.DecodeToken(c)
	if err != nil {
		c.AbortWithStatus(401)
		return
	}
	s.userId = int64(userId)
	if err := h.utils.Validate(c, &s); err != nil {
		h.utils.ValidateSchemaError(c, http.StatusBadRequest, "invalid input body")
		return
	}
	if err := h.validator.ValidateBills(s.Numbers); err != nil {
		h.utils.ValidateSchemaError(c, http.StatusBadRequest, "invalid input body")
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
// @Param 		 number 	query 	string false "bill number or container number"
// @Tags         Schedule Tracking
// @Success      200 {object} GetInfoAboutTrackResponse
// @Failure      400
// @Failure 	 500  {object} BaseResponse
// @Router       /schedule/info [get]
func (h *HttpHandler) GetInfoAboutTracking(c *gin.Context) {
	var s GetInfoAboutTrackRequest
	userId, err := h.utils.DecodeToken(c)
	if err != nil {
		c.AbortWithStatus(401)
		return
	}
	if err := c.ShouldBindQuery(&s); err != nil {
		h.utils.ValidateSchemaError(c, http.StatusBadRequest, "invalid input body")
		return
	}
	s.userId = int64(userId)
	if err := h.validator.ValidateBill(s.Number); err != nil {
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

//GetTimeZone
// @Summary      get timezone information
// @Tags         Schedule Tracking
// @Description  get timezone in format UTC+10, this route is for get time zone, because users want to know in which tz will tracking works
// @Success      200 {object} TimeZoneResponse
// @Failure 	 500  {object} BaseResponse
// @Router       /schedule/timezone [get]
func (h *HttpHandler) GetTimeZone(c *gin.Context) {
	timeZone, err := h.client.GetTimeZone(c.Request.Context())
	if err != nil {
		c.JSON(500, gin.H{"success": false, "error": err.Error()})
		return
	}
	c.JSON(200, timeZone)
	return
}

//UpdateContainers
// @Summary      update container tracking task
// @Security ApiKeyAuth
// @Tags         Schedule Tracking
// @Description  update tracking tasks by input params
// @accept json
// @Produce      json
// @Param input body AddOnTrackRequest true "info"
// @Success      200 {object} BaseResponse
// @Failure 	 500  {object} BaseResponse
// @Router       /schedule/containers [put]
func (h *HttpHandler) UpdateContainers(c *gin.Context) {
	var s *AddOnTrackRequest
	if err := c.ShouldBindJSON(&s); err != nil {
		h.utils.ValidateSchemaError(c, http.StatusBadRequest, "invalid input body")
		return
	}
	for _, b := range s.Numbers {
		if err := h.validator.ValidateBill(b); err != nil {
			h.utils.ValidateSchemaError(c, http.StatusBadRequest, "invalid input body")
			return
		}
	}
	if err := h.validator.ValidateEmails(s.Emails); err != nil {
		h.utils.ValidateSchemaError(c, http.StatusBadRequest, "invalid input body")
		return
	}
	if err := h.validator.ValidateTime(s.Time); err != nil {
		h.utils.ValidateSchemaError(c, http.StatusBadRequest, "invalid input body")
		return
	}
	userId, err := h.utils.DecodeToken(c)
	if err != nil {
		c.JSON(401, gin.H{"message": "cannot decode token"})
		return
	}
	if err := h.client.Update(c.Request.Context(), int(userId), true, s); err != nil {
		c.JSON(500, gin.H{"success": false, "error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"success": true, "error": nil})
	return
}

//UpdateBills
// @Summary      update container tracking task
// @Security ApiKeyAuth
// @Tags         Schedule Tracking
// @Description  update tracking tasks by input params
// @accept json
// @Produce      json
// @Param input body AddOnTrackRequest true "info"
// @Success      200 {object} BaseResponse
// @Failure 	 500  {object} BaseResponse
// @Router       /schedule/bills [put]
func (h *HttpHandler) UpdateBills(c *gin.Context) {
	var s *AddOnTrackRequest
	if err := c.ShouldBindJSON(&s); err != nil {
		h.utils.ValidateSchemaError(c, http.StatusBadRequest, "invalid input body")
		return
	}
	for _, b := range s.Numbers {
		if err := h.validator.ValidateBill(b); err != nil {
			h.utils.ValidateSchemaError(c, http.StatusBadRequest, "invalid input body")
			return
		}
	}
	if err := h.validator.ValidateEmails(s.Emails); err != nil {
		h.utils.ValidateSchemaError(c, http.StatusBadRequest, "invalid input body")
		return
	}
	if err := h.validator.ValidateTime(s.Time); err != nil {
		h.utils.ValidateSchemaError(c, http.StatusBadRequest, "invalid input body")
		return
	}
	userId, err := h.utils.DecodeToken(c)
	if err != nil {
		c.JSON(401, gin.H{"message": "cannot decode token"})
		return
	}
	if err := h.client.Update(c.Request.Context(), int(userId), false, s); err != nil {
		c.JSON(500, gin.H{"success": false, "error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"success": true, "error": nil})
	return
}
