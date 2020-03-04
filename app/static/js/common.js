$(document).ready(function() {
    $("select").selectpicker();
    if ($(".button-delete").length){
        $(".button-delete").on("click", function(event) {
            event.preventDefault()
        }) 
    }     
});
