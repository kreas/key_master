package ipfs

import (
	"context"
	"io/ioutil"
	"strings"
	"time"

	"github.com/ipfs/go-ipfs/core"
	"github.com/ipfs/go-ipfs/core/coreunix"
)

// AddToIPFS takes the contents of
func AddToIPFS(ipfs *core.IpfsNode, contents string) (string, error) {
	c := strings.NewReader(contents)
	return coreunix.Add(ipfs, c)
}

// FetchFromIPFS fetches data from an IPFS
func FetchFromIPFS(ipfs *core.IpfsNode, hash string) ([]byte, error) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 100*time.Millisecond)

	// If the request takes too long, which happens more often then I'd like,
	// cancel the transaction.
	defer cancel()

	resp, err := coreunix.Cat(ctx, ipfs, hash)

	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(resp)
}
