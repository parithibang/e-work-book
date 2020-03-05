package validations

import (
	"math"
	"net/http"
	"strconv"

	"github.com/astaxie/beego/validation"
	"github.com/parithibang/e-work-book/app/models"
)

// UserProjectvalidate validate the user create
func UserProjectvalidate(req *http.Request) (validation.Validation, models.UsersProjects) {
	userID, _ := strconv.Atoi(req.FormValue("user"))
	projectID, _ := strconv.Atoi(req.FormValue("project"))
	percentage, _ := strconv.ParseFloat(req.FormValue("work-percentage"), 64)
	percentageRoundOff := math.Floor(percentage*float64(100.00)) / float64(100.00)
	maxPercentage := 100.00

	selectedUser := models.Users{
		ID: userID,
	}

	selectedProject := models.Projects{
		ID: projectID,
	}

	userProjects := models.UsersProjects{
		Percentage: percentageRoundOff,
		Users:      &selectedUser,
		Projects:   &selectedProject,
		IsActive:   1,
	}

	valid := validation.Validation{}
	_, _ = valid.Valid(&userProjects)

	if userID == 0 {
		valid.SetError("User", "User Should be selected")
	}
	if projectID == 0 {
		valid.SetError("Project", "Project Should be selected")
	}

	if percentageRoundOff > maxPercentage || percentageRoundOff == 0 {
		valid.SetError("Percentage", "Percentage is required and should be less than or equal to 100")
	}

	return valid, userProjects
}
