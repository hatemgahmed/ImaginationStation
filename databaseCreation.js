// var mysql = require("mysql");

// var con = mysql.createConnection({
//   host: "localhost",
//   user: "root",
//   password: "",
//   database: "ACML"
// });

// con.connect(function(err) {
//   if (err) throw err;
//   console.log("Connected!");
//   con.query("DROP DATABASE ACML", function(err, result) {
//     if (err) throw err;
//     console.log("Database Created");
//   });
// });
// con = mysql.createConnection({
//   host: "localhost",
//   user: "root",
//   password: ""
// });
// con.connect(function(err) {
//   var sql =
//     "CREATE TABLE Users(username varchar(20) PRIMARY KEY,password varchar(15),admin bool not null)";
//   con.query(sql, function(err, result) {
//     if (err) throw err;
//     console.log("Users created");
//   });
//   sql =
//     "CREATE TABLE Notifications(id int(8) UNSIGNED AUTO_INCREMENT PRIMARY KEY,title varchar(50),notificationDate date,content varchar(4096))";
//   con.query(sql, function(err, result) {
//     if (err) throw err;
//     console.log("Notifications created");
//   });
// for (var i = 0; i < 24; i++) {
//   sql =
//     'insert into Notifications(title,content,notificationDate) values("title' +
//     i +
//     '", "content' +
//     i +
//     '",CURDATE());';
//   con.query(sql, function(err, result) {
//     if (err) throw err;
//     console.log("Inserted Notification" + i);
//   });
// }
// });
