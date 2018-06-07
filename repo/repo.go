package repo // import "github.com/kreas/key_master/repo"
import (
	"os/user"

	"github.com/syndtr/goleveldb/leveldb"
)

// DB pointer to levedb database
type DB = *leveldb.DB

// Result represents the data as it's stored in levelDB. Where ID is the RIPEMD
// of the users public key and Hash is the most recent version of the of the
// records hash as it's stored on IPFS.
type Result struct {
	ID   string
	Hash string
}

// Upsert updates a record to leveldb. If the record doesn't exist create it.
func Upsert(id string, hash string) (Result, error) {
	s := func(db DB) Result {
		db.Put([]byte(id), []byte(hash), nil)

		return Result{id, hash}
	}

	return transact(s)
}

// Get fetchs a record from leveldb.
func Get(id string) (Result, error) {
	s := func(db DB) Result {
		data, _ := db.Get([]byte(id), nil)

		return Result{id, string(data)}
	}

	return transact(s)
}

// transact opens a levelDB from the users' .key_master directory and performs a
// database operation f. Once the operation is complete the levelDB connection
// is closed.
func transact(f func(DB) Result) (Result, error) {
	user, _ := user.Current()
	db, err := leveldb.OpenFile(user.HomeDir+"/.key_master", nil)

	defer db.Close()

	return f(db), err
}
