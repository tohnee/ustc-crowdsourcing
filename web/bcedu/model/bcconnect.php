<?php	
	if(isset($_COOKIE["user"])
	{
		$ser=$_COOKIE["user"];

		$query="select user from registrants where user='$user' and pass='$pass'";
		$result=$conn->query($query);
		$check1=$result->fetch_row();
	}
	else if (isset($_POST["user"]))
	{
		$user=$_POST["user"];
		$query="select user from registrants where user='$user' and pass='$pass'";
		$result=$conn->query($query);
		$check2=$result->fetch_row();
	}

?>