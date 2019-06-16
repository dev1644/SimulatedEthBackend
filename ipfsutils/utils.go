package ipfsutils

import (
	"fmt"
	"io"
	"io/ioutil"
	"time"

	ipfsapi "github.com/ipfs/go-ipfs-api"
)

// IpfsManager is our helper wrapper for IPFS
type IpfsManager struct {
	shell       *ipfsapi.Shell
	nodeAPIAddr string
}

// NewManager is used to instantiate IpfsManager with a connection to an ipfs api.
// if token is provided, we use it to establish an authentication, direct connection
// to an ipfs node api, which involves skipping multiaddr parsing. This is useful
// in situations such as interacting with Nexus' delegator to talk with private ipfs
// networks which use non-standard connection methods.
func NewManager(ipfsURL string, timeout time.Duration) (*IpfsManager, error) {
	var sh *ipfsapi.Shell

	sh = ipfsapi.NewShell(ipfsURL)

	// validate we have an active connection
	if _, err := sh.ID(); err != nil {
		return nil, fmt.Errorf("failed to connect to ipfs node at '%s': %s", ipfsURL, err.Error())
	}
	// set timeout
	sh.SetTimeout(timeout)
	// instantiate and return manager
	return &IpfsManager{
		shell:       sh,
		nodeAPIAddr: ipfsURL,
	}, nil
}

// NodeAddress returns the node the manager is connected to
func (im *IpfsManager) NodeAddress() string { return im.nodeAPIAddr }

// DagPut is used to store data as an ipld object
func (im *IpfsManager) DagPut(data interface{}, encoding, kind string) (string, error) {
	return im.shell.DagPut(data, encoding, kind)
}

// DagGet is used to get an ipld object
func (im *IpfsManager) DagGet(cid string, out interface{}) error {
	return im.shell.DagGet(cid, out)
}

// Cat is used to get cat an ipfs object
func (im *IpfsManager) Cat(cid string) ([]byte, error) {
	var (
		r   io.ReadCloser
		err error
	)
	r, err = im.shell.Cat(cid)
	if err != nil {
		return nil, err
	}
	defer r.Close()
	return ioutil.ReadAll(r)
}

// Stat is used to retrieve the stats about an object
func (im *IpfsManager) Stat(hash string) (*ipfsapi.ObjectStats, error) {
	return im.shell.ObjectStat(hash)
}
