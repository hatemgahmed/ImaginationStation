const express = require("express");
const router = express.Router();
const notificationsController = require("../controllers/notifications");

// get all content (with db)
router.get("/getAllNotifications", notificationsController.allNotifications);
// crearting new content (with db)
router.post(
  "/getAllNotifications",
  notificationsController.createNotifications
);

module.exports = router;
