package database

import "database/sql"

type Key struct {
	Id          int           `json:"id,omitempty"`
	User        string        `json:"user"`
	Title       string        `json:"title"`
	Fingerprint string        `json:"fingerprint"`
	Key         string        `json:"key"`
}

func initKeysDatabase(db *sql.DB) {
	c := "CREATE TABLE IF NOT EXISTS `keys` (`id` INTEGER PRIMARY KEY AUTOINCREMENT, `user` char, `title` char, `fingerprint` char, `key` char, UNIQUE(`title`));"
	if _, err := db.Exec(c); err != nil {
		glog.Panic("Could not create keys database", err)
	}
}

func GetUserKeys(user string) ([]Key) {
	db := getDb()
	defer db.Close()
	rows, err := db.Query("select id, user, title, fingerprint, key from keys where user = ?;", user)
	checkErr(err)
	defer rows.Close()
	keys := []Key{}
	for rows.Next() {
		var k Key
		err := rows.Scan(&k.Id, &k.User, &k.Title, &k.Fingerprint, &k.Key)
		glog.Infof("key: %+v\n", rows)
		checkErr(err)
		keys = append(keys, k)
	}
	glog.Infof("Returning [ %s ] keys: %+v\n", user, keys)
	return keys
}

func AddUserKey(key *Key) (err error){
	glog.Infof("Storing User Key: %+v\n", key)
	db := getDb()
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO keys (user, title, fingerprint, key) VALUES (?, ?, ?, ?)")
	checkErr(err)

	_, err = stmt.Exec(key.User, key.Title, key.Fingerprint, key.Key)
	//_, err = db.Query("INSERT INTO keys (user, title, fingerprint, key) VALUES ('?', '?', '?', '?');", key.User, key.Title, key.Fingerprint, key.Key)
	return err
}