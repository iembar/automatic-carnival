package helpers

import (
	"github.com/spf13/viper"
	"os"
	"fmt"
	"models"
)

func InitConfig() (viper.Viper, error){
	gopath := os.Getenv("GOPATH")
	var config = viper.New()
	config.SetConfigType("yaml")
	config.SetConfigName("costconfig")
	config.AddConfigPath(gopath + "/resources")
	err := config.ReadInConfig()

	if err != nil {
		fmt.Printf("Fatal error config file: %s \n", err)
		return *config, err
	} else {
		return *config, nil
	}
}


func GetServerCosts(config viper.Viper) (models.ServerCost) {
	servers := make(map[string]map[string]float64)
	east := make(map[string]float64)
	east["l"] = config.GetFloat64("us-east.large")
	east["xl"] = config.GetFloat64("us-east.xlarge")
	east["2xl"] = config.GetFloat64("us-east.2xlarge")
	east["4xl"] = config.GetFloat64("us-east.4xlarge")
	east["8xl"] = config.GetFloat64("us-east.8xlarge")
	east["10xl"] = config.GetFloat64("us-east.10xlarge")

	west := make(map[string]float64)
	west["l"] = config.GetFloat64("us-west.large")
	west["2xl"] = config.GetFloat64("us-west.2xlarge")
	west["4xl"] = config.GetFloat64("us-west.4xlarge")
	west["8xl"] = config.GetFloat64("us-west.8xlarge")
	west["10xl"] = config.GetFloat64("us-west.10xlarge")

	servers["east"] = east
	servers["west"] = west
	return models.ServerCost{servers}
}