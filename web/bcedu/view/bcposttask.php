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

<?php
//导航条
include(V.'bctop.php');
?>  

  <?php
 // var_dump($_POST);  //超全局变量，$_GET  $POST  (数组类型)
if(isset($_POST["ac"]) && $_POST["ac"]=="login")  //防止灌水，判断表单确实是自己的表单
{
  $taskname=$_POST["taskname"];
  $duration=$_POST["duration"];
  $starttime=$_POST["starttime"];
  $taskbrif=$_POST["taskbrif"];
  $exceptedprice=$_POST["exceptedprice"];
  echo("任务名是：{$taskname}");


  $data=array(
    "taskname" => $taskname,
    "duration" => $duration,
    "starttime" => $starttime,
    "taskbrif" => $taskbrif,
    "exceptedprice" => $exceptedprice
    );

    $ch = curl_init();

    curl_setopt($ch, CURLOPT_URL, "http://47.100.79.216:3000/posttask");
    curl_setopt($ch, CURLOPT_POST, 1);
    //The number of seconds to wait while trying to connect. Use 0 to wait indefinitely.
    curl_setopt($ch, CURLOPT_CONNECTTIMEOUT, 60);
    curl_setopt($ch, CURLOPT_POSTFIELDS , http_build_query($data));
    curl_setopt($ch, CURLOPT_RETURNTRANSFER, 1);

    $output = curl_exec($ch);

    echo $output;

    curl_close($ch);


}
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
<input type="submit" value="提交任务" />
  </form>
<br>
      <div class="alert alert-info" role="alert">
        <p class="text-center">
   <strong><a href="?page=bcstudent">返回个人主页</a></strong>
      </div>
        </div>


 </BODY>
</HTML>
