package actions

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/kreas/key_master/ipfs"
	"github.com/kreas/key_master/repo"
)

func TestGetDataFileHandler(t *testing.T) {
	ripemdHash := "5e52fee47e6b070565f74372468cdc699de89107"
	w := httptest.NewRecorder()
	vars := map[string]string{"id": ripemdHash}

	createRecord(ripemdHash)

	r := httptest.NewRequest("GET", "/", nil)
	r = mux.SetURLVars(r, vars)

	GetDataHandler(w, r)

	if w.Code != 200 {
		t.Fail()
	}

	var result Result
	json.Unmarshal(w.Body.Bytes(), &result)

	if string(result.Data) != "Short Test" {
		t.Fail()
	}
}

func TestGetDataHandler404(t *testing.T) {
	ripemdHash := "5e52fee47e6b070565f74372468cdc699de89105"
	w := httptest.NewRecorder()
	vars := map[string]string{"id": ripemdHash}

	r := httptest.NewRequest("GET", "/", nil)
	r = mux.SetURLVars(r, vars)

	GetDataHandler(w, r)

	if w.Code != 404 {
		t.Fail()
	}
}

type PostBody struct {
	Data string `json:"data"`
}

func TestAddDataHandler(t *testing.T) {
	ripemdHash := "5e52fee47e6b070565f74372468cdc699de89104"
	w := httptest.NewRecorder()
	vars := map[string]string{"id": ripemdHash}
	data := PostBody{Data: "Some ciphertext"}
	jsonBody, _ := json.Marshal(data)

	r := httptest.NewRequest("POST", "/", bytes.NewReader(jsonBody))
	r = mux.SetURLVars(r, vars)

	AddDataHandler(w, r)

	var result Result
	json.Unmarshal(w.Body.Bytes(), &result)

	if result.Hash[:2] != "Qm" {
		t.Fail()
	}
}

// Helper functions =============================================>

func createRecord(id string) (repo.Result, error) {
	hash, _ := ipfs.Add("Short Test")
	return repo.Upsert(id, hash)
}
