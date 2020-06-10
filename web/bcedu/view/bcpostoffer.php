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
include(V.'bctopteacher.php');
?>  

  <?php
 // var_dump($_POST);  //超全局变量，$_GET  $POST  (数组类型)
if(isset($_POST["ac"]) && $_POST["ac"]=="login")  //防止灌水，判断表单确实是自己的表单
{
  $arrivaltime=$_POST["arrivaltime"];
  $departuretime=$_POST["departuretime"];
  $cost=$_POST["cost"];
  $taskname=$_POST["taskname"];
  echo("任务名是：{$taskname}");


  $data=array(
    "arrivaltime" => $arrivaltime,
    "departuretime" => $departuretime,
    "cost" => $cost,
    "taskname" => $taskname
    );

    $ch = curl_init();

    curl_setopt($ch, CURLOPT_URL, "http://47.100.79.216:3000/postoffer");
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
到达时间：<input type="text" name="arrivaltime" />
<br>
离开时间：<input type="text" name="departuretime" />
<br>
报价：<input type="text" name="cost" />
<br>
任务名：<input type="text" name="taskname" />
<input type="hidden" name="ac" value="login" />
<br>
<input type="submit" value="提交请求" />
  </form>
<br>
      <div class="alert alert-info" role="alert">
        <p class="text-center">
   <strong><a href="?page=bcstutor">返回</a></strong>
      </div>
        </div>


 </BODY>
</HTML>
