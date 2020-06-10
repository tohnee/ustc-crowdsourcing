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
  	  	欢迎注册区块链版的众包平台
  	<br>
  	  请输入用户名与您的类型(worker或者poster)
  	  <br>
用户名：<input type="text" name="username" />
<br>
类型：<input type="password" name="userpwd" />
<input type="hidden" name="ac" value="login" />
<br>
<input type="submit" value="提交表单" />
  </form>
<br>
			<div class="alert alert-info" role="alert">
				<p class="text-center">
   <strong><a href="?page=bclogin">返回登录页</a></strong>
			</div>
        </div>

  <?php
 // var_dump($_POST);  //超全局变量，$_GET  $POST  (数组类型)
if(isset($_POST["ac"]) && $_POST["ac"]=="login")  //防止灌水，判断表单确实是自己的表单
{
	$username=$_POST["username"];
	$userpwd=$_POST["userpwd"];
	echo("您的用户名是：{$username}");


	$data=array(
    "username" => $username,
    "userrole" => $userpwd,
    );

    $ch = curl_init();

    curl_setopt($ch, CURLOPT_URL, "http://47.100.79.216:3000/register");
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


 </BODY>
</HTML>
