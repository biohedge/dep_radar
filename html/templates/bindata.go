// Code generated by go-bindata.
// sources:
// html/templates/apps.html
// html/templates/libs.html
// html/templates/main.html
// DO NOT EDIT!

package templates

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

var _htmlTemplatesAppsHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x95\x6f\x6f\xe2\x38\x10\xc6\x5f\x37\x9f\xc2\x17\x55\xda\x56\xdd\xc4\x50\x28\x2d\xab\x24\x12\xd7\x2d\x6c\xa1\xe5\xf8\xb3\x6d\xe9\x9e\xee\x85\x13\x9b\xc4\xe0\xd8\xa9\x3d\x50\x38\x94\xef\x7e\x4a\xa0\xb4\xbb\x6a\xa5\xbb\x7b\x03\x9e\x67\x3c\xf6\x6f\x46\x8f\x1c\xef\x37\xaa\x22\x58\x67\x0c\x25\x90\x8a\xc0\xf2\x8a\x3f\x24\x88\x8c\x7d\x9b\x49\xbb\x10\x18\xa1\x81\x75\xe0\xa5\x0c\x08\x8a\x12\xa2\x0d\x03\xdf\x5e\xc0\xd4\xb9\xb0\xf7\xba\x24\x29\xf3\xed\x25\x67\xcf\x99\xd2\x60\xa3\x48\x49\x60\x12\x7c\xfb\x99\x53\x48\x7c\xca\x96\x3c\x62\x4e\x19\x7c\x46\x5c\x72\xe0\x44\x38\x26\x22\x82\xf9\xd5\xcf\xc8\x24\x9a\xcb\xb9\x03\xca\x99\x72\xf0\xa5\x2a\xcf\x15\x5c\xce\x91\x66\xc2\xb7\x0d\xac\x05\x33\x09\x63\x60\xa3\x44\xb3\xa9\x6f\x27\x00\x99\xf9\x82\x71\x4a\x56\x11\x95\x6e\xa8\x14\x18\xd0\x24\x2b\x82\x48\xa5\x78\x2f\xe0\xba\x5b\x71\x2b\x38\x32\xe6\x55\x73\x53\x2e\xdd\xc8\x18\x1b\x71\x09\x2c\xd6\x1c\xd6\xbe\x6d\x12\x52\xbb\xa8\x3b\x1d\x79\x56\xbb\xa8\xaf\x9e\x86\x55\xa2\x1e\x26\xad\x93\xca\xd9\xc5\x68\x32\x58\x0d\xe2\xc6\x74\x5d\xbf\x7e\x58\x7e\xef\x27\x95\xab\xd3\x46\x6d\x92\xb6\xa3\xae\x18\xb7\x9e\x79\x27\x6e\xb7\x1e\x30\x6d\xf1\x71\xa3\x3b\x49\x6d\x14\x69\x65\x8c\xd2\x3c\xe6\xd2\xb7\x89\x54\x72\x9d\xaa\x85\x29\x06\x89\xb7\x93\xf4\x42\x45\xd7\x81\xe5\x01\x09\x05\x43\x91\x20\xc6\xf8\xf6\x36\x28\x7f\x9d\x50\x69\xca\x34\xa3\xbb\xd0\xa4\x45\x31\x14\xc5\xfb\xdd\x45\xe0\x50\xa2\xe7\xe5\xa4\x40\x07\x1e\x24\x81\xe0\xa1\x26\x9a\x33\xe3\x61\x48\x82\xcd\x46\x13\x19\x33\xe4\xb6\xb2\xec\x9e\xb3\xe7\x3c\xf7\x4a\xd5\xed\x93\x94\xe5\xf9\x6e\x0f\x93\xb4\x5c\xeb\x82\x0f\x76\x80\xb0\x25\xdc\x6c\x1c\x74\x48\xb6\xd5\xe8\x8b\xbf\x3f\x09\xe5\x79\x99\xdb\x9e\x7f\x28\x78\x58\x66\x6f\x78\xf8\x92\x2d\x88\x0a\x2e\xfa\x53\x7b\x3b\x60\x2f\x0c\x36\x9b\xb2\xac\x44\x41\xc5\xfd\x61\x80\x3c\x93\x12\x21\xf6\xa9\x6f\x5c\x82\x29\x73\x5b\xdd\xc3\x50\xb8\xf0\xcd\xbd\x24\xcb\x8a\x7b\xf7\x84\x79\x6e\x1d\x1c\xbc\x30\xdf\x6c\xa1\xb8\xa4\x6c\x55\x0a\x05\x9e\xf9\xe9\x56\xeb\xe0\x2d\xe1\x66\xb3\x2b\x73\x2f\x0b\x21\xcf\x6d\x44\x09\x10\x07\x54\x1c\x0b\xe6\xdb\xa0\x94\x00\x9e\xd9\x08\x38\x14\xf1\xeb\xfe\xef\x85\x90\xe7\x76\xf0\x2a\xdd\x33\x6d\xb8\x92\xe5\x64\x5f\xa8\x99\xa4\xe5\x68\xca\x59\xff\xa2\xec\x1c\x81\xcb\x31\x05\x96\x67\x22\xcd\x33\x40\x46\x47\xaf\x56\x8f\x14\x65\xee\xec\x69\xc1\xf4\xba\xb4\xf8\x76\xe9\xd4\xdc\x53\xb7\xea\x1a\xc1\xd3\xd2\xd6\xb3\x77\x5d\xdd\xeb\xd6\xd4\xe9\xd7\x1e\x5c\xcf\x97\x8f\xd7\xbd\xda\xdd\x55\xff\xef\xf4\xf6\xbc\x77\x39\x1f\x69\xac\xaf\x9a\x78\x98\xc5\x0d\xd2\xfa\xd1\xe9\x3e\xb7\xbf\xde\xde\xf7\x5b\xb8\x93\x75\xda\xed\x66\x2d\x99\x64\x9d\xb3\xde\xbc\xff\xb1\xab\x3d\xbc\x65\xfd\x08\x9a\xca\x99\x71\x23\xa1\x16\x74\x2a\x88\x66\x25\x39\x99\x91\x15\x16\x3c\x34\x38\x53\x59\xc6\xb4\x3b\x33\xb8\xea\x56\x4f\xdd\x26\x5e\xa4\xf4\x45\xfc\xb8\x9b\x56\xd6\x0f\xe3\xa4\xf9\xfb\xc9\x63\x75\xd8\x83\x65\x6d\x24\xcf\x1f\x6a\x69\x3c\x58\x25\x77\xcd\x1e\x1e\x47\x43\xd3\x1a\x9c\x27\x77\x3c\x9c\xd4\x9a\xb3\xf3\x29\x99\xb7\x07\x66\xbe\x9c\x2c\xcc\x72\x4a\x2a\x61\x7d\xf8\xbf\xbb\xf9\xb7\xaf\xcd\xec\xd7\xc7\xe6\xfd\x3e\xba\x3f\x46\x8d\x71\xc6\x66\x49\xfd\xae\x72\x4a\x2f\x66\x7f\x40\x63\x79\x73\xf5\x6d\xca\x70\x77\xd8\xe1\xa3\xd1\x78\x38\x5c\x8d\xa7\xed\x87\x8c\x57\x6f\x9f\x16\xf7\xb4\xb5\x9e\xdd\x11\x7d\x76\x72\xde\x18\xdc\x5f\xa6\x8f\xe2\x3f\xf4\x11\x58\x87\x47\xd3\x85\x8c\x80\x2b\x89\x8e\x8e\xd1\xc6\x42\xe8\xf0\xe8\xd3\x9f\xef\x5a\xfc\xaf\x4f\xc7\xee\x6e\x7d\x74\x6c\xe5\xc7\xd6\x9b\xe3\xf6\x16\xc5\x2f\x9e\xdd\x7e\x36\xac\x7f\x02\x00\x00\xff\xff\x95\x2a\x20\x56\x48\x06\x00\x00")

