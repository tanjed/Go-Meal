package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	// "github.com/gorilla/mux"
)

type Meal struct {
	Id       int    `json:"meal_id"`
	MemberId int    `json:"member_id"`
	Count    int    `json:"meal_count"`
	Date     string `json:"meal_date"`
}

var meals []Meal

func GetMeals(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	dateFrom := r.URL.Query().Get("dateFrom")
	dateTo := r.URL.Query().Get("dateTo")

	if (len(dateFrom) == 0) || (len(dateTo) == 0) {
		json.NewEncoder(w).Encode(generateResponse("Date Params Required", http.StatusUnprocessableEntity))
		return
	}
	if meals == nil {
		json.NewEncoder(w).Encode(generateResponse("No Meal Found", http.StatusNotFound))
		return
	}
	fmt.Println(dateFrom, dateTo)
	json.NewEncoder(w).Encode(meals)
}

func StoreMeal(w http.ResponseWriter, r *http.Request) {

}

func UpdateMeal(w http.ResponseWriter, r *http.Request) {

}

func DeleteMeal(w http.ResponseWriter, r *http.Request) {

}

func generateResponse(message string, status int) map[string]interface{} {
	response := make(map[string]interface{})
	response["message"] = message
	response["status"] = status
	return response
}
