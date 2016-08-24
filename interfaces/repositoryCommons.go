package interfaces

// DbHandler interface for querying
type DbHandler interface {
	Execute(statement string, args ...interface{}) (Result, error)
	Query(statement string, args ...interface{}) (Row, error)
}

// Result interface for result processing
type Result interface {
	LastInsertId() (int64, error)
	RowsAffected() (int64, error)
}

// Row interface for row processing
type Row interface {
	Scan(dest ...interface{})
	Next() bool
	Close()
}

// DbRepo struct holding an instance of DbHandler
type DbRepo struct {
	dbHandler DbHandler
}
