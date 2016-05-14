package helpers

import (
	"models"
)

func GetPlan(servercost models.ServerCost, hours int, cpus int, price float64, region string)( map[string]int , float64, int) {
	var servers = make(map[string]int)
	var totalCPUS int
	numofcpus := cpus
	var tenxcost, eightxcost, fourxcost, twoxcost, xcost, cost, totalCost float64

	switch region {
	case "east":
		tenxcost = servercost.RegionCost["east"]["10xl"]
		eightxcost = servercost.RegionCost["east"]["8xl"]
		fourxcost = servercost.RegionCost["east"]["4xl"]
		twoxcost = servercost.RegionCost["east"]["2xl"]
		xcost = servercost.RegionCost["east"]["xl"]
		cost   = servercost.RegionCost["east"]["l"]

	case "west":
		tenxcost = servercost.RegionCost["west"]["10xl"]
		eightxcost = servercost.RegionCost["west"]["8xl"]
		fourxcost = servercost.RegionCost["west"]["4xl"]
		twoxcost = servercost.RegionCost["west"]["2xl"]
		cost   = servercost.RegionCost["west"]["l"]
		xcost = 0
	}

	tenx := countCPUS(cpus, 10)
	tenxCost := float64(tenx) *tenxcost * float64(hours)
	if (totalCost + tenxCost <= price || price == 0) && (tenx!=0){
			servers["10xlarge"] = tenx
			totalCPUS = totalCPUS + tenx * 10
			cpus = cpus - tenx * 10
			totalCost = totalCost + tenxCost
	}
	eightx := countCPUS(cpus, 8)
	eightxCost := float64(eightx) * eightxcost * float64(hours)
	if( totalCost + eightxCost <= price  || price == 0) && eightx!=0{
			servers["8xlarge"] = eightx
			totalCPUS = totalCPUS + eightx * 8
			cpus = cpus - eightx * 8
			totalCost = totalCost + eightxCost
	}

	fourx := countCPUS(cpus, 4)
	fourxCost := float64(fourx) * fourxcost * float64(hours)
	if( totalCost + fourxCost <= price  || price == 0) && fourx !=0 {
			servers["4xlarge"] = fourx
			totalCPUS = totalCPUS + fourx * 4
			cpus = cpus - fourx * 4
			totalCost = totalCost + fourxCost
	}

	twox := countCPUS(cpus, 2)
	twoxCost := float64(twox) * twoxcost * float64(hours)
	if( totalCost + twoxCost <= price  || price == 0) && twox !=0{
			servers["2xlarge"] = twox
			totalCPUS = totalCPUS + twox * 2
			cpus = cpus - twox * 2
			totalCost = totalCost + twoxCost
	}

	if(region == "east"){
		if ((float64(cpus) * xcost * float64(hours) + totalCost) <= price  || price == 0 ) && (cpus != 0 && totalCPUS < numofcpus){
			servers["xlarge"] = cpus
			totalCPUS = totalCPUS + cpus
			totalCost =  totalCost + (float64(cpus) * xcost) * float64(hours)
		}
	}

	if ((float64(cpus) * cost * float64(hours) + totalCost) <= price  || price == 0) && (cpus !=0 && totalCPUS < numofcpus){
			servers["large"] = cpus
			totalCPUS = totalCPUS + cpus
			totalCost = totalCost + (float64(cpus) * cost) * float64(hours)
	}

	return servers, totalCost, totalCPUS
}

func countCPUS(cpus int, cpuconfig int) int{
	cpumod := cpus % cpuconfig
	cpus = cpus / cpuconfig
	if cpumod < cpuconfig {
		return cpus
	}
	return cpus + countCPUS(cpumod, cpuconfig)
}

func getRemCPUs(cpus int, calculatedCPUs int, cpuconfig int) int {
	if cpus != 0{
		remCPUs := cpus - (calculatedCPUs * cpuconfig)
		return remCPUs
	}else{
		return calculatedCPUs
	}
}

