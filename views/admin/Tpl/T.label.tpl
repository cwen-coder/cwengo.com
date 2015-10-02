<link rel="stylesheet"  type="text/css" href="/static/css/admin.css">

<div class="row-fluid">
	<div class="btn-toolbar">
		<button type="button" class="btn btn-primary" data-toggle="modal" data-target="#addLabel"><span class="glyphicon glyphicon-plus" aria-hidden="true"  ></span>添加标签</button>
	</div>
	<div class="well">
		<table class="table table-hover">
		 	<thead>
		 		<th>标签ID</th>
		 		<th>标签名称</th>
		 		<th>文章数</th>
		 		<th>操作</th>
		 	</thead>
		 	<tbody>
		 		{{range .Labels}}
		 		<tr>
		 			<td>{{.Id}}</td>
		 			<td>{{.Title}}</td>
		 			<td>{{.TopicCount}}</td>
		 			<td>
		 				<a href="javascript:;" class="label_edit" section="{{.Id}}"><span class="glyphicon glyphicon-pencil" aria-hidden="true"></span></a>&nbsp;&nbsp;
		 				<a href="javascript:;" class="label_del" section="{{.Id}}" ><span class="glyphicon glyphicon-trash" aria-hidden="true"></span></a>
		 			</td>
		 		</tr>
		 		{{end}}
		 	</tbody>
		</table>
	</div>
</div>

<div class="modal fade" id="addLabel" tabindex="-1" role="dialog" aria-labelledby="myModalLabel">
  <div class="modal-dialog" role="document">
    <div class="modal-content">
      <div class="modal-header">
        <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
        <h4 class="modal-title" id="myModalLabel">添加标签</h4>
      </div>
      <div class="modal-body">
       		<form  method="post">
       		  <div class="form-group">
       		    <label>新添标签名</label>
       		    <input type="text" class="form-control" name="NewLabel" id="NewLabel" placeholder="Email">
       		  </div>
      </div>
      <div class="modal-footer">
        <button type="button" class="btn btn-default" data-dismiss="modal">关闭</button>
        <button type="submit" class="btn btn-primary" onclick="" >添加</button>
        <script>
        function checkInput() {
	        var NewCategory = document.getElementById("NewLabel");
	        if(NewCategory.value.length == 0) {
	            alert("标签名不能为空");
	            return false;
	        }
      	}
        </script>
        </form>
      </div>
    </div>
  </div>
</div>

<div class="modal fade" id="editLabel" tabindex="-1" role="dialog" aria-labelledby="myModalLabel">
  <div class="modal-dialog" role="document">
    <div class="modal-content">
      <div class="modal-header">
        <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
        <h4 class="modal-title" id="myModalLabel">标签修改</h4>
      </div>
      <div class="modal-body">
       		<form  method="post">
       		  <div class="form-group">
       		    <label>标签名</label>
       		    <input type="hidden" name="EditCategoryId" id="EditLabelId">
       		    <input type="text" class="form-control" name="EditNewLabel" id="EditNewLabel" placeholder="Email">
       		  </div>
      </div>
      <div class="modal-footer">
        <button type="button" class="btn btn-default" data-dismiss="modal">关闭</button>
        <button type="button" class="btn btn-primary" onclick="return checkEditInput();" >保存</button>
        <script>
        function checkEditInput() {
	        var EditNewLabel = document.getElementById("EditNewLabel");
	        if(EditNewLabel.value.length == 0) {
	            alert("分类名称不能为空");
	            return false;
	        }
      	}
        </script>
        </form>
      </div>
    </div>
  </div>
</div>
<script type="text/javascript" src="/static/js/label.js"></script>