package validations

import (
	"net/http"
	"strconv"

	"github.com/astaxie/beego/validation"
	"github.com/parithibang/e-work-book/app/models"
)

// Teamvalidate validate the team create
func Teamvalidate(req *http.Request) (validation.Validation, models.Teams) {
	podID, _ := strconv.Atoi(req.FormValue("pods"))

	selectedPod := models.Pods{
		ID: podID,
	}

	team := models.Teams{
		Name:     req.FormValue("team-name"),
		Pods:     &selectedPod,
		IsActive: 1,
	}

	valid := validation.Validation{}
	valid.Valid(&team)

	if podID == 0 {
		valid.SetError("Pods", "Pod Should be selected")
	}

	return valid, team
}
