{{template "layout/application.tpl" .}}
{{define "head"}}
<title>Home page</title>
{{end}}
{{define "body"}}
<div class="ui two column centered grid">
<div class="row"></div>
  <div class="column">
    <div class="ui segment">
      <h2 class="ui header">{{.post.Title}}</a></h2>
      <p>{{.post.Content}}</p>
    </div>
  </div>

</div>
{{ if gt (.comments|len) 0 }}
  <div class="ui two column centered grid">
    <div class="column">
      <div class="ui segment">
        {{range $index, $elem := .comments}}
          <p>{{$elem.Content}}</p>
        {{end}}
      </div>
    </div>
  </div>
{{end}}
  <div class="ui two column centered grid">
    <div class="column">
      <div class="ui segment">
        <h2 class="ui header">发表评论</h2>
        <form class="ui form" action="/posts/{{.post.Id}}/comments" method="post">
          <div class="field">
            <label>评论内容</label>
            <textarea name="content"></textarea>
          </div>
          <button class="ui button primary" type="submit">提交</button>
      </div>
    </div>
  </div>
{{end}}
