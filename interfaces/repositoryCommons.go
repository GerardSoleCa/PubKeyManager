package interfaces

type DbHandler interface {
	Execute(statement string, args ...interface{}) (Result, error)
	Query(statement string, args ...interface{}) (Row, error)
}

type Result interface {
	LastInsertId() (int64, error)
	RowsAffected() (int64, error)
}

type Row interface {
	Scan(dest ...interface{})
	Next() bool
	Close()
}

type DbRepo struct {
	dbHandler DbHandler
}


