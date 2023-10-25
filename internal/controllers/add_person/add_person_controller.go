package add_person

import (
	"context"
	"effective-mobile/internal/domain/dto"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AddPersonCtrl struct {
	addPersonSvc AddPersonSvc
	ioSvc        IoSvc
}

type AddPersonSvc interface {
	AddPerson(ctx context.Context, req dto.AddPersonRequest) (int, error)
}

type IoSvc interface {
	FillAddPersonRequest(ctx context.Context, req dto.AddPersonRequest) (dto.AddPersonRequest, error)
}

func NewAddPersonCtrl(apSvc AddPersonSvc, ioSvc IoSvc) *AddPersonCtrl {
	return &AddPersonCtrl{
		addPersonSvc: apSvc,
		ioSvc:        ioSvc,
	}
}

func (apc *AddPersonCtrl) AddPerson(c *gin.Context) {
	var req dto.AddPersonRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	req, err := apc.ioSvc.FillAddPersonRequest(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(req)

	personId, err := apc.addPersonSvc.AddPerson(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, personId)

}
