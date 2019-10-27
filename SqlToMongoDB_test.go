package ConvSQLToMongoDB

import (
	"ConvSqlToMongo/userconfig"
	"database/sql"
	"testing"

	"gopkg.in/mgo.v2"

	_ "github.com/go-sql-driver/mysql"
)

var (
	sqlConnStr string
)

// func init() {
// 	sqlConnStr = userconfig.SQLConnectionString
// }

func TestSQL(t *testing.T) {
	sqlConnStr = userconfig.CreateSQLConnString("sqlConfig.json")

	sqlDB, err := sql.Open("mysql", sqlConnStr)
	defer sqlDB.Close()
	if err != nil {
		t.Fatal("could not connect to sql db", sqlConnStr)
	}
}

func TestMongo(t *testing.T) {
	mgoConfig := userconfig.GetMongoConfig("mongoConfig.json")
	session, err := mgo.Dial(mgoConfig.ConnURL)
	defer session.Close()
	if err != nil {
		t.Fatal(err)
	}

}

func TestSQLToMongo(t *testing.T) {
	sqlConfig := userconfig.GetSQLConfig("sqlConfig.json")
	sqlConnStr = userconfig.CreateSQLConnString("sqlConfig.json")
	sqlDB, err := sql.Open("mysql", sqlConnStr)
	defer sqlDB.Close()
	if err != nil {
		t.Fatal("could not connect to sql db", sqlConnStr)
	}
	mgoConfig := userconfig.GetMongoConfig("mongoConfig.json")
	mgoSession, err := mgo.Dial(mgoConfig.ConnURL)
	defer mgoSession.Close()
	if err != nil {
		t.Fatal(err)
	}
	CopySQLTOMongo(sqlDB, mgoSession, mgoConfig, sqlConfig)
}
