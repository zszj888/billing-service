package rest

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/samz/billing/domain"
)

type BillService interface {
	GetByIdSvc(c context.Context, id int64) (domain.BillEntity, error)
	Save(c context.Context, bill *domain.BillEntity) error
}
type BillHandler struct {
	svc BillService
}

func RegisterHandlers(r *gin.Engine, svc BillService) {
	h := &BillHandler{svc: svc}
	r.GET("/bills/:id", h.getById)
	r.POST("/bills", h.save)
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
	res, err := h.svc.GetByIdSvc(c, uri.ID)
	switch {
	case errors.Is(err, domain.ErrNotFound):
		c.JSON(http.StatusNotFound, gin.H{"error": domain.ErrNotFound.Error()})
	case err != nil:
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	default:
		c.JSON(http.StatusOK, res)
	}

}

func (h *BillHandler) save(c *gin.Context) {
	entity := domain.BillEntity{}
	if err := c.ShouldBindJSON(&entity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.svc.Save(c, &entity); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, entity)
}
