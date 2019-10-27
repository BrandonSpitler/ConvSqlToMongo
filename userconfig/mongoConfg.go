package userconfig

import (
	"encoding/json"
	"fmt"
	"os"
)

type MongoDBConfig struct {
	ConnURL string `json:",omitempty"`
	URL     string `json:"MONGO_URL"`
	Port    string `json:"MONGO_PORT"`
	DB      string `json:"MONGO_DB"`
}

func GetMongoConfig(filepath string) MongoDBConfig {
	mongo := new(MongoDBConfig)
	configFileReader, err := os.Open(filepath)
	defer configFileReader.Close()
	if err != nil {
		panic(fmt.Sprintf("Could not find config file %s", filepath))
	}
	err = json.NewDecoder(configFileReader).Decode(mongo)
	if err != nil {
		panic(fmt.Sprintf("Could not load the config file @%s into SqlConfig struct", filepath))
	}
	mongo.ConnURL = mongo.createConn()
	return *mongo
}
func (m MongoDBConfig) createConn() string {
	return fmt.Sprintf("mongodb://%s:%s", m.URL, m.Port)
}
