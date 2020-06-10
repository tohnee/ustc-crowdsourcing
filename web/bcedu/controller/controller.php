<?php
//determine which page to show
if (isset($_GET['page']))
    $page = $_GET['page'];
else
    $page = 'home';

// show page
switch ($page)
{
    case 'home':
		include(V."home.php");
        break;

    case 'student':
		include(V."student.php");
        break;

    case 'tutor':
		include(V."tutor.php");
        break;

    case 'contact':
		include(V."contact.php");
        break;

    case 'about':
		include(V."about.php");
        break;

    case 'login':
		include(V."login.php");
        break;

    case 'logout':
		include(V."logout.php");
        break;

    case 'register':
		include(V."register.php");
        break;

    case 'registered':
		include(V."registered.php");
        break;

    case 'myinfo':
		include(V."myinfo.php");
        break;
    case 'modify':
		include(V."modify.php");
        break;

    case 'update':
		include(M."update.php");
        break;

    case 'appoint':
		include(V."appoint.php");
        break;

    case 'appointed':
		include(V."appointed.php");
        break;
    case 'connect':
		include(M."connect.php");
        break;

	case 'more':
		include(V."more.php");
        break;

    case 'bcregister':
        include(V."bcregister.php");
        break;  

    case 'bclogin':
        include(V."bclogin.php");
        break;    

    case 'bcposttask':
        include(V."bcposttask.php");
        break;    

    case 'bctutor':
        include(V."bctutor.php");
        break;    

    case 'bcstudent':
        include(V."bcstudent.php");
        break;    

    case 'bchome':
        include(V."bchome.php");
        break;  

    case 'bccontact':
        include(V."bccontact.php");
        break;  

    case 'bcpostoffer':
        include(V."bcpostoffer.php");
        break;

    case 'bcassigntask':
        include(V."bcassigntask.php");
        break;

    case 'bcbonuspay':
        include(V."bcbonuspay.php");
        break;

    case 'bcqueryuser':
        include(V."bcqueryuser.php");
        break;

    case 'bcquerytask':
        include(V."bcquerytask.php");
        break;

    case 'bcqueryalltask':
        include(V."bcqueryalltask.php");
        break;

    case 'bcqueryusertask':
        include(V."bcqueryusertask.php");
        break;

    case 'bcqueryworkeroffers':
        include(V."bcqueryworkeroffers.php");
        break;

    case 'bcqueryworkeroffer':
        include(V."bcqueryworkeroffer.php");
        break;

    case 'bcqueryassignresult':
        include(V."bcqueryassignresult.php");
        break;
    case 'bcqueryalltaskT':
        include(V."bcqueryalltaskT.php");
        break;

}


?>