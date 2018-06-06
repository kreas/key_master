package repo // import "github.com/kreas/key_master/repo"
import (
	"os/user"

	"github.com/gobuffalo/uuid"
	"github.com/syndtr/goleveldb/leveldb"
)

// DB pointer to levedb database
type DB = *leveldb.DB

// Result represents the data as it's stored in levelDB. Where ID is the RIPEMD
// of the users public key and Hash is the most recent version of the of the
// records hash as it's stored on IPFS.
type Result struct {
	ID   uuid.UUID
	Hash string
}

// Create adds a new record to leveldb
// ** TODO: the UUID should be replaced with the RIPEMD-160(pk) of the user
// making the request
func Create(hash string) (Result, error) {
	s := func(db DB) Result {
		id := uuid.Must(uuid.NewV4())
		db.Put(id.Bytes(), []byte(hash), nil)

		return Result{id, hash}
	}

	return transact(s)
}

// Upsert updates a record to leveldb. If the record doesn't exist create it.
// ** TODO: the UUID should be replaced with the RIPEMD-160(pk) of the user
// making the request
func Upsert(id uuid.UUID, hash string) (Result, error) {
	s := func(db DB) Result {
		db.Put(id.Bytes(), []byte(hash), nil)

		return Result{id, hash}
	}

	return transact(s)
}

// AddFileToIPFS takes the contents of
// func AddFileToIPFS(ipfs core.IpfsNode, contents string) {
// 	fmt.Printf(ipfs)
// 	fmt.Printf(contents)
// }

// transact opens a levelDB from the users' .key_master directory and performs a
// database operation f. Once the operation is complete the levelDB connection
// is closed.
func transact(f func(DB) Result) (Result, error) {
	user, _ := user.Current()
	db, err := leveldb.OpenFile(user.HomeDir+"/.key_master", nil)

	defer db.Close()

	return f(db), err
}
