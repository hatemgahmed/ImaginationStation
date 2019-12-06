<?php
$connection=new mysqli('localhost','root','');
if($connection->connect_error)
    die("Failed to connect" . $connection->connect_error);
else
    echo 'Connected';
echo '</br>';
if($connection->query("CREATE DATABASE ACML")==true)
    echo 'Database created';
else
    echo 'Database not created ' . $connection->error;
echo '</br>';
$connection->close();

//creating table
$connection=mysqli_connect('localhost','root','','ACML');
if($connection->query("CREATE TABLE Users(
    username varchar(20) PRIMARY KEY,
    password varchar(15),
    admin bool not null
)"))
    echo 'Users table created\n';
else{
    echo 'Users Table not created \n' .  $connection->error;
}
if($connection->query("CREATE TABLE Notifications(
    id int(8) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    title varchar(50),
    notificationDate date,
    content varchar(4096)
)"))
    echo 'Notifications table created\n';
 else{
     echo 'Notifications Table not created \n' .  $connection->error;
 }

$connection->close();
?>