package database

import (
	"database/sql"
	"github.com/op/go-logging"
	_ "github.com/xeodou/go-sqlcipher"
)

const DB_NAME = "PubKeyManager.db"

var glog = logging.MustGetLogger("database")

func init() {
	glog.Infof("Creating database %s", DB_NAME)
	db := getDb()
	defer db.Close()
	initKeysDatabase(db)
	initUsersDatabase(db)
}

func getDb() (*sql.DB) {
	db, err := sql.Open("sqlite3", DB_NAME)
	if err != nil {
		glog.Panic("Database error on open", err)
	}
	//p := "PRAGMA key = '123456';"
	//_, err = db.Exec(p)
	checkErr(err)
	return db
}

func checkErr(err error) {
	if err != nil {
		glog.Panic(err)
	}
}