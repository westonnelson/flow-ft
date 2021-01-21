// Code generated by go-bindata. DO NOT EDIT.
// sources:
// ../../../transactions/burn_tokens.cdc (1.093kB)
// ../../../transactions/create_forwarder.cdc (1.836kB)
// ../../../transactions/mint_tokens.cdc (895B)
// ../../../transactions/privateForwarder/create_private_forwarder.cdc (961B)
// ../../../transactions/privateForwarder/deploy_forwarder_contract.cdc (278B)
// ../../../transactions/privateForwarder/setup_and_create_forwarder.cdc (1.954kB)
// ../../../transactions/privateForwarder/transfer_private_many_accounts.cdc (1.021kB)
// ../../../transactions/scripts/get_balance.cdc (468B)
// ../../../transactions/scripts/get_supply.cdc (229B)
// ../../../transactions/setup_account.cdc (1.258kB)
// ../../../transactions/transfer_many_accounts.cdc (1.215kB)
// ../../../transactions/transfer_tokens.cdc (1.387kB)

package assets

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data, name string) ([]byte, error) {
	gz, err := gzip.NewReader(strings.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)

	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}

	clErr := gz.Close()
	if clErr != nil {
		return nil, clErr
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _burn_tokensCdc = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x52\xc1\x6e\xdb\x3a\x10\xbc\xeb\x2b\xe6\xe5\xf0\x60\x1f\x22\xa5\x40\xd1\x83\xe1\x36\x75\x12\xbb\x28\x5a\xb8\x40\xe3\xb4\x67\x4a\x5a\x5b\x44\x25\x52\x58\xae\x6a\x07\x41\xfe\xbd\x20\x29\xab\x52\x81\xc4\x07\x13\x22\x67\x67\x67\x66\x37\xcb\xb0\xab\xb4\x83\xb0\x32\x4e\x15\xa2\xad\x81\x76\x50\x10\x6a\xda\x5a\x09\x61\x6f\xd9\x7f\x8e\xde\xa5\x52\x92\x64\x19\x0a\xdb\xd5\x25\x72\x42\xe7\xa8\x44\xfe\x08\xa9\x08\xaa\x6c\xb4\x81\x2a\x0a\xdb\x19\x81\x58\xe4\x1d\x1b\x88\xfd\x45\xc6\xf9\xa2\x3d\xdb\xc6\x03\x35\xc3\x89\x65\x2a\xf1\x43\x75\xb5\xe7\x4b\x82\x16\x0a\x05\xda\x1c\xa0\x9a\x40\x71\x3c\x77\x51\x68\x15\xab\x86\x84\xd8\xf3\xfa\x66\x23\x55\x49\xa2\x9b\xd6\xb2\x60\xd3\x99\x83\xce\x6b\xda\xf9\x96\xb1\xdd\xd5\x69\xf3\xb0\xfd\xf4\xf9\xe6\xeb\x7a\xf7\xed\xcb\x7a\xbb\xba\xbb\xfb\xbe\xbe\xbf\x3f\x17\xac\x4f\xaa\x69\xff\xc1\x4f\x70\xc9\xa8\xcd\x2c\xaa\x5a\xe0\x61\xa3\x4f\xef\xde\xce\xf1\x94\x24\x00\x90\x65\xd1\x07\x98\x9c\xed\xb8\xa0\x90\x12\x2a\x5b\x97\x2e\x4a\x0d\x09\xc4\x5b\xc5\x84\x9c\xbc\x47\xef\x95\xca\xc0\x50\x93\xe0\xb7\xa7\x58\xe0\xe3\xc4\x43\x1a\x03\x1a\x40\x21\xe1\x05\xfe\x1f\xeb\x4e\x57\xfe\x52\x3b\x61\x25\x96\x23\xb6\x65\x6a\x15\xd3\xcc\xe9\x83\x21\x5e\x60\xd5\x49\xb5\x8a\x73\x19\x64\xf7\xd2\x7f\x6a\xa9\x4a\x56\x47\xbc\xb9\x3a\x0b\x3d\xcf\xa9\x1f\x68\x50\x06\x6d\xc2\xd0\xd4\x81\x86\x6a\x47\xf5\x3e\x8d\xaf\xcb\x4b\xc4\x5e\x69\x6e\x99\xed\x71\x39\x95\x18\x6c\x7c\x98\x79\xe2\x05\xb2\x9e\x27\xa3\x11\x24\x20\xe6\xff\x0d\xdc\xfe\x97\x1e\x7b\x6d\x43\xf2\xf1\x9c\x4f\x0c\xdc\x32\xf9\x5d\x55\x60\xda\x13\x93\xf1\xf9\xdb\xf1\x3e\x86\xff\x61\x36\x2f\xf9\x88\xb0\xf7\xaf\xda\x98\x24\xfd\xaa\x9d\x80\x9c\x4f\xdc\x5c\x5f\xa3\x55\x46\x17\xb3\x8b\xdb\xb0\xd6\xc6\x0a\x62\x97\x97\xb5\x9f\x55\x5f\x44\xaa\xe7\x68\x9c\x4e\x54\x74\x42\x78\x1a\xf8\xfd\x6e\x84\x7d\xe2\x30\x89\xc1\x4f\x5a\x84\x70\xb6\x74\xbc\x09\xaf\xb3\x51\x74\x11\x9f\xfa\x23\x48\x76\xbd\x9f\xe5\xe5\xdf\xb9\x8e\xe0\x25\x39\x61\xfb\xd8\x97\xf5\x72\x9e\x93\x3f\x01\x00\x00\xff\xff\xef\x73\x4d\xa0\x45\x04\x00\x00"

func burn_tokensCdcBytes() ([]byte, error) {
	return bindataRead(
		_burn_tokensCdc,
		"burn_tokens.cdc",
	)
}

func burn_tokensCdc() (*asset, error) {
	bytes, err := burn_tokensCdcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "burn_tokens.cdc", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x53, 0x65, 0xc2, 0xb4, 0xe1, 0x29, 0xef, 0x51, 0xff, 0xc4, 0xd6, 0xb8, 0x58, 0x43, 0xdb, 0x3d, 0xca, 0xce, 0xa9, 0x3d, 0x1, 0xa1, 0xbc, 0xa3, 0x50, 0x6, 0x6e, 0x66, 0x1, 0x29, 0xc2, 0x6a}}
	return a, nil
}

