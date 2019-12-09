const Content = require("../../models/notifications");
const mongoose = require("mongoose");
const Users = require("../../models/users");

exports.allUsers = async function(req, res) {
  Users.find({}, (err, users) => {
    if (!err) {
      res.json(users:users});
    }
  });
};
exports.createUsers = async function(req, res) {
  Users.create(
    { username: req.body.username, password: req.body.password },
    (err, user) => {
      if (err) {
        console.log(err);
      } else {
        res.send(user);
      }
    }
  );
};
