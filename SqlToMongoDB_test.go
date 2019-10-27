package ConvSQLToMongoDB

import (
	"ConvSqlToMongo/userconfig"
	"database/sql"
	"testing"

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
	if err != nil {
		t.Fatal("could not connect to sql db", sqlConnStr)
	}
	defer sqlDB.Close()
}

func TestMongo(t *testing.T) {

}

func TestSQLToMongo(t *testing.T) {

}