var _create_forwarderCdc = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x54\x4d\x6f\xe3\x36\x10\xbd\xf3\x57\xcc\xa9\xf5\x06\x8e\xdc\xcf\x8b\x91\x16\x70\x37\xc9\x22\x68\x91\x05\x9c\xb4\x3d\x76\xc7\xe4\xc8\x64\x23\x91\x02\x39\xb2\xd6\x58\xe4\xbf\x17\x24\x25\x46\x0a\x8a\xac\x4f\x96\x38\xf3\xe6\xbd\x37\x4f\xdc\x5c\x5c\x08\xf1\xa8\x4d\x00\xf6\x68\x03\x4a\x36\xce\x82\x09\x80\xc0\xd4\x76\x0d\x32\x41\xed\x7c\x7c\x9c\x9d\xb3\x46\x06\xe9\xfa\x46\xc1\x81\xa0\x0f\xa4\x04\x3b\x08\xc4\xd0\x77\x80\x16\x50\x4a\xd7\x5b\x06\x76\xb1\x79\x40\xaf\x40\x51\xe7\x82\x61\x52\xc0\xee\x89\x6c\x88\x67\x68\x1d\x6b\xf2\xe0\x49\x92\x39\x91\xaf\x84\xb8\xab\x01\xed\xd9\x59\x82\x40\x56\x85\x79\x71\x9c\xe3\xbf\x0d\x70\x9b\x11\xc9\xc3\x7e\xec\x5b\x0b\xd6\x54\x9e\x60\x30\x4d\x03\xff\xf6\x81\xcb\x70\xd6\x2e\xd0\x0c\x2b\x96\xff\x85\x7d\xc3\x59\x89\xc6\x00\x07\x22\x2b\xa2\x02\x0c\xe9\xd8\x93\x34\x9d\x21\xcb\x80\x56\x01\xb5\x26\xfe\x01\x3a\xc5\x37\xa9\xc9\x58\x65\x24\x32\x05\x31\x68\x23\x75\x62\x37\x0d\x8c\x2a\xf5\x34\xb0\x1a\x0d\x1e\xf0\xbc\x06\x13\xf5\x81\xab\xeb\x4b\xa9\xd1\x58\x08\xe4\x4f\x46\x12\x0c\x68\x39\x51\x6b\x9d\x35\xec\x3c\x0c\xda\xc5\x35\x8c\x80\xc6\x1e\xc5\x0b\x7d\xc3\x6b\x30\x0c\x12\x2d\x0c\xc8\x52\x67\x5a\xe9\x28\x10\xc1\xa0\xc9\xd3\x8c\x00\x48\x6c\x09\x6a\xef\xda\x4a\x88\x07\xa6\x6e\xac\xcc\xdb\xca\xab\x0a\x30\x18\xd6\xb9\xa1\xa8\xf0\x5b\x21\xbe\xaf\xe0\x51\x13\xdc\xf6\xf6\x68\x0e\x0d\xc1\x63\xaa\x90\xce\xb2\x47\x19\x5d\x60\xf2\x35\x4a\x82\xa0\x53\x1e\xb0\xf1\x84\xea\x1c\x73\xa1\xa8\x6b\xdc\x99\x14\x04\xd7\x52\x22\x25\x7e\xc8\x68\xd8\x75\x8d\x91\x18\xf1\x78\x89\x37\xa2\xcc\xba\x2b\xf1\x63\x6e\x9a\x6d\x64\x8c\xd7\x58\xac\xf1\x44\x80\xe3\x42\x63\x58\x39\xe5\x39\x03\x7b\x42\x26\x25\x00\x20\x2d\x32\xb0\xf3\xa4\xc0\x58\x30\x1c\xd2\x13\x1e\x29\x6b\x47\xe8\xfa\x43\x63\x82\x26\x55\xb2\x24\x7e\xaa\xe0\x3a\x11\x49\x7e\x7e\x4a\xea\x6f\xcb\x4e\x2a\xa9\xe4\xa7\x17\xf2\x29\xa5\xca\xd4\x35\xf9\x19\x4d\xf1\x73\x15\x33\x0b\x08\x96\x06\xd8\xe5\x97\x5b\x78\x9f\x98\x25\xd8\x49\x8f\x75\xbe\xc5\xa6\x39\xaf\x13\x5d\xd6\x64\xc1\xf7\x36\x4f\xce\x42\xfe\x29\xab\xc9\xa3\x67\x1f\x65\x6e\x3a\x12\xb3\xb1\x47\x58\x7c\x10\x71\xf5\x8b\x41\x39\xc0\xaf\x82\x5e\x89\x8b\x8d\x10\xa6\xed\x9c\xe7\xb2\xef\xbc\xee\x04\xf0\xdd\xe7\xdb\x3f\xef\x3f\xdc\xfd\xf6\xc7\xcd\xe3\xc7\xdf\x6f\xee\x77\xd7\xd7\xfb\x9b\x87\x87\xa9\xe1\xe6\x33\xb6\xdd\xab\xfa\xff\xab\x7b\xe5\x60\x81\xfe\xb8\xff\x7b\xb7\xbf\xbe\xbb\xff\x30\xd5\x8b\x99\xb6\xd5\x74\x43\x6c\x61\xa7\x94\xa7\x10\xde\xc1\x17\x91\x04\x77\x9e\x3a\xf4\xb4\x42\x29\x79\x0b\xbb\x9e\xf5\xe8\x70\xac\x80\xf1\xd7\x10\xcf\xe2\xf3\x4b\x74\x69\xac\x2a\xc8\xef\x4a\x71\xfc\x55\x47\xe2\xf7\xd8\xe1\xc1\x34\x86\xcf\x57\xdf\x7c\x59\xf8\x51\x4d\xce\x3e\xff\xba\xda\xa4\xd0\xc8\x0d\xcd\xf4\xef\x0b\xe6\x82\xc1\x29\x05\xf4\xea\xf2\xb5\x07\x55\xde\xed\x3d\x0d\xe5\x6a\x5b\x15\xb6\xdb\x17\xe2\x2f\x14\xa3\xd8\x2a\xe0\x89\x56\x57\x97\x09\x75\x0d\xec\xb6\xb0\x19\xf3\xbc\x20\x53\x30\x67\x6c\xe2\x2d\x14\x21\x16\x2a\xdf\x96\x52\x49\x4d\xf2\xe9\x2d\x27\xe6\x86\x17\x92\xbd\x6d\x8c\x7d\xfa\x8a\x4b\x53\xc7\xb3\x58\x2a\x8c\xad\x6f\x4d\x5c\x8c\x7b\x6b\xc4\x7a\x51\xc9\xe8\x8f\xc4\x5f\xb3\xab\xb4\x64\x7e\xcf\xe2\xf9\xbf\x00\x00\x00\xff\xff\xa7\x39\xe1\x44\x2c\x07\x00\x00"

func create_forwarderCdcBytes() ([]byte, error) {
	return bindataRead(
		_create_forwarderCdc,
		"create_forwarder.cdc",
	)
}

func create_forwarderCdc() (*asset, error) {
	bytes, err := create_forwarderCdcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "create_forwarder.cdc", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x6b, 0xff, 0x9c, 0x63, 0x34, 0x73, 0xfe, 0xcf, 0xa5, 0x23, 0x2c, 0x88, 0x99, 0xa5, 0x37, 0x84, 0xcc, 0x24, 0xbf, 0xdf, 0xe6, 0x7e, 0xc8, 0xe1, 0x94, 0x12, 0x19, 0xa4, 0x93, 0xdd, 0x7d, 0x53}}
	return a, nil
}

