const mongoose = require("mongoose");

var notificationsSchema = new mongoose.Schema({
  title: { type: String },
  date: { type: Date },
  content: { type: String }
});
mongoose.model("Notifications", notificationsSchema);

module.exports = mongoose.model("Notifications");
