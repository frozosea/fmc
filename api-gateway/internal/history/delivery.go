package history

import (
	"fmc-gateway/pkg/utils"
	"github.com/gin-gonic/gin"
)

type HttpHandler struct {
	tasksClient IHistoryTasksClient
	utils       *utils.HttpUtils
}

func NewHttpHandler(tasksClient IHistoryTasksClient, utils *utils.HttpUtils) *HttpHandler {
	return &HttpHandler{tasksClient: tasksClient, utils: utils}
}

// GetTasksArchive
// @Summary      get user's tasks archive with all info about containers and bills
// @Security ApiKeyAuth
// @Description   get user's tasks archive with all info about containers and bills
// @accept json
// @Produce      json
// @Tags         History
// @Success      200  {object}  TasksArchive
// @Failure      400
// @Failure      500  {object}  BaseResponse
// @Router       /tasks [get]
func (h *HttpHandler) GetTasksArchive(c *gin.Context) {
	userId, err := h.utils.DecodeToken(c)
	if err != nil {
		c.AbortWithStatus(401)
		return
	}
	result, err := h.tasksClient.Get(c.Request.Context(), int(userId))
	if err != nil {
		c.JSON(500, gin.H{"success": false, "error": err.Error()})
		return
	}
	c.JSON(200, result)
}
