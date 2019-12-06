<?php
 $connection=mysqli_connect('localhost','root','','ACML');
        
//  $query="Select * from Users;";
  
//  $outputString="";

//  if($result=mysqli_query($connection,$query)){
     
//     $rows = array();
//     while($r = mysqli_fetch_assoc($result)) {
//         $rows[] = $r;
//     }
//     print '{ "Users": ' . json_encode($rows) . '}';
//  }
//  else{
//      echo 'Query failed';
//  }
 $query="Select * from Notifications;";
  
 $outputString="";

 if($result=mysqli_query($connection,$query)){
     
    $rows = array();
    while($r = mysqli_fetch_assoc($result)) {
        $rows[] = $r;
    }
    print '{ "Notifications" : ' . json_encode($rows) .'}';
 }
 else{
     echo 'Query failed';
 }
 mysqli_close($connection);
 ?>