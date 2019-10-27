package userconfig

import (
	"encoding/json"
	"os"
	"reflect"
	"testing"
)

func CreateTestFile(t *testing.T, testFileCnfgFilePath string, writeConf interface{}) {

	//create test file

	testConfFile, err := os.Create(testFileCnfgFilePath)
	defer testConfFile.Sync()
	defer testConfFile.Close()
	if err != nil {
		t.Fatal("could not create testconfig.json for testing file config")
	}
	json.NewEncoder(testConfFile).Encode(&writeConf)
}

func TestSQLConfigStruct(t *testing.T) {
	testSQLConfigToFile := SQLConfig{
		SQLURL:      "localhost",
		SQLPort:     "3000",
		SQLUsername: "testUserName",
		SQLPWD:      "testPWD",
		SQLDB:       "testDB",
	}
	testFileCnfgFilePath := writeConfigToTestFile(t, testSQLConfigToFile)
	defer os.Remove(testFileCnfgFilePath)
	testSQLConfigFromFile := GetSQLConfig(testFileCnfgFilePath)
	if !reflect.DeepEqual(testSQLConfigFromFile, testSQLConfigToFile) {
		t.Fatal("user config sql structs were not equal ")
	}
}

func TestMongoConfigStruct(t *testing.T) {
	testMongoToFile := MongoDBConfig{
		URL:  "localhost",
		Port: "27017",
		DB:   "personnel",
	}
	testFileCnfgFilePath := writeConfigToTestFile(t, testMongoToFile)
	// defer os.Remove(testFileCnfgFilePath)
	testMongoFromFile := GetMongoConfig(testFileCnfgFilePath)
	testMongoToFile.ConnURL = testMongoToFile.createConn()

	if !reflect.DeepEqual(testMongoToFile, testMongoFromFile) {
		t.Fatal("user mongdb config structs were not equal ")
	}
}

func writeConfigToTestFile(t *testing.T, write interface{}) string {
	testFileCnfgFilePath := "testconfig.json"
	CreateTestFile(t, testFileCnfgFilePath, write)
	return testFileCnfgFilePath
}
