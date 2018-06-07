package ipfs

import (
	"testing"

	coremock "github.com/ipfs/go-ipfs/core/mock"
)

var mockNode, _ = coremock.NewMockNode()

func TestAddToIPFS(t *testing.T) {
	if hash, _ := AddToIPFS(mockNode, "test"); hash[:2] != "Qm" {
		t.Fail() // IPFS hashes of files always start with Qm
	}
}

func TestFetchFromIPFS(t *testing.T) {
	hash, err := AddToIPFS(mockNode, "Random")

	if err != nil {
		t.Error(err)
	}

	if result, _ := FetchFromIPFS(mockNode, hash); string(result) != "Random" {
		t.Fail()
	}
}
