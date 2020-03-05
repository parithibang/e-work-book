<div class="panel panel-headline">
    <div class="panel-heading">
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
        <h3 class="panel-title">Projects</h3>
        <a href="{{ urlfor "ProjectController.AddProject"}}">
            <span class="pull-right button-add">Add
                <i class="fa fa-plus-circle"></i>
            </span>
        </a>
        <div class="clearfix"></div>
    </div>
    <div class="panel-body">
        <table class="table">
            <thead>
                <tr>
                    <th>#</th>
                    <th>Name</th>
                    <th>Action</th>
                </tr>
            </thead>
            <tbody>
                {{range $key,$project := .projectList}}
                    <tr>
                        <td>{{ calculate $.pageStart $key "+" }}</td>
                        <td>{{$project.Name}}</td>
                        <td>
                            <a href="{{ urlfor "ProjectController.EditProject" ":id" $project.ID}}">
                                <span class="button-edit">Edit<i class="fa fa-edit"></i>
                                </span>
                            </a>
                            <a href="#">
                                <span class="button-delete" data-toggle="modal" data-target="#myModal{{ $project.ID}}">Delete
                                    <i class="fa fa-trash"></i>
                                </span>
                            </a>
                            {{ template "base/delete-record.tpl" (dynamicMap  "RecordId" $project.ID "method" $.deleteMethod "deleteUrl" (urlfor "ProjectController.DeleteProject" ":id" $project.ID) ) }}

                        </td>
                    </tr>
                {{end}}
            </tbody>
        </table>
        <span class="badge text-uppercase">
            Total Record :
            {{ .count }}</span>
        {{template "base/paginator.tpl" .}}
    </div>

</div>
