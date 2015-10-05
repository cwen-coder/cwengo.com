<!-- <link rel="stylesheet" type="text/css" href="/static/css/demo.css" /> -->
<link rel="stylesheet" type="text/css" href="/static/css/topic.css" />
<!-- <link rel="stylesheet" type="text/css" href="/static/css/bootstrap-markdown.min.css"> -->
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
          <!-- <textarea class="form-control" name="summery" id="summery" data-provide="markdown"  rows="12"></textarea> -->
           <script type="text/javascript" charset="utf-8">
                      window.UEDITOR_HOME_URL = "/static/ueditor/";
                    </script>
           <script type="text/plain" id="summery" name="summery"></script>
      </div>
        <div class="form-group topic">
			       <label>博客内容</label>
           <script type="text/javascript" charset="utf-8">
                      window.UEDITOR_HOME_URL = "/static/ueditor/";
                    </script>
                     <script type="text/plain" id="content" name="content"></script>
                    
                   
                    <script type="text/javascript" src="/static/ueditor/ueditor.config.js"></script>
                    <script type="text/javascript" src="/static/ueditor/ueditor.all.min.js"></script>
                 
                    <script type="text/javascript" charset="utf-8">
                      var options = {"fileUrl":"/admin/article/upload","filePath":"","imageUrl":"/admin/article/upload","imagePath":"","initialFrameWidth":"90%","initialFrameHeight":"400"};
                      var ue = UE.getEditor("content", options);
                      var ue = UE.getEditor("summery", options);
                    </script>
        </div>
        <div class="btn-toolbar">
        	<button class="btn btn-primary" type="submit" onclick="return checkInput();"><i class="icon-save" ></i> 提交</button>
        
		</div>
    </form>
     <script type="text/javascript" src="http://cdn.staticfile.org/jquery/2.1.1-rc2/jquery.min.js"></script>
 <script type="text/javascript" src="/static/js/jquery.base64.js"></script>
    <script type="text/javascript">
        function checkInput() {
          var title = document.getElementById("title");
          if(title.value.length == 0) {
              alert("博客名称不能为空");
              return false;
          }
          var summery = document.getElementById("summery");
          console.log(summery.value);
          if(summery.value.length == 0) {
              alert("博客概要不能为空");
              return false;
          }
          var val = document.getElementById("content");   
          if(val.value.length == 0) {
              alert("博客内容不能为空");
              return false;
          }
          console.log(val.value);
          /*return false;*/
         /* $("#content").val($.base64.btoa(val.value));*/
        /*  console.log($("#content").val());
            return false;*/
         /* console.log(val);
          return false;*/
        }
    </script>
      </div>
  </div>

</div>
  </div>


