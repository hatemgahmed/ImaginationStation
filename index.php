<?php
if( $_GET["request"] ) {
    if($_GET['request']=='Users')
        include 'getUsers.php';
    if($_GET['request']=='Notifications')
        include 'getNotifications.php';
    exit();
 }
//  if( $_POST["request"])

?>