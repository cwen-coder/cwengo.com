<div class="panel panel-default">
<!-- <div class="panel-heading">浏览排名</div>
  	List group
  	<ul class="list-group">
	    <li class="list-group-item"><a href="/">Cras justo odio</a></li>
	    <li class="list-group-item"><a href="/">Cras justo odio</a></li>
	    <li class="list-group-item"><a href="/">Cras justo odio</a></li>
	    <li class="list-group-item"><a href="/">Cras justo odio</a></li>
	    <li class="list-group-item"><a href="/">Cras justo odio</a></li>
  	</ul> -->
  	<script>
  	window.onload = function ()
  	{
  		var oLi = document.getElementById("tab").getElementsByTagName("li");
  		var oUl = document.getElementById("ms-main").getElementsByTagName("div");
  		
  		for(var i = 0; i < oLi.length; i++)
  		{
  			oLi[i].index = i;
  			oLi[i].onmouseover = function ()
  			{
  				for(var n = 0; n < oLi.length; n++) oLi[n].className="";
  				this.className = "cur";
  				for(var n = 0; n < oUl.length; n++) oUl[n].style.display = "none";
  				oUl[this.index].style.display = "block"
  			}	
  		}
  	}
  	</script>
  	<div class="panel-heading padding-tab">
  	      <div class="ms-top ">
  	        <ul class="hd" id="tab">
  	          <li class="cur"><a href="/">点击排行</a></li>
  	          <li><a href="/">最新文章</a></li>
  	          <li><a href="/">最新评论</a></li>
  	        </ul>
  	      </div>
  	      </div>
  	      <div class="ms-main" id="ms-main">
  	        <div style="display: block;" class="bd bd-news" >
  	          <ul>
  	            <li class="list-group-item"><a href="/" target="_blank">住在手机里的朋友</a></li>
  	            <li class="list-group-item"><a href="/" target="_blank">教你怎样用欠费手机拨打电话</a></li>
  	            <li class="list-group-item"><a href="/" target="_blank">原来以为，一个人的勇敢是，删掉他的手机号码...</a></li>
  	            <li class="list-group-item"><a href="/" target="_blank">手机的16个惊人小秘密，据说99.999%的人都不知</a></li>
  	            <li class="list-group-item"><a href="/" target="_blank">你面对的是生活而不是手机</a></li>
  	            <li class="list-group-item"><a href="/" target="_blank">豪雅手机正式发布! 在法国全手工打造的奢侈品</a></li>
  	          </ul>
  	        </div>
  	        <div  class="bd bd-news">
  	          <ul>
  	            <li class="list-group-item"><a href="/" target="_blank">原来以为，一个人的勇敢是，删掉他的手机号码...</a></li>
  	            <li class="list-group-item"><a href="/" target="_blank">手机的16个惊人小秘密，据说99.999%的人都不知</a></li>
  	            <li class="list-group-item"><a href="/" target="_blank">住在手机里的朋友</a></li>
  	            <li class="list-group-item"><a href="/" target="_blank">教你怎样用欠费手机拨打电话</a></li>
  	            <li class="list-group-item"><a href="/" target="_blank">你面对的是生活而不是手机</a></li>
  	            <li class="list-group-item"><a href="/" target="_blank">豪雅手机正式发布! 在法国全手工打造的奢侈品</a></li>
  	          </ul>
  	        </div>
  	        <div class="bd bd-news">
  	          <ul>
  	            <li class="list-group-item"><a href="/" target="_blank">手机的16个惊人小秘密，据说99.999%的人都不知</a></li>
  	            <li class="list-group-item"><a href="/" target="_blank">你面对的是生活而不是手机</a></li>
  	            <li class="list-group-item"><a href="/" target="_blank">住在手机里的朋友</a></li>
  	            <li class="list-group-item"><a href="/" target="_blank">豪雅手机正式发布! 在法国全手工打造的奢侈品</a></li>
  	            <li class="list-group-item"><a href="/" target="_blank">教你怎样用欠费手机拨打电话</a></li>
  	            <li class="list-group-item"><a href="/" target="_blank">原来以为，一个人的勇敢是，删掉他的手机号码...</a></li>
  	          </ul>
  	        </div>
  	      </div>
  	      <!--ms-main end --> 
</div>

<div class="panel panel-default">
<div class="panel-heading">分类目录</div>
  	<!-- List group -->
  	<ul class="list-group">
	    <li class="list-group-item"><a href="/">Cras justo odio</a></li>
	    <li class="list-group-item"><a href="/">Cras justo odio</a></li>
	    <li class="list-group-item"><a href="/">Cras justo odio</a></li>
	    <li class="list-group-item"><a href="/">Cras justo odio</a></li>
	    <li class="list-group-item"><a href="/">Cras justo odio</a></li>
  	</ul>
</div>

<div class="panel panel-default">
	<div class="panel-heading">标签云</div>
	<div class="cloud">
	      <ul>
	        <li><a href="/">个人博客</a></li>
	        <li><a href="/">web开发</a></li>
	        <li><a href="/">前端设计</a></li>
	        <li><a href="/">Html</a></li>
	        <li><a href="/">CSS3</a></li>
	        <li><a href="/">Html5+css3</a></li>
	        <li><a href="/">百度</a></li>
	        <li><a href="/">Javasript</a></li>
	        <li><a href="/">web开发</a></li>
	        <li><a href="/">前端设计</a></li>
	        <li><a href="/">Html</a></li>
	        <li><a href="/">CSS3</a></li>
	        <li><a href="/">Html5+css3</a></li>
	        <li><a href="/">百度</a></li>
	      </ul>

	</div>
	<br>
</div>