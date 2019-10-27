package userconfig

import (
	"encoding/json"
	"os"
	"reflect"
	"testing"
)

func CreateTestFile(t *testing.T, testFileCnfgFilePath string, testSQLConfig sqlConfig) {

	//create test file

	testConfFile, err := os.Create(testFileCnfgFilePath)
	defer testConfFile.Sync()
	defer testConfFile.Close()
	if err != nil {
		t.Fatal("could not create testconfig.json for testing file config")
	}
	json.NewEncoder(testConfFile).Encode(&testSQLConfig)
}

func TestSQLConfigStruct(t *testing.T) {
	testSQLConfigToFile := sqlConfig{
		SQLURL:      "localhost",
		SQLPort:     "3000",
		SQLUsername: "testUserName",
		SQLPWD:      "testPWD",
		SQLDB:       "testDB",
	}
	testFileCnfgFilePath := "testconfig.json"
	defer os.Remove(testFileCnfgFilePath)
	CreateTestFile(t, testFileCnfgFilePath, testSQLConfigToFile)
	testSQLConfigFromFile := getSQLConfig(testFileCnfgFilePath)
	if !reflect.DeepEqual(testSQLConfigFromFile, testSQLConfigToFile) {
		t.Fatal("user config structs were not equal ")
	}
}
