package helpers

import (
	"math"
	"models"
)

func GetCPUCount(servercost models.ServerCost, hours int, price float64, region string) ((map[string]int), float64, int){

	var servers = make(map[string]int)
	var totalCPUs int
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
	}

	tenxcpus := int(getCPUs(tenxcost, price, hours))
	if price >= 0  && tenxcpus != 0{
			servers["10xlarge"] = tenxcpus
			totalCost = totalCost + (float64(tenxcpus) * tenxcost )
			price = price - (float64(tenxcpus) * tenxcost)
			totalCPUs = totalCPUs + tenxcpus * 10
	}
	eightxcpus := int(getCPUs(eightxcost, price, hours))
	if price >= 0 && eightxcpus !=0{
			servers["8xlarge"] = eightxcpus
			totalCost = totalCost + (float64(eightxcpus) * eightxcost )
			price = price - (float64(eightxcpus) * eightxcost)
			totalCPUs = totalCPUs + eightxcpus * 8
	}

	fourxcpus := int(getCPUs(fourxcost, price, hours))
	if price >= 0 && fourxcpus!=0 {
		servers["4xlarge"] = fourxcpus
		totalCost = totalCost + (float64(fourxcpus) * fourxcost )
		price = price - (float64(fourxcpus) * fourxcost)
		totalCPUs = totalCPUs + fourxcpus * 4
	}

	twoxcpus := int(getCPUs(twoxcost, price, hours))
	if price >= 0 && twoxcpus != 0{
		servers["2xlarge"] = twoxcpus
		totalCost = totalCost + (float64(twoxcpus) * twoxcost )
		price = price - (float64(twoxcpus) * twoxcost)
		totalCPUs = totalCPUs + twoxcpus * 2
	}

	if( region == "east"){
		xcpus := int(getCPUs(xcost, price,hours))
		if(price >= 0 && xcpus != 0){
			servers["xlarge"] = xcpus
			totalCost = totalCost + (float64(xcpus) * xcost )
			price = price - (float64(xcpus) * xcost)
			totalCPUs = totalCPUs + xcpus * 1
		}
	}

	cpus := int(getCPUs(cost, price, hours))
	if price >= 0 && cpus!=0 {
		servers["large"] = cpus
		totalCost = totalCost + (float64(cpus) * cost )
		price = price - (float64(cpus) * cost)
		totalCPUs = totalCPUs + cpus * 1
	}
	return servers , totalCost, totalCPUs
}

func getCPUs(cpucost float64, price float64, hours int) float64{
	if price < 0{
		return 0
	}
	cpumod := math.Mod(price, cpucost*float64(hours))
	cpus := int(price / (cpucost* float64(hours)))
	if cpumod < cpucost {
		return float64(cpus)
	}
	return float64(cpus) + getCPUs(cpumod, cpucost, hours)
}
