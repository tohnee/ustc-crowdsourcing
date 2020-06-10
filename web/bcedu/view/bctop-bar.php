		<nav class="navbar navbar-inverse navbar-fixed-top" role="navigation">
			<div class='container'>
				<div class="navbar-header">
					<button type="button" class="navbar-toggle"
					data-toggle="collapse" data-target="#mainmenu">
						<span class="sr-only">导航条</span>
						<span class="icon-bar"></span>
						<span class="icon-bar"></span>
						<span class="icon-bar"></span>
					</button>
					<a class="navbar-brand" href="?page=about">量子众包</a>
				</div><!--navbar-header-->

				<div class="collapse navbar-collapse" id="mainmenu">
					<ul class="nav navbar-nav navbar-left">
						<li <?php if($page=="bchome") echo "class=\"active\"";?>><a href="?page=bchome">首页</a></li>
						<li <?php if($page=="bcstudent") echo "class=\"active\"";?>><a href="?page=bcstudent">我是发布者(学生)</a></li>
						<li <?php if($page=="tbcutor") echo "class=\"active\"";?>><a href="?page=bctutor">我是工人(老师)</a></li>
						<li <?php if($page=="bccontact") echo "class=\"active\"";?>><a href="?page=bccontact">联系我们</a></li>		<li <?php if($page=="login") echo "class=\"active\"";?>><a href="?page=login">传统版入口</a></li>

					</ul>
				</div><!--collapse-->

			</div><!--container-->
		</nav><!--navbar-->