<link rel="stylesheet" type="text/css" href="/static/css/demo.css" />
<link rel="stylesheet" type="text/css" href="/static/css/topic.css" />
        
<script type="text/javascript" src="/static/js/Markdown.Converter.js"></script>
<script type="text/javascript" src="/static/js/Markdown.Sanitizer.js"></script>
<script type="text/javascript" src="/static/js/Markdown.Editor.js"></script>
<div class="row-fluid">
<div class="well">
    <div id="myTabContent" class="tab-content">
      <div class="tab-pane active in" id="home">
    <form id="tab" method="post" action="/admin/topic/EditTopicAct">
        <div class="form-group topic">
            <label>博客标题</label>
            <input type="text" name="title" value="{{.Topic.Title}}" id="title" class="form-control">
            <input type="hidden" name="tid" value="{{.Topic.Id}}" class="form-control">
        </div>
        <div class="form-group topic">
            <label>博客分类</label>
            <select class="form-control" name="category">
              {{range .Categories}} 
                <option name="category" value="{{.Id}}" {{if .IsSelected}} selected="selected" {{end}} >{{.Title}}</option>
              {{end}}
            </select>
        </div>
        <div class="form-group topic">
            <label>博客标签</label>
            <div>
            {{range .Labels}}
                <label class="checkbox-inline">
                  <input type="checkbox" name="label[]" value="{{.Id}}" {{if .IsSelected}} checked="checked" {{end}}> {{.Title}}
                </label>
            {{end}}
            </div>
        </div>
        <div class="wmd-panel form-group">
             <label>博客内容</label>
            <div id="wmd-button-bar"></div>
            <textarea name="content" class="wmd-input" id="wmd-input">{{.Topic.Content}}</textarea>
        </div>
        <div id="wmd-preview" class="wmd-panel wmd-preview"></div>
        <div class="btn-toolbar">
          <button class="btn btn-primary" type="submit" onclick="return checkInput();"><i class="icon-save" ></i> 提交</button>
        
    </div>
    </form>
    <script type="text/javascript">
        function checkInput() {
          var title = document.getElementById("title");
          if(title.value.length == 0) {
              alert("博客名称不能为空");
              return false;
          }
          var content = document.getElementById("wmd-input");
          if(content.value.length == 0) {
              alert("博客内容不能为空");
              return false;
          }
        }
    </script>
      </div>
  </div>

</div>
  </div>
<script type="text/javascript">
(function () {
      var converter1 = Markdown.getSanitizingConverter();
      var editor1 = new Markdown.Editor(converter1);
      editor1.run();
})();
</script>