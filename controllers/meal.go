package controllers

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	util "tanjed.me/mealmanagement/utils"
)

type Meal struct {
	Id       int    `json:"meal_id"`
	MemberId int    `json:"member_id"`
	Count    int    `json:"meal_count"`
	Date     string `json:"meal_date"`
}

var mealsContainer []Meal

func GetMeals(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	dateFrom := r.URL.Query().Get("dateFrom")
	dateTo := r.URL.Query().Get("dateTo")

	if (len(dateFrom) == 0) || (len(dateTo) == 0) {
		json.NewEncoder(w).Encode(util.GenerateApiResponse("Date Params Required", http.StatusUnprocessableEntity))
		return
	}
	if mealsContainer == nil {
		json.NewEncoder(w).Encode(util.GenerateApiResponse("No Meal Found", http.StatusNotFound))
		return
	}
	var filteredMeals []Meal
	for _, meal := range mealsContainer {
		if meal.Date >= dateFrom && meal.Date <= dateFrom {
			filteredMeals = append(filteredMeals, meal)
		}
	}
	if len(filteredMeals) <= 0 {
		json.NewEncoder(w).Encode(util.GenerateApiResponse("No Meal Found",http.StatusNotFound))
		return	
	}
	json.NewEncoder(w).Encode(filteredMeals)
	return
}

func StoreMeal(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var meals []Meal
	json.NewDecoder(r.Body).Decode(&meals)
	rand.Seed(time.Now().UnixNano())
	for _, meal := range meals {
		meal.Id = rand.Intn(100)
		mealsContainer = append(mealsContainer, meal)
	}
	json.NewEncoder(w).Encode(meals)
	return
}

func UpdateMeal(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	mealId, _ := strconv.Atoi(params["id"])
	for index, meal := range mealsContainer {
		if meal.Id == mealId {
			mealsContainer = append(mealsContainer[:index], mealsContainer[index+1:]...)
			var meal Meal
			json.NewDecoder(r.Body).Decode(&meal)
			meal.Id = mealId
			mealsContainer = append(mealsContainer, meal)
		}
	}

	json.NewEncoder(w).Encode(util.GenerateApiResponse("Updated Successfully", 200))
	return
}

func DeleteMeal(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	mealId, _ := strconv.Atoi(params["id"])
	for index, meal := range mealsContainer {
		if meal.Id == mealId {
			mealsContainer = append(mealsContainer[:index], mealsContainer[index+1:]...)
		}
	}

	json.NewEncoder(w).Encode(util.GenerateApiResponse("Deleted Successfully", 200))
	return
}
