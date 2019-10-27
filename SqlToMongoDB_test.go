package ConvSQLToMongoDB

import (
	"fmt"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

var (
	sqlURL, sqlPort, sqlUsername, sqlPWD         string
	mongoURL, mongoPort, mongoUsername, mongoPWD string
)

func initSQL() {
	//test
	// sqlPort = getENVElsePanic("MYSQL_PORT")
	// sqlUsername = getENVElsePanic("MYSQL_USER")
	// sqlPWD = getENVElsePanic("MYSQL_PWD")
}

func init() {
	// sqlURL :=
	// mongoURL :=
	// sql
	// sqlPwd :=
	// mongoPwd :=
	// initSql()
	sqlURL = getENVElsePanic("MYSQL_URL")
	fmt.Println(sqlURL)
	// 	initMongo()
}

// func initMongo() {
// 	mongoURL = getENVElsePanic("MYSQL_URL")
// 	sqlPort = getENVElsePanic("MYSQL_PORT")
// 	mongoPort = getENVElsePanic("MYSQL_USER")
// 	mongoPWD = getENVElsePanic("MYSQL_PWD")
// }

func getENVElsePanic(env string) string {
	returnEnvVar, exists := os.LookupEnv(env)
	if !exists {
		fmt.Println("Variable does not exists")
		panic("variable does not exist")
	}
	return returnEnvVar
}

func TestSQL(t *testing.T) {

}

func TESTMongo(t *testing.T) {

}

func TestSQLToMongo(t *testing.T) {

}
