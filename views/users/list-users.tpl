<div class="panel panel-headline">
    <div class="panel-heading">
        <h3 class="panel-title">Users</h3>
        <a href="{{ urlfor "UserController.AddUser"}}"><span class="pull-right button-add">Add <i class="fa fa-plus-circle"></i></span></a>
    </div>
    <table class="table">
        <thead>
            <tr><th>#</th><th>Name</th><th>Pod</th><th>Actions</th></tr>
        </thead>
        <tbody>
            {{range $user := .userList}}
                <tr>
                    <td>{{ $user.Id}} </td>
                    <td>{{ $user.FullName}}</td>
                    <td>{{ $user.Pods.Name}}</td>
                    <td>
                        <span class="button-edit">Edit <i class="fa fa-edit"></i>
                        </span>
                        <span class="button-delete">Delete <i class="fa fa-trash"></i>
                        </span>
                    </td>
                </tr>
            {{end}}
            
        </tbody>
    </table>
    {{template "base/paginator.tpl" .}}
</div>