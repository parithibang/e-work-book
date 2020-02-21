package validations

import (
	"net/http"

	"github.com/aravindkumaremis/e-work-book/models"
	"github.com/astaxie/beego/validation"
)

// Projectvalidate validate the project
func Projectvalidate(req *http.Request) (validation.Validation, models.Projects) {
	project := models.Projects{
		Name:     req.FormValue("project-name"),
		IsActive: 1,
	}

	valid := validation.Validation{}
	valid.Valid(&project)

	return valid, project
}
