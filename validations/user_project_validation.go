package validations

import (
	"math"
	"net/http"
	"strconv"

	"github.com/parithibang/e-work-book/models"
	"github.com/astaxie/beego/validation"
)

// UserProjectvalidate validate the user create
func UserProjectvalidate(req *http.Request) (validation.Validation, models.UsersProjects) {
	userId, _ := strconv.Atoi(req.FormValue("user"))
	projectId, _ := strconv.Atoi(req.FormValue("project"))
	percentage, _ := strconv.ParseFloat(req.FormValue("work-percentage"), 64)
	percentageRoundOff := math.Floor(percentage*float64(100.00)) / float64(100.00)
	maxPercentage := 100.00

	selectedUser := models.Users{
		Id: userId,
	}

	selectedProject := models.Projects{
		Id: projectId,
	}

	userProjects := models.UsersProjects{
		Percentage: percentageRoundOff,
		Users:      &selectedUser,
		Projects:   &selectedProject,
		IsActive:   1,
	}

	valid := validation.Validation{}
	valid.Valid(&userProjects)

	if userId == 0 {
		valid.SetError("User", "User Should be selected")
	}
	if projectId == 0 {
		valid.SetError("Project", "Project Should be selected")
	}

	if percentageRoundOff > maxPercentage || percentageRoundOff == 0 {
		valid.SetError("Percentage", "Percentage is required and should be less than or equal to 100")
	}

	return valid, userProjects
}
