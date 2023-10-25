package get_persons

import (
	"context"
	"effective-mobile/internal/domain/models"

	"github.com/gin-gonic/gin"
)

type GetPersonsCtrl struct {
	getPersonSvc GetPersonsSvc
}

type GetPersonsSvc interface {
	GetPersonsByFilter(ctx context.Context) ([]models.Person, error)
}

func (gpc *GetPersonsCtrl) GetPersonsByFilter(c *gin.Context) {

}

// minAge maxAge, gender, nation, name, surname, patronymic
