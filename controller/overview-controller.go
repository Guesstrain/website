package controller

import (
	"github.com/Guesstrain/website/database"
	"github.com/Guesstrain/website/dto"
)

type OverviewController interface {
	FindAll() dto.PersonalInfo
}

type controller struct {
	service database.DatabaseService
}

func New(service database.DatabaseService) OverviewController {
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() dto.PersonalInfo {
	overview := c.service.All()
	return overview
}