var _mint_tokensCdc = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x52\xcd\x6e\xdb\x3c\x10\xbc\xeb\x29\x16\x3e\x04\x32\xf0\x45\xfe\x0e\x45\x0f\x86\x93\x40\x6d\xec\xa2\x68\xeb\x02\x71\xdc\x3b\x45\xad\x95\x45\x29\x92\x58\xae\x6a\x07\x41\xde\xbd\xa0\x7e\x1a\xc9\x6d\xc2\x8b\x00\x6a\x66\x76\x66\x87\x54\x7b\xc7\x02\x9b\xc6\x56\x54\x18\xbc\x77\x3f\xd1\xc2\x81\x5d\x0d\xff\x9f\x36\xfb\xed\xa7\xcf\x1f\xbe\xae\xef\xbf\x7f\x59\x6f\xf3\xdb\xdb\xbb\xf5\x6e\x97\xf4\x84\xf5\x49\xd5\xfe\x0c\x3f\xc1\x25\xc2\xca\x06\xa5\x85\x9c\x4d\x19\x35\x79\x42\x2b\x4b\xc8\xcb\x92\x31\x84\xff\x40\xd5\xae\x89\x17\xfb\x0d\x9d\xde\xbf\x9b\xc3\x53\x02\x00\x60\x50\x40\xa2\x68\x5e\xd6\x64\x97\x70\x31\x1e\x94\xb5\x97\x14\x84\x95\x38\x9e\xe2\xef\x50\x23\xfd\x42\x5e\xc2\xc5\xd3\x24\x4d\x36\xfc\x79\x4e\x5a\x86\x67\xf4\x8a\x31\x0d\x54\xd9\x08\xcf\x1b\x79\xc8\xb5\x8e\x66\x06\x13\xf1\x04\x34\x87\xec\xc5\x09\x5c\x41\x47\xc8\x0a\xc7\xec\x8e\xab\x37\x8c\x5d\xa7\x71\x21\x4b\x58\x04\x71\xac\x2a\x5c\xe0\x08\xda\x22\xe7\x7f\xc6\xc4\x73\x73\x03\x5e\x59\xd2\xe9\x6c\xd7\x8e\x00\x0a\x60\x9d\x80\x3c\x60\x97\x0d\x54\x24\xcd\xe6\xc9\x3f\xdc\x0d\xe9\xe0\x0a\x2a\x94\x3e\xc8\xcb\xc2\xa7\x93\xb2\x0a\xe5\xa3\xf2\xaa\x20\x43\xf2\x98\x2e\x7c\x53\x18\xd2\x13\x7f\x83\xde\x19\x71\x48\xfd\xda\x6e\xaf\xd3\xd7\x32\xed\xad\x2a\x4c\x0c\x02\x9d\x06\xf0\xe0\x98\xf1\x80\x8c\x56\xe3\xac\xe3\xf6\x05\xe1\x09\x75\x23\x38\xea\x22\x96\x5c\x93\x15\x64\x58\x5d\x9e\x37\x93\x69\x46\x25\xb8\xc5\xe3\xb7\x16\x92\x2a\x63\xdc\x11\xcb\xbc\x7f\x5f\xdd\x3b\x9b\xff\x2d\x56\xfe\x50\x8d\x91\xa8\xd8\x69\x67\xf1\xd3\xc6\x0a\xa9\x3a\x23\xbf\xb1\xf8\xac\x44\xef\x02\x49\x5f\xfa\xea\x72\x24\x3e\x22\x96\x18\x84\xdd\x63\x3f\xab\xcf\xfb\xfc\x3b\x00\x00\xff\xff\x5e\x98\x79\x6e\x7f\x03\x00\x00"

func mint_tokensCdcBytes() ([]byte, error) {
	return bindataRead(
		_mint_tokensCdc,
		"mint_tokens.cdc",
	)
}

func mint_tokensCdc() (*asset, error) {
	bytes, err := mint_tokensCdcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "mint_tokens.cdc", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x0, 0xed, 0x9c, 0x6f, 0x6f, 0xea, 0x95, 0xac, 0x93, 0x24, 0xf, 0x68, 0x49, 0xcc, 0x54, 0x71, 0x4c, 0xff, 0xe2, 0xa2, 0x9, 0xcc, 0xff, 0xe8, 0x31, 0x37, 0x16, 0xfc, 0x93, 0x6d, 0xdc, 0x89}}
	return a, nil
}

var _privateforwarderCreate_private_forwarderCdc = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x52\x4d\x8f\xda\x30\x10\xbd\xfb\x57\xcc\xa9\x65\xa5\xdd\xa4\x67\x44\x2b\xa5\x25\xac\x50\x2b\x16\x05\xba\x3d\x0f\x66\x4a\xac\x0d\x76\x34\x99\x84\x45\x88\xff\x5e\x39\x21\x28\x49\x41\xaa\xea\x43\x64\xd9\xef\xc3\xef\x4d\xcc\x3e\x77\x2c\x30\x2b\xed\xce\x6c\x32\x5a\xbb\x37\xb2\xf0\x9b\xdd\x1e\x3e\xbd\xcf\x7e\x2e\x9e\xe7\x5f\x7f\xc4\xeb\x97\xef\xf1\x22\x9a\x4e\x93\x78\xb5\x52\x17\x42\xfc\x8e\xfb\x7c\x80\xbf\x85\x5b\xb2\xa9\x50\x28\x21\x4d\xa6\x22\x9e\x39\x3e\x20\x6f\x89\x5b\xce\x32\x99\xbf\x46\xeb\x78\xf6\x92\xfc\x8a\x92\xe9\x7c\xf1\xdc\xf2\x55\x18\xc2\x3a\x35\x05\x08\xa3\x2d\x50\x8b\x71\x16\x34\x13\x0a\x15\x80\x60\xe9\x00\x79\xa3\x0d\x7c\x11\x07\x63\x01\x2d\xa0\xd6\xae\xb4\x02\x92\xa2\x80\x97\xd9\x3a\x2a\xec\x47\x01\xcc\x98\x70\x7b\x84\x14\x2b\x02\xfc\x9b\xee\xd8\x9f\x96\x9b\xcc\x68\x90\x3a\x58\x7b\xe5\x55\x36\xa5\xd4\x4a\x43\x99\x57\x2c\x33\x51\xaa\xfb\xcc\x93\x52\x00\x00\x39\x53\x8e\x4c\x23\xd4\x5a\xc6\x10\x95\x92\x46\xcd\xd3\x1e\xe0\x54\x03\xfc\x6a\x2d\xbe\x61\x8e\x1b\x93\x19\x39\xc2\x67\x28\xcc\xce\x12\x07\x99\xb1\x6f\x93\x0f\xdd\xaa\x83\xda\xed\xd4\x1b\x57\xd0\x96\x7b\xfe\x32\xba\xca\xfa\x15\x5e\x12\x86\xd4\x51\x68\xc1\x8f\x3d\xa8\x20\xef\x48\xc6\x10\x16\xe2\x18\x77\x7d\x4a\x93\xb0\x85\x3e\xa8\xeb\x36\x23\x81\xca\x5f\xc2\xe4\xe9\xee\xa4\x83\x66\x68\x0b\x3a\x5c\x8f\x46\x4c\xda\xe4\x86\xac\x8c\x6f\xe4\xef\x18\xf8\xe6\x82\x02\x2b\x1a\x4d\x9e\x6a\xa3\x47\x10\x37\xbe\x6f\x35\xb8\x58\x35\x61\x96\x28\x69\x47\xb4\x57\xee\xe9\xae\xd6\x75\x37\xac\xf5\x5f\xdd\x97\xf5\x9f\xe4\xcd\x6f\x77\xfd\x1f\x29\x3a\x43\xf0\xdf\xb3\x3a\xab\x3f\x01\x00\x00\xff\xff\x32\x7d\x0f\x3e\xc1\x03\x00\x00"

func privateforwarderCreate_private_forwarderCdcBytes() ([]byte, error) {
	return bindataRead(
		_privateforwarderCreate_private_forwarderCdc,
		"privateForwarder/create_private_forwarder.cdc",
	)
}

func privateforwarderCreate_private_forwarderCdc() (*asset, error) {
	bytes, err := privateforwarderCreate_private_forwarderCdcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "privateForwarder/create_private_forwarder.cdc", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xcc, 0xbb, 0xa3, 0x44, 0xe6, 0x7e, 0xf8, 0xc4, 0x3c, 0x8, 0xf, 0x2a, 0xdc, 0x33, 0x93, 0xe1, 0xe3, 0xab, 0x1c, 0x9c, 0x7b, 0xb8, 0xea, 0x9b, 0xb3, 0x8d, 0x50, 0x17, 0xa8, 0x74, 0x3d, 0xd1}}
	return a, nil
}

