// Code generated by go-bindata.
// sources:
// views/index.html
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

var _viewsIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x92\x3d\x53\xf3\x30\x0c\xc7\xf7\x7c\x0a\x3f\x59\xba\x3c\xb1\x13\xde\x86\x36\xc9\x02\x9d\xe1\xb8\x2e\x8c\xaa\x2d\xa8\x72\x76\x6c\x62\xa5\x6d\xbe\x3d\x79\x29\x03\x30\x71\xc7\xe4\x93\x2c\xff\xfc\xff\x4b\x2a\xff\x3d\x3c\xde\xef\x5e\x9e\xb6\xe2\xc0\xce\xd6\x49\xf9\x79\x20\x98\x3a\x11\xa2\x74\xc8\x20\x5a\x70\x58\xa5\x47\xc2\x53\xf0\x1d\xa7\x42\xfb\x96\xb1\xe5\x2a\x3d\x91\xe1\x43\x65\xf0\x48\x1a\xb3\x39\xf8\x2f\xa8\x25\x26\xb0\x59\xd4\x60\xb1\x2a\x64\x9e\xaa\x99\x14\x75\x47\x81\x45\xec\x74\x95\x2a\xa5\x4d\xdb\x44\xa9\xad\xef\xcd\xab\x85\x0e\xa5\xf6\x4e\x41\x03\x67\x65\x69\x1f\x55\xf3\xde\x63\x37\xa8\x6b\x99\xcb\x3c\x03\x1b\x0e\x50\x5c\x72\xd2\x51\x2b\x9b\x98\xd6\xa5\x5a\x88\xbf\x87\x77\x08\x9a\x55\x2e\x8b\x1b\x79\xb5\x04\x7f\x0e\xcc\x8c\x77\x3f\xa1\x4c\x6c\xb1\x7e\x9e\x0a\x44\x26\x76\x43\xc0\xcb\x37\xdb\x33\xb8\x60\xb1\x54\x4b\x45\x52\xaa\x65\x00\x49\xb9\xf7\x66\x98\xdf\x1a\x3a\x0a\x32\xd5\xea\xd2\xfb\x95\x88\x3c\x8c\xfd\x4d\x1d\x74\x6f\xd4\xae\xa1\x67\xbf\x11\xf3\x08\xd6\x77\x79\x1e\xce\x1b\xb1\xdc\x64\xec\xc3\xba\xb8\x1d\x33\x93\x98\x91\x32\xd1\x17\x6a\xf2\xd5\x65\xe8\xf7\x96\xb4\x6a\xa2\x1a\xdd\x05\xb2\x68\x54\x9c\x65\x7d\x33\x32\x8a\x9b\x96\xe4\x23\x00\x00\xff\xff\xdf\x52\xdf\xb6\x3b\x02\x00\x00")

func viewsIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_viewsIndexHtml,
		"views/index.html",
	)
}

func viewsIndexHtml() (*asset, error) {
	bytes, err := viewsIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "views/index.html", size: 571, mode: os.FileMode(420), modTime: time.Unix(1472426701, 0)}
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
	"views/index.html": viewsIndexHtml,
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
	"views": &bintree{nil, map[string]*bintree{
		"index.html": &bintree{viewsIndexHtml, map[string]*bintree{}},
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
