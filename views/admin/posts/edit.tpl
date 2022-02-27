{{template "layout/admin/application.tpl" .}}
{{define "head"}}
<title>Admin</title>
{{end}}
{{define "body"}}
<div class="ui grid container">
  <div class="row"></div>
  <div class="three wide column">
  </div>
  <div class="ten wide column">
    <form class="ui form" action="/admin/posts/{{.post.Id}}" method="POST">
      <div class="field">
        <label>标题</label>
        <input type="text" name="title" placeholder="标题" value="{{.post.Title}}">
      </div>
      <div class="field">
        <label>内容</label>
        <textarea name="content" rows="20">{{.post.Content}}</textarea>
      </div>
      <button class="ui primary button" type="submit">更新</button>
      <a class="ui secondary button" href="/admin/posts">取消</a>
    </form>
  </div>
  <div class="three wide column">
  </div>
</div>
{{end}}
