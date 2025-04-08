package assessment

import (
	"net/http"

	"github.com/gin-gonic/gin"
	types "github.com/teamcubation/teamcandidates/pkg/types"
	dto "github.com/teamcubation/teamcandidates/projects/teamcandidates-api/internal/assessment/handler/dto"
)

func (h *Handler) GenerateLink(c *gin.Context) {
	assessmentID := c.Param("id")

	linkId, err := h.ucs.GenerateLink(c.Request.Context(), assessmentID)
	if err != nil {
		apiErr, errCode := types.NewAPIError(err)
		c.Error(apiErr).SetMeta(errCode)
		return
	}

	c.JSON(http.StatusCreated, dto.GenerateLinkResponse{
		Message: "Assessment link successfully generated",
		LinkID:  linkId,
	})
}

func (h *Handler) SendLink(c *gin.Context) {
	assessmentLinkID := c.Param("id")
	if err := h.ucs.SendLink(c.Request.Context(), assessmentLinkID); err != nil {
		apiErr, errCode := types.NewAPIError(err)
		c.Error(apiErr).SetMeta(errCode)
		return
	}

	c.JSON(http.StatusCreated, dto.SendLinkResponse{
		Message: "Unique link successfully sent",
		Link:    assessmentLinkID,
	})
}
