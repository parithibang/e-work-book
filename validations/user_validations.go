package validations

import (
	"net/http"
	"strconv"

	"github.com/parithibang/e-work-book/models"
	"github.com/astaxie/beego/validation"
)

// Uservalidate validate the user create
func Uservalidate(req *http.Request) (validation.Validation, models.Users) {
	podId, _ := strconv.Atoi(req.FormValue("pods"))
	isPodLead, _ := strconv.Atoi(req.FormValue("is-pod-lead"))

	selectedPod := models.Pods{
		Id: podId,
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

	if podId == 0 {
		valid.SetError("Pods", "Pod Should be selected")
	}

	return valid, user
}
