package interfaces

import (
	"github.com/GerardSoleCa/PubKeyManager/domain"
)

// DbKeyRepository inheriting from DbRepo
type DbKeyRepository DbRepo

// NewDbKeyRepository creates a new instance
func NewDbKeyRepository(dbHandler DbHandler) *DbKeyRepository {
	dbKeyRepo := new(DbKeyRepository)
	dbKeyRepo.dbHandler = dbHandler
	dbKeyRepo.CreateTable()
	return dbKeyRepo
}

// CreateTable function from DbKeyRepository
func (repo *DbKeyRepository) CreateTable() {
	_, err := repo.dbHandler.Execute("CREATE TABLE IF NOT EXISTS Keys (id INTEGER PRIMARY KEY AUTOINCREMENT, user TEXT NOT NULL, title TEXT NOT NULL, fingerprint TEXT NOT NULL, key TEXT NOT NULL)")
	if err != nil {
		panic(err)
	}
}

// Store function from DbKeyRepository
func (repo *DbKeyRepository) Store(key *domain.Key) error {
	res, err := repo.dbHandler.Execute("INSERT INTO keys (user, title, fingerprint, key) VALUES (?, ?, ?, ?)", key.User, key.Title, key.Fingerprint, key.Key)
	key.Id, _ = res.LastInsertId()
	return err
}

// Delete function from DbKeyRepository
func (repo *DbKeyRepository) Delete(id int64) error {
	_, err := repo.dbHandler.Execute("DELETE FROM keys where id=?", id)
	return err
}

// GetKeys function from DbKeyRepository
func (repo *DbKeyRepository) GetKeys() []domain.Key {
	rows, err := repo.dbHandler.Query("SELECT * FROM keys")
	return repo.processKeyRows(rows, err)
}

// GetUserKeys function from DbKeyRepository
func (repo *DbKeyRepository) GetUserKeys(user string) []domain.Key {
	rows, err := repo.dbHandler.Query("SELECT * FROM keys WHERE user=?", user)
	return repo.processKeyRows(rows, err)
}

func (repo *DbKeyRepository) processKeyRows(rows Row, err error) []domain.Key {
	var keys []domain.Key
	defer rows.Close()
	if err != nil {
		return keys
	}
	for rows.Next() {
		k := domain.Key{}
		rows.Scan(&k.Id, &k.User, &k.Title, &k.Fingerprint, &k.Key)
		keys = append(keys, k)
	}
	return keys
}
