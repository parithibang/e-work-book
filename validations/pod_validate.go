package validations

import (
	"net/http"
	"strconv"

	"github.com/aravindkumaremis/e-work-book/models"
	"github.com/astaxie/beego/validation"
)

// Podvalidate validate the pod create
func Podvalidate(req *http.Request) (validation.Validation, models.Pods) {
	unitId, _ := strconv.Atoi(req.FormValue("pods"))

	selectedUnit := models.Units{
		Id: unitId,
	}

	pod := models.Pods{
		Name:     req.FormValue("pod-name"),
		Units:    &selectedUnit,
		IsActive: 1,
	}

	valid := validation.Validation{}
	valid.Valid(&pod)

	if unitId == 0 {
		valid.SetError("Units", "Unit Should be selected")
	}

	return valid, pod
}
