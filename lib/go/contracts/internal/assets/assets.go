// Code generated by go-bindata. DO NOT EDIT.
// sources:
// ../../../contracts/ExampleNFT.cdc (7.169kB)
// ../../../contracts/MetadataViews.cdc (7.242kB)
// ../../../contracts/NonFungibleToken.cdc (4.832kB)

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

var _examplenftCdc = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x59\x4f\x6f\xe3\xba\x11\xbf\xfb\x53\xcc\xcb\x61\x2b\xa3\x89\xd5\x02\x45\x0f\x46\xbc\xd9\xc5\xe6\x19\xcd\xe1\x05\x8b\xac\xdb\x1e\x16\x41\x1f\x2d\x8e\x63\x22\x12\x69\x90\x94\x5d\x35\xcf\xdf\xbd\x18\x52\x7f\x28\x89\x4a\xf2\xba\x97\x0a\x8b\x5d\x4b\x9c\x19\xce\xff\xf9\x91\x9b\xa6\xb0\xd9\x0b\x03\xc2\x00\x93\x80\xff\x66\xc5\x21\x47\x10\xf4\x77\x81\xd2\x32\x2b\x94\x04\xb5\x03\x06\xeb\x5c\x9d\xe0\x5e\xc9\xab\x75\x29\x9f\xc4\x36\x47\xd8\xa8\x67\x94\xb3\x34\x85\x3b\x4b\xfc\x52\x59\x38\x30\x6d\x89\xdc\xee\x11\xd4\x6e\x27\x32\xc1\x72\x30\x96\x49\xce\x34\x87\x6d\x69\x41\x58\x60\xc6\x94\x05\x72\xb0\x0a\xb6\x48\xfc\x47\xd4\x15\x18\x51\x88\x9c\x69\xfa\xba\x57\x27\x28\x98\xac\xe0\x7e\xbd\x31\x70\x52\x65\xce\x3b\x95\x9c\xec\x4c\x69\x84\x5d\x29\x33\xd2\x8f\xe5\xc2\x56\x8b\xd9\x4c\x14\x07\xa5\x2d\xe9\xd8\xa8\xe8\x34\x84\x9d\x56\x05\x5c\x2c\xd2\xe1\xc2\x22\xe3\xd9\x45\xc3\xf5\x0b\x5a\xc6\x99\x65\xff\x10\x78\x32\x2d\x4b\xef\xab\xa7\x9f\x1d\xca\x2d\x64\x4a\x5a\xcd\x32\x0b\x3f\x7b\x8f\xdd\xaf\x37\xcb\xf1\xc6\x2f\xb3\x19\x00\x00\x31\x1c\x9d\x65\x96\xe5\xdf\xca\xc3\x21\xaf\x96\xf0\xf7\x3b\x69\xff\xfa\x97\x8e\x00\x8f\x64\xdb\x97\x5a\xee\x9d\x14\x56\xb0\x5c\xfc\x07\x79\x32\x1f\xd0\xfc\x53\xd8\x3d\xd7\xec\x94\x08\xde\x88\xb9\x74\x0a\x2f\xe1\x33\xe7\x1a\x8d\xb9\x19\xb2\xdc\xe2\x41\x19\x61\x7b\x1c\x56\x85\xf4\x2d\x43\x8e\xa4\x45\x9e\xa3\x73\xed\x37\xab\x34\x7b\xc2\xaf\xcc\xee\x97\x10\xbc\x4c\x90\x7f\x2d\xb7\xb9\xc8\x3c\x75\xf7\xbb\x47\xfc\x8b\x90\x16\xf5\xa4\xdc\x96\x56\xa3\x51\xa5\xce\x10\xa2\xae\x5d\xdc\xdd\xaf\x37\x97\xfd\xa0\x2d\x1e\xd0\xa8\xfc\x88\x1a\x5e\x9c\x94\x70\xd7\xce\xf0\xd9\x68\x4d\xb2\x02\x49\x09\x2d\xe4\xd3\x68\x91\xa3\xc9\xb4\x38\x90\x71\x93\x34\x76\x5f\x16\x5b\xc9\x44\x3e\xa2\x60\x59\x86\xc6\x24\x06\xf3\xdd\xdc\x91\x6a\x55\xb1\xdc\x0a\x34\x4b\xf8\x3e\x50\xde\xad\x54\x8f\x9d\x7e\x42\x0a\x9b\xb4\x6f\xee\x4b\x17\xbe\xde\xf7\xd0\x82\xfe\x4a\x44\xfd\x3e\xc1\x48\xf7\xfe\xf2\x3b\xf4\x6d\x48\xe7\x81\xdb\xe9\x21\xa3\x17\x82\xc3\x0a\x04\x1f\x2f\x90\xca\xb0\x72\x9a\x8f\x17\x03\xad\x61\x15\xda\x30\x26\x6d\xf5\x87\x55\x67\xcb\x98\xac\xb5\x03\x56\x9d\x4d\x2d\xd9\xd9\xfd\xea\xc5\x75\x57\x4a\x78\x42\xeb\x6c\x4d\xe6\x4b\xf8\xbe\xa9\x0e\xf8\x38\x30\x51\xa3\x2d\xb5\x84\xef\xbd\x8f\xf4\x10\xf1\x75\xdf\x5f\xb7\xc2\x1c\x72\x56\x7d\x4c\xe6\x97\xef\x21\x7f\x68\x94\xfc\x58\xd7\x7f\xf3\x3c\x06\x5a\x8f\x34\xd6\xbe\x02\x48\x44\xf2\x2f\x38\x0a\x3c\x2d\x9d\xf0\xf9\x12\x3e\xcb\xea\x9b\xd5\x65\x66\x6f\x86\x71\x3a\x09\x9b\xed\x1d\xf1\x60\x85\x9e\x8c\x19\x7c\xdd\x9c\xe5\x88\x27\x70\x4d\x94\x29\x89\x72\x40\x9b\xc8\x6d\x82\x8c\x3d\xd5\x3c\xbd\xbc\x1e\xe6\xcc\x34\x5b\x90\xed\x7d\xcd\xfe\xb6\xd9\x7c\x5d\x8b\x1c\xa7\x55\xa3\xa7\xd4\xf9\x72\x90\x76\x93\xf4\xf3\xe8\xca\xf8\xeb\x94\x83\xc3\x04\xf8\x1d\x2e\x6e\xd9\xa6\x2d\xe9\x17\xc4\x3b\xd4\x3c\xc7\x92\x5e\x06\xb6\x7b\x82\x73\xa4\x79\xbb\x66\xbf\x63\x19\x06\x83\x72\x38\x2e\x06\xed\x9a\xf2\x98\xd7\xf3\xca\x52\xaf\x5f\xc2\xa7\x51\xfb\xbf\x5f\x6f\xe6\xb1\x7a\xbd\xbb\xf5\xd5\xea\xbb\xe4\xe3\x88\x64\xab\xb4\x56\xa7\xfb\xf5\x26\x18\x85\xf3\x25\x7c\x88\x6d\x30\xc1\xdc\x19\x32\x90\xd1\x2d\x10\xf7\xb0\xcc\x0e\xca\xd8\x48\x7d\x25\x1a\x4d\x99\x5b\x58\xad\xc8\xa3\x73\xf8\xed\xb7\xe6\xd3\x8d\x6b\x9e\xd4\x3d\x27\xe2\x7f\xf1\x85\x49\xc2\x5c\x5e\xad\xc0\xc1\xa0\x71\x87\x1a\x65\x86\x4b\x07\x96\xee\x6e\x1b\x48\xe6\x63\x87\xbc\xa3\x20\xe0\x26\x64\xa6\xb4\xc6\xcc\x5e\x4c\x84\x7d\x3a\xbe\x5d\x2c\x97\xaf\x44\xf8\x72\x3c\xbf\xbf\x6a\x75\x14\x1c\x75\x64\xe9\x01\x33\x14\xc7\xe8\xd2\x58\x70\x1c\x01\x74\x74\x81\xcb\xd3\x14\xb8\xf0\x80\x51\x57\xe4\x11\x72\x55\xa6\xe4\x4e\xe9\x42\xc8\x27\x70\xc9\x66\x42\x72\x22\x20\x60\xdc\xd9\x6b\xab\x03\xc2\x49\xd8\x3d\xa1\xe5\x5f\x7d\xec\x7f\x25\x07\xef\x04\xe6\xbc\x97\x31\x84\xf8\xd4\x49\x22\x27\x10\xbb\x84\x4f\x2f\x9e\x3a\x82\x65\xee\xd7\x9b\x73\x7f\xec\x43\x12\x9d\xa6\xad\x38\xb8\xbe\x82\x97\x73\x6c\x12\xa4\xa9\x53\x8f\xf0\x21\x68\x2c\xd4\x11\x1d\xb0\x27\x4b\x1c\xa6\xf5\xe0\xb9\xf5\x0e\x93\x1c\x3c\x91\xb0\x84\xbc\xdd\x32\xcb\x73\xd4\xa3\xec\x6f\xc4\x26\xcd\x8f\xbb\xdb\x20\xfb\xa3\x25\x3a\xb0\xc1\x61\x25\x07\x8c\xaf\xaf\x06\x06\x2d\xbc\xae\xc9\x33\x56\x4b\xe8\x36\x98\xc3\xcd\x0d\x1c\x98\x14\x59\x72\x51\x08\x63\x28\x4c\xf7\xeb\xcd\xc5\x7c\xd6\x13\x8c\x85\x18\xc0\x62\xb7\xcd\x42\xf0\x06\x18\xb7\xbb\xe9\x9b\x05\xf3\xa0\x77\x20\xa3\x6e\x6b\xd7\x57\x8e\x75\xc2\xb5\x75\x5f\x02\xcb\x9e\xc9\xaf\xce\xad\xe4\x42\xc6\x79\xcf\x83\xad\x83\x4d\x90\x72\xa1\xa0\x96\xa9\xae\xcf\x9a\x51\x70\x60\x5a\xb3\xea\x7f\x6b\x88\xaf\xb9\xdb\xff\x60\xe6\x27\xf8\xd4\xef\x53\xb3\x11\x4f\xd7\xd5\x08\x53\xd5\x8e\xec\x93\x91\x05\x9c\x3b\x95\x25\x9e\x6a\xe1\xb5\x0d\x41\x8d\x9d\xf6\x22\xdb\xb7\x69\xe8\xce\x84\x39\x07\x25\x71\xb4\xa7\xca\xf9\x26\x9e\x19\xdf\x05\x7f\x6c\x0d\x88\x84\x3d\x3c\xda\x50\xbc\xe9\x58\xf3\x76\xb4\x39\x1a\xab\x55\xd5\xee\x3b\x11\x6f\x3f\x51\xea\xdc\x70\x85\xe4\xc2\xd3\xb4\x53\x5a\xb3\x7b\x66\x81\x69\x1a\x75\x83\xd8\xbf\x63\x3e\xc5\xf1\xe4\xa0\x34\x9e\xb1\x32\x13\xfa\xb5\xe3\x8c\x64\xfb\x46\xd5\xf4\x75\xab\x9a\xba\x9f\x56\x2c\x4d\xc1\x28\x6f\x41\x57\xf8\x90\x31\x02\x92\x8c\x83\xb0\x06\x8a\xba\xbf\xba\x8c\x25\x82\xe6\xeb\x5e\x71\xf3\x43\xe3\x35\x6e\xfb\x87\x48\xf4\x99\x79\x63\x40\x9f\x67\x63\xf0\xfe\x43\xc3\x5a\xec\x62\x59\xf8\x93\x9b\xd1\x91\x21\x9e\xa6\xf0\x45\x23\xb3\xe8\x32\xa4\xb4\x7b\xa5\xe9\xe4\x3e\x88\x46\x9e\xab\x13\x70\x75\x92\x19\x33\x36\x3c\x25\x86\x85\xa0\x71\x07\xab\x29\x2f\x90\xe8\x37\x5c\x31\x70\x27\x89\xa3\xa2\x1f\xd8\x3b\x98\xf3\x6f\xe3\xbb\x09\xf7\xd2\xd4\x6d\x66\xee\xc0\xc1\x9f\x65\xf5\x50\x4f\xcd\x97\xf8\x90\x3e\x47\xfa\x95\xdc\xd9\x1f\x36\x9f\xe4\x60\x07\x86\x56\x4e\xe8\x5b\x4e\xa8\xad\x0e\xf8\x28\xeb\xde\x61\x44\xcc\x49\x75\x77\x19\x0d\xf1\xa6\xeb\xf4\xcd\x8b\xc3\xac\x34\x25\x5f\x13\x3c\x6e\xae\xb9\xea\x56\x23\x2b\x25\xd1\x15\xa9\x2b\x47\xab\x20\xab\x73\xcf\xf5\x62\x2c\x0e\xb6\x1a\x16\x7b\x13\x35\x4f\xf9\x33\x91\x74\x10\x29\x89\x8e\xef\x28\x84\x6a\x87\x64\xb3\x67\x28\x65\xa0\xfd\x43\x8b\x99\xbc\xda\xc0\x78\x21\x24\x28\x0d\x46\x51\xff\xa0\x59\xde\xdc\xf9\xf9\x2b\x3e\x75\x92\xf5\x9d\x60\x2d\x82\x6d\x73\x57\x3a\x85\x90\xd6\x19\xd7\xba\x2b\x4d\xa3\x17\x45\xfe\x72\xa9\xb9\x77\xab\xa5\x10\x37\x05\x94\xfe\x35\xb5\x97\xe8\xdd\xc3\x38\xf7\x7a\x77\x3b\x1c\xce\xcd\xa4\xa7\x3f\xb2\xc6\xce\x99\x38\x08\x24\x19\x01\x80\x2a\x1d\x26\xb1\x7b\x14\x3a\xfc\xdc\x56\xfe\xa8\x70\x6a\x6d\x92\x41\xf6\xd5\xb2\x97\xf0\xe1\xe5\x4d\xd4\x7b\xfe\x3f\xbd\xfd\x19\x42\x84\x5e\x5e\x0e\x0b\x8e\xe0\xb1\x44\x37\xbb\xba\x74\x1a\x39\x06\xea\x5b\xaf\xa0\x72\x83\x5b\xd4\xf1\x41\xdf\xfb\x22\x7e\x77\xd0\xf3\xc6\xab\xd7\x05\x81\x5b\xda\x9f\x63\xaa\xc0\x3b\xf1\xd3\xf4\x7c\xe4\x91\x57\xb2\xea\x0f\x06\x58\x96\xa9\x52\xda\x5e\x4e\x8d\x13\x09\xc2\x7c\x59\x0c\x80\xe1\xf5\x95\xf7\xe9\x60\xeb\xb8\xfb\x60\x35\xb5\xf0\xc7\xba\x91\x27\x7f\x9e\xc7\xbb\x93\xbb\x9b\x9c\xf7\x0f\x57\xdd\x9d\xb5\xb3\xcc\xc9\x03\xe3\x04\xb6\x64\xfe\xf2\xa4\xa7\xc2\x9f\x7a\xc5\xfa\x0d\x3d\x16\xa1\x10\x72\x38\x30\xbb\x37\x7d\xe6\xe8\xdd\x34\xac\x20\x35\xfe\x35\xc5\xc8\x31\x74\x4a\x44\x77\x47\x4d\x12\x7c\xbf\x7d\x87\x80\xd1\x1d\x76\x7c\x7f\x4f\xd6\x33\xaf\x81\x09\x41\xe3\xec\x1a\x18\xf5\x1d\xc3\x8e\x58\x9f\x24\x6a\x81\x2d\x3b\xcd\xb5\xa0\xc5\xbc\xd2\x84\x5b\x45\xeb\x8c\x5a\x90\xd4\xe4\xfa\xaa\xe3\x0e\x70\x72\xd4\xa1\xf3\x9e\xd6\x6d\x21\xd7\x13\x29\x63\x07\xb6\x15\xb9\xb0\x15\xec\x94\x9e\x42\x97\x3d\x0d\x72\x21\x9f\xaf\xc3\x01\xdc\x6d\xfb\x76\xc3\xbb\x0c\xf3\x74\xfa\x9a\xe1\xfc\x31\x19\x9f\x99\x63\xc1\x1e\x34\x41\xa6\x9f\xd0\xbe\xe6\x8d\x59\xa4\xa2\xc3\x60\xd6\x63\xe7\xf7\x04\xb2\xf0\x2c\xbd\xd6\xe7\xc5\xbc\x11\x43\xcf\x18\xc4\x6f\x94\x8c\x81\x92\xee\x84\x34\xfd\x7f\x4a\xe7\xd9\x79\xf6\xdf\x00\x00\x00\xff\xff\xfd\x10\x37\x2a\x01\x1c\x00\x00"

