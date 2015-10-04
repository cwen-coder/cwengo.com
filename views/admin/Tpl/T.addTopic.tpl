<!-- <link rel="stylesheet" type="text/css" href="/static/css/demo.css" /> -->
<link rel="stylesheet" type="text/css" href="/static/css/topic.css" />
<link rel="stylesheet" type="text/css" href="/static/css/bootstrap-markdown.min.css">
<div class="row-fluid">
<div class="well">
    <div id="myTabContent" class="tab-content">
      <div class="tab-pane active in" id="home">
    <form id="tab" method="post">
        <div class="form-group topic">
            <label>博客标题</label>
            <input type="text" name="title" value="" id="title" class="form-control">
        </div>
        <div class="form-group topic">
            <label>博客分类</label>
            <select class="form-control" name="category">
              {{range .Categories}} 
                <option name="category" value="{{.Id}}">{{.Title}}</option>
              {{end}}
            </select>
        </div>
        <div class="form-group topic">
            <label>博客标签</label>
            <div>
            {{range .Labels}}
                <label class="checkbox-inline">
                  <input type="checkbox" name="label[]" value="{{.Id}}"> {{.Title}}
                </label>
            {{end}}
            </div>
        </div>
      <div class="form-group topic">
          <label>博客概要</label>
          <textarea class="form-control" name="summery" id="summery" data-provide="markdown"  rows="12"></textarea>
      </div>
        <div class="form-group topic">
			       <label>博客内容</label>
            <textarea  class="form-control" name="content"  data-provide="markdown"  rows="12"  ></textarea>
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

          var summery = document.getElementById("summery");
          if(summery.value.length == 0) {
              alert("博客概要不能为空");
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
  <script type="text/javascript" src="/static/js/markdown.js"></script>
  <script  type="text/javascript"  src="/static/js/to-markdown.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap-markdown.js"></script>
