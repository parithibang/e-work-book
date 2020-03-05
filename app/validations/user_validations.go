package validations

import (
	"net/http"
	"strconv"

	"github.com/astaxie/beego/validation"
	"github.com/parithibang/e-work-book/app/models"
)

// Uservalidate validate the user create
func Uservalidate(req *http.Request) (validation.Validation, models.Users) {
	podID, _ := strconv.Atoi(req.FormValue("pods"))
	isPodLead, _ := strconv.Atoi(req.FormValue("is-pod-lead"))

	selectedPod := models.Pods{
		ID: podID,
	}

	user := models.Users{
		FirstName: req.FormValue("first-name"),
		LastName:  req.FormValue("last-name"),
		Password:  req.FormValue("password"),
		UserName:  req.FormValue("user-name"),
		Pods:      &selectedPod,
		IsPodLead: isPodLead,
		IsActive:  1,
	}

	valid := validation.Validation{}
	valid.Valid(&user)

	if podID == 0 {
		valid.SetError("Pods", "Pod Should be selected")
	}

	return valid, user
}
