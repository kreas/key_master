package repo

import (
	"testing"

	"github.com/gobuffalo/uuid"
)

func TestCreate(t *testing.T) {
	if result, err := Create("test"); err != nil && result.Hash != "test" {
		t.Fail()
	}
}

func TestUpdatingARecordWithUpsert(t *testing.T) {
	record, _ := Create("Awww yeah")
	newHash := "Yeah Buddy"

	// Updating an already
	if result, _ := Upsert(record.ID, newHash); result.Hash != newHash {
		t.Fail()
	}
}

func TestCreatingARecordWithUpsert(t *testing.T) {
	hash := "test"
	id := uuid.Must(uuid.NewV4())

	if result, _ := Upsert(id, hash); result.Hash != hash && result.ID != id {
		t.Fail()
	}
}
