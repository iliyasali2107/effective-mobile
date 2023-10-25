package update_person

import (
	"context"
	"effective-mobile/internal/domain/dto"
	"effective-mobile/internal/services"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UpdatePersonCtrl struct {
	updatePersonSvc UpdatePersonSvc
}

type UpdatePersonSvc interface {
	UpdatePerson(ctx context.Context, req dto.UpdatePersonRequest) error
}

func NewUpdatePersonCtrl(upSvc UpdatePersonSvc) *UpdatePersonCtrl {
	return &UpdatePersonCtrl{
		updatePersonSvc: upSvc,
	}
}

func (upc *UpdatePersonCtrl) UpdatePerson(c *gin.Context) {
	var req dto.UpdatePersonRequest
	err := c.BindJSON(&req)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = upc.updatePersonSvc.UpdatePerson(c, req)
	if err != nil {
		if errors.Is(err, services.ErrNotFound) {
			c.AbortWithError(http.StatusNotFound, err)
			return
		}
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": "updated successfully"})
}
