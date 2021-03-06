package go_db
import (
	"database/sql"
	"strings"
	"fmt"
)

func ExecLog(db *sql.DB, sqlStr string, args ...interface{}) (sql.Result,string,error) {
	stmt,err := db.Prepare(sqlStr)
	defer stmt.Close()
	if err != nil {
		return nil,sqlStr,err
	}
	res,err := stmt.Exec(args...)

	return res, SqlLog(sqlStr,args...), err
}

func QueryLog(db *sql.DB, sqlStr string, args ...interface{}) (*sql.Rows, string, error) {
	rows,err := db.Query(sqlStr, args...)
	return rows, SqlLog(sqlStr,args...), err
}

func SqlLog(sqlStr string, args ...interface{}) string {
	for _, e := range args {
		sqlStr = strings.Replace(sqlStr, "?", "'"+fmt.Sprintf("%v",e)+"'", 1)
	}
	return sqlStr
}