var _privateforwarderDeploy_forwarder_contractCdc = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\xce\xb1\x0a\x83\x40\x0c\x06\xe0\xfd\x9e\x22\xa3\x82\x38\x17\x37\xc7\x2e\x45\x90\x4e\xa5\x43\x9a\x0b\x7a\xd0\xe6\x24\x17\xa7\xe2\xbb\x17\x6b\x15\x0b\xed\x12\x12\xfe\x1f\xf2\x99\xa2\x24\x24\x0b\x51\x32\x8a\x62\x8a\x64\x27\x7c\x70\x05\xad\x69\x90\xae\x00\x8a\x9e\x2b\xb8\x9c\x8f\x62\x87\x6b\x01\x89\xc5\xb3\xb6\x16\x15\x3b\x6e\xd0\xfa\xb9\xb9\x1d\x05\xa4\xbf\xc9\x30\xde\xee\x81\x96\xa0\xd9\xf6\x1c\x9e\xce\x01\x0c\xca\x03\x2a\x67\x29\x74\xc2\x5a\x41\x3d\x5a\x5f\x13\xc5\x51\xec\xd3\x00\x58\xb2\x72\x65\xa6\x12\xbd\xcf\xe4\x8d\xdd\xd3\x57\xf2\x3c\x7f\x78\xbf\x88\x7b\x55\x3e\x7f\x99\xdc\xe4\xc0\xbd\x02\x00\x00\xff\xff\x1f\xdb\x7f\x3e\x16\x01\x00\x00"

func privateforwarderDeploy_forwarder_contractCdcBytes() ([]byte, error) {
	return bindataRead(
		_privateforwarderDeploy_forwarder_contractCdc,
		"privateForwarder/deploy_forwarder_contract.cdc",
	)
}

func privateforwarderDeploy_forwarder_contractCdc() (*asset, error) {
	bytes, err := privateforwarderDeploy_forwarder_contractCdcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "privateForwarder/deploy_forwarder_contract.cdc", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x6f, 0x1b, 0x6d, 0x82, 0x8e, 0xce, 0xb0, 0x1e, 0xa3, 0x72, 0xe3, 0x7, 0x68, 0x37, 0xba, 0x88, 0x1f, 0x2c, 0xbf, 0x92, 0x53, 0x24, 0x47, 0x4f, 0x9c, 0xd3, 0x48, 0x97, 0xbf, 0x94, 0x84, 0xb}}
	return a, nil
}

var _privateforwarderSetup_and_create_forwarderCdc = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x54\xdd\x6e\xe2\x3c\x10\xbd\xcf\x53\xcc\x55\x45\x24\xda\x7c\xd7\x15\xad\x44\xdb\x50\x55\xdf\x8a\x22\x60\xbb\xd7\x26\x0c\x89\xd5\x60\x47\xce\x84\x1f\xa1\xbc\xfb\xca\xc6\x09\x71\x02\xb4\xda\xdd\x5c\x25\x99\x33\x67\x7e\xce\xb1\x3d\xbe\xce\xa4\x22\x18\x15\x22\xe6\x8b\x14\xe7\xf2\x13\x05\xac\x94\x5c\xc3\x7f\xbb\xd1\xcf\xf1\xeb\xdb\xd3\x8f\x70\xfe\xfe\x7f\x38\x1e\xbe\xbc\x4c\xc3\xd9\xac\x4a\x08\x77\x6c\x9d\xb5\xf0\xe7\x70\x13\xc5\x37\x8c\x70\x8a\x11\xf2\x0d\xaa\x91\x54\x5b\xa6\x96\xa8\xaa\x9c\xc9\xf4\xed\x63\x38\x0f\x47\xef\xd3\x5f\xc3\xe9\xcb\xdb\xf8\xb5\xca\xf7\x82\x00\xe6\x09\xcf\x81\x14\x13\x39\x8b\x88\x4b\x01\x6c\xb9\xcc\x81\xc1\x07\x2b\x52\xea\x03\x83\xec\xc8\x0e\xca\xd2\xc3\xaa\xe2\xd7\xe9\x0c\x16\x2c\x65\x22\x42\x88\x58\xc6\x16\x3c\xe5\xb4\xef\x03\x13\x4b\x9d\x59\x2c\x52\x1e\x35\x02\x3a\x15\x28\x39\x71\x79\x5e\xb3\xf2\xc1\xf3\x00\x00\x32\x85\x19\x53\xd8\xcb\x79\x2c\x50\xdd\xc3\xb0\xa0\x64\x18\x45\xb2\x10\xe4\x57\x18\xfd\xf0\x15\x1c\x21\x77\x0b\xa9\x94\xdc\x0e\x6e\x9a\x0b\xbb\x33\x03\x3c\xf6\xf4\x0e\xee\x21\xc8\x49\x2a\x16\x63\x80\x0d\x88\x41\xf8\xf0\xf0\x00\x82\xa7\x70\xa8\x89\xf5\x13\x04\xf0\xac\x50\xcf\xcd\x40\xe0\xd6\xd5\xc2\x24\x9a\x21\xb3\x82\x80\x13\x70\x01\xb6\x80\x43\x62\xdb\xcb\xd9\x06\x7b\x4e\x40\x3f\x83\x5b\xa7\xdd\xc8\x54\x0b\xd7\x19\xed\x0d\x7d\xcf\xef\x77\x52\x48\x5e\x9b\xc4\x81\xfb\xf5\x57\x79\xda\x98\xed\x27\xe5\xe2\x73\x70\x73\x70\xfc\x78\x57\xb9\xa7\x7c\x74\x5b\x0d\xac\xfe\x4e\xbd\x0a\xec\xb6\x48\x4c\xc5\x48\xdf\x6a\xd1\x3f\x35\x95\x22\xd5\x7e\x78\x3e\x59\xe5\xa1\xea\x36\x46\x3a\xfd\xbe\xd6\xf6\xd5\x4e\x1b\x05\x9b\xd2\x76\x2d\x4a\xd2\x38\xf4\x28\x31\x25\x8c\x40\x8a\x74\x0f\xb8\xcb\x64\x8e\x79\x93\x44\xc3\x2a\xf3\xaf\x38\xa6\x4b\xa0\x44\xc9\x22\x4e\x4c\xe4\xc9\x46\xb8\x20\x54\x2b\x16\xe1\x79\x15\xba\x96\x6d\x0d\x68\x79\xba\xb2\x98\xce\x9d\x59\x2d\xf4\x5f\x89\x52\x9f\x73\x18\xdc\x5e\xbc\x63\xac\x6f\xc7\xb8\xad\x7f\xf5\x14\x46\x3c\xe3\x28\xe8\xfe\x8c\xb0\x7e\xc7\x8e\xe6\x78\x0c\x6e\xeb\x72\x7d\xe3\xf3\x8b\x05\x5b\x81\xd9\x71\xac\x09\xa3\xa4\x4b\x7d\xdc\xf1\x45\xaa\xfa\xad\xb5\xdc\xef\xd6\x9e\x18\x0d\x74\xe9\xf3\x3b\xff\x83\x19\x3c\xf7\x00\x97\x5e\x69\x6e\xe9\xd6\x35\x19\x04\x06\xf3\xd5\x4d\x69\x61\x7f\x2d\xa8\xef\x50\x5d\x95\xad\xb6\x99\xfd\x7f\x36\xd5\x5e\x40\x5f\xeb\xa2\x5d\xdf\x48\x6f\x1a\xdf\xf2\xf7\xdb\xf1\x8e\xdf\x2d\xb0\x89\xf3\xab\x8f\x52\xbf\x94\xde\xef\x00\x00\x00\xff\xff\x5b\x2c\x92\xbe\xa2\x07\x00\x00"

