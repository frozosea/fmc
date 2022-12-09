package tracking

import (
	"fmc-gateway/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HttpHandler struct {
	client    IClient
	validator *Validator
	*utils.HttpUtils
}

func NewHttpHandler(client *Client, httpUtils *utils.HttpUtils) *HttpHandler {
	return &HttpHandler{client: client, HttpUtils: httpUtils}
}

// TrackByContainerNumber
// @Summary      Track by container number
// @Description  tracking by container number
// @accept 		 json
// @Produce      json
// @Param        scac  	 query     string   false  "scac code"       	default(SKLU) Enums(AUTO, FESO, SKLU,SITC,HALU,MAEU,MSCU,COSU,ONEY,KMTU) Pattern([a-zA-Z]{4})
// @Param        number  query     string   false  "container number"   	default(TEMU2094051) minlength(10)  maxlength(11) Pattern([a-zA-Z]{4}\d{6,7})
// @Tags 	     Tracking
// @Success      200  {object}  ContainerNumberResponse
// @Success      204
// @Failure      400
// @Router       /tracking/container [get]
func (h *HttpHandler) TrackByContainerNumber(c *gin.Context) {
	var schema Track
	if err := c.ShouldBindQuery(&schema); err != nil {
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
// @Description  tracking by bill number, if eta not found will be 0
// @Tags         Tracking
// @Param        scac  query     string     false  "scac code"       default(FESO) Enums(AUTO, FESO, SKLU,SITC,HALU, ZHGU)
// @Param        number  query     string   false  "bill number"   	default(FLCE405711) minlength(9)  maxlength(30)
// @Success      200  {object}  BillNumberResponse
// @Success      204
// @Failure      400
// @Router       /tracking/bill [get]
func (h *HttpHandler) TrackByBillNumber(c *gin.Context) {
	var schema *Track
	if err := c.ShouldBindQuery(&schema); err != nil {
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
	if err != nil {
		c.JSON(204, gin.H{})
		return
	}
	c.JSON(200, response)
	return
}

// GetContainerScac
// @Summary      get all scac codes for frontend
// @Description   get all scac codes for frontend
// @Tags         Tracking
// @Success      200  {object}  []Scac
// @Failure      500 {object}   BaseResponse
// @Router       /scac/containers [get]
func (h *HttpHandler) GetContainerScac(c *gin.Context) {
	response, err := h.client.GetContainerScac(c.Request.Context())
	if err != nil {
		c.JSON(500, gin.H{`success`: false, `error`: err.Error()})
		return
	}
	c.JSON(200, response)
}

// GetBillScac
// @Summary      get all scac codes for frontend
// @Description   get all scac codes for frontend
// @Tags         Tracking
// @Success      200  {object}  []Scac
// @Failure      500 {object}   BaseResponse
// @Router       /scac/bills [get]
func (h *HttpHandler) GetBillScac(c *gin.Context) {
	response, err := h.client.GetBillScac(c.Request.Context())
	if err != nil {
		c.JSON(500, gin.H{`success`: false, `error`: err.Error()})
		return
	}
	c.JSON(200, response)
}
