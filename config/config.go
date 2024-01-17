package config

import (
	"os"
	"log"
	"fmt"
	"gopkg.in/yaml.v2"
)

func GetDSN() string {

	configFile, err := os.Open("config.yaml")
	
	if err != nil {
		log.Fatal("Error opening config file:", err)
	}

	defer configFile.Close()

	var myConfig map[string]map[string]string

	decoder := yaml.NewDecoder(configFile)
	

	err = decoder.Decode(&myConfig)

	if err != nil {
		log.Fatal("Error decoding config file :",err)
	}

	dsn := fmt.Sprintf(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		myConfig["database"]["username"],
		myConfig["database"]["password"],
		myConfig["database"]["host"],
		myConfig["database"]["port"],
		myConfig["database"]["dbname"]))

	return dsn
}