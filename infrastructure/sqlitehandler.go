package infrastructure

import (
	"database/sql"

	"github.com/GerardSoleCa/PubKeyManager/interfaces"
	// Blank import for go-sqlcipher
	_ "github.com/xeodou/go-sqlcipher"
)

// SqliteHandler struct
type SqliteHandler struct {
	Conn *sql.DB
}

// Execute function contained on SqliteHandler
func (handler *SqliteHandler) Execute(statement string, args ...interface{}) (interfaces.Result, error) {
	result, err := handler.Conn.Exec(statement, args...)
	if err != nil {
		return new(SqliteResult), err
	}
	r := new(SqliteResult)
	r.Result = result
	return r, nil
}

// Query function contained on SqliteHandler
func (handler *SqliteHandler) Query(statement string, args ...interface{}) (interfaces.Row, error) {
	rows, err := handler.Conn.Query(statement, args...)
	if err != nil {
		return new(SqliteRow), err
	}
	row := new(SqliteRow)
	row.Rows = rows
	return row, nil
}

// SqliteResult struct
type SqliteResult struct {
	Result sql.Result
}

// LastInsertId function contained on SqliteResult
func (r SqliteResult) LastInsertId() (int64, error) {
	return r.Result.LastInsertId()
}

// RowsAffected function contained on SqliteResult
func (r SqliteResult) RowsAffected() (int64, error) {
	return r.Result.RowsAffected()
}

// SqliteRow struct
type SqliteRow struct {
	Rows *sql.Rows
}

// Scan function contained on SqliteResult
func (r SqliteRow) Scan(dest ...interface{}) {
	r.Rows.Scan(dest...)
}

// Next function contained on SqliteRow
func (r SqliteRow) Next() bool {
	return r.Rows.Next()
}
// Close function contained on SqliteRow
func (r SqliteRow) Close() {
	r.Rows.Close()
}

// NewSqliteHanlder creates a new Ciphered SqliteHandler
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
