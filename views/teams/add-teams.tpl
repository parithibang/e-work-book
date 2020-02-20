<div class="panel">
    <div class="panel-heading">
        <h3 class="panel-title">Add New Team</h3>
    </div>
    <div class="panel-body">
        {{if .flash.error}}
            <div class="alert alert-danger alert-dismissible" role="alert">
                <button type="button" class="close" data-dismiss="alert" aria-label="Close"><span aria-hidden="true">×</span></button>
                {{.flash.error}}
            </div>
        {{end}}
        {{if .flash.success}}
            <div class="alert alert-success alert-dismissible" role="alert">
                <button type="button" class="close" data-dismiss="alert" aria-label="Close"><span aria-hidden="true">×</span></button>
                <i class="fa fa-check-circle"></i> {{.flash.success}}
            </div>
        {{end}}

        <form role="form" id="user" class="user-create" method="POST">
            <input type="hidden" name="_method" value="{{ .method }}" />
            <div class="form-group col-sm-12 col-md-12">
                <div class="{{if .Errors.Name}}has-error{{end}}  col-sm-6 col-md-6">
                    <label for="team-name">Team name*</label>
                    <input name="team-name" id="team-name" type="text" value="{{.Team.Name}}" class="form-control" tabindex="1" />
                    {{ template "base/form-valid.tpl" .Errors.Name }}
                </div>    
                <div class="{{if .Errors.Pods}}has-error{{end}} col-sm-6 col-md-6">
                    <label for="pods">Pods </label>
                    <select class="form-control selectpicker show-tick" id="pods" name="pods" data-live-search="true" title="Choose one of the following..." data-dropup-auto="false" noneSelectedText>
                        {{range $pod := .pods}}
                            <option value="{{ $pod.Id }}"  {{if compare ($pod.Id) ($.Team.Pods.Id) }} selected {{end}} >{{ $pod.Name }}</option>
                         {{ end }}
					</select>
                    {{ template "base/form-valid.tpl" .Errors.Pods }}
                </div>     
            </div>
            <div class="form-group col-xs-10 col-sm-12 col-md-12 text-center">
                {{ if .create }}
                <button type="submit" class="btn btn-success">Create</button>
                {{ end }}
                {{ if .update }}
                <button type="submit" class="btn btn-success">Update</button>
                {{ end }}
                <a href="{{ urlfor "TeamController.ListTeams"}}">
                    <button type="button" class="btn btn-danger">Cancel</button>
                </a>
            </div>
        </form>
    </div>
</div>