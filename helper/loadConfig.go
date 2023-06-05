package helper

import (
	"io/ioutil"
	"log"

	"github.com/gadhittana01/go-sqlc/config"
	"gopkg.in/yaml.v2"
)

func LoadConfig(c *config.GlobalConfig) {
	path := "config/sqlc-http.yaml"
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

}
