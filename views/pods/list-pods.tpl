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
        <h3 class="panel-title">Pods</h3>
        <a href="{{ urlfor "PodController.AddPod"}}">
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
                    <th>Pod Name</th>
                    <th>Unit</th>
                    <th>Action</th>
                </tr>
            </thead>
            <tbody>
                {{range $key,$pod := .podList}}
                    <tr>
                        <td>{{ calculate $.pageStart $key "+" }}</td>
                        <td>{{$pod.Name}}</td>
                        <td>{{$pod.Units.Name}}</td>
                        <td>
                            <a href="{{ urlfor "PodController.EditPod" ":id" $pod.Id}}">
                                <span class="button-edit">Edit<i class="fa fa-edit"></i>
                                </span>
                            </a>
                            <a href="#">
                                <span class="button-delete" data-toggle="modal" data-target="#myModal{{ $pod.Id}}">Delete
                                    <i class="fa fa-trash"></i>
                                </span>
                            </a>
                            {{ template "base/delete-record.tpl" (dynamicMap  "RecordId" $pod.Id "method" $.deleteMethod "deleteUrl" (urlfor "PodController.DeletePod" ":id" $pod.Id) ) }}

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
