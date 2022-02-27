{{template "layout/admin/application.tpl" .}}
{{define "head"}}
<title>Admin</title>
{{end}}
{{define "body"}}
<div class="ui two column centered grid">
  <div class="row"></div>
  <div class="column">
    <table class="ui celled table">
      <thead>
        <tr>
          <th>ID</th>
          <th>内容</th>
          <th>创建时间</th>
          <th>操作</th>
        </tr></thead>
      <tbody>
        {{range $index, $elem := .comments}}
        <tr>
          <td data-label="title">{{$elem.Id}}</td>
          <td data-label="title">{{$elem.Content}}</td>
          <td data-label="created_at">{{$elem.CreatedAt}}</td>
          <td data-label="operation">
            <a href="/admin/comments/{{$elem.Id}}/destroy">删除</a>
          </td>
        </tr>
        {{end}}
      </tbody>
    </table>
  </div>
  {{end}}