func privateforwarderSetup_and_create_forwarderCdcBytes() ([]byte, error) {
	return bindataRead(
		_privateforwarderSetup_and_create_forwarderCdc,
		"privateForwarder/setup_and_create_forwarder.cdc",
	)
}

func privateforwarderSetup_and_create_forwarderCdc() (*asset, error) {
	bytes, err := privateforwarderSetup_and_create_forwarderCdcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "privateForwarder/setup_and_create_forwarder.cdc", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xf6, 0x80, 0x4b, 0x3d, 0xf9, 0x4d, 0x5f, 0xae, 0xf8, 0x82, 0xce, 0xb0, 0x85, 0x5c, 0xa3, 0x75, 0x1, 0xfb, 0x5d, 0x6, 0x40, 0xd0, 0xf0, 0x28, 0x34, 0x24, 0xde, 0x1, 0x4a, 0x6b, 0x3f, 0xc0}}
	return a, nil
}

var _privateforwarderTransfer_private_many_accountsCdc = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x93\x41\x6f\x9b\x40\x10\x85\xcf\xe6\x57\x4c\x73\x48\x6d\x29\xc5\x3d\x54\x3d\xa0\xb4\x11\xad\xc1\xb2\x5a\x39\x16\x76\xd2\x43\xd5\xc3\x9a\x1d\x60\x55\xbc\x8b\x66\x07\x63\xa9\xf2\x7f\xaf\x60\xb1\x45\x9a\x44\xbe\xf4\xb4\x12\xbc\xf7\x78\xdf\xcc\xa2\x76\x95\x21\x86\xb8\xd6\xb9\xda\x96\xb8\x31\xbf\x51\x43\x46\x66\x07\xef\x0f\xf1\xc3\x72\xbe\xf8\xf2\x3d\xda\xdc\x7f\x8b\x96\xe1\x6c\x96\x44\xeb\xb5\xd7\x1b\xa2\x83\xd8\x55\xff\xe8\x5f\xd2\xad\x48\xed\x05\x63\x82\x29\xaa\x3d\x52\x6c\xa8\x11\x24\x91\x4e\x9e\x55\xb2\x78\x0c\x37\x51\x7c\x9f\xfc\x08\x93\xd9\x62\x39\x3f\xf9\x3d\x26\xa1\xad\x48\x59\x19\x3d\x16\x3b\x53\x6b\x0e\xe0\x21\x56\x87\x8f\x1f\x6e\x80\x4d\x00\x3f\x43\x29\x09\xad\xfd\x35\x81\x3f\x9e\x07\x00\x30\x9d\xc2\xa6\x40\x78\x14\x75\xc9\x40\x68\x4d\x4d\x29\x02\x17\x82\xa1\x30\xa5\xb4\xc0\x05\x02\xb7\x8d\xad\x7b\x2a\x08\x61\x8b\x4a\xe7\xd0\x7d\x2c\x43\x22\x94\x5d\x54\x89\x0c\xfb\x36\x27\xc1\x2c\x80\xeb\x21\xad\xdf\xe5\x7b\x67\x59\xe5\x08\x7b\x32\xa5\xf3\x35\x6a\x89\x14\xc0\xf5\x6b\xec\xbe\x53\xb8\x88\x8a\xb0\x12\x84\x63\xab\x72\xdd\xba\xc2\x9a\x8b\x30\x4d\x5b\xde\x33\x58\x0f\x37\x47\x06\x01\x84\x19\x12\xea\x96\xcc\x74\x44\xce\xf9\xd6\x82\x65\x43\x28\x5d\xef\xb3\xcf\x62\x99\xf9\x27\x14\xf8\xd4\xab\xfd\xad\x21\x32\xcd\xed\x0b\x64\x9f\xc7\xed\x6a\x02\x98\xb6\x71\x22\xc7\x29\x0e\x24\x9d\x62\xe2\x8d\x46\xa3\xbb\x3b\xa8\x84\x56\xe9\xf8\xea\xab\xa9\x4b\x09\xda\x30\xb8\xd0\xe7\x0d\x4d\xe3\x0a\x76\xee\x37\x57\x13\xef\x69\xbb\x57\x26\xf8\xbc\xec\x85\x81\x9e\x9a\x5f\x90\xad\x1d\xd7\x4a\x70\xf1\x9f\x48\x8e\xee\xc0\x03\xa6\x35\xe3\x70\x6b\x99\x21\x10\xee\x9e\x82\xd2\x6d\xca\xe0\xe5\xa5\x01\xf8\x16\xb5\xec\x61\xba\xe9\xdb\x71\x9f\x75\xd3\x5f\xe3\x00\x6e\xdf\x3d\x59\xb0\xdf\x28\x2e\x24\x89\xe6\xfc\xcb\xb8\x73\x32\x98\xf9\xb1\xef\x7c\xf4\xfe\x06\x00\x00\xff\xff\x0d\x84\xe4\x91\xfd\x03\x00\x00"

func privateforwarderTransfer_private_many_accountsCdcBytes() ([]byte, error) {
	return bindataRead(
		_privateforwarderTransfer_private_many_accountsCdc,
		"privateForwarder/transfer_private_many_accounts.cdc",
	)
}

func privateforwarderTransfer_private_many_accountsCdc() (*asset, error) {
	bytes, err := privateforwarderTransfer_private_many_accountsCdcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "privateForwarder/transfer_private_many_accounts.cdc", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xc3, 0x39, 0xa5, 0x34, 0x9c, 0x6d, 0x13, 0xd1, 0x6d, 0x17, 0xde, 0x85, 0x18, 0x42, 0x31, 0xf4, 0x28, 0x20, 0xa3, 0x69, 0x44, 0x2d, 0x50, 0xb2, 0x1d, 0xfb, 0xdb, 0x2b, 0x7e, 0x99, 0x2b, 0xe0}}
	return a, nil
}

var _scriptsGet_balanceCdc = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x91\xcb\x4e\xf3\x30\x14\x84\xf7\x7e\x8a\x51\x17\xff\x9f\x6e\x12\x16\x88\x45\x45\xa9\x7a\x49\x10\x02\x15\xa9\x17\xf6\x4e\x72\xd2\x5a\x38\xb6\xe5\xd8\xb4\xa8\xea\xbb\xa3\xdc\x4a\x8b\x57\x96\xce\x37\xa3\x99\x73\xa2\x08\x9b\xbd\xa8\x50\x65\x56\x18\x07\x4b\x3c\xaf\xe0\xf6\x84\x94\x4b\xae\x32\x42\x21\x48\xe6\xd0\x05\xb8\x02\xcf\x32\xed\x95\xfb\x5f\x21\x91\xfa\xb0\xd1\x9f\xa4\x30\x6b\x39\xc6\x44\x69\xb4\x75\x48\xbc\xda\x89\x54\x52\x3b\x2d\xac\x2e\x71\x77\x4c\xb6\xcb\xe7\x97\xd9\x5b\xbc\x79\x7f\x8d\x97\xd3\xc5\x62\x15\xaf\xd7\xbd\x20\x3e\xf2\xd2\xfc\xe1\x6f\x38\x66\x7c\x8a\xc2\x2b\x94\x5c\xa8\xa0\xcb\x30\xc2\x34\xcf\x2d\x55\xd5\x70\x84\x6d\x22\x8e\x0f\xf7\x38\x31\x00\x90\xe4\xea\x9c\x0e\x63\xec\xc8\x4d\x5b\xba\x57\x0d\x2f\xc8\x17\xf7\xd2\xad\xa8\xc0\xb8\xa1\xc3\x1d\xb9\x39\x37\x3c\x15\x52\xb8\xef\x20\x32\x3e\x95\x22\x8b\xe8\x2a\x5b\x57\xb4\xb5\xa8\x5f\x98\x6a\x6b\xf5\xe1\xf1\xdf\x75\x83\xf0\xa3\x36\x3e\xdd\x6c\x21\xec\xa4\xe7\xa7\xe0\x57\x3d\x99\xc0\x70\x25\xb2\x60\x30\xd7\x5e\xe6\x50\xda\xa1\x35\xec\x57\x0a\x4b\x05\x59\xaa\x7f\x4e\x37\x37\x69\xbc\x07\x43\xd6\x98\x58\x72\xde\xaa\x4b\x91\xb0\x3b\x18\x3b\xb3\x9f\x00\x00\x00\xff\xff\x61\x86\x36\x18\xd4\x01\x00\x00"

