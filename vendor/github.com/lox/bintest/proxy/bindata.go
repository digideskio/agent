// Code generated by go-bindata.
// sources:
// data/client.go
// DO NOT EDIT!

package proxy

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

var _dataClientGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x95\x51\x6f\xdb\x36\x10\xc7\x9f\xc9\x4f\x71\x13\x90\x41\x1a\x34\x09\xc3\x86\x61\x70\x91\x87\x36\x36\x82\xee\xa1\x0d\x92\x01\x7d\x28\x8a\x42\x11\x4f\x36\x37\x9b\x54\x8e\x94\x55\x63\xf0\x77\x1f\x8e\x94\x6c\x25\xf6\x1a\xbf\x24\xe2\x1d\x79\xfc\xff\xee\x2f\x9f\xda\xaa\xfe\xa7\x5a\x22\x6c\x2a\x6d\xa4\xd4\x9b\xd6\x92\x87\x54\x8a\xe4\x71\xe7\xd1\x25\x52\x24\x68\x6a\xab\xb4\x59\x96\x7f\x3b\x6b\x38\xd0\x6c\x3c\xff\xd3\x96\xff\x1a\xf4\xe5\xca\xfb\x96\x9f\x6d\xd8\xef\x76\xa6\x4e\x64\x26\xe5\xb6\x22\xae\xe4\x90\xb6\x48\xe0\x3c\x69\xb3\xe4\xb8\xdf\xb5\x08\x84\x4f\x1d\x3a\xcf\xe1\xae\xf6\xf0\xaf\x14\x6f\x69\xe9\xe0\xf3\x97\x61\x9f\x58\x98\x2d\x4c\x96\x73\x4d\x30\xd6\xd8\x1f\x6a\xb8\xd6\x1a\x87\x93\x22\xef\xe7\xa0\x8d\xff\xfd\x37\xde\xd3\x74\xa6\x0e\x5c\x69\xc6\xa9\x0e\x66\xd7\xd0\x6c\x7c\xf1\xd0\x92\x36\xbe\x49\x13\xd6\x3d\x2b\xcb\x2b\x57\x26\x39\x44\x99\x99\x94\xa2\x57\x39\x20\x11\x6f\xb7\xae\xb8\x45\xdf\xab\x34\x93\x42\x37\x21\xfa\xc3\x35\x18\xbd\xe6\x82\xa2\xad\x8c\xae\x53\x24\xca\xa4\xd8\x4b\x29\x58\xcf\xe1\x28\x77\xeb\xce\x3a\x9f\x76\xf9\x08\xcb\x67\x98\x72\xc6\x75\xf9\xe1\xf3\x2f\xb3\x2f\xb9\x14\xcc\x3a\x03\x0e\x2e\xcc\x56\x93\x35\x69\xc6\xd1\xb9\xa6\x19\x40\xaf\x72\x29\xf6\xac\xeb\x75\x01\xdc\xf1\x7e\x09\xec\x40\xf1\xa9\xd2\xfe\x96\x6c\xd7\x4a\xd1\x2f\x8b\xb7\x4a\xa5\xbf\x72\x91\xb2\x84\x55\x65\xd4\x9a\x9b\xa6\xb4\x91\x62\x69\x81\x1b\x15\x7b\x24\x14\x36\xc8\x35\x8a\xb9\x35\xc8\xd4\x82\x72\xe8\x99\x47\xdb\xe2\x4e\xb7\x1c\x93\xe2\xc5\x21\xf1\x35\x87\xda\xb6\xbb\x45\x24\xd7\xb6\xb8\xb1\xed\x2e\xed\x73\x46\x7a\xe0\x6b\xb8\x10\xeb\x1f\x77\x4d\x18\x84\xe8\x8b\x9b\xb5\x75\xf8\x49\xfb\xd5\x82\xc8\xd2\x00\x24\xc4\x5e\x1e\xb3\x41\xcb\x3e\xde\x1e\x94\xdf\xe3\x53\x1e\x19\x86\x6b\xd9\xcd\xe2\x03\xf6\xf7\xb1\xd9\x69\x72\xf7\xf1\xe1\xaf\x24\x7f\x6e\xf9\x95\xbb\x52\x65\x38\x95\xe4\x10\x9c\x71\x6d\xf1\x7e\x9e\xe5\x10\xae\xd4\xcd\xb1\xe4\x54\x63\x6c\xf4\x98\x0a\x52\x58\xc8\xd7\xe8\xf6\x70\xf7\x1c\x9b\xaa\x5b\xfb\x9b\xb5\x46\xe3\x8b\xb9\x4d\x47\x9d\x43\x65\x3c\x5b\x14\xc7\x7a\x03\xdd\x33\x83\x6c\xe7\x5f\x3a\x14\xa3\xf9\x90\x9d\xb2\xdf\xa2\x4f\xcf\xc2\xda\xce\x3f\xa7\x3d\xa2\x0e\x25\xce\xb2\xc6\xdc\x01\xf6\x85\xe7\xa3\xcb\xd1\xe2\x89\xa4\xe2\x9d\x55\xbb\xe0\xdf\x64\x3d\x31\x51\x4c\xdf\x2e\x66\x3e\x07\x8e\x44\x67\xc0\x91\x28\x1f\xb2\x17\x81\x23\xd1\xff\x81\x0f\x25\xce\x82\xc7\xdc\x65\xe0\x13\x49\xcf\xc0\xc7\xf5\x05\xe0\xfd\x32\xfc\x58\xc3\x33\x7e\xd3\xfe\xc6\x2a\xbc\x9f\x8e\x92\xef\x40\xf2\xfe\xda\x2a\x3c\xc1\xbc\x6c\x58\x8c\xd7\xf1\xd0\x3c\x9c\x89\xb3\x8b\x7f\x4a\x73\xe4\xda\x94\x4e\x55\x45\xca\x22\xa6\xd2\x1f\xc7\x54\xf6\xe6\xb5\xfb\x78\xbc\x7d\xd3\xfe\x50\x2c\x3b\xcc\xe8\xe3\xa8\x1c\x26\x7c\x98\x98\xe3\xd4\xcc\x20\xfd\x69\x9c\xf4\xa1\x25\x96\x82\x13\x8f\x56\xed\xb8\x3b\x06\xfb\x34\x7c\xac\x8a\x77\x5d\xd3\xf0\x08\x1f\x41\x66\x47\x92\x85\x89\x24\x8f\x41\x7c\x5c\xa5\x84\x4f\xa7\xb2\x09\x7d\x47\x86\x97\xe1\xb2\xa8\xbd\x2c\x81\x05\x42\x05\x7f\x3e\x7c\xfc\x00\xca\xd6\xdd\x06\xb9\x65\x74\xe2\xd3\x38\xf3\x93\xaa\x6d\xd7\xba\xae\xbc\xb6\x26\x7c\x3a\xdf\x40\xbd\xaa\xc8\xa1\xbf\xee\x7c\xf3\xf3\x1f\x49\x0e\x8f\xf1\x7d\x39\xb5\xea\x8c\x86\x61\x32\xd3\xe8\xc0\xe1\xbd\x0a\xe2\xee\xb1\x46\xbd\x45\xf0\x2b\x0c\x55\xa1\x72\x41\x69\x74\x59\x05\xaf\xd4\xe1\x7b\xf9\x1d\xa7\xe9\xd4\xe1\xe1\xf4\x85\x9d\x1a\x82\xe3\xa9\x9c\xd3\x72\x2f\xff\x0b\x00\x00\xff\xff\xb4\x14\x88\x1f\x67\x08\x00\x00")

func dataClientGoBytes() ([]byte, error) {
	return bindataRead(
		_dataClientGo,
		"data/client.go",
	)
}

func dataClientGo() (*asset, error) {
	bytes, err := dataClientGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/client.go", size: 2151, mode: os.FileMode(420), modTime: time.Unix(1505205622, 0)}
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
	"data/client.go": dataClientGo,
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
	"data": &bintree{nil, map[string]*bintree{
		"client.go": &bintree{dataClientGo, map[string]*bintree{}},
	}},
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

