package ipfs

import (
	"context"
	"io/ioutil"
	"strings"
	"time"

	"github.com/ipfs/go-ipfs/core/coreunix"
	coremock "github.com/ipfs/go-ipfs/core/mock"
)

// Node is jus a mock
var Node, _ = coremock.NewMockNode()

// Add takes the contents of
func Add(contents string) (string, error) {
	c := strings.NewReader(contents)
	return coreunix.Add(Node, c)
}

// Fetch fetches data from an IPFS
func Fetch(hash string) ([]byte, error) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 100*time.Millisecond)

	// If the request takes too long, which happens more often then I'd like,
	// cancel the transaction.
	defer cancel()

	resp, err := coreunix.Cat(ctx, Node, hash)

	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(resp)
}
