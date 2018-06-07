package ipfs

import (
	"testing"
)

func TestAdd(t *testing.T) {
	if hash, _ := Add("test"); hash[:2] != "Qm" {
		t.Fail() // IPFS hashes of files always start with Qm
	}
}

func TestFetch(t *testing.T) {
	hash, err := Add("Random")

	if err != nil {
		t.Error(err)
	}

	if result, _ := Fetch(hash); string(result) != "Random" {
		t.Fail()
	}
}
