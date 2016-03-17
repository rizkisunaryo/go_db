package go_db
import (
	"database/sql"
	"strings"
	"fmt"
)

func ExecLog(db *sql.DB, sqlStr string, args ...interface{}) (sql.Result,string,error) {
	stmt,err := db.Prepare(sqlStr)
	if err != nil {
		return nil,sqlStr,err
	}
	res,err := stmt.Exec(args...)
	for _, e := range args {
		sqlStr = strings.Replace(sqlStr, "?", "'"+fmt.Sprintf("%v",e)+"'", 1)
	}
	return res,sqlStr,err
}