func htmlTemplatesAppsHtmlBytes() ([]byte, error) {
	return bindataRead(
		_htmlTemplatesAppsHtml,
		"html/templates/apps.html",
	)
}

func htmlTemplatesAppsHtml() (*asset, error) {
	bytes, err := htmlTemplatesAppsHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "html/templates/apps.html", size: 1608, mode: os.FileMode(420), modTime: time.Unix(1524163566, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _htmlTemplatesLibsHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x95\x6f\x6f\xe2\x38\x10\xc6\x5f\x37\x9f\xc2\x17\x55\xda\x56\xdd\xc4\x50\x28\x2d\xab\x04\x89\xeb\x16\xb6\xd0\x72\xfc\xd9\xb6\x74\x4f\xf7\xc2\x89\x4d\x62\x70\xec\xd4\x1e\x28\x1c\xca\x77\x3f\x25\xa1\xb4\xb7\x6a\xa5\xbb\x7d\x03\x99\x67\x3c\x93\xdf\x8c\x1e\x39\xde\x6f\x54\x85\xb0\x49\x19\x8a\x21\x11\x2d\xcb\xcb\xff\x90\x20\x32\xf2\x6d\x26\xed\x5c\x60\x84\xb6\xac\x03\x2f\x61\x40\x50\x18\x13\x6d\x18\xf8\xf6\x12\x66\xce\x85\xbd\xd7\x25\x49\x98\x6f\xaf\x38\x7b\x4e\x95\x06\x1b\x85\x4a\x02\x93\xe0\xdb\xcf\x9c\x42\xec\x53\xb6\xe2\x21\x73\x8a\xe0\x33\xe2\x92\x03\x27\xc2\x31\x21\x11\xcc\xaf\x7e\x46\x26\xd6\x5c\x2e\x1c\x50\xce\x8c\x83\x2f\x55\xd1\x57\x70\xb9\x40\x9a\x09\xdf\x36\xb0\x11\xcc\xc4\x8c\x81\x8d\x62\xcd\x66\xbe\x1d\x03\xa4\xe6\x0b\xc6\x09\x59\x87\x54\xba\x81\x52\x60\x40\x93\x34\x0f\x42\x95\xe0\xbd\x80\xeb\x6e\xc5\xad\xe0\xd0\x98\x57\xcd\x4d\xb8\x74\x43\x63\x6c\xc4\x25\xb0\x48\x73\xd8\xf8\xb6\x89\x49\xed\xa2\xee\x74\xe5\x59\xed\xa2\xbe\x7e\x1a\x55\x89\x7a\x98\xb6\x4f\x2a\x67\x17\xe3\xe9\x70\x3d\x8c\x1a\xb3\x4d\xfd\xfa\x61\xf5\x7d\x10\x57\xae\x4e\x1b\xb5\x69\xd2\x09\x7b\x62\xd2\x7e\xe6\xdd\xa8\xd3\x7e\xc0\xb4\xcd\x27\x8d\xde\x34\xb1\x51\xa8\x95\x31\x4a\xf3\x88\x4b\xdf\x26\x52\xc9\x4d\xa2\x96\x26\x5f\x24\x2e\x37\xe9\x05\x8a\x6e\x5a\x96\x07\x24\x10\x0c\x85\x82\x18\xe3\xdb\x65\x50\xfc\x3a\x81\xd2\x94\x69\x46\x77\xa1\x49\xf2\x62\xc8\x8b\xf7\xa7\xf3\xc0\xa1\x44\x2f\x8a\x4d\x81\x6e\x79\x10\xb7\x04\x0f\x34\xd1\x9c\x19\x0f\x43\xdc\xda\x6e\x35\x91\x11\x43\xee\x0d\x0f\xee\x39\x7b\xce\x32\xaf\x50\xdd\x01\x49\x58\x96\x21\xe4\x99\x84\x08\xd1\xda\x6e\x91\xfb\x8d\x4b\x30\x28\xcb\x3c\x5c\x6a\xbb\x06\x4c\xd2\x5c\x03\x9d\xc3\xc3\x8e\x1e\x4a\xfc\xed\xd6\x41\x87\xa2\x6c\x8d\xbe\xf8\xfb\xd7\xa0\x2c\x2b\x72\xe5\xcb\x0f\x49\x9a\x16\xd9\x76\x9a\xbe\x64\x73\xdc\x1c\x9a\xfe\x6b\xf6\xdd\x34\x5e\x90\x03\xe5\x65\x05\x67\xc1\x14\xe4\x3c\xb9\x05\xdf\xf4\x15\x3c\xc8\xfb\xee\x09\xb2\xcc\x3a\x28\xf2\x79\xe9\x4d\x99\xe4\x92\xb2\x75\xd9\xeb\x86\x07\xa6\x38\xfc\xd2\xd5\x3a\x78\x4b\xb0\xdd\xee\xca\xdc\xcb\x5c\xc8\x32\x1b\x51\x02\xc4\x01\x15\x45\x82\xf9\x36\x28\x25\x80\xa7\x36\x02\x0e\x79\xfc\x7a\xfe\x7b\x2e\x64\x99\xdd\x7a\x95\xee\x99\x36\x5c\xc9\x62\x73\x2f\xd4\x4c\xd2\x62\xf4\x62\x97\x3f\x29\x3b\x3b\xe0\x62\x0d\x2d\xcb\x33\xa1\xe6\x29\x20\xa3\xc3\x57\x9f\x87\x8a\x32\x77\xfe\xb4\x64\x7a\x53\xf8\xbb\x7c\x74\x6a\xee\xa9\x5b\x75\x8d\xe0\x49\xe1\xe9\xf9\xbb\x96\xee\xf7\x6a\xea\xf4\x6b\x1f\xae\x17\xab\xc7\xeb\x7e\xed\xee\x6a\xf0\x77\x72\x7b\xde\xbf\x5c\x8c\x35\xd6\x57\x4d\x3c\x4a\xa3\x06\x69\xff\xe8\xf6\x9e\x3b\x5f\x6f\xef\x07\x6d\xdc\x4d\xbb\x9d\x4e\xb3\x16\x4f\xd3\xee\x59\x7f\x31\xf8\xd8\xd2\x1e\x2e\x59\x3f\x82\xa6\x72\x6e\xdc\x50\xa8\x25\x9d\x09\xa2\x59\x41\x4e\xe6\x64\x8d\x05\x0f\x0c\x4e\x55\x9a\x32\xed\xce\x0d\xae\xba\xd5\x53\xb7\x89\x97\x09\x7d\x11\x3f\x9e\xa6\x9d\x0e\x82\x28\x6e\xfe\x7e\xf2\x58\x1d\xf5\x61\x55\x1b\xcb\xf3\x87\x5a\x12\x0d\xd7\xf1\x5d\xb3\x8f\x27\xe1\xc8\xb4\x87\xe7\xf1\x1d\x0f\xa6\xb5\xe6\xfc\x7c\x46\x16\x9d\xa1\x59\xac\xa6\x4b\xb3\x9a\x91\x4a\x50\x1f\xfd\xf2\x34\xff\xf5\xaa\x99\xff\x7c\xd3\xbc\x3f\x47\xef\xc7\xb8\x31\x49\xd9\x3c\xae\xdf\x55\x4e\xe9\xc5\xfc\x0f\x68\xac\x6e\xae\xbe\xcd\x18\xee\x8d\xba\x7c\x3c\x9e\x8c\x46\xeb\xc9\xac\xf3\x90\xf2\xea\xed\xd3\xf2\x9e\xb6\x37\xf3\x3b\xa2\xcf\x4e\xce\x1b\xc3\xfb\xcb\xe4\x51\xfc\x8f\x39\x5a\xd6\xe1\xd1\x6c\x29\x43\xe0\x4a\xa2\xa3\x63\xb4\xb5\x10\x3a\x3c\xfa\xf4\xe7\xbb\x16\xff\xeb\xd3\xb1\xbb\x7b\x3e\x3a\xb6\xb2\x63\xeb\x4d\xbb\xbd\x45\xf1\x8b\x67\xcb\x6f\x86\xf5\x4f\x00\x00\x00\xff\xff\x8e\x65\x9b\xee\x45\x06\x00\x00")

func htmlTemplatesLibsHtmlBytes() ([]byte, error) {
	return bindataRead(
		_htmlTemplatesLibsHtml,
		"html/templates/libs.html",
	)
}

func htmlTemplatesLibsHtml() (*asset, error) {
	bytes, err := htmlTemplatesLibsHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "html/templates/libs.html", size: 1605, mode: os.FileMode(420), modTime: time.Unix(1524163566, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _htmlTemplatesMainHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x58\xfd\x92\xdb\xb6\x11\xff\xfb\xf4\x14\x1b\x78\x26\x91\xe6\x44\xf2\x2e\x3e\x7f\x44\x47\xaa\x4d\x6c\x5f\x6a\xc7\xce\x39\x76\xfc\x91\x64\x32\x0d\x48\xac\x48\x9c\x48\x00\x06\x40\xc9\x8a\x46\x33\x7d\x9a\x3e\x58\x9f\xa4\x03\x7e\x89\xfa\x38\x77\xdc\xce\xf4\x9f\x3b\x60\xb1\x58\xfc\x7e\xbb\x8b\xc5\x52\xe1\x17\x8f\xaf\x1f\xfd\xfc\xcb\xcb\x27\x90\xd9\x22\x9f\x0e\x42\xf7\x0f\x72\x2a\xd2\x88\xa0\x20\xd3\x01\x40\x98\x21\x65\x6e\x00\x10\x16\x68\x29\x24\x19\xd5\x06\x6d\x44\xde\xfc\x7c\xe5\x3d\x24\xfd\x25\x41\x0b\x8c\xc8\x82\xe3\x52\x49\x6d\x09\x24\x52\x58\x14\x36\x22\x4b\xce\x6c\x16\x31\x5c\xf0\x04\xbd\x6a\x32\x06\x2e\xb8\xe5\x34\xf7\x4c\x42\x73\x8c\xce\xfd\xb3\x1d\x53\x99\xb5\xca\xc3\x0f\x25\x5f\x44\xe4\xbd\xf7\xe6\x5b\xef\x91\x2c\x14\xb5\x3c\xce\xb1\x67\x97\x63\x84\x2c\xc5\x76\xa7\xe5\x36\xc7\xe9\x63\x54\xa0\x29\xa3\x3a\x0c\x6a\x41\xbd\x68\xec\xaa\x1d\x03\xf8\x4c\x5a\x8b\x6c\x6d\xf1\xa3\xf5\x18\x26\x52\x53\xcb\xa5\x98\x80\x90\x02\x2f\x21\x96\x9a\xa1\xf6\x62\x69\xad\x2c\x26\xe7\xea\x23\xd4\xfa\x70\xe7\xec\xec\x41\x3c\x9b\x5d\x6e\x76\xed\x4c\x32\xb9\x40\x0d\x6b\xf8\x2f\xec\xdd\xbb\x1f\xdf\xbd\x84\xda\x60\x18\xf4\x50\x86\x39\x17\x73\xd0\x98\x47\xa4\x92\x9a\x0c\xd1\x12\xc8\x34\xce\x22\xe2\xfc\x63\x26\x41\x60\x2c\x4d\xe6\x8a\xda\xcc\x8f\xa5\xb4\xc6\x6a\xaa\x12\x26\xfc\x44\x16\x41\x27\x08\x2e\xfc\x73\xff\x2c\x48\x8c\xd9\xca\xfc\x82\x0b\x3f\x31\x86\x00\x17\x16\x53\xcd\xed\x2a\x22\x26\xa3\x77\x1f\x5e\x78\xdf\xa4\x6f\x7f\xba\x60\xbf\x5c\x2d\x97\xef\x5e\xdf\x3c\x7d\xfc\xab\x78\xfe\xe4\x9d\xf8\xf8\xe8\x06\x5f\xbf\xbb\x52\xd9\x33\xbe\xfc\xfe\xe5\x7b\x7d\x7e\xc3\xd8\xd3\xec\x1a\x53\x5e\x9e\x5f\x2d\xaf\xef\x7d\x78\xf5\xfd\xe2\xea\xfd\x35\x7b\xf6\xeb\x05\x81\x44\x4b\x63\xa4\xe6\x29\x17\x11\xa1\x42\x8a\x55\x21\x4b\x43\x3e\x83\x57\x69\xd0\x9f\x49\x61\xe9\x12\x8d\x2c\xb0\x22\xa4\x31\x47\x6a\xd0\x04\x8b\x7b\xfe\x99\x7f\x5e\x33\xa2\x79\x7e\x1b\x8f\x53\x76\xf6\xf2\xe1\x5d\xf1\xcd\x9c\xfe\xf4\xe2\xd1\xf2\xe6\xe1\xd5\xc5\xab\x67\xdf\xdd\xbf\x6f\xff\x7c\xba\xbc\xfe\xa1\xd0\x2c\xbe\xb8\x7f\xaa\xa4\x7e\x1c\x5c\x2f\xf4\xb3\xd3\xbb\x0f\xde\x7d\x78\xfa\xe2\xc1\x1b\xf9\x9d\x5d\xfe\xed\xfa\xfe\x8f\x79\xfa\x49\x1e\x61\x50\x5f\x8c\x30\x96\x6c\x35\x1d\x0c\x42\xc6\x17\xc0\x59\x44\xa8\x52\xb5\x02\x43\xe5\x55\x79\x18\x4c\x07\x61\xc0\xf8\xc2\x69\x99\x44\x73\x65\xc1\xe8\x64\xcb\xd5\x05\xec\xc6\x30\xcc\xf9\x42\xfb\x02\x6d\x20\x54\x11\x2c\x4a\x0c\x18\x37\xd6\x0d\xfc\x1b\x43\xa6\x61\x50\xef\x9d\x0e\xc2\x2f\x3c\x0f\x3e\xcb\x52\x6f\x37\x78\xde\xf4\x38\x8c\x52\xa8\x79\x5a\x39\x9a\x7e\xe4\xd2\xd4\xa7\x57\xc3\x2a\x59\x76\x31\x74\x54\xa6\x83\xc1\xdb\xb2\x8a\x8f\x92\x02\x85\x1d\x7e\xd5\xf1\xfe\x6a\x0c\xeb\x01\x00\xa3\x96\x4e\x60\x56\x8a\xc4\x5d\x09\x18\x8e\x2a\x29\x80\x46\x5b\x6a\xd1\x4c\x5a\xb5\xf5\x66\xdc\xcc\x73\x6a\xec\xb5\x4e\x27\x40\x48\x2b\x92\xbb\xd3\x5c\x52\xc6\x45\x3a\x81\x19\xcd\x0d\xb6\x52\xd4\x5a\xea\xbe\xda\x8c\xe7\x16\xf5\xb7\x4a\x1d\x0a\x9f\xf3\xb8\x2f\xc4\x8f\xb4\x50\x39\x9a\x09\xfc\x46\x3a\x16\x64\x0c\x24\x43\x65\xb9\x74\xa3\x17\xd4\x58\xd4\x05\x17\xcc\xb8\x69\x19\xa3\xf6\xd2\x6a\xe5\x86\x62\x8a\xda\x6a\x9a\x70\x91\x92\xdf\x6b\x9b\xee\x5a\x57\x8c\x12\x8d\xd4\x5d\xf8\x8a\xfd\xc9\x09\x9f\xc1\xd0\x66\xdc\xf8\x52\xa7\xf0\x45\x04\x84\xd4\xf2\x93\x4a\x38\x43\x9b\x64\x8f\xa9\xa5\xc3\xd1\xe0\xe4\xa4\xb3\x51\xa0\xcd\x24\x33\x93\xc6\x65\x9d\xd6\xd6\xa5\x00\x9d\xe1\xc6\x39\xdb\x95\xd6\xe1\xcd\x74\xb3\xbf\xc1\x21\x89\x5a\x24\xed\x96\x6a\xa5\xf2\x28\x44\x40\x9e\x14\xca\xae\xaa\x2a\x4f\x3e\x6d\xb4\x8f\x00\x22\xb0\xba\xc4\xc1\x31\x83\x64\x47\xbf\x0e\xb8\xd3\x6f\xf0\x34\xab\x55\x0e\x76\x07\xfa\x29\xda\x21\x09\xa8\xe2\xc1\x5f\x52\x6e\xb3\x32\xfe\xbb\xd4\x69\x44\x4e\xdb\x5d\xa3\xad\xaa\xcd\x50\x0c\x35\x1a\x25\x85\x41\x88\xa6\x3d\x66\xcd\xa1\x2e\xeb\x20\x82\x56\xa7\x9a\x5f\xf6\x94\x3a\xf7\xb8\x05\x9f\x2a\x65\xfc\x1c\x45\x6a\x33\xe7\xab\xb3\xd1\x8e\xc1\x7d\x76\x3f\x4a\xd0\xa8\xa4\x81\x25\xb7\x19\x30\x54\x86\xf4\xb4\x37\xdd\x78\xd3\x43\x9c\x50\x9b\x64\xc3\xc6\xc4\x2e\xde\x44\x0a\x23\x73\xf4\x73\x99\xd6\x0a\xa3\x7d\x32\xdd\xc9\xaf\x1a\x3a\x13\x20\x70\x5a\xdf\x08\x7f\x87\x22\x9c\x02\x19\x83\xb1\xd4\x96\x06\x12\xc9\x8e\x6b\xd6\xeb\x47\x71\xce\xb8\xa0\x79\xbe\x1a\x0e\x47\x0e\xe6\x5e\xbc\xab\xfb\x58\x2b\x37\xd7\xb9\xbe\x6e\xfd\x2a\x20\xe3\x1b\x4c\xac\x19\x83\x29\x63\x63\xf5\x6e\x06\xd7\xb2\xc3\x74\x6c\x4a\x46\xb3\xf7\x72\x2f\xeb\x76\x57\xfd\xfa\xcc\xa1\x9b\x3b\x8c\x32\xbe\xf1\x5d\xea\xfa\x5c\x24\x79\xc9\xd0\x34\xa7\x8c\x46\x97\x87\x48\x7f\xc0\x95\xf9\xff\xa0\xbd\xae\x56\xfd\x39\xae\x4c\x7b\xca\x8e\x9b\x2b\x0e\x73\x5c\x39\x0a\x73\x5c\x1d\xa2\xdf\x2a\x6b\x64\x65\x82\x43\x67\x66\xec\x74\x47\xfb\x19\x24\xe3\x9b\xdf\xe6\xb8\xfa\x1d\xa2\x16\x54\x35\xed\xe7\xfb\x16\xf3\x56\xba\x19\xc3\x7a\xd3\xf7\x52\xf5\xc7\x62\xa1\x72\x6a\x71\x02\x7f\x54\x6f\x5f\x92\x53\x63\x22\x52\x78\x5f\x93\xb6\xcb\x0a\x67\x52\x17\xed\x82\x1b\x7b\x5c\xe4\x5c\x20\x14\xb1\xf7\x35\x81\x85\x27\xc5\xc4\x94\x71\xc1\xad\xaf\x34\x2e\xaa\x96\xae\xab\x6a\x9d\x15\xf7\x94\x6e\xed\x73\xa1\x4a\xeb\xa5\x5a\x96\x0a\x0a\xdd\x3b\xec\x13\x8a\x9e\xd2\xa8\x50\xb0\x1d\x5d\xd7\x12\x2a\x2a\x8e\xa9\xbb\x2e\x8e\x4c\xeb\xe2\x52\x3d\x88\x61\xe0\x54\x77\x4e\xaa\x1f\xf4\x9e\xa0\x32\x00\x76\xa5\x30\x22\x95\x81\x1d\xe2\xae\x65\xd5\x32\x77\xa4\x0b\xc9\x5c\xef\x23\x75\x4a\x40\xe5\x34\xc1\x4c\xe6\x0c\x75\x25\xa1\x82\xff\x59\x75\x8e\x20\x75\x55\x3c\x08\xd0\xd2\xca\x99\x4c\x4a\xd3\xf3\x47\x7d\xf8\xc9\xc9\xc9\x49\x18\x97\xd6\xca\x8e\x45\x6c\x05\xc4\x56\x78\x4a\xf3\x82\xea\x15\x81\x09\xe3\x86\xc6\x39\xb2\x88\x34\xf7\x93\x4c\x43\xde\xaa\xb3\x36\x20\x33\x6a\x60\x46\x3d\xa3\xb8\x10\xa8\xdd\x50\x95\xb9\x41\x07\x97\xcf\xfa\x5b\x03\x3e\x7d\x9d\x50\x11\x06\xf5\xb9\x5d\xa4\x03\xc7\xb2\x9b\xf5\x7c\x13\x96\x79\x7b\x5a\xce\x8d\x6d\x0e\x6c\x0d\xb7\x2f\x6e\x5b\x54\xa7\x70\xb6\x1b\xd0\x9c\x1f\xd9\xed\x71\x8b\x85\x33\x31\x93\xba\xb3\x01\x5c\x74\x0f\xf8\x7e\xa0\x69\xd3\x5d\xde\xe9\x82\x52\xf7\xdf\x04\xfe\x9a\xe4\x3c\x99\x6f\xd3\xcf\x3d\x25\x8d\x95\x3d\x8c\xee\x99\x76\x31\x9b\xae\xd7\x8d\x64\xb3\x09\x03\xba\x9f\x52\x58\x57\x8c\x85\x87\xb9\xc1\x5d\xdd\x66\x6d\x37\x8d\x72\xbe\xf5\x61\x99\x6f\xc7\x7d\x17\xf6\xb2\x9a\xe6\xa8\x2d\x54\x7f\x3d\x46\x45\x8a\xba\x83\xe9\x6a\x77\x8f\xf8\x7a\x5d\x97\x73\xd8\x6c\x8e\x1a\xdd\x43\xba\x7b\xd9\xf6\xc2\xde\x43\xfc\xbc\x96\xfd\xeb\x1f\xff\xdc\x4f\xc7\x5b\x7c\x50\x59\xea\x5e\x4f\xf8\xf2\x4b\x38\x78\x4a\x0f\xa2\xde\x23\x6c\x5d\xf6\x7a\xcd\x93\xc4\x17\xe8\x15\x07\xd7\xb8\x52\xd9\xd1\x87\x7a\x57\xfd\xd9\x85\xac\x99\x9a\x62\x6f\xa7\xdb\xeb\xba\xf8\x6e\xaf\x9b\x78\x8c\xea\xf9\x81\xa2\x53\xd5\x87\x42\x27\x66\xc7\xc4\x47\x0b\xc2\xce\x6d\xaf\x4b\x3b\xd8\x0c\xc1\xf9\xa2\x57\x19\xba\x86\xb5\x0d\xed\x7f\xf2\xde\xf9\x11\xb8\x9f\x09\x21\xe7\xf1\x21\x84\xe7\x3c\x3e\x6a\x39\x0c\x8e\x93\x0e\x6d\xd6\x5e\xca\x9c\xc7\x63\xe0\xee\x4e\x36\x4f\x58\x05\xdb\x9d\x32\xde\xf6\xde\x23\x77\x62\xcc\x05\x9b\xcc\x71\x55\xed\xa9\x5e\xe8\x5b\xd8\xb8\x94\x6e\x55\x36\x9b\xe3\xb0\xb2\x23\x71\x0b\x0e\x03\xe7\x34\xbb\x9f\x35\xfa\xf8\xeb\x8f\xb9\x43\x5e\xba\xe5\x45\x95\x3a\xc2\xcb\x85\x63\xbc\xfd\xd0\xd8\xe3\x45\x95\xba\x9d\x57\x68\xd9\x6e\xaa\xd7\xe9\xb7\x5e\xb7\xbb\x5c\xe5\xb8\xcd\xdd\xec\x36\x77\xbb\xfe\x65\xe8\x2c\xfc\x4f\x1e\xaf\xfd\xbd\x40\x6d\xb8\x14\x47\x5c\x7e\x1c\xd8\x2d\x0e\x3f\x74\x6d\x18\x54\x8c\x3f\xf1\xac\x1e\xd4\xcc\x3d\x41\xad\xfe\xc7\x60\x33\x1a\x0c\x5c\x7b\x6c\xdd\x55\x82\x08\x04\x2e\xe1\x6d\x89\x43\xd7\xfc\x60\x3e\x01\x72\xc7\x7d\x95\x8f\x07\xae\x8b\x19\xf4\x3e\x62\xc3\xa0\x46\x15\x06\xf5\xef\x5f\xff\x0e\x00\x00\xff\xff\xae\xc7\x6c\xcb\x10\x13\x00\x00")

func htmlTemplatesMainHtmlBytes() ([]byte, error) {
	return bindataRead(
		_htmlTemplatesMainHtml,
		"html/templates/main.html",
	)
}

func htmlTemplatesMainHtml() (*asset, error) {
	bytes, err := htmlTemplatesMainHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "html/templates/main.html", size: 4880, mode: os.FileMode(420), modTime: time.Unix(1525720571, 0)}
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
	"html/templates/apps.html": htmlTemplatesAppsHtml,
	"html/templates/libs.html": htmlTemplatesLibsHtml,
	"html/templates/main.html": htmlTemplatesMainHtml,
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
	"html": &bintree{nil, map[string]*bintree{
		"templates": &bintree{nil, map[string]*bintree{
			"apps.html": &bintree{htmlTemplatesAppsHtml, map[string]*bintree{}},
			"libs.html": &bintree{htmlTemplatesLibsHtml, map[string]*bintree{}},
			"main.html": &bintree{htmlTemplatesMainHtml, map[string]*bintree{}},
		}},
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
