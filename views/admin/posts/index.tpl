{{template "layout/application.tpl" .}}
{{define "head"}}
<title>Admin</title>
{{end}}
{{define "body"}}
<div class="ui grid container">
  <div class="row">
    <div class="two wide column">
    </div>
    <div class="twelve wide column">
          <a class="ui primary button" href="/admin/posts/new">新建</a>

    </div>
  </div>
  <div class="row">
  <div class="two wide column">
  </div>
  <div class="twelve wide column">
    <table class="ui celled table">
      <thead>
        <tr>
          <th>标题</th>
          <th>创建时间</th>
          <th>操作</th>
        </tr></thead>
      <tbody>
        {{range $index, $elem := .posts}}
        <tr>
          <td data-label="title">{{$elem.Title}}</td>
          <td data-label="created_at">{{$elem.CreatedAt}}</td>
          <td data-label="operation">
            <a href="/admin/posts/{{$elem.ID}}/edit">编辑</a>
            <a href="/admin/posts/{{$elem.ID}}/destroy">删除</a>
          </td>
        </tr>
        {{end}}
      </tbody>
    </table>
    <div class="two wide column">
    </div>
    </div>
  </div>
  {{end}}
