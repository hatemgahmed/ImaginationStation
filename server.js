const express = require("express");
const mongoose = require("mongoose");
const cors = require("cors");
const dotenv = require("dotenv");
var path = require("path");
//creating app
const app = express();
app.use(express.json());

// Connect to mongo
dotenv.config();
mongoose
  .connect(
    //url here
    `mongodb+srv://${process.env.MONGO_ATLAS_USER}:${process.env.MONGO_ATLAS_PASSWORD}@imaginationstation-bucxx.gcp.mongodb.net/ACML`
  )
  .then(() => console.log("Connected to MongoDB"))
  .catch(err => {
    console.log(err);
  });

// Require Router Handlers
const notifications = require("./api/routes/notifications");
const users = require("./api/routes/users");
const bodyParser = require("body-parser");
app.use(bodyParser.json());
app.use("/api/notifications", notifications);
app.use("/api/users", users);

app.use((req, res) => {
  res.status(404).send({ err: "We can not find what you are looking for" });
});

const port = process.env.PORT || 5000;
app.listen(port, () => console.log(`Server on ${port}`));
