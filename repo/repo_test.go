package repo

import (
	"testing"
)

func TestUpdatingARecordWithUpsert(t *testing.T) {
	record, _ := Upsert("cab60fa36c2598f187912634adf2a534ae28c423", "Awww yeah")
	newHash := "Yeah Buddy"

	// Updating an already
	if result, _ := Upsert(record.ID, newHash); result.Hash != newHash {
		t.Fail()
	}
}

func TestCreatingARecordWithUpsert(t *testing.T) {
	hash := "test"
	id := "e7d4f9258f8de09802a83fcb5b728d821576410d"

	if result, _ := Upsert(id, hash); result.Hash != hash && result.ID != id {
		t.Fail()
	}
}

func TestGet(t *testing.T) {
	hash := "test"
	id := "bab60fa36c2598f187912634adf2a534ae28c423"

	Upsert(id, hash)

	if result, err := Get(id); err != nil && result.Hash != hash {
		t.Fail()
	}
}