func scriptsGet_balanceCdcBytes() ([]byte, error) {
	return bindataRead(
		_scriptsGet_balanceCdc,
		"scripts/get_balance.cdc",
	)
}

func scriptsGet_balanceCdc() (*asset, error) {
	bytes, err := scriptsGet_balanceCdcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "scripts/get_balance.cdc", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xe1, 0x1, 0x19, 0xb, 0x49, 0xe3, 0xc6, 0x11, 0x8e, 0x82, 0x8c, 0xb4, 0x29, 0x92, 0x11, 0xad, 0xe3, 0x33, 0xbc, 0x4b, 0xd9, 0x5a, 0x9c, 0x4a, 0xc1, 0x2b, 0xb6, 0x6d, 0x74, 0xc, 0x79, 0x47}}
	return a, nil
}

var _scriptsGet_supplyCdc = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x8e\xbd\x4a\xc5\x40\x10\x85\xfb\x79\x8a\x53\x26\x8d\x6b\x21\x16\x82\x85\x90\xd8\x08\x0a\x26\x3e\xc0\x9a\xec\x9a\xc5\xfd\x63\x76\x16\x22\x72\xdf\xfd\x42\x92\x5b\xa4\x9d\xef\x9b\x73\x8e\x52\x18\x17\x57\x50\x26\x76\x59\xc0\x46\xcf\x05\xb2\x18\x48\x12\xed\x51\x6a\xce\xfe\x0f\xd6\x19\x3f\x93\x52\x48\x76\x83\xfd\xaa\x43\xf6\x66\x4c\xbf\x26\xa2\x04\xcd\x82\x29\x45\x61\x3d\x09\x91\x0b\x39\xb1\x9c\x1d\xcb\x29\xe0\x7e\x1d\x3f\xde\xfa\xf7\x97\xae\xfb\xec\x87\x81\x28\xd7\x6f\xd8\x1a\x11\xb4\x8b\x4d\xfb\x84\xaf\x57\xb7\x3e\x3e\xe0\x9f\x08\x00\xbc\x91\x5b\xfd\xf3\x29\xec\x6e\x9b\x36\x6c\xe8\x50\xd3\x4f\xb3\xab\xed\x7e\x60\x23\x95\xe3\xf1\x4e\x97\x6b\x00\x00\x00\xff\xff\xce\x23\xa0\xa3\xe5\x00\x00\x00"

func scriptsGet_supplyCdcBytes() ([]byte, error) {
	return bindataRead(
		_scriptsGet_supplyCdc,
		"scripts/get_supply.cdc",
	)
}

func scriptsGet_supplyCdc() (*asset, error) {
	bytes, err := scriptsGet_supplyCdcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "scripts/get_supply.cdc", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xa5, 0xda, 0x29, 0xbe, 0x3f, 0x6c, 0xe7, 0xd3, 0x2b, 0x69, 0x4f, 0x10, 0xe8, 0x42, 0x81, 0x51, 0xb7, 0x6c, 0x34, 0x33, 0xc6, 0x40, 0x28, 0x93, 0x3a, 0xf8, 0x9b, 0x98, 0xa5, 0xed, 0x1d, 0x52}}
	return a, nil
}

var _setup_accountCdc = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x53\xc1\x6e\xdb\x3a\x10\xbc\xeb\x2b\xe6\x5d\x1e\x6c\xa0\xb5\x7a\x0e\xd2\x00\x4e\xe3\x14\x45\x8b\x14\x70\xdc\xde\xd7\xd2\xca\x26\x42\x93\xc4\x72\x99\x58\x30\xfc\xef\x05\x65\xd9\xb5\x5a\xa3\x28\x90\x43\x75\xe3\xee\x70\x76\x66\xb8\x2a\xca\x12\x8b\xb5\x89\x50\x21\x17\xa9\x52\xe3\x1d\x4c\x04\x41\x79\x13\x2c\x29\xa3\xf1\x92\x8f\x3f\xfb\xf9\x8e\x7a\x50\x5d\x83\xf0\x9d\x92\x55\x08\x47\x9f\xa4\xe2\x5c\xd7\x35\x1b\x01\x55\x95\x4f\x4e\x33\x36\xe6\x1a\x69\x6e\xb4\xa8\xc8\x21\x45\xce\x07\xf0\x96\x36\xc1\xf2\xc2\x3f\xb1\x2b\x0a\xb3\x09\x5e\x14\xf7\xc9\xad\xcc\xb2\xaf\xa2\x11\xbf\xc1\xbb\xed\xfd\xb7\x87\x8f\x9f\x6e\xbf\xcc\x16\x5f\x3f\xcf\x1e\xa6\x77\x77\xf3\xd9\xe3\xe3\xf1\xc2\xec\x8c\xe5\x88\x1f\xe0\x8a\x73\x6f\xbb\xa2\x00\x80\x20\x1c\x48\x78\x14\xcd\xca\xb1\x5c\x61\x9a\x74\x3d\x3d\x48\x1e\x1f\x31\xf9\x2b\x4b\xcc\x59\x93\x38\x30\x89\x6d\x61\x9a\x4e\x79\xef\x0e\x64\x85\xa9\x6e\x11\xd5\x0b\xe7\xd4\x06\x62\xba\x6c\x4e\x54\xa6\xc1\x61\xda\x64\xe9\x45\xfc\xcb\xf5\xff\xe7\xe0\x49\x07\xbe\x19\x65\x03\x57\x28\x33\x21\xad\xb8\x3c\x8f\xa8\x43\x8c\xf1\xdf\x7b\x38\x63\xb1\x3b\x11\xe7\x4f\x3a\x91\xa7\xd2\x7e\xe0\xe0\x83\x70\x7e\x48\x82\xe3\x97\x0b\x0a\x41\xae\x46\x48\x0a\xa3\x30\x0e\xfd\xe8\x13\x41\x2f\x3a\xd2\x33\x8f\x06\x33\xaf\xdf\x0e\x0c\x54\xdd\x94\xd9\x26\x68\xdb\xd1\x8e\xc6\x6f\x06\x70\xf5\x7f\xf2\x75\x82\x8e\x2f\x4b\x0f\x69\x69\x4d\x85\x8a\x02\x2d\x8d\x35\xda\xf6\xab\xd6\x5b\xe8\x16\xcc\x3b\xdb\x82\xb7\xc1\x47\x8e\xe7\x24\x19\x56\x73\xf0\xd1\x28\x9a\xe4\x0e\x8b\xa0\x6b\xf1\x69\xb5\xee\x9a\x73\xae\xd8\x3c\xb3\xc0\x38\x65\x69\xa8\xfa\xcd\xbd\x35\xee\xe9\xd2\x83\xed\x06\xeb\x3a\x39\x12\xed\x6f\x86\x51\x95\x07\xf9\x03\xdb\x47\xec\x2f\x29\x91\xac\x58\xff\x65\x52\x4b\xb2\xe4\x2a\x46\x63\xd8\xd6\x83\x98\x6e\xfb\xce\x6b\x53\xea\x79\xfe\x26\xa4\x1e\xfa\x8a\x8c\x0e\xbf\xc3\xbe\xf8\x11\x00\x00\xff\xff\xdc\x74\x82\x22\xea\x04\x00\x00"

