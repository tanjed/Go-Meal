package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	controllers "tanjed.me/mealmanagement/controllers"
)


func main(){
	bootRoutes()
}



func bootRoutes()  {
	r := mux.NewRouter()
	r.HandleFunc("/meals",controllers.GetMeals).Methods("GET")
	r.HandleFunc("/meal",controllers.StoreMeal).Methods("POST")
	r.HandleFunc("/meal/{id}",controllers.UpdateMeal).Methods("PUT")
	r.HandleFunc("/meal/{id}",controllers.DeleteMeal).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":4000",r))

}