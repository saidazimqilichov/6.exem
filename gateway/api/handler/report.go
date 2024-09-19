package handler

import (
	"gateway/proto/report"
	"log"

	"github.com/gin-gonic/gin"
)


type ErrorResponseRE struct {
    Error string `json:"error"`
}

// GetReport godoc
// @Summary Get report
// @Description Retrieve the report based on the provided criteria
// @Tags reports
// @Produce json
// @Success 200 {object} report.ReportResponse "Report details"
// @Failure 500 {object} ErrorResponse "Unable to get report"
// @Router /reports [get]
func (h *Handler) GetReport(c *gin.Context) {
	var empty report.ReportEmpty
	res, err := h.Report.GetReport(c.Request.Context(), &empty)
	if err != nil {
		log.Println("Error getting report:", err)
		c.AbortWithStatusJSON(500, ErrorResponse{Error: "Unable to get report"})
		return
	}
	c.IndentedJSON(200, res)
}