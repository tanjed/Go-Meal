package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Meal struct {
	Id    int    `json:"meal_id"`
	MemberId   int    `json:"member_id"`
	Count int    `json:"meal_count"`
	Date  string `json:"meal_date"`
}

var meals []Meal

func GetMeals(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET ALL MEALS")
	w.Header().Set("Content-Type", "application/json")
	if meals == nil {
		json.NewEncoder(w).Encode(noDataResponse())
		return
	}
	json.NewEncoder(w).Encode(meals)
}

func StoreMeal(w http.ResponseWriter, r *http.Request) {

}

func UpdateMeal(w http.ResponseWriter, r *http.Request) {

}

func DeleteMeal(w http.ResponseWriter, r *http.Request) {

}

func noDataResponse() map[string]interface{}{
	response := make(map[string]interface{})
	response["message"] = "No Meal Found"
	response["status"] = 404
	return response
}
