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

// PatchLink is used to link two objects together
// path really means the name of the link
// create is used to specify whether intermediary nodes should be generated
func (im *IpfsManager) PatchLink(root, path, childHash string, create bool) (string, error) {
	return im.shell.PatchLink(root, path, childHash, create)
}

// AppendData is used to modify the raw data within an object, to a max of 1MB
// Anything larger than 1MB will not be respected by the rest of the network
func (im *IpfsManager) AppendData(root string, data interface{}) (string, error) {
	return im.shell.PatchData(root, false, data)
}

// SetData is used to set the data field of an ipfs object
func (im *IpfsManager) SetData(root string, data interface{}) (string, error) {
	return im.shell.PatchData(root, true, data)
}

// NewObject is used to create a generic object from a template type
func (im *IpfsManager) NewObject(template string) (string, error) {
	return im.shell.NewObject(template)
}
