<div class="panel">
    <div class="panel-heading">
        <h3 class="panel-title">Search User</h3>
        <a href="{{ urlfor "UsersProjectsController.AddUserProjectDetail"}}">
            <span class="pull-right button-add">Add
                <i class="fa fa-plus-circle"></i>
            </span>
        </a>
    </div>
    <div class="panel-body">
        {{if .flash.custom_error}}
            <div class="alert alert-danger alert-dismissible" role="alert">
                <button type="button" class="close" data-dismiss="alert" aria-label="Close">
                    <span aria-hidden="true">×</span>
                </button>
                {{.flash.custom_error}}
            </div>
        {{end}}

        {{if .flash.custom_success}}
            <div class="alert alert-success alert-dismissible" role="alert">
                <button type="button" class="close" data-dismiss="alert" aria-label="Close">
                    <span aria-hidden="true">×</span>
                </button>
                <i class="fa fa-check-circle"></i>
                {{.flash.custom_success}}
            </div>
        {{end}}
        <div class="col-md-6">
            <form action={{ urlfor "SearchController.UserSearchResults"}} role="form" id="user" class="user-search" method="GET">
                <div class="form-group">
                    <div id="custom-search-input">
                        <div class="input-group">
                            <input type="text" name="user-name" class="form-control input-lg" placeholder="Tom Johnson" value="{{if .userName}}{{.userName}}{{ end }}"/>
                            <span class="input-group-btn">
                                <button class="btn btn-info btn-lg" type="submit">
                                    <i class="glyphicon glyphicon-search"></i>
                                </button>
                            </span>
                        </div>
                    </div>
                </div>
            </form>
        </div>
        {{if .userProjectList}}
            <table class="table">
                <thead>
                    <tr>
                        <th>#</th>
                        <th>Name</th>
                        <th>Pod</th>
                        <th>Team</th>
                        <th>Project</th>
                        <th>% of work</th>
                        <th>Actions</th>
                    </tr>
                </thead>
                <tbody>
                    {{range $key,$userproject := .userProjectList }}
                        <tr>
                            <td>{{ calculate $.pageStart $key "+" }}</td>
                            <td>{{ $userproject.Users.FullName }}</td>
                            <td>{{ $userproject.Users.Pods.Name }}</td>
                            {{if $userproject.Users.Teams}}
                                <td>{{ $userproject.Users.Teams.Name }}</td>
                            {{ else }}
                                <td></td>
                            {{end}}
                            <td>{{ $userproject.Projects.Name }}</td>
                            <td>{{ $userproject.Percentage }}</td>
                            <td>
                                <a href="{{ urlfor "UsersProjectsController.EditUserProjects" ":id" $userproject.Id}}">
                                    <span class="button-edit">Edit
                                        <i class="fa fa-edit"></i>
                                    </span>
                                </a>
                                <a href="#">
                                    <span class="button-delete" data-toggle="modal" data-target="#myModal{{ $userproject.Id}}">Delete
                                        <i class="fa fa-trash"></i>
                                    </span>
                                </a>
                                {{ template "base/delete-record.tpl" (dynamicMap  "RecordId" $userproject.Id "method" $.deleteMethod "deleteUrl" (urlfor "UsersProjectsController.DeleteUserProject" ":id" $userproject.Id) "requestParams"  $.requestParams) }}
                            </td>
                        </tr>
                    {{end}}
                </tbody>
            </table>
            <span class="badge text-uppercase">
                Total Record :
                {{ .count }}</span>
            {{template "base/paginator.tpl" .}}
        {{end}}
    </div>
    {{ define "js" }}
        <script src="/static/js/search.js"></script>
    {{ end}}
</div>
