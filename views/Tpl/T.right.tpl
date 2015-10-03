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
  	          <li class="cur"><a href="#" >点击排行</a></li>
  	          <li><a href="#">最新文章</a></li>
  	          <li><a href="#" >最新评论</a></li>
  	        </ul>
  	      </div>
  	      </div>
  	      <div class="ms-main" id="ms-main">
  	        <div style="display: block;" class="bd bd-news" >
  	          <ul>
  	           {{ range .VIewsTopics}}
                        <li class="list-group-item"><a href="/topic?topicId={{.Id}}" target="_blank">{{.Title}}</a></li>
                 {{end}}
  	          </ul>
  	        </div>
  	        <div  class="bd bd-news">
  	          <ul>
                {{ range .NewTopics}}
  	            <li class="list-group-item"><a href="/topic?topicId={{.Id}}" target="_blank">{{.Title}}</a></li>
  	           {{end}}
  	          </ul>
  	        </div>

  	        <div class="bd bd-news ds-recent-comments" data-num-items="6" data-show-avatars="0" data-show-time="1" data-show-title="0" data-show-admin="1" data-excerpt-length="70">
  	        
  	        </div>
            <script type="text/javascript">
            var duoshuoQuery = {short_name:"cwen"};
              (function() {
                var ds = document.createElement('script');
                ds.type = 'text/javascript';ds.async = true;
                ds.src = (document.location.protocol == 'https:' ? 'https:' : 'http:') + '//static.duoshuo.com/embed.js';
                ds.charset = 'UTF-8';
                (document.getElementsByTagName('head')[0] 
                 || document.getElementsByTagName('body')[0]).appendChild(ds);
              })();

              </script>
            <!-- 多说公共JS代码 end -->
  	      </div>
  	      <!--ms-main end --> 
</div>

<div class="panel panel-default">
<div class="panel-heading">分类目录</div>
  	<!-- List group -->
  	<ul class="list-group">
	{{range .Categories}}
	    <li class="list-group-item"><a href="/?cate={{.Id}}">{{.Title}}</a></li>
	{{end}}
  	</ul>
</div>

<div class="panel panel-default">
	<div class="panel-heading">标签云</div>
	<div class="cloud">
	      <ul>
        {{range .Labels}}
	        <li><a href="/?label={{.Id}}">{{.Title}}</a></li>
          {{end}}

	      </ul>

	</div>
	<br>
</div>



