<!-- Modal -->
<div class="modal fade" id="myModal{{.RecordId}}" tabindex="-1" role="dialog" aria-labelledby="myModalLabel">
    <div class="modal-dialog" role="document">
        <div class="modal-content">
            <div class="modal-header">
                <button type="button" class="close" data-dismiss="modal" aria-label="Close">
                    <span aria-hidden="true">&times;</span>
                </button>
                <h4 class="modal-title" id="myModalLabel">Are you sure?</h4>
            </div>
            <div class="modal-body">
                Do you really want to delete these records? This process cannot be undone.
            </div>
            <div class="modal-footer">
                <form action="{{.deleteUrl}}" role="form" id="user" class="user-create" method="POST">
                    <input type="hidden" name="_method" value="{{ .method }}"/>
                    <input type="hidden" name="request-params" value="{{ .requestParams }}"/>

                    <button type="button" class="btn btn-default" data-dismiss="modal">Close</button>
                    <button type="submit" class="btn btn btn-danger">
                        <i class="fa fa-trash-o"></i>
                        Delete
                    </button>
                </form>
            </div>
        </div>
    </div>
</div>
