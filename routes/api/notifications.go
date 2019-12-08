package api

import (
	"Impact/server/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	// "github.com/api/app/model"
	// "github.com/gorilla/mux"
	// "github.com/jinzhu/gorm"
)

//GetAllNotifications ...
func GetAllNotifications(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	Notifications := []models.Notification{}
	db.Find(&Notifications)
	respondJSON(w, http.StatusOK, Notifications)
}

// getNotificationOr404 gets a notification instance if exists, or respond the 404 error otherwise
func getNotificationOr404(db *gorm.DB, id string, w http.ResponseWriter, r *http.Request) *models.Notification {
	notification := models.Notification{}
	if err := db.First(&notification, models.Notification{Title: id}).Error; err != nil { //VERY IMPORTANT FIND BY ID
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &notification
}


//GetNotification ...
func GetNotification(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	notification := getNotificationOr404(db, id, w, r)
	if notification == nil {
		return
	}
	respondJSON(w, http.StatusOK, notification)
}

//CreateNotification ...
func CreateNotification(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	Notification := models.Notification{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&Notification); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&Notification).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, Notification)
}

