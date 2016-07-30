package infrastructure

import (
	"database/sql"

	"github.com/GerardSoleCa/PubKeyManager/interfaces"
	_ "github.com/xeodou/go-sqlcipher"
)

type SqliteHandler struct {
	Conn *sql.DB
}

func (handler *SqliteHandler) Execute(statement string, args ...interface{}) (interfaces.Result, error) {
	result, err := handler.Conn.Exec(statement, args...)
	if err != nil {
		return new(SqliteResult), err
	}
	r := new(SqliteResult)
	r.Result = result
	return r, nil
}

func (handler *SqliteHandler) Query(statement string, args ...interface{}) (interfaces.Row, error) {
	rows, err := handler.Conn.Query(statement, args...)
	if err != nil {
		return new(SqliteRow), err
	}
	row := new(SqliteRow)
	row.Rows = rows
	return row, nil
}

type SqliteResult struct {
	Result sql.Result
}

func (r SqliteResult) LastInsertId() (int64, error) {
	return r.Result.LastInsertId()
}

func (r SqliteResult) RowsAffected() (int64, error) {
	return r.Result.RowsAffected()
}

type SqliteRow struct {
	Rows *sql.Rows
}

func (r SqliteRow) Scan(dest ...interface{}) {
	r.Rows.Scan(dest...)
}

func (r SqliteRow) Next() bool {
	return r.Rows.Next()
}
func (r SqliteRow) Close() {
	r.Rows.Close()
}

func NewSqliteHandler(dbfileName string, password string) *SqliteHandler {
	conn, _ := sql.Open("sqlite3", dbfileName)

	p := "PRAGMA key = '" + password + "';"
	_, err := conn.Exec(p)
	if err != nil {
		panic(err)
	}

	sqliteHandler := new(SqliteHandler)
	sqliteHandler.Conn = conn
	return sqliteHandler
}
