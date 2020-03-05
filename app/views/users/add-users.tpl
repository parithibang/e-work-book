<div class="panel">
    <div class="panel-heading">
        <h3 class="panel-title">Add New User</h3>
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

        <form role="form" id="user" class="user-create" method="POST">
            <input type="hidden" name="_method" value="{{ .method }}"/>
            <div class="form-group col-sm-12 col-md-12">
                <div class="{{if .Errors.FirstName}}has-error{{end}} col-sm-6 col-md-6">
                    <label for="first-name">First name*</label>
                    <input name="first-name" id="first-name" type="text" value="{{.User.FirstName}}" class="form-control" tabindex="1"/>
                    {{ template "base/form-valid.tpl" .Errors.FirstName }}
                </div>
                <div class="{{if .Errors.LastName}}has-error{{end}} col-sm-6 col-md-6">
                    <label for="last-name">Last name</label>
                    <input name="last-name" id="last-name" type="text" value="{{.User.LastName}}" class="form-control" tabindex="1"/>
                    {{ template "base/form-valid.tpl" .Errors.LastName }}
                </div>
            </div>
            <div class="form-group col-xs-10 col-sm-12 col-md-12">
                <div class="{{if .Errors.UserName}}has-error{{end}} col-sm-6 col-md-6">
                    <label for="user-name">User name*
                    </label>
                    <input name="user-name" id="user-name" type="text" value="{{.User.UserName}}" class="form-control" tabindex="1"/>
                    {{ template "base/form-valid.tpl" .Errors.UserName }}
                </div>
                <div class="{{if .Errors.Password}}has-error{{end}} col-sm-6 col-md-6">
                    <label for="password">Password*
                    </label>
                    <input name="password" id="password" type="password" value="{{.User.Password}}" class="form-control" tabindex="1"/>
                    {{ template "base/form-valid.tpl" .Errors.Password }}
                </div>
            </div>
            <div class="form-group col-xs-10 col-sm-12 col-md-12">
                <div class="{{if .Errors.Pods}}has-error{{end}} col-sm-6 col-md-6">
                    <label for="pods">Pods
                    </label>
                    <select class="form-control selectpicker show-tick" id="pods" name="pods" data-live-search="true" title="Choose one of the following..." data-dropup-auto="false" noneSelectedText="noneSelectedText">
                        {{range $pod := .pods}}
                            <option value="{{ $pod.ID }}" {{if compare ($pod.ID) ($.User.Pods.ID) }} selected="selected" {{end}}>{{ $pod.Name }}</option>
                        {{ end }}
                    </select>
                    {{ template "base/form-valid.tpl" .Errors.Pods }}
                </div>
                <div class="col-sm-6 col-md-6">
                    <div>
                        <label>
                            Is Pod Lead</label>
                    </div>
                    <label class="fancy-radio visible-lg-inline-block">
                        <input name="is-pod-lead" value="1" type="radio" {{if compare (1) (.User.IsPodLead) }} checked="checked" {{end}}>
                            <span>
                                <i></i>Yes</span>
                        </label>
                        <label class="fancy-radio visible-lg-inline-block">
                            <input name="is-pod-lead" value="0" type="radio" {{if compare (0) (.User.IsPodLead) }} checked="checked" {{end}}>
                                <span>
                                    <i></i>No</span>
                            </label>
                        </div>
                    </div>
                    <div class="form-group col-xs-10 col-sm-12 col-md-12 text-center">
                        {{ if .create }}
                            <button type="submit" class="btn btn-success">Create</button>
                        {{ end }}
                        {{ if .update }}
                            <button type="submit" class="btn btn-success">Update</button>
                        {{ end }}
                        <a href="{{ urlfor "UserController.ListUsers"}}">
                            <button type="button" class="btn btn-danger">Cancel</button>
                        </a>
                    </div>
                </form>
            </div>
        </div>