func examplenftCdcBytes() ([]byte, error) {
	return bindataRead(
		_examplenftCdc,
		"ExampleNFT.cdc",
	)
}

func examplenftCdc() (*asset, error) {
	bytes, err := examplenftCdcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "ExampleNFT.cdc", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x8f, 0x44, 0x5e, 0x4a, 0xf7, 0xbc, 0xba, 0x5e, 0x92, 0x41, 0x30, 0xf5, 0xf8, 0xf6, 0x9, 0xf3, 0x66, 0xe4, 0xd5, 0x0, 0x12, 0x11, 0xcd, 0xd2, 0xd8, 0xe6, 0xc9, 0x2, 0x6f, 0x17, 0xe3, 0x40}}
	return a, nil
}

var _metadataviewsCdc = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x59\x6d\x6f\xdc\x36\x12\xfe\xbe\xbf\x62\xe0\x03\x02\xaf\x6f\x77\xe5\xe4\x5a\xa3\xb7\xa8\x5b\x18\x76\x9d\xf3\xa1\xc9\x19\x8e\xd3\x2f\x41\x50\x53\xd2\x68\x35\x17\x8a\xd4\x91\xd4\xae\x85\xc0\xff\xfd\x30\xa4\x5e\x77\x65\x5f\x5a\x9c\x3f\xb4\x2b\x89\x1c\xce\xeb\x33\xcf\x30\xd1\xc9\xc9\x6c\x76\x9f\x93\x85\x44\x2b\x67\x44\xe2\x80\x8a\x52\x62\x81\xca\x59\x70\x39\x42\x81\x4e\xa4\xc2\x09\xb0\x4e\xa8\x54\x98\x14\x4a\xa3\x4b\x6d\x31\x9d\x91\x82\xeb\x5f\x6f\x6e\x97\xa7\x67\x7f\x3b\x5b\xcd\x66\x77\x98\xad\x21\x77\xae\xb4\xeb\x28\xda\x90\xcb\xab\x78\x95\xe8\x22\xd2\x2a\x93\x7a\x17\xf9\xff\xc4\x52\xc7\x51\x21\xac\x43\x13\x65\x92\x4a\x1b\xbd\x39\x7d\xf3\xfa\xf4\xef\xaf\xcf\x96\x2a\x73\xcb\xf6\xb0\x55\x91\xce\x66\x1f\x9c\xa9\x12\x67\x41\xa8\x14\x0c\x5a\x5d\x99\x04\x2d\x24\x42\xf5\x2a\x82\x56\x08\xda\x40\xa1\x0d\xce\x3a\x4d\x5d\x5d\xa2\x5d\x40\x22\xa4\xc4\x14\xb6\x84\x3b\xbb\x82\x5f\x44\x92\xfb\xdf\xfe\x33\x18\x2c\x0d\x5a\xb6\x72\x26\x20\xa5\x2c\x43\xc3\xf2\xbe\x90\x4a\x41\x67\x9d\xd5\x0b\xb0\x55\x92\x83\xb0\x20\x20\x31\x28\x9c\x36\x10\x93\xde\x18\x51\xe6\xf5\x4c\x1b\x10\xf0\xcf\xdb\x5f\xde\x02\x15\x62\x83\x90\x91\xc4\xd5\xec\x24\x9a\xcd\xa8\x28\xb5\x71\x70\x5d\xa9\x0d\xc5\x12\xef\xf5\x17\x54\x90\x19\x5d\xc0\xe9\x23\xe2\x0f\x6f\x7e\xf8\xfe\x2c\xce\xde\x9c\xe2\x1b\x21\xce\x66\xb3\xb2\x8a\x7b\xff\xbf\x6b\x8e\xfe\x8d\xf5\x86\xaf\xb3\x19\x00\x40\x14\x45\x70\x01\x77\x68\xb5\xdc\xa2\xe1\x10\x6c\x29\x45\x0b\x22\x49\xd0\x5a\x70\x1a\x04\x58\x74\x43\xd5\x1b\xc3\xdb\xed\x03\x31\xd6\x3b\x96\xfd\xd6\xba\x15\x8e\x71\xb5\x59\x81\x50\xf0\xfe\xfa\x7e\xbe\xe7\x63\xc7\xe9\x41\xca\xa1\xc9\x44\x82\x9d\x1c\xa7\x5b\x35\x06\x5a\x70\xc6\xf8\x73\xc1\xe5\xc2\x01\x39\xb0\x55\xc9\xae\xd8\x53\x84\x2d\xee\x0e\xef\x64\xf7\x06\x7e\xf5\xab\xda\x95\x59\xa5\x60\x83\xce\x7b\xe4\x78\xbe\x86\x4f\xf7\x75\x89\x9f\x0f\x96\x98\xb0\x9b\x97\x1d\xff\xee\xd5\x58\x03\xaf\x9c\xaf\xe1\x42\xd5\x21\x9d\x7e\xf6\xbb\x9e\xa6\xbc\x7a\xa9\xa5\xc4\xc4\x91\x56\x40\x1c\xef\x8d\xd1\x55\xc9\x1e\xf5\x59\xd3\x08\x37\xec\x8a\x14\x1f\x21\xae\xe1\xe6\xea\x0f\x19\x35\x90\x7f\x68\x5e\xac\x8d\xd1\x3b\x56\xbd\x5d\x7e\x4c\xe9\x1a\x3e\xde\x28\x77\xf6\xdd\x7c\x0d\xaf\xbe\xb6\xef\x9f\xa6\x5c\x73\x73\x15\x1c\x13\xd6\x7f\xde\x37\xf2\x8a\x6c\x29\x45\x1d\xec\x8a\x85\xa5\xa4\x29\x05\x1f\x24\x95\xc8\x8a\x93\x89\x83\xa7\x44\x81\x0b\x48\xd1\x26\x86\x4a\xaf\xab\x50\x69\x1f\xf3\xbc\x2a\x62\x25\x48\x42\xc6\xb9\xaf\x40\xc7\xff\xc6\xc4\xad\xe0\x9d\xb6\xae\x79\xb0\x60\x73\x5d\xc9\x74\x3f\x83\xf8\xc0\x43\x7f\x35\xb9\xd8\x2a\xd8\xa4\x7b\x7b\xde\x7d\xa3\x11\x47\x81\xb5\x6b\x8f\x1b\x2e\xda\xdb\x40\x16\x32\x42\x99\xc2\x8e\xa4\x84\x18\x21\x0d\xa2\x31\x05\x52\x20\xc9\x36\x88\xe2\x72\x34\x98\x69\x83\x8d\xba\x23\x31\xb1\x7f\x6b\x1c\x9b\x98\x68\x95\x90\xc5\xd5\xe4\x99\x6c\x82\x44\xe7\x95\x5c\xc3\x07\x67\x48\x6d\xc6\x26\x5c\xc0\xce\x90\x73\xa8\x46\x4e\xfd\x7f\xd9\x23\x20\x45\x27\xa8\xc5\xb9\xb1\xdc\xc5\x48\x94\xd5\xbe\xae\x63\xf4\x68\x09\x5b\x34\xb1\xb6\x5d\xe5\x43\x29\x8c\xf0\xb0\x06\xa4\xac\x43\xe1\x61\x50\x80\x25\xb5\x91\x08\x92\x14\xce\x5f\x76\xc1\xc0\xbc\xe7\x3c\x61\x0b\x21\xe5\x20\x89\x3a\x10\x16\x13\x4e\xf9\x16\x9f\x34\x99\x16\x23\x08\xd8\x61\xbc\xcc\x0c\xa1\x4a\x65\xed\x91\x18\x8e\x69\x85\x1e\x9e\x17\x70\xfb\xfe\xed\x7c\x24\xc4\x67\x7e\xe3\x8f\xc3\x0c\x59\xb0\xc1\x5f\xa0\x34\xe8\xc1\x6c\x01\xe8\x92\x97\xad\xef\x8c\x1a\x60\xcd\xd7\x6b\x92\xf8\xd4\x3b\x81\x14\xb9\xe3\xee\x89\xff\x86\x69\xb3\x18\x7d\x99\xf0\xe6\x78\xc1\x0b\x07\xb6\x4b\xe6\x03\x9c\xe1\x3f\x8b\x32\x5b\xf9\x72\x3a\xf7\x27\x1f\x7e\x1c\xa6\xe8\xf9\x50\x87\xc3\xa5\x7d\x14\xcf\x7b\x5d\xba\x65\x4f\xfb\x08\xc4\x9a\x35\xb0\x8a\x0a\x0d\x25\x03\x80\xf4\xb1\xe8\x1b\x32\x88\x10\x3e\xeb\xb4\xc1\x14\x38\x31\x0c\xe8\x2c\x83\x24\x17\xa4\x0e\x1b\x1a\x8b\xb6\x6d\x2c\x2b\x8b\x29\x37\x22\x83\xbe\xa3\x33\x63\xf0\xbd\xd9\x2e\x80\x5b\x95\x0e\xc5\xaf\xb9\xfa\xa1\xc0\x94\xc4\xb3\x90\xd4\xeb\xe7\x75\x3f\x84\xec\xca\x10\x63\x6e\x93\xea\x7b\xf6\xfe\xe3\xfe\xfe\xb6\xb7\xd9\xdb\x13\xc0\xb6\xed\xd9\x4c\x0b\x40\x78\x84\xe1\xb5\x70\xac\x8d\xff\xf1\x61\x0e\x1f\xef\x7e\x6d\x20\x61\x42\xad\x56\xf0\x7a\x4a\x2d\xce\xc4\xca\xc8\xc3\xfa\xf3\xa9\x37\xf8\x32\x99\x1a\x95\xe1\x60\x56\x66\x18\xc6\x97\xad\xde\x93\x62\xd0\x55\x46\x75\xc2\x9e\xcf\x86\x9b\xdb\xeb\x0f\x5e\xfd\xb0\x83\x5d\xd4\xe7\x53\xc3\xa5\x86\x0d\xa6\xdb\xd8\xe4\x04\xf3\x31\x35\x20\x5d\x5c\xb8\x2c\xf3\x30\x39\xf8\x6d\x93\x01\x20\x0c\xf6\x79\x91\x72\x07\x77\x39\x92\xf1\xd4\x8b\xdb\x14\xa5\xa8\x1c\x65\x84\x06\x8e\x2f\x6f\xae\xe6\x9d\x10\x23\x7c\xbe\xb8\x5c\x78\xc4\x25\x83\x89\x83\x8f\x77\x37\x2b\xb8\x80\x44\x12\xef\x15\x65\x29\x29\x09\x38\xc6\xa9\x58\x59\x0c\x6d\xef\xf2\xe6\x6a\x48\x99\x32\x26\x98\x9c\x82\x52\x0b\xdf\x84\x1a\x2b\xb6\x24\xd8\x24\xaf\xee\x46\x38\xdc\x89\xfa\xd9\xcc\x6c\xbd\xd7\xa5\xc0\x08\xd9\x2e\x6f\xae\x38\xcb\x58\xf4\x84\x61\xec\x56\xaf\x97\x3f\x29\x10\xd6\xc1\xee\x91\xa4\x11\xa1\x4f\x75\x62\x57\x54\x66\x76\x45\x3a\xe2\x9e\x88\xa5\xb3\x51\x73\xc2\x52\xa4\xa9\xe1\xa4\x56\x9b\xe8\x45\x98\x4c\x98\xd4\x4c\x35\x87\x5b\xe1\x72\x5f\x1c\x0a\xb4\x47\x1c\x21\xa1\xe4\x77\x0d\xa9\xcc\x42\xb6\x74\xe4\xaa\x73\x56\x88\x86\x36\xf5\x37\x35\x0c\xb2\xa0\x95\xac\x41\x21\xa6\x8c\xf7\x59\x2f\xdc\x93\x5c\xeb\xe9\xec\xb7\x08\xfd\x06\xe7\xb0\xd8\xa5\xad\xad\xc3\xc2\xbe\xec\x16\xb6\xb4\xf5\xcb\xcf\x7b\x55\x3b\x70\xd9\x62\xbc\x70\xb2\x88\x13\x4a\xe1\x9c\xfd\x7c\xf8\xc9\xfb\xf3\xdc\xcb\x98\xaa\xf0\xde\x55\x95\x0a\x34\xb5\xad\x4e\xf6\x91\x77\xb6\x12\x8e\xb6\xc8\x00\xd5\x27\xd2\x9f\xcd\xa1\x5c\xef\x96\x4e\x47\x4d\xe6\x2c\xf9\xf5\x52\xab\xe5\x0e\xe3\xe8\x2f\xe1\x9c\x65\x65\xa4\x7d\xd6\x6f\xff\x0b\x8c\x28\xeb\x5c\x0b\xe7\x03\x07\x8c\x57\x41\x0f\x5b\x47\xac\xc2\x3a\x8a\x8e\x56\x1c\x41\xe1\x8e\x5b\x7f\xce\xdb\x17\x47\xd1\x51\xf7\x9b\x65\xcd\x47\xa2\x9e\x46\x4f\x53\xc0\xf8\xfc\x09\xcf\x40\xe5\x89\xff\xdf\x09\xc0\x9d\xae\x85\x74\x35\xf8\x11\xa8\x7d\x79\x85\x19\x29\x6c\x4b\xbd\x28\xb5\x15\xdc\x54\x4c\xb3\xb6\x1b\xd8\x7d\xe3\xd9\xd0\x16\x2d\x14\xc2\x7c\x41\x57\x4a\xc1\x8d\x51\x40\xa5\x18\x15\xd2\xbd\xe9\xee\x04\xb8\xea\x9a\xc1\x8d\x27\xc2\x46\x24\x61\x33\xc5\x75\x6a\xbd\x1b\x8a\x63\xd8\xfb\x4f\x85\xa6\x0e\x79\xf1\x70\xd7\x6e\x7a\x68\x81\xcb\x0f\xc0\xef\xaf\xef\x2d\xb4\x02\x18\x09\x19\x96\xf1\xb1\xc4\xc4\x85\xd6\x5d\x8a\xba\x3f\x90\xa7\x95\x40\x02\x5c\x8e\x16\xc1\x96\x98\x50\xd6\x20\xed\x58\x9d\x68\xd0\x5f\x7a\x6e\xd1\x4f\xa9\x81\x29\xdb\x96\x66\x36\x6e\x5a\xed\xa3\x6b\xa7\xf6\x3e\xb0\x5e\x18\x23\xea\xe0\x4c\x9e\xd4\xbf\x04\xc7\x93\x4a\x69\x4b\x69\x25\x64\xaf\x74\xb7\x2d\xb4\x7a\x1f\xe6\x79\xc0\xbf\xca\xdd\xa8\x4c\xdb\x35\x7c\x6a\x42\xfa\x79\xdc\x62\x7d\xc9\xff\x3e\xb5\x6e\xbf\xda\xa3\x08\x7e\x13\x92\x52\xe1\x1a\x6a\x61\xab\xc2\x73\x76\x29\x79\x3b\x14\x95\x74\x54\x4a\xe2\x91\xb5\xe1\xc9\x4a\x3b\xe6\x48\x1b\x83\xc2\xb5\xfd\xec\xf5\xea\x74\x24\x76\x2b\x0c\x38\xed\x84\xbc\xac\x1c\x9c\xc3\xe9\xde\x67\x2e\xfb\x36\xc1\x48\x75\x7a\x4e\x54\xd5\x40\x48\xf7\xf3\xaf\x9d\xd7\x93\xca\xbd\x50\x3c\xc2\x5a\x34\xee\xb8\xdb\xf7\xe3\x39\xeb\xb9\x80\x02\xad\x15\x1b\x5c\xc3\xd1\x87\x60\x6c\x77\xfe\xb7\x5b\x7b\x34\xdf\x77\xe3\x85\xb5\xb4\x51\xa1\x8c\x1a\x79\x13\xb8\xda\x9e\x74\x7e\xb8\x68\x0f\x44\xef\x42\xb5\x0f\xe5\xf9\xb1\x62\x6a\x70\xef\xb2\xcd\x8f\xef\x6d\xac\x5f\xe0\x55\x13\x87\xc3\x98\x5c\x85\x59\xc0\x97\x30\xb3\xa5\x61\xde\x77\xf3\x5c\x1b\x42\x4e\x14\xcf\xb4\x3c\x3e\xf8\xeb\x9f\xe9\x7a\x38\x18\xcd\xdf\x36\x34\x7e\x7c\xc5\x75\x87\x09\xd2\xb6\xe3\x19\x08\x31\x2a\xcc\x28\x21\x61\xea\xbd\xc2\x1b\x93\x16\xe1\xbd\xd1\xb2\x96\xc4\x20\x67\x75\x5d\x76\x93\xbf\x69\x05\xef\xc8\xe5\xdd\xd3\x6a\x83\xee\xbe\x2e\xf1\x78\xbe\x17\x80\x44\x17\x05\xaa\x34\xf0\xb1\x25\x7c\xb4\x83\xbc\xf0\xf7\x78\xdc\xe5\x15\xee\xc2\x9c\x17\x3c\x70\x2d\xf5\x2e\x58\x61\xc6\x56\x90\x85\x8a\xfd\x06\x0f\x5d\xc0\xea\xd6\xd0\xdb\x2a\x96\x94\x30\x7b\x39\x9e\x3f\x8c\x07\x6e\xc6\x36\xce\xc1\x40\x05\x39\x0e\x99\xa8\xa4\x9b\x38\x67\x75\xd8\x83\xfd\xb4\x2f\xa4\xd4\x3b\xde\x6f\xfc\xe5\x5a\x55\x36\xc5\x8e\x90\x88\x52\xc4\x24\x29\x54\xa1\xe7\x30\x95\xab\x8c\xc7\x3b\x3e\x4f\x84\x21\xbf\x9d\xb5\xfa\xe5\x07\xd4\xa3\xd5\x61\x0d\x97\xdd\xa2\x1f\x5f\x5d\xa8\xfa\xae\x21\x5b\x5f\x47\x11\x5e\xb5\x86\x3f\xfd\x34\xce\x87\x77\x5d\x01\x76\x33\x58\x22\x64\x52\xc9\x56\x65\x51\xe8\x4a\xf9\x1b\x4a\x2b\x24\xc2\x56\xc8\x0a\x19\x48\x95\xcd\xd0\x98\x66\x6a\x6b\x72\x6d\xda\x31\xef\xb5\x43\x58\xc2\x8d\x1b\xcc\xfd\x31\xba\x1d\xa2\x62\x9c\xf2\x0e\x7f\xbd\x3a\x1d\xdd\xa6\xc0\x2f\x8f\xbc\x25\x24\xd1\xe0\x60\xb2\xf0\xe8\x37\xf4\xc0\xc1\xef\x4e\x57\xdf\x9f\xf1\x52\x35\xcc\xd4\x66\xcb\xae\x3d\xd3\x2f\x3a\x81\xc7\xe7\x59\x8f\x2f\x0e\x21\x65\x0d\x25\x9a\x04\x95\x63\x8a\xbf\xc1\xc1\x94\x1b\x2e\x1b\x1c\x9a\xc2\x17\x66\x2c\x2c\x59\x28\x35\x29\x37\x62\x3c\xbc\xc8\x6a\x49\x29\x47\x3a\xb4\x42\x5b\x08\xe3\xba\xeb\x62\x0b\xbb\x9c\xd9\x6b\x22\x52\x9e\x6a\x78\x4e\xe6\x6c\x79\xf8\x78\x4d\x8f\x67\xdf\x3d\x84\xce\x20\xa4\x41\x91\xd6\xdd\x5d\xec\xde\x4d\x08\x8e\x8e\xf7\xf9\x93\x08\xcb\xbe\x4d\x04\x3f\x90\xb3\xa0\x4b\x34\xa1\xe7\x8e\x73\x9c\x7b\xb7\x72\x64\x50\xd6\x0c\x34\x68\x0a\x52\x64\x5d\x33\xdf\x6f\xd0\x0c\x76\x7a\x7f\xb7\xb4\xa2\x2a\x39\xe0\x3f\xb4\x87\xea\x0c\x4a\x83\x09\x59\xd2\x6a\x75\x38\x36\x54\x6e\x0d\xc1\xa4\x71\xd6\xfd\xab\x9d\x15\x46\xf7\x25\xe1\x1f\x34\xc2\x85\x40\xa8\x16\x36\x83\x8f\x10\x35\xd7\xf1\x20\xb6\x8b\x03\x5f\x18\x94\x41\xdd\x9c\xca\x2e\xbd\xf8\xc3\xc3\x4e\x48\x89\xee\xa1\xbd\x3e\x64\xb0\x5c\x80\x1f\x53\x6b\x97\xb3\x5c\x94\xb6\xe9\xc5\xfe\x1e\x6b\xa7\xd0\x40\x41\x9b\xdc\xc1\x4e\xa8\x80\xc8\x9e\xbe\x1c\x56\xe1\x8b\x97\x67\x9e\x13\x70\x49\x94\x3c\x6a\xfe\xf1\x42\x5d\x0c\xfd\xb7\x98\x3a\x6b\x9f\x5b\x94\x06\x27\x3a\x3a\x37\x8b\x9f\x3c\x25\x80\x57\xaf\xfc\x53\xe8\xcb\xb0\x86\x23\xee\xd2\xa1\x4c\xfa\xda\x24\xc5\xaf\x28\x05\x23\xd4\x06\x81\x56\x08\x9f\x4e\x17\xaf\x3f\x1f\xbd\xd0\xf9\x7d\x8f\xeb\xf0\xf7\x1c\x3a\xb3\x27\x3b\x72\x68\xc6\x7f\xe6\x1e\xeb\xa0\x6b\xbe\x6d\xba\x4f\x18\x3a\x3d\x9f\xea\x0c\xf1\x80\xe6\x99\x8f\x57\x8c\x43\x3d\xe6\x7a\x1d\x6c\xfb\xdb\x9e\x5e\x84\x47\x71\xdc\xa2\x72\x95\x47\x83\xa1\xac\xfe\x26\xcc\xee\xc8\x25\x79\xac\x99\xab\xb7\xa6\x2f\x3a\xb9\xb9\xaf\xeb\xf6\x6e\x1e\xe2\xaa\x11\xeb\xe7\xd8\x91\x72\x1d\x71\xe5\x27\xa5\x77\x3d\xb5\x1d\x91\x8d\xc9\xde\xb5\x86\xfe\x69\x10\xfa\x86\x79\x44\xa5\xff\x18\x35\x4d\xff\xfa\xbe\x15\xd1\x78\xf1\x69\x06\xff\x0d\x00\x00\xff\xff\x65\x8e\xd4\x10\x4a\x1c\x00\x00"

