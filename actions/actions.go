package actions

import (
	"io/ioutil"
	"net/http"

	"encoding/json"

	"github.com/gorilla/mux"
	"github.com/kreas/key_master/ipfs"
	"github.com/kreas/key_master/repo"
)

// Result is the data that is returned to client
type Result struct {
	Hash string `json:"hash"`
	Data []byte `json:"data,omitempty"`
}

type postBody struct {
	Data string
}

// GetDataHandler fetches a given datafile from IPFS and returns it to the
// user.
// ** TODO: Add authentication
func GetDataHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	record, err := repo.Get(id)

	if err != nil {
		http.Error(w, "Database breakdown", http.StatusInternalServerError)
		return
	}

	if record.Hash == "" {
		http.NotFound(w, r)
		return
	}

	data, _ := ipfs.Fetch(record.Hash)

	result := Result{
		Hash: record.Hash,
		Data: data,
	}

	response, _ := json.Marshal(result)

	w.Write(response)
}

// AddDataHandler takes a post request and saves the `data` portion of the
// request to IPFS and updates the LevelDB database pointing value of a given
// key to the hash returned from IPFS.
// ** TODO: Add authentication
func AddDataHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	body, _ := ioutil.ReadAll(r.Body)

	var t postBody
	err := json.Unmarshal(body, &t)

	if err != nil {
		http.Error(w, "Invalid Request", http.StatusUnprocessableEntity)
		return
	}

	hash, _ := ipfs.Add(t.Data)
	record, _ := repo.Upsert(id, hash)

	result := Result{
		Hash: record.Hash,
	}

	response, _ := json.Marshal(result)

	w.Write(response)
}
