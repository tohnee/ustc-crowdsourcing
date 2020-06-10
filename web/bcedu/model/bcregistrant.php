<?php
    header("content-type:text/html;charset=utf-8");//
    session_start();
    
    
    $name=$_POST['user'];
    $type=$_POST['type'];

       //下面判断信息是不是输入完整
    if(empty($name)||empty($type)){
        echo "<script>alert('信息输入没完整');</script>";
    }
    else{  
            //通过php连接到服务器
            $data=array(
            "username" => "$name",
            "userrole" => "$type"
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
            
            if(! $output == "注册成功" ){
                echo "恭喜你注册成功\n";
            }else{
                echo "注册失败，请重新注册"
            }
?>