const mongoose = require("mongoose");

var userSchema = new mongoose.Schema({
  username: { type: String, required: true, unique: true },
  password: { type: String }
});
mongoose.model("Users", userSchema);

module.exports = mongoose.model("Users");
