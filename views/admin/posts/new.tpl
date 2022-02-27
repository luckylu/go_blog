{{template "layout/admin/application.tpl" .}}
{{define "head"}}
<title>Admin</title>
{{end}}
{{define "body"}}
<div class="ui grid centered container">
  <div class="row"></div>
  <div class="three wide column">
  </div>
  <div class="ten wide column">
    <form class="ui form" action="{{urlfor "admin.PostsController.Create" }}"
      method="post">
      <div class="field">
        <label>标题</label>
        <input type="text" name="title" placeholder="标题">
      </div>
      <div class="field">
        <label>内容</label>
        <textarea name="content" rows="20"></textarea>
      </div>
      <button class="ui primary button" type="submit">创建</button>
      <a class="ui secondary button" href="/admin/posts">取消</a>
    </form>
  </div>
  <div class="three wide column">
  </div>
</div>
{{end}}
