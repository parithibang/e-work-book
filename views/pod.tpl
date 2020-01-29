<div class="panel panel-headline">
    <div class="panel-heading">
        <h3 class="panel-title">POD</h3>
        <div class="panel">
            <div class="panel-heading">
                <h3 class="panel-title">Active PODs</h3>
                <span class="pull-right button-add">Add <i class="fa fa-plus-circle"></i></span>
                <div class="clearfix"></div>
            </div>
            <div class="panel-body">
                <table class="table">
                    <thead>
                        <tr><th>#</th><th>POD Name</th><th>Team</th><th>Created@</th><th class="pull-right">Action</th></tr>
                    </thead>
                    <tbody>
                        {{range $val := .POD}}
                        <tr>
                            <td>{{$val.Id}}</td>
                            <td>{{$val.Name}}</td>
                            <td>{{$val.Name}}</td>
                            <td>{{$val.Name}}</td>
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
            </div>
        </div>
    </div>
</div>