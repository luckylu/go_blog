{{template "layout/login.tpl" .}}
{{define "head"}}
<title>登录</title>
{{end}}
{{define "body"}}
<div class="ui middle aligned center aligned grid">
  <div class="column">
    <h2 class="ui image header">
      <div class="content">
        MyBlog
      </div>
    </h2>
    <form url="/login" method="post" class="ui large form">
      <div class="ui stacked secondary  segment">
        <div class="field">
          <div class="ui left icon input">
            <i class="user icon"></i>
            <input type="text" name="name" placeholder="用户名">
          </div>
        </div>
        <div class="field">
          <div class="ui left icon input">
            <i class="lock icon"></i>
            <input type="password" name="password" placeholder="密码">
          </div>
        </div>
        <button class="ui fluid large teal submit button">登录</button>
      </div>

      <div class="ui error message"></div>

    </form>

    <div class="ui message">
      新用户? <a href="/register">注册</a>
    </div>
  </div>
</div>
{{end}}
<style>
   body > .grid {
      height: 100%;
    }
    .image {
      margin-top: -100px;
    }
    .column {
      max-width: 450px;
    }
</style>
