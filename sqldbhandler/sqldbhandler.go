package sqldbhandler

import (
	"ConvSqlToMongo/userconfig"
	"database/sql"
	"log"
)

type CustSQLDB struct {
	*sql.DB
}

func (db *CustSQLDB) GetTables() []string {
	var columnName []string
	var tableName string
	findTableQuery := "SELECT TABLE_NAME " +
		" FROM INFORMATION_SCHEMA.TABLES " +
		" WHERE TABLE_TYPE = 'BASE TABLE'" +
		" AND TABLE_SCHEMA='dbName' "
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
