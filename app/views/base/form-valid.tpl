{{range $error := .}}
    <span class="text-danger visible-lg"> {{ $error.Message }} </span>
{{ end}}
