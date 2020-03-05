package validations

import (
	"net/http"
	"strconv"

	"github.com/astaxie/beego/validation"
	"github.com/parithibang/e-work-book/app/models"
)

// Podvalidate validate the pod create
func Podvalidate(req *http.Request) (validation.Validation, models.Pods) {
	unitID, _ := strconv.Atoi(req.FormValue("pods"))

	selectedUnit := models.Units{
		ID: unitID,
	}

	pod := models.Pods{
		Name:     req.FormValue("pod-name"),
		Units:    &selectedUnit,
		IsActive: 1,
	}

	valid := validation.Validation{}
	valid.Valid(&pod)

	if unitID == 0 {
		valid.SetError("Units", "Unit Should be selected")
	}

	return valid, pod
}
