<link rel="stylesheet"  type="text/css" href="/static/css/admin.css">
<div class="row-fluid">

<div class="btn-toolbar">
    <a type="button" href="/admin/topic" class="btn btn-primary" ><span class="glyphicon glyphicon-plus" aria-hidden="true"  ></span>添加博客</a>
  <div class="btn-group">
  </div>
</div>
<div class="well">
    <table class="table">
      <thead>
        <tr>
          <th>博客id</th>
          <th>博客标题</th>
          <th>博客发布时间</th>
          <th>跟新时间</th>
          <th>浏览数</th>
          <th>操作</th>
        </tr>
      </thead>
      <tbody>
        {{range .Topics}}
          <tr>

          <td>{{.Id}}</td>
          <td>{{.Title}}</td>
          <td>{{.Created}}</td>
          <td>{{.Updated}}</td>
          <th>{{.Views}}</th>
          <td>
                <a href="/admin/topic/EditTopicShow?tid={{.Id}}" class="topic_edit" section="{{.Id}}"><span class="glyphicon glyphicon-pencil" aria-hidden="true"></span></a>&nbsp;&nbsp;
                <a href="javascript:;" class="topic_del" section="{{.Id}}" ><span class="glyphicon glyphicon-trash" aria-hidden="true"></span></a>
          </td>
        </tr>
        {{end}}
      </tbody>
    </table>
</div>
  </div>

<div class="pagination">

    {{ str2html .PageStr }}<!-- |markdown -->

</div>
<script type="text/javascript" src="/static/js/topic.js"></script>