func setup_accountCdcBytes() ([]byte, error) {
	return bindataRead(
		_setup_accountCdc,
		"setup_account.cdc",
	)
}

func setup_accountCdc() (*asset, error) {
	bytes, err := setup_accountCdcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "setup_account.cdc", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x7a, 0xd7, 0x69, 0x70, 0x86, 0xf8, 0xdd, 0xd3, 0x10, 0xb9, 0x9, 0xc2, 0x7e, 0xd6, 0x4a, 0x17, 0x97, 0x7c, 0xb5, 0x4b, 0x42, 0xf8, 0x4c, 0x66, 0xb, 0x48, 0xe4, 0x42, 0x88, 0x65, 0xa7, 0x5d}}
	return a, nil
}

var _transfer_many_accountsCdc = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x53\x4d\x6f\xd3\x40\x10\x3d\xc7\xbf\x62\xc8\x81\x3a\x12\x75\x38\x20\x0e\x51\xda\x2a\x34\x49\x85\x40\x45\x4a\x53\x38\x20\x0e\x6b\x7b\x6c\x2f\x38\xbb\xd6\xec\xb8\x09\xaa\xf2\xdf\xd1\xee\xda\x96\xd3\x58\x25\x97\x48\xf6\xbc\x8f\x99\xf7\x2c\x77\x95\x26\x86\x75\xad\x72\x19\x97\xb8\xd5\x7f\x50\x41\x46\x7a\x07\xef\x0f\xeb\xc7\xfb\xbb\xcf\x9f\xbe\xae\xb6\xdf\xbe\xac\xee\x17\xcb\xe5\x66\xf5\xf0\x10\x34\x80\xd5\x41\xec\xaa\x17\xf3\x27\x73\x01\x93\x50\x46\x24\x2c\xb5\x0a\xc5\x4e\xd7\x8a\x67\xf0\xb8\x96\x87\x8f\x1f\xde\x01\xeb\x19\xfc\x5c\xa4\x29\xa1\x31\xbf\x26\xf0\x1c\x04\x00\x00\xd3\x29\x6c\x0b\x84\xef\xa2\x2e\x19\x08\x8d\xae\x29\x41\xe0\x42\x30\x14\xba\x4c\x0d\x70\x81\xc0\x56\xd1\xf8\xa7\x82\x10\x62\x94\x2a\x07\x27\x96\x21\x11\xa6\x8e\xaa\x44\x86\x27\xcb\xb3\xc1\x6c\x06\x6f\xfb\x6e\x23\xc7\xef\x15\x2b\xc2\x4a\x10\x86\x46\xe6\x0a\x69\x06\x8b\x9a\x8b\x45\x92\x58\xb3\x9d\xab\xc6\xd9\x1d\x32\x08\x20\xcc\x90\x50\x59\x5b\xda\xd9\xf1\xc8\x0b\x03\x86\x35\x61\xea\x45\x3b\x9c\xc1\x32\x8b\x5a\x1f\x70\xd5\x4c\x47\xb1\x26\xd2\xfb\xf9\x80\xad\xeb\xd0\xde\x72\x06\x53\x4b\x27\x72\x9c\x62\x6f\xc4\x4d\x4c\x82\xd1\x68\x74\x73\x03\x95\x50\x32\x09\xc7\xb7\xba\x2e\x53\x50\x9a\xc1\x93\x9e\x3b\xd4\x7b\x6f\xd0\xa1\xdf\x8c\x27\xce\xdc\xd1\xef\x86\x07\x4c\x6a\xc6\xfe\xaa\x99\x26\x10\x3e\x19\x90\xca\x92\xf4\x5e\x36\xb7\xf8\x21\xb9\x48\x49\xec\xdb\x30\x5c\xfe\xff\xbf\x46\x1b\x8c\x41\xc5\x3e\xe4\xf9\xe5\xe9\x89\xa2\x7d\xc3\xdc\x35\xc6\xff\x4f\xce\x2c\xd8\x38\xac\x22\x61\x22\x2b\x89\x8a\x2f\x0c\x54\x75\x5c\xca\x04\x84\x0f\x10\x74\xfc\x1b\x93\x73\xf5\x0e\x01\x57\x90\x23\x37\x71\x87\xcd\xce\xc3\x4a\x03\xc1\xf7\x85\x37\x98\xa0\x7c\x42\x1a\xd2\x72\x2f\x7c\xfa\x1d\x24\xca\x91\x6f\x45\x25\x62\x59\x4a\xfe\x1b\x4e\xbd\xf1\x93\xac\x5b\xca\xc9\x09\xa7\xfd\x75\xed\x79\x3e\xf9\x68\xa3\x16\x71\xbc\x0e\xcf\x41\xaf\x16\xc6\xe3\x5e\x5f\xd1\xe5\x35\x3e\xbf\xce\x12\x2b\x6d\xa4\xcf\xa2\x0d\x4f\xb5\xbd\xb0\xfd\x79\xc1\x43\x43\xa7\xea\x9d\x29\x4a\x3d\x61\xf3\x1d\xcc\x2f\xbb\xb2\xf4\xb4\x8f\x4d\x87\x8f\xc1\xbf\x00\x00\x00\xff\xff\xe4\xb3\xe6\xcc\xbf\x04\x00\x00"

func transfer_many_accountsCdcBytes() ([]byte, error) {
	return bindataRead(
		_transfer_many_accountsCdc,
		"transfer_many_accounts.cdc",
	)
}

func transfer_many_accountsCdc() (*asset, error) {
	bytes, err := transfer_many_accountsCdcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "transfer_many_accounts.cdc", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x14, 0xde, 0x24, 0x74, 0xef, 0x77, 0x97, 0xf3, 0xcd, 0xba, 0xbc, 0xae, 0x97, 0x86, 0xfd, 0x5d, 0x4e, 0x78, 0xbe, 0x2a, 0x47, 0x3a, 0xe9, 0xdf, 0x87, 0x83, 0x8b, 0xdc, 0xb9, 0x6a, 0x92, 0xfb}}
	return a, nil
}

