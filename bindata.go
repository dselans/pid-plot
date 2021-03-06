// Code generated by go-bindata.
// sources:
// index.template
// DO NOT EDIT!

package main

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

var _indexTemplate = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x94\x56\x4b\x6f\xdb\x38\x10\xbe\xe7\x57\xcc\xaa\x97\xb6\x08\x45\x5b\x0a\xd2\xd4\x95\x85\x05\x16\xbd\x6d\xb1\x8b\x14\xbb\xd7\x05\x4d\x32\x26\x53\xbe\x96\xa4\xfc\x80\x91\xff\x5e\x88\x92\x63\x57\x56\x0c\xfb\x24\x72\xf8\x3d\x66\x3c\xe4\xc0\x95\x88\x5a\xd5\x37\x00\x95\xe0\x84\xb5\x0b\x80\x2a\xca\xa8\x78\xed\x24\x43\x4e\xd9\x58\xe1\x6e\xdf\x9d\xfd\x86\x10\x54\x4a\x9a\x1f\xe0\xb9\x9a\x67\x21\x6e\x15\x0f\x82\xf3\x98\x81\xf0\xfc\x69\x9e\x89\x18\x5d\x98\x61\xac\xc9\x86\x32\x93\x2f\xac\x8d\x21\x7a\xe2\xda\x0d\xb5\x1a\xbf\x06\x70\x99\x97\xf9\x27\x4c\x43\x38\xc4\x72\x2d\x4d\x4e\x43\xc8\x6a\x40\xa8\x77\xbc\xc4\x2c\x09\xac\x49\xa4\x22\x79\xd0\xed\xc2\xfa\xe5\x98\xec\x4d\xa7\x99\x94\x3a\xfd\xe8\x21\x8a\xdb\xf4\x61\xb0\x4b\x21\x00\xda\xf8\x60\xfd\x0c\x9c\x95\x26\x72\xff\x25\x85\x5f\x3a\x32\xfe\x08\xdf\x35\x51\x0a\x18\x5f\x49\xca\x03\xbc\x8f\x64\xa1\x78\x0c\xb7\xf0\xe9\xfe\xc1\x6d\x80\x18\x06\x8d\xfb\x00\x1f\x71\xc2\xff\xae\x39\x93\x04\xde\x6b\x69\xd0\x5a\xb2\x28\x66\x1d\xee\x03\xec\xe0\x5d\xe0\xc4\x53\x81\xa4\x71\x4d\x84\x1d\xf4\xe7\x77\x93\x89\xdb\x7c\x81\x97\x23\xcb\x6f\x9c\xc9\x46\x1f\x3c\x19\x0f\x3f\xa2\x75\xe1\x16\x3e\x7f\x2e\x2e\x32\x4d\xb8\x33\xa6\xf7\xa7\xa6\x7f\x12\xbf\xe4\x07\x4f\xd5\x6f\xf7\xce\xd3\xa2\xa5\x5c\x60\xdd\x01\xcf\x78\x3f\x1c\xbc\x53\x7f\x70\xdf\xa0\xbe\x5b\xd4\x4b\x17\x21\x78\x7a\xe8\x38\xb5\x8c\xe7\xcf\xff\x37\xdc\x6f\x53\xcb\xbb\x25\x2a\xf2\x69\x5e\xa6\x76\x3f\x87\xac\xae\x70\x47\xad\xdf\xd6\xb9\xf4\x9a\x3e\x0f\x6f\xe9\x50\xbf\xc2\xdd\x0b\x6a\x97\x0b\xcb\xb6\xbd\x27\x93\x2b\xa0\x8a\x84\x30\xcf\xa8\x35\x91\x48\xc3\xfd\xfe\x1a\x0e\x01\xde\xae\xb3\xfa\xf5\x64\x48\x56\x48\x2d\xd1\xb4\x80\x76\xa5\x19\x9a\x4e\xd2\x2a\x68\xf4\x70\xac\x97\x78\x86\xbc\xf2\x0c\x59\x2d\x88\x87\xee\x83\xa4\x59\x71\x1f\xf8\x2f\x2e\x6f\xa6\x89\x9e\x54\x23\xd9\x09\xf6\x57\x74\x2f\xdc\xd6\x9e\x0a\x1b\x62\x01\x2a\x32\xc0\x2e\x3c\x31\x6c\xff\x82\xdf\x65\x47\x83\x86\x8c\x78\x61\x26\x57\xe7\x53\xa0\x56\x29\xe2\x02\xdf\x17\xb9\xdf\x67\x20\xd9\x3c\x5b\x04\xc4\x37\x44\x3b\xc5\xd1\xe0\x1c\x15\xe3\xf9\x36\xea\x28\xe1\xbd\xe8\xd1\xd2\xcb\xa5\x88\xa3\xd4\x34\xad\xea\x8a\xf4\xb5\xe1\x95\xe4\x6b\x5c\x4e\x74\x56\x97\x13\xdd\x96\x57\x61\x25\x2f\x26\x4e\x45\x56\x4f\xc5\xd5\xb4\x52\x64\x75\x79\x3d\xed\x5e\x64\xf5\xfd\xf5\xb4\x69\xd1\x66\x59\x5c\x4f\x2c\xee\x44\x56\x17\x77\x67\x89\x15\x6e\xd4\x85\x57\x62\x24\x58\x61\x43\x56\xc3\xa7\x71\x74\x6f\x9e\x1b\xbd\xb0\xd1\x5b\x93\x41\x1a\x37\xf3\x4c\x13\xbf\x94\x06\x45\xeb\x66\x50\xa4\x89\xe4\x08\x63\xd2\x2c\xdb\x91\xec\x36\x27\x3d\xdf\xed\x40\x3e\x41\xfe\xd5\x7b\xeb\xe1\xe5\x65\x98\x92\xab\xbf\x3e\x3e\xfe\xf5\x38\x6b\x71\xaf\xa0\x0a\xbb\x11\x19\xae\x02\x1f\x51\x10\x65\xdd\x72\xbf\x2b\x49\xdb\x63\xf8\x57\xf2\x75\x85\x45\x39\xa8\x2a\x79\xfd\xf1\xf7\x3f\x23\xda\x95\xd4\xcb\x6e\xe2\x75\x3f\xfb\xb1\x1c\xa6\xae\xf9\x4f\xf3\xe8\x25\x0d\xb9\x33\xcb\xd3\xb1\xe0\xea\x6f\x5c\x5b\xbf\xbd\x56\x58\x73\x7d\x56\xb8\xad\xd8\xb0\x61\xc1\x7d\x0f\x6f\x4e\x22\xa3\xdb\x03\xba\xc2\xdd\xc0\xad\x70\xf7\x6f\xe6\x67\x00\x00\x00\xff\xff\x0b\x94\xf0\x4f\xd5\x08\x00\x00")

func indexTemplateBytes() ([]byte, error) {
	return bindataRead(
		_indexTemplate,
		"index.template",
	)
}

func indexTemplate() (*asset, error) {
	bytes, err := indexTemplateBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "index.template", size: 2261, mode: os.FileMode(420), modTime: time.Unix(1485755423, 0)}
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
	"index.template": indexTemplate,
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
	"index.template": &bintree{indexTemplate, map[string]*bintree{}},
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

