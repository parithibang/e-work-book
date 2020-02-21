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
        {{if .flash.error}}
            <div class="alert alert-danger alert-dismissible" role="alert">
                <button type="button" class="close" data-dismiss="alert" aria-label="Close">
                    <span aria-hidden="true">×</span>
                </button>
                {{.flash.error}}
            </div>
        {{end}}
        <div class="col-md-6">
            <form role="form" id="user" class="user-search" method="POST">
                <div class="form-group">
                    <div id="custom-search-input">
                        <div class="input-group">
                            <input type="text" name="user-name" class="form-control input-lg" placeholder="Tom Johnson" value="{{if .userName}} {{ .userName}} {{ end }}"/>
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
        {{if .userList}}
            <table class="table">
                <thead>
                    <tr>
                        <th>#</th>
                        <th>Name</th>
                        <th>Pod</th>
                        <th>Team</th>
                        <th>Project</th>
                        <th>% of work</th>
                    </tr>
                </thead>
                <tbody>
                    {{range $key,$user := .userList }}
                        <tr>
                            <td>{{ calculate $key 1 "+" }}</td>
                            <td>{{ $user.Users.FullName }}</td>
                            <td>{{ $user.Users.Pods.Name }}</td>
                            {{if $user.Users.Teams}}
                                <td>{{ $user.Users.Teams.Name }}</td>
                            {{ else }}
                                <td></td>
                            {{end}}
                            <td>{{ $user.Projects.Name }}</td>
                            <td>{{ $user.Percentage }}</td>
                        </tr>
                    {{end}}
                </tbody>
            </table>
        {{end}}
    </div>
</div>
