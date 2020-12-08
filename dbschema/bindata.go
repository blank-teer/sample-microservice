// Code generated for package dbschema by go-bindata DO NOT EDIT. (@generated)
// sources:
// migrations/1_create_orders_table.sql
// migrations/2_create_emission_table.sql
// migrations/3_create_coverage_table.sql
package dbschema

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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _migrations1_create_orders_tableSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x94\xcf\x6e\x9c\x30\x10\x87\xef\x3c\xc5\x28\x97\x65\xd5\x46\xea\x9f\x34\x97\x9c\x1c\xf0\x36\x56\x89\x4d\x8d\xdd\x6e\x4e\x96\x8b\xdd\xca\x52\x02\xc8\x98\x6e\xf7\xed\x2b\x02\x69\xb2\x40\x45\xd5\x3d\x5a\xfe\xbe\xfd\x0d\x33\x9e\xf3\x73\x78\xf5\xe0\x7e\x78\x1d\x2c\xc8\x26\x4a\x38\x46\x02\x83\x40\xd7\x19\x06\xb2\x03\xca\x04\xe0\x3d\x29\x44\x01\x67\xb5\x37\xd6\xb7\x67\x51\x1c\x01\x00\x38\x03\xf3\x9f\x94\x24\x3d\x39\xe8\x79\x2a\xb3\x0c\x24\x25\x9f\x25\x86\x9c\x93\x5b\xc4\xef\xe0\x13\xbe\x7b\xfd\xa8\x09\xc7\xc6\xce\x34\x5f\x10\x4f\x6e\x10\x8f\x2f\x2f\xb6\xa7\x9a\x81\x69\x83\x0e\x5d\xfb\x17\xe6\xfd\xbb\x09\x03\x29\xde\x21\x99\x09\xd8\x54\xf6\xb0\x19\x0c\x8d\xaf\x7f\x3a\x63\xfd\xa2\xe1\xed\xe5\xe2\xbf\x76\xad\xf5\x6a\x5a\xf5\x35\xf9\x48\xa8\x18\x2e\x1c\xf4\xfd\xbd\x0d\x4a\x1b\xe3\x6d\xdb\x4e\xa4\x17\x1f\x16\xa5\xdf\x74\x6b\x55\xd9\x79\x6f\xab\xf2\xb8\x1a\xe4\xb9\x14\x2c\xf9\x66\xc1\xa0\xf4\x43\xdd\x55\x61\x8c\xb5\xd0\x87\x25\xa6\xf1\xae\xb4\x00\x29\x93\x7d\xd3\x73\x8e\x13\x52\x10\x46\x27\x4c\xa3\x8f\xb3\xa0\x6b\x9f\xec\x25\xf3\x14\x6d\x2d\xdb\x09\x33\x46\x5b\xcb\xd6\x4f\xaf\xea\xc1\x50\xab\xbe\xb6\x47\xed\x0a\x53\xd6\xae\x9a\xce\xd0\x5a\xb6\xef\x76\x3e\xab\x6b\x4c\xe9\xad\x0e\xd6\x28\x1d\xfe\x85\xf9\xd3\x61\xbc\x17\x1c\x25\x22\xc6\x39\x4b\x6e\x60\xc7\xd9\x2d\x50\xf6\x35\xde\x6e\xc7\x61\x6c\xcc\xff\x58\xdf\x44\xdb\xab\xe8\xe9\x91\x13\x9a\xe2\x3d\x0c\xaf\x5a\x39\xa3\x9c\xf9\x05\x8c\x8e\x07\x10\x3b\xd3\xdf\x7d\xb9\x1f\xd2\xfa\x50\x45\x29\x67\xf9\xf3\x7e\x98\xec\x06\x48\x50\x91\xa0\x14\x5f\x45\xbf\x03\x00\x00\xff\xff\x02\x3f\x29\xcd\x58\x04\x00\x00")

func migrations1_create_orders_tableSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations1_create_orders_tableSql,
		"migrations/1_create_orders_table.sql",
	)
}

