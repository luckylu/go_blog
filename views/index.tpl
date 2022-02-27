{{template "layout/application.tpl" .}}
{{define "head"}}
<title>Home page</title>
{{end}}
{{define "body"}}
<div class="ui two column centered grid">
  <div class="row"></div>
  <div class="column">
    {{range $index, $elem := .Posts}}
    <div class="ui segment">
      <h2 class="ui header"><a href="posts/{{$elem.Id}}">{{$elem.Title}}</a></h2>
      <p>{{$elem.Content}}</p>
    </div>
    {{end}}
  </div>
</div>
{{end}}
