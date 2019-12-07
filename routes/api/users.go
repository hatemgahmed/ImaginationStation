package api

import (
	"Impact/server/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	// "github.com/api/app/model"
	// "github.com/gorilla/mux"
	// "github.com/jinzhu/gorm"
)

// getUserOr404 gets a user instance if exists, or respond the 404 error otherwise
func getUserOr404(db *gorm.DB, id string, w http.ResponseWriter, r *http.Request) *models.User {
	user := models.User{}
	if err := db.First(&user, models.User{Username: id}).Error; err != nil { //VERY IMPORTANT FIND BY ID
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &user
}

//GetUser ...
func GetUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	user := getUserOr404(db, id, w, r)
	if user == nil {
		return
	}
	respondJSON(w, http.StatusOK, user)
}

//GetAllUsers ...
func GetAllUsers(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	Users := []models.User{}
	db.Find(&Users)
	respondJSON(w, http.StatusOK, Users)
}

//CreateUser ...
func CreateUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	fmt.Println(user)
	fmt.Println("between")
	decoder := json.NewDecoder(r.Body)
	fmt.Println(decoder)
	if err := decoder.Decode(&user); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()

	if err := db.Save(&user).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, user)
}
