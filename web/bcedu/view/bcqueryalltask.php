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
    $data=array(
 //   "taskname" => "haha"
    );

    $ch = curl_init();

    curl_setopt($ch, CURLOPT_URL, "http://47.100.79.216:3000/queryalltask");
    curl_setopt($ch, CURLOPT_POST, 1);
    //The number of seconds to wait while trying to connect. Use 0 to wait indefinitely.
    curl_setopt($ch, CURLOPT_CONNECTTIMEOUT, 60);
    curl_setopt($ch, CURLOPT_POSTFIELDS , http_build_query($data));
    curl_setopt($ch, CURLOPT_RETURNTRANSFER, 1);

    $output = curl_exec($ch);

    echo $output;

    curl_close($ch);
?>


 </BODY>
</HTML>