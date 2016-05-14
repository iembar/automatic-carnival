package handlers

import (
	"net/http"
	"github.com/gorilla/mux"
	"helpers"
	"encoding/json"
	"models"
	"strconv"
)

func CalculateCost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	hours, _ :=  strconv.Atoi(vars["hrs"])
	cpus, _ := strconv.Atoi(vars["cpus"])
	price, _ := strconv.ParseFloat(vars["price"], 64)

	config, err := helpers.InitConfig()
	if err!=nil{
		response, _ := json.Marshal(&models.FatalError{false,503,"Technical Difficulties. Apologize!"})
		w.Write(response)
	}
	if hours <= 0{
		response, _ := json.Marshal(&models.FatalError{false,422,"hours  should be valid"})
		w.Write(response)
		return
	}

	if price <= 0 && cpus <=0{
		response, _ := json.Marshal(&models.FatalError{false,422,"price or cpus should be valid"})
		w.Write(response)
		return
	}
	if price < 0 {
		response, _ := json.Marshal(&models.FatalError{false,422,"price should be valid"})
		w.Write(response)
		return
	}

	servercost := helpers.GetServerCosts(config)
	if(cpus == 0){
		cost_details := calculatePlans(servercost, hours, cpus, price, "cpus")
		response, _ := json.Marshal(cost_details)
		w.Write(response)
		return
	}

	cost_details := calculatePlans(servercost, hours, cpus, price, "cost")
	response, _ := json.Marshal(cost_details)
	w.Write(response)
}

func calculatePlans(servercost models.ServerCost, hours int, cpus int, price float64, plan string) models.ServerResponse{
	var eastservers, westservers map[string]int
	var eastcost, westcost float64
	var ecpus, wcpus int

	if plan == "cpus"{
		eastservers, eastcost, ecpus= helpers.GetCPUCount(servercost, hours,  price, "east")
		westservers, westcost, wcpus= helpers.GetCPUCount(servercost, hours,  price, "west")
	}else if plan == "cost"{
		eastservers, eastcost, ecpus= helpers.GetPlan(servercost, hours, cpus, price, "east")
		westservers, westcost, wcpus= helpers.GetPlan(servercost, hours, cpus, price, "west")
	}
	east_details :=  models.CalculatedCost{"us-east", eastcost, ecpus, eastservers}
	west_details :=  models.CalculatedCost{"us-west", westcost,wcpus, westservers}

	cost_details :=  []models.CalculatedCost{east_details, west_details}
	return models.ServerResponse{true, 200, "OK", cost_details}
}
