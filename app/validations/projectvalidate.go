package validations

import (
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/parithibang/e-work-book/app/models"
)

// Projectvalidate validate the project
func Projectvalidate(req *http.Request) (validation.Validation, models.Projects) {
	project := models.Projects{
		Name:     req.FormValue("project-name"),
		IsActive: 1,
	}

	valid := validation.Validation{}
	_, _ = valid.Valid(&project)

	return valid, project
}
