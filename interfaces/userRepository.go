package interfaces

import (
	"github.com/GerardSoleCa/PubKeyManager/domain"
	"github.com/op/go-logging"
)

var glog = logging.MustGetLogger("UserRepository")

type DbUserRepository DbRepo

func NewDbUserRepository(dbHandler DbHandler) *DbUserRepository {
	dbUserRepo := new(DbUserRepository)
	dbUserRepo.dbHandler = dbHandler
	dbUserRepo.CreateTable()
	return dbUserRepo
}

func (repo *DbUserRepository) CreateTable() {
	_, err := repo.dbHandler.Execute("CREATE TABLE IF NOT EXISTS Users (id INTEGER PRIMARY KEY AUTOINCREMENT, username TEXT, password TEXT NOT NULL, UNIQUE (id, username))")
	if err != nil {
		panic(err)
	}
}

func (repo *DbUserRepository) Store(user *domain.User) error {
	res, err := repo.dbHandler.Execute("INSERT INTO users (username, password) VALUES (?, ?)", user.Username, user.Password)
	if err == nil {
		user.Id, _ = res.LastInsertId()
	}
	return err
}

func (repo *DbUserRepository) GetUser(user string) (domain.User, error) {
	row, err := repo.dbHandler.Query("SELECT * FROM users WHERE username=?", user)
	u := domain.User{}
	if err != nil {
		glog.Debugf("Error > UserRepository > GetUser :: %s", err.Error())
		return u, err
	}
	for row.Next() {
		row.Scan(&u.Id, &u.Username, &u.Password)
		break
	}
	row.Close()
	return u, nil
}

func (repo *DbUserRepository) Update(user *domain.User) error {
	_, err := repo.dbHandler.Execute("UPDATE users SET password = ? WHERE username = ?", user.Username, user.Password)
	return err
}

func (repo *DbUserRepository) Delete(username string) error {
	_, err := repo.dbHandler.Execute("DELETE FROM users where username=?", username)
	return err
}

func (repo *DbUserRepository) Count() (int, error) {
	row, err := repo.dbHandler.Query("SELECT COUNT(*) as count FROM users")
	defer row.Close()
	var u int
	if err != nil {
		glog.Debugf("Error > UserRepository > Count :: %s", err.Error())
		return u, err
	}
	for row.Next() {
		row.Scan(&u)
		break
	}
	return u, nil
}