var _transfer_tokensCdc = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x53\xcd\x6e\xdb\x3c\x10\x3c\x47\x4f\xb1\x5f\x0e\x5f\x64\xa0\x91\x7a\x28\x7a\x30\xf2\x53\x37\x71\x82\xa2\x45\x0a\xe4\xa7\x3d\x53\xd2\xda\x62\x2b\x91\xc4\x72\x15\x3b\x08\xfc\xee\x05\x49\x51\x95\x9c\xa2\xa9\x2f\x86\x56\xb3\xb3\xb3\x33\xab\x3c\x87\xfb\x5a\x5a\x60\x12\xca\x8a\x92\xa5\x56\x20\x2d\x08\x60\x6c\x4d\x23\x18\x61\xa5\xc9\x3d\x8e\xde\x73\x2d\x38\xc9\x73\x28\x75\xd7\x54\x50\x20\x74\x16\x2b\x28\x9e\x40\xa8\x27\xad\x10\x58\x83\x45\x55\x01\xeb\x9f\xa8\xac\x7b\x14\x4a\x73\x8d\x04\xa2\x2c\x75\xa7\x7c\xb3\x23\x81\x5a\x58\x28\x10\x15\x58\x64\xe8\x8c\x83\x12\x96\x28\x1f\xb1\x6f\xce\x92\x3c\x4f\xbc\x46\x84\x8d\xe4\xba\x22\xb1\x01\xd1\x3a\x12\x10\x6e\x44\x8d\x91\x14\x56\xa4\x5b\x58\x23\x2f\x7e\x0f\xd9\x44\x85\x0e\x67\x04\x89\x16\x19\xc9\x4b\x72\x95\xd1\x52\x49\x22\x5b\xa3\x89\xe1\xaa\x53\x6b\x59\x34\x78\xef\xe6\x07\xce\xb7\xdb\xab\x87\x9b\xeb\x4f\x1f\xbf\x2c\xef\xbf\x7e\x5e\xde\x2c\x2e\x2f\x6f\x97\x77\x77\xb1\x61\xb9\x15\xad\xd9\xc3\x4f\x70\xc9\x68\x4c\x1a\xb4\xcf\xe1\xe1\x4a\x6e\xdf\xbf\x7b\x03\xac\xe7\xb0\xa8\x2a\x42\x6b\x67\xf0\x9c\x24\x00\x00\xfd\xbe\xdf\x44\xd7\x30\x10\x5a\xdd\x51\x89\xbd\x61\xba\xa9\x6c\xd0\xde\x9b\xeb\xaa\x82\x10\x0a\x94\x6a\x1d\x36\x5a\x21\x11\x56\x9e\xaa\x41\x76\x59\xb0\xe7\x9a\xc3\x87\xc9\x76\x99\xaf\x86\x99\x86\xd0\x08\xc2\xd4\xca\xb5\x42\x9a\xc3\xa2\xe3\xba\x37\x72\xd0\xd5\x6b\xbb\x46\x06\x01\x84\x2b\x24\x54\x25\x46\x33\x43\xe7\x91\x05\xcb\x9a\xb0\x82\x47\x4f\x1e\xfb\x9c\x10\x5f\xb9\xc5\x15\x9c\xf6\xe0\xac\xd0\x44\x7a\x73\xf2\xff\xd8\xc3\xa0\xea\x2c\x75\x56\xce\x21\x77\x6c\x62\x8d\x39\x8e\x20\x1e\x31\x4b\x0e\x0e\x0e\xce\xcf\xc1\x08\x25\xcb\xf4\xf0\xc2\x47\xad\x34\x43\x20\x7d\x29\x50\x6f\x82\x3e\xdf\xfd\xdf\xe1\x6c\xb2\xd4\xf7\x78\x5c\xbd\xaf\x3e\xc8\xd7\xd7\xb2\xd8\xac\xb2\xc1\x60\x38\x39\x1e\x96\xcc\xe2\xb9\x0e\x91\x87\xff\x99\xef\xdd\x85\xe1\xb8\xc5\xb2\x63\xfc\x83\xc1\x6e\x34\x61\x29\x8d\x44\xc5\x47\x16\x4c\x57\x34\xb2\x1c\x6e\x5d\x17\x3f\xb0\x9c\xba\x3b\xa0\xe1\x74\xf4\x15\xa4\xac\x67\xff\x92\xde\x78\xd6\x6d\xf8\x04\x69\x9f\xde\x17\x43\x7e\x03\x3c\x5b\x23\x5f\x08\x23\x0a\xd9\x48\x7e\x4a\xf3\xa0\x73\x92\x56\xa4\x9b\x0d\x7c\xee\x37\x64\xff\x3c\xbd\xc9\x88\xde\x9d\xa5\xaf\x27\x1c\xa0\x7f\xdf\xc6\x27\xb3\x97\xf6\x25\x1a\x6d\x65\x70\x39\xe6\xa4\x62\xf4\x52\xbd\xe0\xa0\x7d\x47\x46\x6e\x64\x55\x20\xeb\x0f\xf6\xe4\x78\x7a\x13\x31\xef\x5d\xf2\x2b\x00\x00\xff\xff\x28\xad\x7c\x0d\x6b\x05\x00\x00"

func transfer_tokensCdcBytes() ([]byte, error) {
	return bindataRead(
		_transfer_tokensCdc,
		"transfer_tokens.cdc",
	)
}

func transfer_tokensCdc() (*asset, error) {
	bytes, err := transfer_tokensCdcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "transfer_tokens.cdc", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xa2, 0xbd, 0xf, 0x8, 0xff, 0x1b, 0x57, 0x63, 0xc1, 0x37, 0xe4, 0xa8, 0x72, 0xad, 0x66, 0x2a, 0xc, 0x8e, 0xcb, 0xb7, 0x45, 0x7d, 0xef, 0x8b, 0xa7, 0x2, 0x48, 0xe6, 0x35, 0x24, 0x5e, 0x4f}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"burn_tokens.cdc":      burn_tokensCdc,
	"create_forwarder.cdc": create_forwarderCdc,
	"mint_tokens.cdc":      mint_tokensCdc,
	"privateForwarder/create_private_forwarder.cdc":       privateforwarderCreate_private_forwarderCdc,
	"privateForwarder/deploy_forwarder_contract.cdc":      privateforwarderDeploy_forwarder_contractCdc,
	"privateForwarder/setup_and_create_forwarder.cdc":     privateforwarderSetup_and_create_forwarderCdc,
	"privateForwarder/transfer_private_many_accounts.cdc": privateforwarderTransfer_private_many_accountsCdc,
	"scripts/get_balance.cdc":                             scriptsGet_balanceCdc,
	"scripts/get_supply.cdc":                              scriptsGet_supplyCdc,
	"setup_account.cdc":                                   setup_accountCdc,
	"transfer_many_accounts.cdc":                          transfer_many_accountsCdc,
	"transfer_tokens.cdc":                                 transfer_tokensCdc,
}

// AssetDebug is true if the assets were built with the debug flag enabled.
const AssetDebug = false

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"burn_tokens.cdc": {burn_tokensCdc, map[string]*bintree{}},
	"create_forwarder.cdc": {create_forwarderCdc, map[string]*bintree{}},
	"mint_tokens.cdc": {mint_tokensCdc, map[string]*bintree{}},
	"privateForwarder": {nil, map[string]*bintree{
		"create_private_forwarder.cdc": {privateforwarderCreate_private_forwarderCdc, map[string]*bintree{}},
		"deploy_forwarder_contract.cdc": {privateforwarderDeploy_forwarder_contractCdc, map[string]*bintree{}},
		"setup_and_create_forwarder.cdc": {privateforwarderSetup_and_create_forwarderCdc, map[string]*bintree{}},
		"transfer_private_many_accounts.cdc": {privateforwarderTransfer_private_many_accountsCdc, map[string]*bintree{}},
	}},
	"scripts": {nil, map[string]*bintree{
		"get_balance.cdc": {scriptsGet_balanceCdc, map[string]*bintree{}},
		"get_supply.cdc": {scriptsGet_supplyCdc, map[string]*bintree{}},
	}},
	"setup_account.cdc": {setup_accountCdc, map[string]*bintree{}},
	"transfer_many_accounts.cdc": {transfer_many_accountsCdc, map[string]*bintree{}},
	"transfer_tokens.cdc": {transfer_tokensCdc, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory.
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
