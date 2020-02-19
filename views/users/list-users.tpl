<div class="panel panel-headline">
    <div class="panel-heading">
        {{if .flash.custom_error}}
        <div class="alert alert-danger alert-dismissible" role="alert">
            <button type="button" class="close" data-dismiss="alert" aria-label="Close"><span aria-hidden="true">×</span></button>
            {{.flash.custom_error}}
        </div>
        {{end}} 

        {{if .flash.custom_success}}
        <div class="alert alert-success alert-dismissible" role="alert">
            <button type="button" class="close" data-dismiss="alert" aria-label="Close"><span aria-hidden="true">×</span></button>
            <i class="fa fa-check-circle"></i> {{.flash.custom_success}}
        </div>
        {{end}} 

        <h3 class="panel-title">Users</h3>
        <a href="{{ urlfor "UserController.AddUser"}}"><span class="pull-right button-add">Add <i class="fa fa-plus-circle"></i></span></a>
    </div>
    <div class="panel-body">
<table class="table">
        <thead>
            <tr><th>#</th><th>Name</th><th>Pod</th><th>Actions</th></tr>
        </thead>
        <tbody>
            {{range $key,$user := .userList}}
                <tr>
                    <td>{{ calculate $.pageStart $key "+"}} </td>
                    <td>{{ $user.FullName}}</td>
                    <td>{{ $user.Pods.Name}}</td>
                    <td>
                        <a href="{{ urlfor "UserController.EditUser" ":id" $user.Id}}"><span class="button-edit">Edit <i class="fa fa-edit"></i>
                        </span></a>
                        <a href="#">
                        <span class="button-delete" data-toggle="modal" data-target="#myModal{{ $user.Id}}">Delete <i class="fa fa-trash"></i>
                        </span>
                        </a>
                        {{ template "base/delete-record.tpl" (dynamicMap  "UserId" $user.Id "method" $.deleteMethod "deleteUrl" (urlfor "UserController.DeleteUser" ":id" $user.Id) ) }}
                    </td>
                </tr>
            {{end}}            
        </tbody>
    </table>
    {{template "base/paginator.tpl" .}}
    </div>
</div>