func migrations1_create_orders_tableSql() (*asset, error) {
	bytes, err := migrations1_create_orders_tableSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/1_create_orders_table.sql", size: 1112, mode: os.FileMode(420), modTime: time.Unix(1599659910, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations2_create_emission_tableSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x93\xdf\xef\x9a\x30\x14\xc5\xdf\xf9\x2b\x6e\xbe\x2f\x42\x36\x93\xcd\x4c\x5f\x7c\xaa\x50\x67\x37\x6c\x59\x69\x37\x7d\x22\xcc\x76\xa6\x89\x80\xe1\x47\x36\xb3\xec\x7f\x5f\x50\x61\x28\x90\x25\x5f\x1e\x6f\x38\x9f\x73\x6f\x4f\xce\x74\x0a\x6f\x12\x73\xcc\xe3\x52\x83\x3c\x5b\x2e\xc7\x48\x60\x10\x68\xe5\x63\x20\x6b\xa0\x4c\x00\xde\x91\x50\x84\xf0\xa2\x13\x53\x14\x26\x4b\x5f\x2c\xdb\x02\x00\x48\xab\x04\x7a\x5f\x88\x39\x41\x7e\x67\x50\x13\xa8\xf4\xfd\xb7\x57\x8d\x51\x7d\x09\x48\x49\xbc\x87\x41\xa3\x01\x49\xc9\x17\x89\x21\xe0\x64\x8b\xf8\x1e\x3e\xe3\xfd\x0d\x53\x5e\xce\xba\x87\xf9\x8a\xb8\xbb\x41\xdc\x5e\x7c\x70\x86\xac\x8f\xd9\x49\x45\x55\x6a\xca\x21\xcd\xfb\xd9\x93\x06\x3c\xbc\x46\xd2\x17\x30\x39\xe6\x71\x52\x4c\x3a\x8c\x38\xc9\xaa\xb4\x43\x59\x91\x8f\x84\x8a\xd1\x93\xbf\xc7\x85\x8e\x0e\x55\x9e\xeb\xf4\x70\xe9\xf9\x2e\x46\x7d\xb1\xe4\x93\x01\x42\x63\xff\x1f\xd7\x3a\xd0\xe8\x2a\x2c\xb3\xf6\x6a\x8f\xc9\x3a\xd7\x80\x63\x97\x84\x84\xd1\x67\x8d\xfe\xa1\x6b\x0f\x3d\xf4\x42\xb3\xf9\xdc\x19\xd9\xf4\xbe\xa6\xd2\x65\x6c\x4e\xc5\x63\x28\x9f\x42\x46\x57\x83\xd9\xb6\xf2\xdf\x7f\xee\x80\x43\xae\xe3\x52\xab\x28\xee\x46\x34\x7a\x67\x0b\xc0\x3b\xc1\x91\x2b\x6c\x1c\x30\x77\x03\x6b\xce\xb6\x40\xd9\x37\xdb\x71\x6e\xd4\xea\xac\x5e\x43\x7d\x67\x39\x4b\xab\xe9\x03\xa1\x1e\xde\x41\x53\x80\xc8\xa8\xc8\xa8\x5f\xc0\x68\x3b\x02\xdb\xa8\xfa\xff\x6e\x9d\xbc\xec\x67\x6a\x79\x9c\x05\xff\xea\xd4\xab\x12\xb8\x28\x74\x91\x87\x97\xd6\xdf\x00\x00\x00\xff\xff\x41\x10\x01\x78\x89\x03\x00\x00")

func migrations2_create_emission_tableSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations2_create_emission_tableSql,
		"migrations/2_create_emission_table.sql",
	)
}

func migrations2_create_emission_tableSql() (*asset, error) {
	bytes, err := migrations2_create_emission_tableSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/2_create_emission_table.sql", size: 905, mode: os.FileMode(420), modTime: time.Unix(1599659739, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations3_create_coverage_tableSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x92\x41\x6b\xc2\x30\x14\xc7\xef\xf9\x14\x7f\x3c\xb5\x6c\xc2\xee\x9e\x62\xf3\xba\x85\xd5\x54\xd2\x94\xe9\xa9\x14\x13\x24\x30\x8d\xd4\xba\xed\xe3\x0f\xed\xe2\x0a\x32\x97\xeb\xfb\xe5\x97\xbc\xff\x7b\xd3\x29\x1e\x76\x7e\xdb\xb5\xbd\x43\x7d\x60\x99\x26\x6e\x08\x86\xcf\x0b\xc2\x64\x13\x3e\x5c\xd7\x6e\xdd\x84\x25\x0c\x00\xbc\xc5\xf5\x54\xa4\x25\x2f\xa0\x4a\x03\x55\x17\x05\x96\x5a\x2e\xb8\x5e\xe3\x95\xd6\x8f\x17\xd8\xed\xfc\xf1\xe8\xc3\xbe\xf1\x16\x75\x2d\x05\xf0\x0b\x6b\xca\x49\x93\xca\xa8\xba\x62\x48\xbc\x4d\x51\x2a\x08\x2a\xc8\x10\x32\x5e\x65\x5c\xd0\xe0\x0a\x9d\x75\x5d\x33\x3c\x7f\xc7\x75\xc1\x8e\x77\x4d\xed\x2e\x9c\xf6\xfd\xd0\xc2\x5c\x3e\x4b\x65\xae\xa6\x01\xd8\x74\xae\xed\x9d\x6d\xda\xfe\x06\x80\xa0\x9c\xd7\x85\x01\xad\x8c\xe6\x99\x49\x68\x59\x66\x2f\xc8\x75\xb9\x80\x2a\xdf\x92\x34\x1d\x14\xa7\x83\xfd\x57\xf1\xc4\xd2\x19\x8b\x71\x4b\x25\x68\x85\x98\x76\xb3\x0d\xef\xb6\x19\xc5\xd7\x78\xfb\x75\xee\x27\x02\x48\x46\xc5\x74\xf6\x87\x25\x66\x76\x7b\x3b\x56\xce\x3f\x18\xcf\x5f\x84\xcf\x3d\x13\xba\x5c\xfe\xcc\x5f\xe6\xa0\x95\xac\x4c\x35\xda\x84\x98\xe6\x8c\x7d\x07\x00\x00\xff\xff\xdd\x1f\xf6\xab\x3a\x02\x00\x00")

func migrations3_create_coverage_tableSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations3_create_coverage_tableSql,
		"migrations/3_create_coverage_table.sql",
	)
}

func migrations3_create_coverage_tableSql() (*asset, error) {
	bytes, err := migrations3_create_coverage_tableSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/3_create_coverage_table.sql", size: 570, mode: os.FileMode(420), modTime: time.Unix(1599658206, 0)}
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
	"migrations/1_create_orders_table.sql":   migrations1_create_orders_tableSql,
	"migrations/2_create_emission_table.sql": migrations2_create_emission_tableSql,
	"migrations/3_create_coverage_table.sql": migrations3_create_coverage_tableSql,
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
	"migrations": &bintree{nil, map[string]*bintree{
		"1_create_orders_table.sql":   &bintree{migrations1_create_orders_tableSql, map[string]*bintree{}},
		"2_create_emission_table.sql": &bintree{migrations2_create_emission_tableSql, map[string]*bintree{}},
		"3_create_coverage_table.sql": &bintree{migrations3_create_coverage_tableSql, map[string]*bintree{}},
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
