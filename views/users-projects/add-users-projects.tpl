<div class="panel">
    <div class="panel-heading">
        <h3 class="panel-title">Add User Project</h3>
    </div>
    <div class="panel-body">
        {{if .flash.error}}
            <div class="alert alert-danger alert-dismissible" role="alert">
                <button type="button" class="close" data-dismiss="alert" aria-label="Close">
                    <span aria-hidden="true">×</span>
                </button>
                {{.flash.error}}
            </div>
        {{end}}
        {{if .flash.success}}
            <div class="alert alert-success alert-dismissible" role="alert">
                <button type="button" class="close" data-dismiss="alert" aria-label="Close">
                    <span aria-hidden="true">×</span>
                </button>
                <i class="fa fa-check-circle"></i>
                {{.flash.success}}
            </div>
        {{end}}
        {{if .flash.warning}}
            <div class="alert alert-warning alert-dismissible" role="alert">
                <button type="button" class="close" data-dismiss="alert" aria-label="Close">
                    <span aria-hidden="true">×</span>
                </button>
                <i class="fa fa-warning"></i>
                {{.flash.warning}}
            </div>
        {{end}}

        <form role="form" id="user" class="user-create" method="POST">
            <input type="hidden" name="_method" value="{{ .method }}"/>
            <div class="form-group col-sm-12 col-md-12">
                <div class="{{if .Errors.User}}has-error{{end}} col-sm-12 col-md-12">
                    <label for="user">Users*</label>
                    <select class="form-control selectpicker show-tick" id="user" name="user" data-live-search="true" title="Choose one of the following..." data-dropup-auto="false" noneselectedtext="noneSelectedText">
                        {{range $user := .users}}
                            <option value="{{ $user.Id }}" {{if compare ($user.Id) ($.UserProjects.Users.Id) }} selected="selected" {{end}}>{{ $user.FullName }}</option>
                        {{ end }}
                    </select>
                    {{ template "base/form-valid.tpl" .Errors.User }}
                </div>
            </div>
            <div class="form-group col-sm-12 col-md-12">
                <div class="{{if .Errors.Project}}has-error{{end}} col-sm-12 col-md-12">
                    <label for="project">Projects*</label>
                    <select class="form-control selectpicker show-tick" id="project" name="project" data-live-search="true" title="Choose one of the following..." data-dropup-auto="false" noneselectedtext="noneSelectedText">
                        {{range $project := .projects}}
                            <option value="{{ $project.Id }}" {{if compare ($project.Id) ($.UserProjects.Projects.Id) }} selected="selected" {{end}}>{{ $project.Name }}</option>
                        {{ end }}
                    </select>
                    {{ template "base/form-valid.tpl" .Errors.Project }}
                </div>
            </div>
            <div class="form-group col-sm-12 col-md-12">
                <div class="{{if .Errors.Percentage}}has-error{{end}} col-sm-12 col-md-12">
                    <label for="work-percentage">Work Percentage*</label>
                    <input name="work-percentage" id="work-percentage" type="number" step="0.01" value="{{if .UserProjects.Percentage}}{{ .UserProjects.Percentage}}{{ else }}100.00{{end}}" class="form-control"/>
                    {{ template "base/form-valid.tpl" .Errors.Percentage }}
                </div>
            </div>

            <div class="form-group col-xs-10 col-sm-12 col-md-12 text-center">
                {{ if .create }}
                    <button type="submit" class="btn btn-success">Create</button>
                {{ end }}
                {{ if .update }}
                    <button type="submit" class="btn btn-success">Update</button>
                {{ end }}
                <a href="{{ urlfor "SearchController.GetSearch"}}">
                    <button type="button" class="btn btn-danger">Cancel</button>
                </a>
            </div>
        </form>
    </div>
</div>