func metadataviewsCdcBytes() ([]byte, error) {
	return bindataRead(
		_metadataviewsCdc,
		"MetadataViews.cdc",
	)
}

func metadataviewsCdc() (*asset, error) {
	bytes, err := metadataviewsCdcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "MetadataViews.cdc", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x49, 0x9e, 0x77, 0xf2, 0x3d, 0x24, 0xe, 0x73, 0x98, 0xab, 0xc5, 0x3c, 0x25, 0xb0, 0x7a, 0x69, 0xf2, 0xaf, 0x8e, 0x5c, 0xf3, 0xb3, 0x31, 0xc0, 0x31, 0x6a, 0x97, 0x3e, 0x2, 0x78, 0x4c, 0x2a}}
	return a, nil
}

var _nonfungibletokenCdc = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x57\x41\x8f\xdb\xba\x11\xbe\xeb\x57\xcc\xcb\x03\x9a\xdd\xc0\x6b\xf7\x50\xf4\x60\x20\x68\xda\xb7\x6f\x01\x5f\xb6\x0f\x5b\x17\x3d\x04\x01\x4c\x8b\x23\x9b\x08\x45\x2a\x24\x65\xc7\x0d\xf6\xbf\x17\x33\x24\x25\xca\xf6\x26\x9b\x5b\x73\x89\x57\x22\xbf\x99\xf9\xe6\x9b\x8f\xd4\xe2\xdd\xbb\xaa\xfa\xf5\x57\x58\xef\x11\x1e\xb4\x3d\xc2\xa3\x35\x77\x0f\xbd\xd9\xa9\xad\x46\x58\xdb\xcf\x68\xc0\x07\x61\xa4\x70\x92\x17\x6e\x1e\xad\xc9\xef\xf9\xf5\x06\x6a\x6b\x82\x13\x75\x00\x65\x02\xba\x46\xd4\x58\x55\x84\x37\xfc\x09\x61\x2f\x02\x08\xad\xc1\x58\x73\xd7\x64\xf4\xc0\xe8\x79\xb7\x87\xda\xf6\x5a\xd2\xdf\x8d\x75\x2d\x04\x3b\xaf\x56\x0d\x08\xe8\x3d\x3a\x38\x0a\x13\x3c\x04\x0b\x12\x3b\x6d\x4f\x20\xc0\xe0\x11\x4c\x13\x86\xfd\x33\x08\x7b\x54\x6e\xcc\xe6\xc8\x70\x06\x51\x56\xc1\x82\x6a\x3b\x8d\x2d\x9a\x40\xcb\xe0\xbc\x88\x31\xd7\x39\xe7\x7e\x89\xb3\x17\x07\xca\x18\x1a\xab\x89\x26\x2a\x86\x80\x5c\xaf\xd1\x83\x30\x12\x8c\x68\x95\xd9\x55\x5c\x6a\x98\x54\xef\x3b\xac\x55\xa3\xd0\xcf\x13\x83\x0f\xeb\x0d\x38\xf4\xb6\x77\x99\xaa\xda\x3a\x1c\x1e\x41\x38\x75\x89\x33\x87\x9d\x43\x8f\x54\xbb\x30\xf0\xf8\xb0\x06\x65\x18\xdd\xb7\xc2\x8d\xb5\x27\xe0\xdf\xac\xd6\x58\x07\x65\xcd\x06\x9e\x26\xf8\x23\x34\xa1\xfa\x60\x1d\x65\xcd\xd4\xbe\xf5\x8c\x5b\x0f\x7b\xe7\xd5\x8a\x5a\x59\xeb\x5e\xf2\xa2\x06\x8f\xd0\xf4\x86\xdf\x71\x0b\x04\x33\x40\x59\xd8\xa3\x41\x47\x8f\x50\x78\xa5\x4f\x55\x6b\x0f\xa9\xad\x9e\x12\x25\x5a\x6c\x1f\xc0\x36\xbc\xba\x0c\xc1\xf9\xfe\xe1\xec\x41\x49\x74\x1b\x5e\xb9\x79\xc2\x1a\xd5\x81\xfe\x1c\xd2\x1d\x48\xf4\x5c\x87\x2f\x9f\x80\xc4\x5a\x0b\x87\x45\x72\x47\x15\xf6\xe0\x6d\x8b\xd0\x39\x64\xd0\xce\x7a\xa6\x49\x2a\x5e\x51\x25\x56\xbf\xf4\xca\x21\x27\x35\x72\x56\x74\xb7\x46\x17\x84\x32\xa9\xa7\x0c\xb4\xc5\xbd\x38\x28\xeb\x86\x69\xf0\x51\x29\x27\xa0\x14\x3c\x76\xc2\x89\x80\xb0\xc5\x5a\xf4\x94\x66\x80\x9d\x3a\xa0\xe7\x18\xac\x60\xfa\x21\xb6\x4a\xab\x70\xa2\x48\x7e\x4f\xfb\x04\x38\x6c\xd0\xa1\xa9\x91\x44\x1a\x15\x5c\xa6\x44\xe9\x5a\xa3\x4f\x80\x5f\x3b\xeb\x13\x5e\xa3\x50\xcb\xa8\xba\xb1\x76\x65\xc0\x1a\x04\xeb\xa0\xb5\x0e\xab\xc4\xf9\x48\xd7\x1c\x56\x34\x83\xde\xa6\xc4\x28\x29\x7f\x9e\x55\x2b\x3e\x23\xd4\xbd\x0f\xb6\x1d\x9a\x90\x48\x9b\x0c\xd0\xb4\x11\x34\x96\x16\x0e\xc2\x29\xdb\x13\xa4\x32\xbb\xd4\x0b\x82\x8f\x7a\x98\x57\xd5\x3f\x4e\xd0\x7b\xe2\x73\x40\xe6\x12\x46\xa0\x59\x4a\xca\x36\x2c\xc9\xa9\xc6\x3d\xd4\xc2\x80\x47\x23\x2b\xda\xe5\xa2\x58\xb2\xda\x3a\x44\x77\x17\xec\x1d\xfd\x3f\xe3\xd8\x24\x3c\x6a\x99\xd9\x51\x7e\x1c\x84\xa7\x99\xd2\x12\x50\x23\xa1\x6a\xd0\x28\x77\xe8\xaa\x8b\x71\x5a\x5b\x0e\x95\xa7\x8e\x54\x6f\x6c\xd8\xa3\xe3\x14\x67\x83\x2d\xb1\x37\x78\xe2\xe6\xc4\xd0\xd2\x89\x38\x1a\x8f\x0f\xeb\xaa\x71\xb6\xbd\xe8\x29\xfb\x94\x81\x3a\x3b\x88\xc4\xce\x7a\x15\x86\x4e\x82\x35\x93\x58\x6f\x7d\x35\xd5\x68\x6d\xa9\x13\x21\xca\x37\x38\x61\x7c\x83\x6e\x5e\x55\xef\x16\x55\xb5\x58\xb0\x93\xb7\x24\xde\x38\xd5\xe7\xd6\x3c\x87\x7f\x32\x74\xf9\x96\x9a\xa5\x35\x6d\x56\x6d\x67\x5d\x88\x6d\x29\xfa\xad\x7c\xe1\xed\x8b\x45\xd5\xf5\xdb\x2b\xd0\x97\xae\xfa\xad\xaa\x00\x00\x52\x56\xc1\x06\xa1\xc1\xf4\xed\x16\x1d\x7b\x42\x6c\x1d\x2b\x55\xf9\xe8\x7a\xca\x00\x7e\x55\x3e\xf0\x44\xd0\x5e\x0a\x75\x10\x2e\x6e\xfe\x57\xdf\x75\xfa\xb4\x84\x7f\xaf\x4c\xf8\xeb\x5f\x06\xf0\xdf\x0f\x31\x4d\x11\x00\x5b\x15\x02\x4a\x38\x12\xc7\xa9\x0f\x45\xaa\x54\x87\x0a\x4a\x68\xf5\x5f\x94\x69\xfb\x10\x06\x19\xe6\xb7\xb4\x78\x35\x2e\xbc\xb9\xbd\x16\x4a\xf9\x69\x34\x91\x0e\x34\xe5\x07\x25\x98\x59\xde\xa7\x8c\x54\xb5\x08\xac\xc6\xc1\x38\x2f\x7c\x31\x01\x07\x38\x8a\x02\x04\x48\x47\xf3\x32\xdb\xc5\x02\x56\x17\x7b\x95\x07\x63\x43\xf4\x5d\x10\x75\x6d\x7b\x13\xde\x7a\x36\x7b\xb1\xc3\x19\x6c\x08\x66\xc3\xad\x86\x2d\xc2\xc6\x28\xbd\x99\x5f\xe7\xe0\x3f\x29\xf4\x8d\x92\x99\xec\x19\x67\xb1\x84\xbf\x4b\xe9\xd0\xfb\xbf\x5d\xa5\xe4\x25\x3e\x92\xc6\x51\xf2\x20\x4d\x0e\x82\xb3\xaa\x42\x66\x2a\x59\xdd\x6b\x88\x2a\xd1\x5f\x28\xe8\x3e\x2e\x99\xd4\x13\xec\xb5\x6a\x56\xd3\x4b\x4b\x92\x90\x1f\xce\xff\xf1\x7a\x72\x1e\xe9\xf2\xd0\x82\x15\xa9\xef\x1b\xaf\x28\xe6\xa0\x37\xea\x4b\x8f\xb0\xba\x4f\xa4\x89\x7a\xcf\x32\xdd\x0b\x3f\x2c\x25\x40\x8d\x01\xc6\x84\xf9\xd5\xf3\x90\xe7\x53\x3c\xc3\xda\x81\x7b\xf2\x93\x94\x1c\xa9\xec\x9a\x81\x52\x0d\x79\x3f\x5f\xa5\x1a\x65\xe2\x19\x94\x32\x27\x53\x42\x19\x1d\x8f\x30\x13\x1e\x3b\x3c\xd5\x72\x59\xeb\xe3\xc3\x7a\x79\x5e\xe6\x0f\x73\x2f\x38\xb6\xd0\xa2\x54\x74\x72\x66\xb9\x7b\xc8\xb6\x59\x98\xe6\x2b\xb8\xce\x97\x89\x29\xdf\x83\x27\x3b\xa4\xcb\xc9\x70\x8d\x1a\x62\x14\x9a\x22\xd7\x8b\x8b\x54\x80\x78\x1a\x47\x46\xdc\xa4\xb4\xa6\x37\x03\xec\x4d\xfe\xb1\xba\xcf\xb5\xde\x2e\xe1\xc3\x94\x0f\xde\x48\xf7\x90\xe9\x23\xfa\xe7\xd0\xf7\x3a\xcc\x95\x84\xf7\xef\xa1\xc4\x7a\x43\x42\x59\xdd\x67\xe5\x8f\x5e\x10\x67\xaa\xed\x7d\xa0\x21\xe6\xab\xa0\x68\x11\x44\x1c\x17\xba\xd9\xa0\xa7\x51\x58\xdd\xbf\x99\x44\x7b\xae\xa6\xbf\x7e\xd0\x8d\x34\x53\x3e\xf3\xf0\x53\xad\xc8\x17\xb9\xec\xff\x29\x50\x3e\xe9\x82\xf8\x3c\x36\x42\xf0\x2f\xe1\x76\x3d\x4b\x99\x7a\x20\xa4\x2c\x5b\x70\x16\xba\x08\x5f\x76\x24\x81\xdf\x30\x3f\xb1\x05\xb7\x2f\x17\xca\x03\x33\xb8\x64\x3a\xc6\x6b\xdb\xb6\x7c\xd7\xca\x1b\xba\x7e\xab\x95\xdf\x43\x63\xdd\xf0\x71\x31\xc9\xe5\x85\xfa\xc7\x8c\xff\x20\x84\xfa\x6c\x36\xbe\x9b\x6e\xb9\x68\x87\x61\x75\xef\x6f\x6e\x97\xf0\x31\x6a\xeb\xd3\xc5\x92\xad\x75\xce\x1e\x1f\x1f\xd6\x85\xb5\xdd\x2e\xe1\x4f\x79\x58\xaf\x1b\x46\x2a\x28\x0d\x80\xa9\x1d\x5d\x27\x26\x9f\x1f\x85\x4d\x6c\x31\xdf\xb4\x65\xfe\xfa\x18\xee\x06\xe4\x34\xd9\x5f\x5e\x14\xc6\x48\xc7\x72\x98\xd2\xd9\x20\x92\xd9\x35\xba\x4a\xd9\xdc\x2b\x7e\x27\x1c\xdf\x50\xf7\x56\xcb\xd1\x95\x53\x3e\x57\x24\x92\xef\x0d\x74\x80\x48\x5a\xbb\x84\x0f\xdf\x22\x3f\x4b\xda\xfb\x5c\xfd\x5f\xd8\xc4\xf7\x06\x24\xce\xc7\xe5\x40\x8c\xb9\x78\x90\x03\x39\x25\xd0\xb0\x29\x44\x17\x49\x1b\x95\x04\xe1\x9c\x38\xbd\x4e\x8d\x25\x60\x54\x22\x38\x0c\xbd\x33\x69\x62\x9d\x38\x65\x7b\xa2\x77\x71\xa6\x1c\xe6\x9e\xd4\xd7\x7b\xf2\x82\xae\xcb\x60\x4f\x39\x4a\x52\x37\xca\xf1\x2b\x29\xde\xc4\xcb\x2f\xe1\x2b\x71\x16\x0b\xf0\x76\x3c\xbf\x63\x73\xf8\xf3\xc1\xa1\x90\x20\x45\x10\x4c\x11\xdf\xc1\x5b\x0c\x7b\x2b\xd3\xa9\xa3\xc2\xcf\x4c\xd8\xb9\xc7\x3b\xbc\x62\xf1\x1e\x75\x33\x1f\x54\xf8\x51\xc9\x4f\xf0\xcb\x7b\x30\x4a\x2f\xe1\x0d\x61\x48\x8b\xf1\xe2\xc6\xf7\xde\xcb\xaa\x7e\x79\xad\x8f\xd7\x0e\x45\xc0\xdf\xdb\x2e\x9c\x8a\x0f\x86\xf8\x94\x5b\x86\xf4\xea\xd2\xc9\x21\x7e\x4e\x45\xce\xcf\x25\x5d\x12\x79\x62\x0a\xed\x91\xe9\xf7\x55\x49\xd2\xd5\xd8\xd4\xe0\x0f\x45\x2a\x85\x0b\x5e\x9e\x86\xe9\x24\xcc\xd2\x98\x6b\x34\xbb\xb0\xa7\x63\xf1\xcf\xe9\x34\x8c\x31\x64\x39\x8a\xf9\x18\xe4\xca\x0a\xa2\x32\x35\xcf\xd5\xff\x02\x00\x00\xff\xff\x33\x4d\x81\x27\xe0\x12\x00\x00"

func nonfungibletokenCdcBytes() ([]byte, error) {
	return bindataRead(
		_nonfungibletokenCdc,
		"NonFungibleToken.cdc",
	)
}

func nonfungibletokenCdc() (*asset, error) {
	bytes, err := nonfungibletokenCdcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "NonFungibleToken.cdc", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xdb, 0x61, 0xca, 0x9d, 0xaa, 0x66, 0x36, 0xdf, 0xbc, 0x51, 0xdb, 0x7b, 0x51, 0xd8, 0x3d, 0x6f, 0x4e, 0x9c, 0x8e, 0x50, 0x28, 0x7c, 0x18, 0x1d, 0x2, 0xb2, 0xc2, 0x2b, 0x26, 0xa1, 0xfe, 0x2d}}
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
	"ExampleNFT.cdc":       examplenftCdc,
	"MetadataViews.cdc":    metadataviewsCdc,
	"NonFungibleToken.cdc": nonfungibletokenCdc,
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
	"ExampleNFT.cdc": {examplenftCdc, map[string]*bintree{}},
	"MetadataViews.cdc": {metadataviewsCdc, map[string]*bintree{}},
	"NonFungibleToken.cdc": {nonfungibletokenCdc, map[string]*bintree{}},
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
