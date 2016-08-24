package interfaces

import (
	"github.com/GerardSoleCa/PubKeyManager/domain"
	"github.com/op/go-logging"
)

var glog = logging.MustGetLogger("UserRepository")

// DbUserRepository struct
type DbUserRepository DbRepo

// NewDbUserRepository creates a new UserRepository backed up by a DB
func NewDbUserRepository(dbHandler DbHandler) *DbUserRepository {
	dbUserRepo := new(DbUserRepository)
	dbUserRepo.dbHandler = dbHandler
	dbUserRepo.CreateTable()
	return dbUserRepo
}

// CreateTable Function contained on DbUserRepository
func (repo *DbUserRepository) CreateTable() {
	_, err := repo.dbHandler.Execute("CREATE TABLE IF NOT EXISTS Users (id INTEGER PRIMARY KEY AUTOINCREMENT, username TEXT, password TEXT NOT NULL, UNIQUE (id, username))")
	if err != nil {
		panic(err)
	}
}

// Store Function contained on DbUserRepository
func (repo *DbUserRepository) Store(user *domain.User) error {
	res, err := repo.dbHandler.Execute("INSERT INTO users (username, password) VALUES (?, ?)", user.Username, user.Password)
	if err == nil {
		user.Id, _ = res.LastInsertId()
	}
	return err
}

// GetUser Function contained on DbUserRepository
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

// Update Function contained on DbUserRepository
func (repo *DbUserRepository) Update(user *domain.User) error {
	_, err := repo.dbHandler.Execute("UPDATE users SET password = ? WHERE username = ?", user.Username, user.Password)
	return err
}

// Delete Function contained on DbUserRepository
func (repo *DbUserRepository) Delete(username string) error {
	_, err := repo.dbHandler.Execute("DELETE FROM users where username=?", username)
	return err
}

// Count Function contained on DbUserRepository
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
