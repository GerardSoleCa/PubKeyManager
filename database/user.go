package database

import "database/sql"

type User struct {
	Id       int64 `json:"-"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
}

func initUsersDatabase(db *sql.DB) {
	c := "CREATE TABLE IF NOT EXISTS `users` (`id` INTEGER PRIMARY KEY AUTOINCREMENT, `email` char, `password` char, UNIQUE(`email`));"
	if _, err := db.Exec(c); err != nil {
		glog.Panic("Could not create keys database", err)
	}
}

func (u *User) Save() (err error) {
	u.Id, err = addUser(u)
	return err
}

func (u *User) CheckValidity() (bool) {
	return false
}

func (u *User) hashPassword() {

}

func addUser(u *User) (int64, error){
	db := getDb()
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO users (email, password) VALUES (?, ?)")
	if err != nil{
		return 0, err
	}
	res, err := stmt.Exec(u.Email, u.Password)
	id, err := res.LastInsertId()
	if err != nil{
		return 0, err
	}
	return id, err
}

func CountUsers() (int) {
	db := getDb()
	defer db.Close()
	rows, err := db.Query("SELECT id FROM users")
	checkErr(err)
	defer rows.Close()
	users := []User{}
	for rows.Next() {
		var u	User
		err := rows.Scan(&u.Id)
		glog.Infof("key: %+v\n", rows)
		checkErr(err)
		users = append(users, u)
	}
	return len(users)
}