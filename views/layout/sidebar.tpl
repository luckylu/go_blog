  <div class="ui top fixed inverted menu visible">
    <a class="item" href="/">
    <i class="home icon"></i>
      首页
    </a>

    {{if .user}}
      <div class="right ui dropdown item">
        <div class="default text">{{.user.Name}}</div>
        <input type="hidden" value="">
        <i class="dropdown icon"></i>
        <div class="menu">
        {{if .user.Admin}}
          <a class="item" href="/admin">后台管理</a>
        {{end}}
          <a class="item" href="/logout">退出</a>
        </div>
      </div>
    {{else}}
      <a class="right item" href="/login">登录</a>
    {{end}}
  </div>

