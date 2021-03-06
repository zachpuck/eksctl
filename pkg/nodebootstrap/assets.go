// Code generated by go-bindata.
// sources:
// assets/10-eksclt.al2.conf
// assets/bootstrap.al2.sh
// assets/bootstrap.ubuntu.sh
// assets/kubelet-config.json
// DO NOT EDIT!

package nodebootstrap

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
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

var __10EkscltAl2Conf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x92\x41\x6f\xda\x40\x10\x85\xef\xfe\x15\x2b\x85\x43\x2b\xb1\x76\x42\x68\x0e\x91\x7c\xa0\xc1\x89\x90\x28\x89\xe2\x44\xad\xd4\x56\x68\xd8\x1d\xe8\x94\xf5\xac\xb5\xbb\x86\xa4\x88\xff\x5e\x19\xdb\x29\x55\xa2\xaa\x37\xd8\x6f\xfc\xde\xcc\x9b\x39\x11\xb8\xf6\x2a\x18\xe9\x4b\x54\xb4\x24\x25\xfc\xb3\x0f\x58\x68\xa1\x9d\x2d\x25\xb1\xa8\x98\x82\x58\x5a\x27\xd6\xd5\x02\x0d\x86\xfe\xe1\xcf\xa8\x80\x5f\x96\xc5\x94\xb8\x7a\x12\x03\xf1\x6e\x34\x1d\xbc\x8f\xa2\xaf\x39\xba\x0d\x29\xfc\x1e\x9d\x88\xa9\x55\x60\x44\x81\x01\x34\x04\x10\x25\x38\x28\x30\xa0\xf3\x97\xe2\x3e\xbb\x99\xdc\xce\xfa\x62\xf4\x39\x9f\x8f\xb3\xeb\xd1\xe3\xf4\x61\xde\xbc\x45\x19\x6f\xc8\x59\x2e\x90\xc3\x35\x19\x4c\x13\x0c\x2a\x69\x5a\x4c\x3a\xad\x18\x79\x13\x9d\x88\x1b\x63\x17\x60\x04\xb0\x16\x3e\x40\x20\xf5\x97\xc7\xa7\xd1\x97\xf9\xdd\xed\x38\xef\x8b\xab\xe9\x63\xfe\x90\xdd\xcf\xc7\xb3\xbc\x2f\x66\xb7\xe3\x6c\x3e\x1d\x7d\xcc\xa6\xf9\x3f\xbd\xda\x61\x5b\xab\x66\x16\xb6\x2c\xdf\x70\x3a\x48\x4e\xee\xfa\x62\x32\xcb\x1f\x46\xb3\xab\x6c\x3e\x19\xff\x97\xb6\xa9\x55\x0f\x0e\x51\xf6\x84\x2a\x0f\xe0\x42\x7a\xf4\x33\xa9\xbc\x4b\x16\xc4\xdd\x07\xe2\x5b\x24\x84\x94\x6c\x35\x4a\x2a\xd3\xde\xae\x75\xde\xb7\x40\x99\xca\x07\x74\x52\xb3\x4f\x7b\xbb\xa3\xb1\xbb\x82\x02\x9e\x64\x69\x75\x4d\xbb\x78\xf6\xc7\xa2\x06\x16\x68\x7c\x27\xdc\xa4\xb4\xef\x83\x29\x7f\x40\xdc\xf4\x1e\x93\x4d\x88\x7d\x00\x56\x28\x49\xa7\xbd\xdd\xd1\xd0\x9d\x16\x18\x63\xb7\xb2\x74\xb4\x21\x83\x2b\xd4\x69\x70\x15\xb6\xac\xb4\x5a\x12\x2f\x1d\x48\x65\x39\x00\x31\x3a\x49\x05\xac\x30\xbd\x38\x1d\x0c\x4f\xcf\xce\x86\xe7\xc3\x0f\x83\x58\xaf\x5d\x8c\xca\xc5\xbd\xdd\xeb\x1b\xd9\xc7\x70\x38\x3e\xd8\xfa\x58\xd9\xa2\x4e\x35\x29\xa1\xf2\x28\xa1\xd0\x17\xc3\xcb\xf3\xf8\xec\x25\x10\x5b\x69\x59\x3a\xbb\x21\x8d\x2e\x85\xad\xef\x00\x93\x5c\x10\x4b\x4d\x2e\x4d\x6c\x19\x12\xc5\x54\x27\x7d\x84\x95\xe5\x65\xc3\xeb\xcd\xd5\x9c\x31\xc4\xba\xab\x78\x69\xde\x55\x1c\xa8\xc0\x54\x5b\xb5\x46\xd7\xa5\x89\x61\x6b\xdd\x5a\x96\xa6\x5a\x11\xa7\x8a\xa9\x05\x0e\x57\x74\xd8\x51\x9d\xf7\x71\x2e\xf5\x8a\x6b\x4b\x5a\xbd\x3a\x95\xe6\x39\x7e\x86\xc2\xfc\x71\x7f\xab\xd0\x60\x68\x51\xfc\xd3\x5b\x8e\x7e\x07\x00\x00\xff\xff\x63\xce\x0d\x56\xdc\x03\x00\x00")

