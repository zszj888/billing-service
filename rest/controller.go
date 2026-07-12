package rest

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/samz/billing/domain"
)

type BillService interface {
	GetByIdSvc(c context.Context, id int64) domain.BillEntity
}
type BillHandler struct {
	svc BillService
}

func RegisterHandlers(r *gin.Engine, svc BillService) {
	h := &BillHandler{svc: svc}
	r.GET("/bill/:id", h.getById)
}

type Uri struct {
	ID int64 `uri:"id" binding:"required"`
}

func (h *BillHandler) getById(c *gin.Context) {
	var uri Uri
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, h.svc.GetByIdSvc(c, uri.ID))
}
