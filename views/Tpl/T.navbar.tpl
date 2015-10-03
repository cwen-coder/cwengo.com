
<div class="navbar-header" display="inline" role="navigation">
			<button type="button" class="navbar-toggle" data-toggle="collapse" 
		         data-target="#example-navbar-collapse">
		         <span class="sr-only">切换导航</span>
		         <span class="icon-bar"></span>
		         <span class="icon-bar"></span>
		         <span class="icon-bar"></span>

		     </button>
			 <a class="navbar-brand brand-img" href="/"><img src="/static/img/cwen.png" alt="" ></a>
		     <a class="navbar-brand" href="/" ><big>沉稳，不乏可爱</big></a>	
		</div>
		<!-- <div>
			<p class="navbar-text">CWen留下的痕迹</p>
		</div> -->
		<div class="pull-right collapse navbar-collapse" id="example-navbar-collapse">
			<ul class="nav navbar-nav">
				<li {{if .IsHome}} class="active" {{end}} ><a href="/">主页</a></li>
				<li {{if .IsArchive}} class="active" {{end}}><a href="/archive">文章归档</a></li>
				<li {{if .IsAbout}} class="active" {{end}}><a href="/about">关于</a></li>
			</ul>
		</div>