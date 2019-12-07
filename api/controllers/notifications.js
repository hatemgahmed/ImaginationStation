const Content = require("../../models/notifications");
const mongoose = require("mongoose");
const Notifications = require("../../models/notifications");

exports.allNotifications = async function(req, res) {
  Notifications.find({}, (err, notifications) => {
    if (!err) {
      res.json(notifications);
    }
  });
};
exports.createNotifications = async function(req, res) {
  Notifications.create(
    { title: req.body.title, date: req.body.date, content: req.body.content },
    (err, notification) => {
      if (err) {
        console.log(err);
      } else {
        res.json(notification);
      }
    }
  );
};
