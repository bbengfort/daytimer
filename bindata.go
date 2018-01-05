// Code generated by go-bindata.
// sources:
// templates/agenda.html
// templates/config.yml
// templates/email.txt
// DO NOT EDIT!

package daytimer

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

var _templatesAgendaHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x56\xdf\x93\x9b\x36\x10\x7e\xbf\xbf\x62\xab\x4c\xfb\x66\x64\xf7\xd2\xe4\x7e\x80\x67\x1a\xdf\xa5\xc9\xcc\x35\xbd\x69\x48\x7f\x3c\x0a\xb4\x06\x4d\x85\xe4\x0a\xd9\x3e\x97\xe1\x7f\xef\x48\xc0\xd9\x87\xc1\xb9\xb9\xd8\x0f\xb0\xa0\xfd\x76\xf5\xed\xb7\x8b\xc2\xef\x6e\x7e\x5b\xc4\x7f\xdf\xdf\x42\x6e\x0b\x09\xf7\x5f\xde\xdd\x7d\x5c\x00\x99\x50\xfa\xe7\xf9\x82\xd2\x9b\xf8\x06\xfe\xfa\x10\xff\x7a\x07\xb3\x60\x0a\xb1\x61\xaa\x14\x56\x68\xc5\x24\xa5\xb7\x9f\x08\x90\xdc\xda\xd5\x15\xa5\xdb\xed\x36\xd8\x9e\x07\xda\x64\x34\xfe\x9d\x3e\x38\xac\x99\x73\x6e\x6f\x27\xf6\xc0\x33\xe0\x96\x93\xf9\x59\xe8\x03\x3e\x14\x52\x95\xd1\x00\xcc\xec\xf2\xf2\xb2\xf1\x26\xf3\x33\x80\x30\x47\xc6\xdd\x0d\x40\x58\xa0\x65\xe0\x3c\x26\xf8\xef\x5a\x6c\x22\xb2\xd0\xca\xa2\xb2\x93\x78\xb7\x42\x02\x69\x63\x45\xc4\xe2\x83\xa5\x0e\xe1\x1a\xd2\x9c\x99\x12\x6d\xf4\x25\x7e\x3f\xb9\x20\x40\x5b\x24\x2b\xac\xc4\x79\x48\x9b\x6b\xf3\xac\xb4\x3b\xff\xac\xb9\xba\xd0\xb4\x8b\x1d\x26\x9a\xef\x3a\x57\x96\x48\x84\x44\x1b\x8e\x26\x22\x53\x02\x29\x4a\xb9\x62\x9c\x0b\x95\x3d\xda\xe5\x8a\xa5\x9d\x9d\xa3\xc8\x72\x1b\x91\xd9\x74\xfa\x3d\x81\xad\xe0\x36\xef\x0c\xc1\x23\xe2\xa0\x63\x87\x49\x9a\x00\x2e\x84\xe9\x6e\x9d\xc1\x81\x49\x91\xa9\x88\xa4\xa8\x2c\x1a\x02\x9b\xd6\xb6\x7a\x45\xf6\x0b\xbf\x9a\xda\x4f\xc7\xa9\xb5\xc9\xbc\xbd\x98\x36\xb9\x60\xc1\x84\x74\xa4\x32\xa1\xd0\x3c\x01\xef\xa5\xb5\x4f\x2e\xd5\x0e\x52\x45\xe4\xbc\x17\x6f\x36\x25\x8e\x61\xde\x43\xa1\x7d\x98\x11\xdc\xc3\x5d\xf6\x19\x38\x8c\xd9\xf7\xf5\x95\x34\x5a\x65\xf3\x30\x11\xd9\xbc\xaa\x82\xd8\x15\xb9\xae\x43\xea\x6c\x57\x5f\xff\xf6\xd8\x2d\x31\x9d\x3e\x9e\xa2\x15\x4c\xca\x79\xb8\xd4\xca\xba\xc0\xda\x44\xe4\xd5\xa5\xff\x11\x87\xbe\xd0\x6b\x65\xeb\x1a\x6e\x37\xa8\x6c\x09\x9f\xd3\x1c\xf9\x5a\x22\x87\xa5\x36\x10\x6b\xce\x76\x21\x75\xbe\x2e\xb2\x47\xea\xef\xf5\xe5\x14\x7d\x33\xf5\x55\x65\x98\xca\x10\x82\x8f\x16\x8b\xb2\xae\xfb\x41\xc1\xf7\x42\x44\xf4\xda\x4a\xa1\xf0\x0a\x6c\x2e\x14\x94\x5a\x0a\x0e\xaf\xde\xdf\xba\xff\x35\xb4\x2f\x27\x7a\xb9\x2c\xd1\x5e\xc1\x64\xb6\x7a\xb8\x26\x03\x44\x5a\xde\xe1\x15\x42\x4d\xbc\xf6\xae\x60\x36\x7d\xed\x96\x3f\x96\xbb\x10\x9c\x4b\x24\x90\x64\x2d\xd7\x9e\x62\xa9\x0d\x90\x24\x23\x75\xdd\xd7\xc2\x71\x1c\x80\x27\xb5\xda\xfb\x2f\xbd\xbf\xab\xd9\x3d\x1a\xa1\xb9\x93\x84\xaf\xcc\x71\xaa\x47\xdc\x41\x4f\x95\x5d\x9a\xad\x29\x71\x69\x87\x53\x61\x90\x1b\x5c\xfa\x34\xbc\x42\x82\x0f\xb6\x90\x77\x42\xfd\xd3\xa6\xd2\x3c\xfc\xbc\x2e\x0a\x66\x76\x2e\x23\x36\x08\x33\xac\xcd\x47\x75\x56\x15\xb4\x48\x77\x3a\x65\x6e\xdc\x82\xc3\x1a\x14\xdc\x0b\xb6\x77\x8a\xeb\x81\xf6\xe8\x97\xcc\xb7\x09\x93\xa8\x38\x33\x87\x5b\x3d\xd5\x16\xcf\x6b\x8c\xaa\x42\x59\xe2\xb1\x70\x4f\x0d\x94\xe1\xbd\x8d\x37\xd3\x9b\xe9\x90\x98\x9f\x8e\x83\x37\x3f\x5e\xcc\xde\x8d\xf0\xd3\x8c\x9b\x4f\xfa\xe4\x84\x18\x1d\x4a\x83\x0a\x7d\x26\x35\x8a\x3f\x8f\x99\xd1\xad\xbf\x9e\x92\xf9\x0f\x2a\x29\x57\xd7\xdf\x30\xa5\x46\xbe\x5e\x89\xb6\x56\x17\x27\x78\x9f\x0d\xf2\x3e\x34\x8d\xdf\xfa\xdf\x20\xfd\xbf\xa0\x42\xc3\x2c\x72\x48\x76\x60\x73\xdc\x37\xa4\x3b\x47\x94\x57\x94\x66\xc2\xe6\xeb\x24\x48\x75\x41\x93\x04\x55\xb6\xd4\xc6\x52\xce\x76\x56\x14\x4e\xf2\x37\xed\x1d\xfc\xbc\x5a\xb9\xe6\x84\x4d\x55\x05\x7f\xa0\x29\x85\x56\x75\x3d\xda\x97\xb1\x28\xf0\x3f\xad\x10\xb8\x28\x53\xbd\x41\x83\x1c\x58\xf9\xa8\x86\xaa\x0a\x3a\xe0\xa0\x5b\xea\x1b\xb6\xfb\x72\x8d\xc0\x56\x95\x44\x05\xcd\x37\xa7\xac\x6b\x48\xdb\xae\x2a\x41\xa8\x54\xae\x39\x72\x10\xca\x4d\xe9\x12\x58\xe6\xde\x8c\x09\xea\xc5\x1f\xa3\x90\xfa\x53\xc6\xc1\x01\xe5\xc0\x65\xbf\xf8\x60\x59\x48\x9b\xd3\x53\xe8\x8f\x64\xf3\xb3\xff\x03\x00\x00\xff\xff\x33\x13\xb3\x4e\x7a\x0a\x00\x00")

func templatesAgendaHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesAgendaHtml,
		"templates/agenda.html",
	)
}

func templatesAgendaHtml() (*asset, error) {
	bytes, err := templatesAgendaHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/agenda.html", size: 2682, mode: os.FileMode(420), modTime: time.Unix(1515097267, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesConfigYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x8f\x41\x4b\xc3\x40\x10\x85\xef\xfb\x2b\x1e\x89\xe0\x2d\x37\x51\x72\x13\xab\xe0\x41\x2a\xb6\x17\x4f\xb2\x74\x27\xcd\xe0\x66\x27\xec\xcc\x5a\xf3\xef\x65\xd3\x2a\x78\x7c\x8f\xe1\xfb\xde\xb4\xd8\xf8\xc5\x78\xa2\x8c\x07\x49\x03\x1f\x4b\xf6\xc6\x92\xf0\xc4\x91\x9c\x6b\xb1\x23\x83\x8d\x84\xa1\xc4\x88\xd9\xdb\x08\x93\xb5\x08\x34\xf8\x12\x0d\x46\xdf\x06\x0a\x6c\x92\xb1\x48\xb9\x0e\x88\xfc\x49\xf5\xaa\x28\x75\xae\xc5\xbb\x14\x1c\x7c\x82\x8f\x2a\x50\x32\x5c\x3d\x6e\x9e\xf7\xdb\x37\x70\x5a\x41\x94\xbe\x38\x4b\x9a\x28\x59\x87\xad\x8d\x94\x4f\xac\x84\x13\xc7\x08\xcb\x4b\x25\x0d\x9c\x02\xbc\x6b\xff\xa4\x17\xdf\xaf\xe4\x1c\x7b\x34\x4d\x9d\x7c\x1f\xc2\x0a\xde\xbd\xec\x5f\x71\xf8\xf7\x95\xd5\x05\x29\x80\x26\xcf\x11\xfe\x48\x29\x78\x75\x6b\xea\x1d\x30\x8a\x5a\x8f\x46\x27\x9b\xbb\x63\x2d\xbb\x83\x4c\x8d\x03\x66\xc9\xd6\xe3\xe6\xee\xd6\xa1\x1a\xcf\x2a\x60\xf6\xaa\x27\xc9\xe1\x12\x8b\xd2\x87\x45\xed\x61\xb9\x90\xfb\x09\x00\x00\xff\xff\xd6\x1b\x83\x3f\x5b\x01\x00\x00")

func templatesConfigYmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesConfigYml,
		"templates/config.yml",
	)
}

func templatesConfigYml() (*asset, error) {
	bytes, err := templatesConfigYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/config.yml", size: 347, mode: os.FileMode(420), modTime: time.Unix(1515191243, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesEmailTxt = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x2b\xca\xcf\xb5\x52\xa8\xae\xd6\x03\x31\x6a\x6b\xb9\x42\xf2\xc1\xbc\x90\xfc\xda\x5a\xae\xe0\xd2\xa4\xac\xd4\xe4\x12\xb0\x00\x94\x5d\x5b\xcb\xe5\xeb\xe9\xeb\xaa\x5b\x96\x5a\x54\x9c\x99\x9f\x67\xa5\x60\xa8\x67\xc0\xe5\x9c\x9f\x57\x92\x9a\x57\xa2\x1b\x52\x59\x90\x6a\xa5\x50\x92\x5a\x51\xa2\x9f\x51\x92\x9b\x63\xad\x90\x9c\x91\x58\x54\x9c\x5a\x62\xab\x56\x58\x9a\x5f\x62\x1d\x1a\xe2\xa6\x6b\x01\x61\x72\x71\x71\x55\x57\xeb\x39\xe5\xa7\x54\xd6\xd6\x72\x01\x02\x00\x00\xff\xff\xc1\x9a\x9f\x3e\x83\x00\x00\x00")

func templatesEmailTxtBytes() ([]byte, error) {
	return bindataRead(
		_templatesEmailTxt,
		"templates/email.txt",
	)
}

func templatesEmailTxt() (*asset, error) {
	bytes, err := templatesEmailTxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/email.txt", size: 131, mode: os.FileMode(420), modTime: time.Unix(1513945137, 0)}
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
	"templates/agenda.html": templatesAgendaHtml,
	"templates/config.yml": templatesConfigYml,
	"templates/email.txt": templatesEmailTxt,
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
	"templates": &bintree{nil, map[string]*bintree{
		"agenda.html": &bintree{templatesAgendaHtml, map[string]*bintree{}},
		"config.yml": &bintree{templatesConfigYml, map[string]*bintree{}},
		"email.txt": &bintree{templatesEmailTxt, map[string]*bintree{}},
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

