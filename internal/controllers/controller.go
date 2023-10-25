package controllers

import (
	"effective-mobile/internal/controllers/add_person"
	"effective-mobile/internal/controllers/delete_person"
	"effective-mobile/internal/controllers/update_person"
)

type Controller struct {
	AddPersonCtrl    *add_person.AddPersonCtrl
	DeletePersonCtrl *delete_person.DeletePersonCtrl
	UpdatePersonCtrl *update_person.UpdatePersonCtrl
}

func NewController(apSvc add_person.AddPersonSvc, dpSvc delete_person.DeletePersonSvc, upSvc update_person.UpdatePersonSvc, ioSvc add_person.IoSvc) *Controller {
	return &Controller{
		AddPersonCtrl:    add_person.NewAddPersonCtrl(apSvc, ioSvc),
		DeletePersonCtrl: delete_person.NewDeletePersonCtrl(dpSvc),
		UpdatePersonCtrl: update_person.NewUpdatePersonCtrl(upSvc),
	}
}
