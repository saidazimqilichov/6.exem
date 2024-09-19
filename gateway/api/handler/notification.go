package handler

import (
	"gateway/proto/notification"
	"log"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
    Error string `json:"error"`
}



// GetNotification godoc
// @Summary Get all notifications
// @Description Retrieve a list of all notifications
// @Tags notifications
// @Produce json
// @Success 200 {array} notification.NotifyList "List of notifications"
// @Failure 500 {object} ErrorResponse "Unable to get notification"
// @Router /notifications [get]
func (h *Handler) GetNotification(c *gin.Context) {
	notifyList, err := h.Notification.GetNotfication(c.Request.Context(), &notification.NotifyEmpty{})
	if err != nil {
		log.Println("Error get notification")
		c.AbortWithStatusJSON(500, ErrorResponse{Error: "Unable to get notification"})
		return
	}
	c.IndentedJSON(200, notifyList)
}

// GetUnreadNotifications godoc
// @Summary Get unread notifications
// @Description Retrieve a list of unread notifications
// @Tags notifications
// @Produce json
// @Success 200 {array} notification.NotifyList "List of unread notifications"
// @Failure 500 {object} ErrorResponse "Unable to get notification"
// @Router /notifications/unread [get]
func (h *Handler) GetUnreadNotifications(c *gin.Context) {
	notifyList, err := h.Notification.GetUnreadNotfications(c.Request.Context(), &notification.NotifyEmpty{})
	if err != nil {
		log.Println("Error get notification")
		c.AbortWithStatusJSON(500, ErrorResponse{Error: "Unable to get notification"})
		return
	}
	c.IndentedJSON(200, notifyList)
}
