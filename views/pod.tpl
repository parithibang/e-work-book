<div class="panel panel-headline">
    <div class="panel-heading">
        <h3 class="panel-title">POD</h3>
        <div class="panel">
            <div class="panel-heading">
                <h3 class="panel-title">Active PODs</h3>
            </div>
            <div class="panel-body">
                <table class="table">
                    <thead>
                        <tr><th>#</th><th>POD Name</th><th>Team</th><th>Created@</th></tr>
                    </thead>
                    <tbody>
                        {{range $val := .POD}}
                        <tr><td>{{$val.Id}}</td><td>{{$val.Name}}</td><td>{{$val.TeamId}}</td><td>{{$val.CreatedAt}}</td></tr>
                        {{end}}
                    </tbody>
                </table>
            </div>
        </div>
    </div>
</div>