<div class="panel panel-headline">
    <div class="panel-heading">
        <h3 class="panel-title">Teams</h3>
                <a href="{{ urlfor "TeamController.AddTeam"}}"><span class="pull-right button-add">Add <i class="fa fa-plus-circle"></i></span></a>
                <div class="clearfix"></div>
            <div class="panel-body">
                <table class="table">
                    <thead>
                        <tr><th>#</th><th>Team Name</th><th>Pod</th><th class="pull-right">Action</th></tr>
                    </thead>
                    <tbody>
                        {{range $key,$val := .teamList}}
                        <tr>
                            <td>{{ calculate $.pageStart $key "+" }}</td>
                            <td>{{$val.Name}}</td>
                            <td>{{$val.Pods.Name}}</td>
                            <td> 
                                <span class="pull-right button-edit"> 
                                    Edit <i class="fa fa-edit"></i> 
                                <span>
                                &nbsp; &nbsp;
                                <span class="pull-right button-delete"> 
                                     Delete <i class="fa fa-trash"></i>
                                <span>
                            </td>
                        </tr>
                        {{end}}
                    </tbody>
                </table>
                {{template "base/paginator.tpl" .}}
            </div>
    </div>
</div>