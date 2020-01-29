<div class="panel panel-headline">
    <div class="panel-heading">
        <h3 class="panel-title">Users</h3>
    </div>
    <table class="table">
        <thead>
            <tr><th>#</th><th>Name</th><th>Pod</th></tr>
        </thead>
        <tbody>
            {{range $user := .userList}}
                <tr><td>{{ $user.Id}} </td><td>{{ $user.FullName}}</td><td>{{ $user.Pods.Name}}</td>  </tr>
            {{end}}
            
        </tbody>
    </table>
    {{template "base/paginator.tpl" .}}
</div>