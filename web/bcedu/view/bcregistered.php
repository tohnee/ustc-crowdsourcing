<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN" "http://www.w3.org/TR/html4/loose.dtd">
<HTML>
 

 <BODY>

 <?php
    include(V.'header-content.php');
?>
        <link rel="stylesheet" href="assets/css/style.css">


  </head>
  <body>
<?php
//导航条
include(V.'bctop-bar.php');
?>  

 
  <form name="form1" method="post" action="">
任务名：<input type="text" name="taskname" />
<br>
持续时间：<input type="text" name="duration" />
<br>
开始时间：<input type="text" name="starttime" />
<br>
任务简介：<input type="text" name="taskbrif" />
<br>
期望任务价格：<input type="text" name="exceptedprice" />
<input type="hidden" name="ac" value="login" />
<br>
<input type="submit" value="发布任务" />
  </form>

<br>
      <div class="alert alert-info" role="alert">
        <p class="text-center">
   <strong><a href="?page=bchome">返回首页</a></strong>
      </div>
        </div> <!-- /container -->
  <?php
 // var_dump($_POST);  //超全局变量，$_GET  $POST  (数组类型)

  ?>




 </BODY>
</HTML>
