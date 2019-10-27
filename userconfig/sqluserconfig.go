package userconfig

import (
	"encoding/json"
	"fmt"
	"os"
)

type SQLConfig struct {
	SQLURL      string `json:"MYSQL_URL"`
	SQLPort     string `json:"MYSQL_PORT"`
	SQLUsername string `json:"MYSQL_USER"`
	SQLPWD      string `json:"MYSQL_PWD"`
	SQLDB       string `json:"MYSQL_DB"`
}

func GetSQLConfig(filepath string) SQLConfig {
	sql := new(SQLConfig)
	configFileReader, err := os.Open(filepath)
	defer configFileReader.Close()
	if err != nil {
		panic(fmt.Sprintf("Could not find config file %s", filepath))
	}
	err = json.NewDecoder(configFileReader).Decode(sql)
	if err != nil {
		panic(fmt.Sprintf("Could not load the config file @%s into SqlConfig struct", filepath))
	}
	return *sql
}

func (s SQLConfig) createSQLURL() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", s.SQLUsername, s.SQLPWD, s.SQLURL, s.SQLPort, s.SQLDB)
}

func CreateSQLConnString(filepath string) string {
	sqlCnfg := GetSQLConfig(filepath)
	return sqlCnfg.createSQLURL()
}
