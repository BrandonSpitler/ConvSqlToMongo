package ConvSQLToMongoDB

import (
	"ConvSqlToMongo/sqldbhandle"
	"ConvSqlToMongo/userconfig"
	"database/sql"
	"fmt"
	"reflect"
	"sync"

	"gopkg.in/mgo.v2/bson"

	"gopkg.in/mgo.v2"
)

func rowBoilerPlateFromColType(colTypes []*sql.ColumnType) []interface{} {
	rowBoilerPlate := make([]interface{}, len(colTypes))
	for i, columnType := range colTypes {
		switch columnType.DatabaseTypeName() {
		case "VARCHAR":
			vtype := reflect.TypeOf("")
			fmt.Println(vtype)
			rowBoilerPlate[i] = reflect.New(reflect.TypeOf("")).Interface()

		default:
			rowBoilerPlate[i] = reflect.New(columnType.ScanType()).Interface()
		}
	}
	return rowBoilerPlate
}

func CopySQLTOMongo(sqlDB *sql.DB,
	mongoSession *mgo.Session,
	mongoDBConf userconfig.MongoDBConfig,
	sqlConfig userconfig.SQLConfig) {

	var wgSQLWriter sync.WaitGroup
	tableNames := sqldbhandle.GetTables(sqlDB, sqlConfig)
	wgSQLWriter.Add(len(tableNames))
	for _, tableName := range tableNames {

		go func(tableName string, mongoSession *mgo.Session, mongoDB string) {
			collection := mongoSession.DB(mongoDBConf.DB).C(tableName)
			// mongoCollection := mongoDataBase.C(tableName)
			sqlQuery := fmt.Sprintf("SELECT * FROM %s", tableName)
			sqlTable, err := sqlDB.Query(sqlQuery)
			defer sqlTable.Close()
			if err != nil {
				panic("Query did not work")
			}
			colTypes, _ := sqlTable.ColumnTypes()
			for sqlTable.Next() {
				var mongoDoc bson.D
				row := rowBoilerPlateFromColType(colTypes)
				err := sqlTable.Scan(row...)
				if err != nil {
					panic("faild to load row values ")
				}

				for i, columnName := range colTypes {
					value := reflect.ValueOf(row[i]).Elem().Interface()
					mongoDoc = append(mongoDoc, bson.DocElem{Name: columnName.Name(), Value: value})
				}
				err = collection.Insert(mongoDoc)
				if err != nil {
					panic("Query did not work")
				}
			}
			wgSQLWriter.Done()
		}(tableName, mongoSession.Copy(), mongoDBConf.DB)

	}
	wgSQLWriter.Wait()
}
