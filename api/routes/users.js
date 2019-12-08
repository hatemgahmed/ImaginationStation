const express = require("express");
const router = express.Router();
const usersController = require("../controllers/users");

// get all content (with db)
router.get("/getAllUsers", usersController.allUsers);
// crearting new content (with db)
router.post("/getAllUsers", usersController.createUsers);

module.exports = router;
