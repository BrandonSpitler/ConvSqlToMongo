package sqldbhandle

import (
	"ConvSqlToMongo/userconfig"
	"database/sql"
	"fmt"
	"log"
)

func GetTables(db *sql.DB, sqlConfig userconfig.SQLConfig) []string {
	var columnName []string
	var tableName string
	findTableQuery := fmt.Sprintf("SELECT TABLE_NAME "+
		" FROM INFORMATION_SCHEMA.TABLES "+
		" WHERE TABLE_TYPE = 'BASE TABLE'"+
		" AND TABLE_SCHEMA = '%s'", sqlConfig.SQLDB)
	tableNameRows, err := db.Query(findTableQuery)
	if err != nil {
		log.Fatal("Could not find table columns")
	}

	for tableNameRows.Next() {
		err := tableNameRows.Scan(&tableName)
		if err != nil {
			log.Fatal("could not scan all table names")
		}
		columnName = append(columnName, tableName)
	}
	if err := tableNameRows.Err(); err != nil {
		log.Fatal("error ", err)
	}
	return columnName
}

func init() {
	userconfig.CreateSQLConnString("sqlConfig.json")
}
