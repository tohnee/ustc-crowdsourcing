<?php
	error_reporting(E_ALL ^ E_NOTICE);
    // enable sessions
	session_start();
?>
<!DOCTYPE html>
<html>
	<head>
	<title>量子教育-学员主页</title>
	<meta name=\"description\" content=\"怎样快速地找到合适的老师？只需花费3分钟，在量子教育网发布一份任务，写清楚您的需求，就可以获得免费推荐老师，自动化评估效果\">
	<meta name=\"keywords\" content=\"学习，知识\">
<?php
include(V.'header-content.php');
?>

	</head>
	<body>
<?php
//导航条
include(V.'top-bar.php');
?>

	<?php if($_SESSION["appointed"]==true) { ?>

		<div class="alert alert-success" role="alert">
			您已预约~  马上为您安排，请耐心等待。         
		</div>
	<?php } else { ?>

		<div class="alert alert-warning" role="alert">
			填写老师预约单         
			<a href="?page=appoint"><strong>【填写】</strong><span class="glyphicon glyphicon-edit"></span></a>
		</div>
	<?php } ?>

<?php
//页尾
include(V.'panel-tutor.php');
include(V.'footer-content.php');
?>

	</body>
</html>