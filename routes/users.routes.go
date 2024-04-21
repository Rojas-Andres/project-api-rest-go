package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Rojas-Andres/project-api-rest-go/db"
	"github.com/Rojas-Andres/project-api-rest-go/models"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	db.DB.Find(&users)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&users)

}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Hello World 2"))
}

func PostUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	createdUser := db.DB.Create(&user)
	err := createdUser.Error
	fmt.Println("emtre aca")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // 400
		// json.NewEncoder(w).Encode(err)
		w.Write([]byte(err.Error()))
		return
	}
	fmt.Println("pase aca")
	// w.Write([]byte("POSTS"))
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&user)
	
}


func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Hello World 2"))
}