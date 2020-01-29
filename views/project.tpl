<div class="panel panel-headline">
    <div class="panel-heading">
        <h3 class="panel-title">Project</h3>
        <div class="panel">
            <div class="panel-heading">
                <h3 class="panel-title">Active Projects</h3>
            </div>
            <div class="panel-body">
                <table class="table">
                    <thead>
                        <tr><th>#</th><th>Project Name</th><th>Team</th><th>Created@</th></tr>
                    </thead>
                    <tbody>
                        {{range $val := .PROJECTS}}
                        <tr><td>{{$val.Id}}</td><td>{{$val.Name}}</td><td>{{$val.Name}}</td><td>{{$val.Name}}</td></tr>
                        {{end}}
                    </tbody>
                </table>
            </div>
        </div>
    </div>
</div>