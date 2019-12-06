<?php

$connection=mysqli_connect('localhost','root','','ACML');
for($i=1;$i<24;$i++){
    $query='insert into Notifications
    (title,content,notificationDate) values
    ("title'.$i.'", "content' . $i . '",CURRENT_TIMESTAMP);';
   
    if($connection->query($query)){
        echo 'notificationDate created successfully';
    }
    else
        echo 'Failed to insert ' . mysqli_error($connection);
    echo '</br>';
}
for($i=1;$i<24;$i++){
    $query='insert into Users
    (username,password,admin) values
    ("username'.$i.'", "password' . $i . '",1);';
   
    if($connection->query($query)){
        echo 'Users created successfully';
    }
    else
        echo 'Failed to insert ' . mysqli_error($connection);
    echo '</br>';
}
 mysqli_close($connection);
?>