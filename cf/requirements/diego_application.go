package requirements

import (
	"github.com/liamawhite/cli-with-i18n/cf/api/applications"
	"github.com/liamawhite/cli-with-i18n/cf/errors"
	"github.com/liamawhite/cli-with-i18n/cf/models"
	"github.com/liamawhite/cli-with-i18n/cf/terminal"
)

//go:generate counterfeiter . DiegoApplicationRequirement

type DiegoApplicationRequirement interface {
	Requirement
	GetApplication() models.Application
}

type diegoApplicationRequirement struct {
	appName string
	ui      terminal.UI
	appRepo applications.Repository

	application models.Application
}

func NewDiegoApplicationRequirement(name string, applicationRepo applications.Repository) DiegoApplicationRequirement {
	return &diegoApplicationRequirement{
		appName: name,
		appRepo: applicationRepo,
	}
}

func (req *diegoApplicationRequirement) Execute() error {
	app, err := req.appRepo.Read(req.appName)
	if err != nil {
		return err
	}

	if app.Diego == false {
		return errors.New("The app is running on the DEA backend, which does not support this command.")
	}

	req.application = app

	return nil
}

func (req *diegoApplicationRequirement) GetApplication() models.Application {
	return req.application
}
