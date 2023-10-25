package delete_person

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeletePersonCtrl struct {
	deletePersonSvc DeletePersonSvc
}

type DeletePersonSvc interface {
	DeletePerson(ctx context.Context, personId int) error
}

func NewDeletePersonCtrl(dpSvc DeletePersonSvc) *DeletePersonCtrl {
	return &DeletePersonCtrl{
		deletePersonSvc: dpSvc,
	}
}

const deleteParamKey = "id"

func (dpc *DeletePersonCtrl) DeletePerson(c *gin.Context) {
	personId, err := strconv.Atoi(c.Param(deleteParamKey))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = dpc.deletePersonSvc.DeletePerson(c, personId)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": "successfully deleted"})
}
