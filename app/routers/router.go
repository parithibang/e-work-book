package routers

import (
	"github.com/parithibang/e-work-book/app/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/home", &controllers.MainController{}, "get,post:Get")
	beego.Router("/search", &controllers.SearchController{}, "get:GetSearch")
	beego.Router("/search/user/", &controllers.SearchController{}, "get:UserSearchResults")

	beego.Router("/users", &controllers.UserController{}, "get:ListUsers")
	beego.Router("/users/add", &controllers.UserController{}, "get:AddUser;post:CreateUser")
	beego.Router("/users/edit/:id:int", &controllers.UserController{}, "get:EditUser;put:UpdateUser")
	beego.Router("/users/delete/:id:int", &controllers.UserController{}, "delete:DeleteUser")

	beego.Router("/teams", &controllers.TeamController{}, "get:ListTeams")
	beego.Router("/teams/add", &controllers.TeamController{}, "get:AddTeam;post:CreateTeam")
	beego.Router("/teams/edit/:id:int", &controllers.TeamController{}, "get:EditTeam;put:UpdateTeam")
	beego.Router("/teams/delete/:id:int", &controllers.TeamController{}, "delete:DeleteTeam")

	beego.Router("/pod", &controllers.PodController{}, "get:ListPods")
	beego.Router("/pods/add", &controllers.PodController{}, "get:AddPod;post:CreatePod")
	beego.Router("/pods/edit/:id:int", &controllers.PodController{}, "get:EditPod;put:UpdatePod")
	beego.Router("/pods/delete/:id:int", &controllers.PodController{}, "delete:DeletePod")

	beego.Router("/project", &controllers.ProjectController{}, "get:ListProjects")
	beego.Router("/projects/add", &controllers.ProjectController{}, "get:AddProject;post:CreateProject")
	beego.Router("/projects/edit/:id:int", &controllers.ProjectController{}, "get:EditProject;put:UpdateProject")
	beego.Router("/projects/delete/:id:int", &controllers.ProjectController{}, "delete:DeleteProject")

	beego.Router("/add-users-project", &controllers.UsersProjectsController{}, "get:AddUserProjectDetail;post:CreateUserProjectDetail")
	beego.Router("/users-project/delete/:id:int", &controllers.UsersProjectsController{}, "delete:DeleteUserProject")
	beego.Router("/users-project/edit/:id:int", &controllers.UsersProjectsController{}, "get:EditUserProjects;put:UpdateUserProjects")

	beego.Router("/login", &controllers.MainController{}, "get,post:Login")
	beego.Router("/loginCheck", &controllers.MainController{}, "get,post:LoginCheck")
	beego.Router("/logout", &controllers.MainController{}, "get,post:Logout")
	beego.Router("/profile", &controllers.MainController{}, "get,post:Profile")

}