func _10EkscltAl2ConfBytes() ([]byte, error) {
	return bindataRead(
		__10EkscltAl2Conf,
		"10-eksclt.al2.conf",
	)
}

func _10EkscltAl2Conf() (*asset, error) {
	bytes, err := _10EkscltAl2ConfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "10-eksclt.al2.conf", size: 988, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _bootstrapAl2Sh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\xcd\xc1\x4a\xc4\x40\x0c\xc6\xf1\x7b\x9f\x22\xae\x1e\xf4\x30\x0d\xca\x2a\x28\x28\x88\xbb\x87\x5e\xaa\xa0\xf7\x25\x9d\x7e\xd0\x61\xd3\x99\xa5\x49\x17\x7c\x7b\x51\x8a\xec\x4d\xf0\x14\xf8\x27\xe1\x77\x7e\xc6\x5d\xca\x6c\x03\x05\xcc\x55\x85\x38\x14\x5a\xb5\xaf\x9b\xed\xae\x79\x7b\xbc\xb8\x8c\xf3\xa4\x14\x82\x25\x45\x76\x1a\xdc\x0f\x0f\xcc\xd7\x77\xf7\xf5\xcd\xed\xba\x5e\x26\xab\x38\xcc\x79\x84\x4b\xe8\xc5\x85\xb5\x44\xd1\x90\x0e\xc7\xf5\xd5\x8a\x9e\x88\xe1\x91\xb1\xb7\xe8\xca\xfb\xb9\x83\xc2\xeb\x9f\x93\x1a\xf9\xb8\x90\x4d\xfb\xfe\xf1\xdc\xbe\x6c\x77\xcd\xe6\xdf\x6c\xca\xe6\x92\x23\x42\xea\xbf\xdd\xbf\xe0\xca\x3e\xcd\x31\x46\x57\xea\x05\x63\xc9\x61\x82\x16\xe9\x4f\x3a\xb2\x74\x0a\x5a\x7e\x4f\x16\xe6\x32\xf9\x6f\xff\x0a\x00\x00\xff\xff\x11\xf1\xc5\x0b\x45\x01\x00\x00")

func bootstrapAl2ShBytes() ([]byte, error) {
	return bindataRead(
		_bootstrapAl2Sh,
		"bootstrap.al2.sh",
	)
}

func bootstrapAl2Sh() (*asset, error) {
	bytes, err := bootstrapAl2ShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "bootstrap.al2.sh", size: 325, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _bootstrapUbuntuSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x94\xe1\x6a\x23\x37\x10\xc7\xbf\xfb\x29\xa6\x3e\x43\x2f\x70\xda\x4d\x72\xbe\x40\x03\x5b\xea\xc6\x6e\x09\x4d\x93\x23\xce\x71\x85\x52\xcc\x58\x9a\xf5\x0a\x6b\x35\x8b\x34\x6b\x37\x17\xf2\xee\x45\x5e\xaf\xcf\x09\xb4\x07\xc5\x1f\xbc\xd2\xfc\x34\x33\xd2\xfc\x67\xde\x7c\x97\x2f\xad\xcf\x97\x18\x2b\x50\xd4\x0e\x06\xa4\x2b\x86\xe1\xed\xdd\x74\xb6\xb8\xfe\x58\x8c\xde\xea\x36\x38\x50\x2a\x5a\x47\x5e\xa0\x12\x69\x2e\xf3\xfc\xec\xe2\x87\xec\xfc\xc3\x38\xdb\xff\xe7\x0e\x85\xa2\xe4\x35\x09\x2a\x83\x82\xb9\x63\x8d\x4e\xd9\x66\x33\x3e\x19\xc2\x8f\x90\x93\xe8\x9c\xd6\x51\x8b\xcb\xd7\xed\x92\x1c\x49\xb6\x43\x32\xf2\x9b\x7d\xc8\xeb\xdb\xf9\xc3\xe4\xf6\x6a\xb6\xb8\x9e\xfe\xef\xb0\xd6\x47\x41\xaf\x49\x59\x93\xe2\x7e\x2b\xf0\x20\x7a\x6c\x00\x9d\xc5\x08\x7b\xab\xa2\x75\xcc\xf6\xdf\xfd\xde\x6b\x4c\x8b\x3b\x60\x5a\x5c\xbf\xd7\x61\x51\xb8\x39\x76\x36\x88\x8f\x51\xa8\x4e\x5c\xa0\x48\xa2\x4a\xb4\x8e\xcc\x60\xf0\x76\x00\xf0\x06\x1e\xee\xa6\x77\x97\x20\x15\x45\x82\x58\x71\xeb\x0c\x2c\x09\x1c\xf3\x9a\x0c\xa0\x00\x6d\x28\x3c\x82\xd8\x9a\x7a\xa7\x10\x05\x83\x44\x68\x9b\x77\x3b\x0f\xdb\xca\xea\x0a\x6c\x84\x6d\x85\x02\x5b\x02\xc3\x60\x3d\x4c\x6e\xce\xe1\xed\xc1\xb6\xc4\x48\x06\xd8\x43\xe3\xd0\x7a\xe8\x72\x32\x9d\x03\xf4\x06\x6a\x42\x2f\x20\x9c\x82\x37\x1c\x04\x97\x8e\xd2\xb2\xe6\x28\x3d\x0d\xc6\x46\x09\x1c\x4f\xde\xc1\xb2\x15\xb0\xf2\x7d\xdc\x9d\xf7\x2c\xa0\x1d\x61\x80\x8a\xb7\xe9\x90\x63\x34\xfb\x2b\x95\x81\xeb\xaf\x89\xa7\xf7\xd9\x5a\xa9\xb8\x15\xa8\x70\x63\xfd\x6a\xe7\x40\x18\x74\x1b\x85\x6b\x1b\x29\x9d\xeb\x40\x2b\x91\x5c\x39\x00\x88\xdc\x06\x4d\xdf\x28\xe5\x7f\x62\xff\x0a\x24\xe5\x24\xe1\x74\x6a\x00\x28\x1d\xae\x62\x91\x2a\x03\x30\x44\x63\x02\xc5\x58\x9c\x66\xbb\xdf\xb0\xdb\xf5\x6c\x48\xd9\xa6\x18\x3d\xed\x7b\xe4\x79\x6f\xd0\xae\x8d\x42\x41\x19\x1f\x8b\xd1\xd3\xd5\xcd\xa7\xf9\xc3\xec\x7e\x31\xbd\x9d\xf7\x40\x8d\x7f\xab\x86\x4d\xb2\xfe\x3e\xf9\x63\xf1\xf1\x6e\x7a\x30\xed\x9c\x3a\x5c\x92\x8b\xbd\xe3\x9b\xc9\xcf\xb3\x9b\xf9\xf3\x3b\x74\x4d\x85\x59\x97\x6f\x66\xf9\x58\xe2\xc5\xe8\xe9\xa8\x69\x7a\x5f\xd8\x4a\x45\x5e\xac\x46\xb1\xec\x95\xf0\x9a\xbc\xda\xd2\xb2\x62\x5e\x17\x12\x5a\x3a\xe2\x38\xd8\x2f\x1d\x56\xb3\xa1\xe2\x73\x47\xf5\x80\x73\xbc\x55\x4d\xb0\x1b\xeb\x68\x45\xe6\xf8\x70\xc3\x46\x59\x5f\x06\x54\x9a\xbd\xa0\xf5\x14\x94\xad\x71\x45\xc5\xc5\xe9\xf9\xf8\xf4\xec\x6c\xfc\x7e\xfc\xe1\x3c\x33\xeb\x90\x91\x0e\xd9\xe8\x69\xf2\x79\xbe\x98\xce\x7e\x99\x7c\xba\x79\x58\xdc\xcf\x7e\xbd\xbe\xbb\x7d\xce\xb0\xc6\x2f\xec\x71\x1b\x33\xcd\x75\x2a\x49\xde\x60\x1b\x49\x61\x6d\x2e\xc6\x97\xef\xb3\xb3\xc3\xcb\x72\x6b\x54\x13\x78\x63\x0d\x85\x02\xb7\xf1\xf5\x93\x73\x8d\xd6\x17\xfb\x65\xa7\x8a\x1e\xf1\x56\x2d\xad\x57\xc6\x86\x22\xe7\x46\x72\xed\x6d\x9a\x78\x47\x66\xcd\xbe\xec\xec\x49\x19\xc9\xee\x49\x32\xd3\x13\x87\xfb\x85\xd6\xa7\x3e\x2c\x0c\xeb\x35\x85\xbe\x72\x24\x5b\x0e\x6b\xd5\xb8\x76\x95\x52\xf0\xb6\x3f\xb7\x0a\xdc\x36\xca\x04\xbb\xa1\x50\x74\xab\xb2\x4f\x3c\xd0\xca\xee\x32\x4f\x85\x3f\x7e\xd7\xdd\x28\x61\x5f\xda\x55\xf1\x5a\xc8\xdd\x76\xf6\x88\x75\x7f\xb7\x92\x50\xda\x40\x6a\x95\xc6\x60\x71\xcf\x82\x42\xbf\x75\x92\x9f\x53\xd8\x50\xb8\xa2\x20\xb6\x4c\x4a\x78\x11\x04\x3d\xfb\xc7\x9a\xdb\xa8\x92\x06\x8a\x12\x5d\xa4\xc3\x8b\x5a\xf2\xa2\x34\xaa\xd2\x3a\x7a\x91\x83\xc6\x4c\x07\x49\xdc\x49\x6a\x95\x6e\xd8\x7d\x1d\x92\x69\xd6\xc1\x70\xf4\xb4\x6b\xa1\x3f\x7f\xfa\xeb\x79\x38\x38\x19\xf4\x23\x11\xc3\x0b\x6e\xf0\x4f\x00\x00\x00\xff\xff\x61\xe3\x4c\xe6\x76\x06\x00\x00")

func bootstrapUbuntuShBytes() ([]byte, error) {
	return bindataRead(
		_bootstrapUbuntuSh,
		"bootstrap.ubuntu.sh",
	)
}

func bootstrapUbuntuSh() (*asset, error) {
	bytes, err := bootstrapUbuntuShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "bootstrap.ubuntu.sh", size: 1654, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _kubeletConfigJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x92\x4d\x6f\xf2\x30\x10\x84\xef\xfc\x8a\xc8\x67\x94\x84\xf7\x15\x52\xcb\x8d\x82\xda\x43\x39\x15\xda\x9e\x1d\x67\x03\x56\x1c\x2f\x5a\xaf\xe9\x97\xf8\xef\x95\x3f\x5a\x0a\x3d\x54\x39\x79\xf4\xec\x78\x36\xe3\x8f\x51\x51\x88\x5e\xdb\x56\xcc\x0a\x71\xef\x1b\x30\xc0\x0b\xb4\x9d\xde\x7a\x92\xac\xd1\x8a\x71\x20\xe4\x5e\x3f\x01\xb9\x70\x9e\x15\xa2\x4f\x5c\xa9\x22\x58\xf6\x57\xae\xd4\x58\x1d\x26\x0d\xb0\x9c\xe4\x81\xb6\x25\x70\x2e\xd0\x75\x19\xbf\xac\x7b\xde\x81\x65\xad\x92\xf9\xac\x08\x01\x82\x6e\xd1\xbe\x0d\xe8\xdd\xb7\x54\x14\x02\xac\x6c\x0c\x84\x68\x9d\x34\x0e\xa2\x7c\x1c\xa7\x81\x17\x68\x76\x88\xfd\x4f\x5c\x49\xb5\x83\xcd\x66\x15\x2e\xfd\x37\xd4\x4e\x8c\x7f\x1b\x31\xf9\x73\x9f\xd7\x69\x7d\x7d\x66\x62\x34\x58\x5e\xcc\x6f\xb5\x81\x60\x54\x01\xab\x0a\x7a\xa7\xd8\x54\x4a\x96\x8a\x58\xa4\xf9\x51\xf6\x88\x2b\x21\xe9\xf7\x8b\x8d\x06\x6c\xa3\xc1\x73\x0e\xfa\x47\xee\x79\x76\x81\x36\x6f\x30\x3d\xdb\x20\x32\x8f\x56\x5e\x52\xff\x6b\x77\x99\x47\x19\xef\x18\x68\x89\x83\xd4\xb1\xae\x2c\x94\x06\x95\x34\xa9\x06\xb5\x25\xf4\xfb\x25\xe9\x03\x50\x44\xe2\xb9\x4b\x17\x8a\x0e\x24\x7b\x82\x3b\xc9\x70\xea\x43\x3c\x20\x4b\x86\xfc\x46\xd6\x40\x07\xa0\x05\x10\xeb\x2e\x94\x09\xa7\x7f\x9b\x52\xb8\x08\x6c\x56\xeb\x1b\x44\x76\x4c\x72\xff\x45\x1c\x3f\x03\x00\x00\xff\xff\xbb\x62\xda\x19\x74\x02\x00\x00")

func kubeletConfigJsonBytes() ([]byte, error) {
	return bindataRead(
		_kubeletConfigJson,
		"kubelet-config.json",
	)
}

func kubeletConfigJson() (*asset, error) {
	bytes, err := kubeletConfigJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "kubelet-config.json", size: 628, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
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

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
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
	"10-eksclt.al2.conf": _10EkscltAl2Conf,
	"bootstrap.al2.sh": bootstrapAl2Sh,
	"bootstrap.ubuntu.sh": bootstrapUbuntuSh,
	"kubelet-config.json": kubeletConfigJson,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
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
	"10-eksclt.al2.conf": &bintree{_10EkscltAl2Conf, map[string]*bintree{}},
	"bootstrap.al2.sh": &bintree{bootstrapAl2Sh, map[string]*bintree{}},
	"bootstrap.ubuntu.sh": &bintree{bootstrapUbuntuSh, map[string]*bintree{}},
	"kubelet-config.json": &bintree{kubeletConfigJson, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory
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
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
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
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

