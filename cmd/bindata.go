// Code generated by go-bindata.
// sources:
// data/bash
// data/fish
// data/sh
// data/zsh
// DO NOT EDIT!

package cmd

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

var _dataBash = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x90\xc1\x6a\xc3\x30\x10\x44\xef\xfb\x15\xdb\xad\x0e\xc9\x21\xe4\x03\x82\x0f\xa6\x2e\xc1\x10\xc7\xc2\xa4\xc9\xa1\x94\x45\x49\x95\xd8\x20\x5b\x41\x51\x74\x29\xfd\xf7\x62\xb9\xd0\xd6\x18\xaa\x8b\x24\x56\x33\xf3\x46\x8f\x0f\xcb\x63\xd3\x2d\x8f\xea\x56\x03\x30\x87\x0b\x3b\x6d\xac\x7a\x9f\xcd\xf1\x03\x10\x11\x65\x55\x16\x72\xc7\x4f\x65\x51\xa4\xdb\x2c\x11\xbc\xcf\xab\xdd\x4b\xba\x59\x97\x5c\x56\xf9\x3a\xdf\xa6\x1b\xfe\xfb\x26\xca\x74\x50\x06\x49\xcc\xc2\x65\x38\x2e\x16\xb7\x5a\x1b\x83\x7d\xd0\x9c\xe0\x13\x80\x65\xf5\xbc\xff\xe5\x26\x0f\x59\x42\xf4\x0d\x71\x75\xb6\xbd\x7a\x3e\xd9\xb6\x55\xdd\x0f\x4c\x73\xc6\x57\x24\x21\x0f\x19\x61\x82\x24\x26\x3c\x08\xdf\x56\xe8\x6b\xdd\x45\x41\xbf\x9c\xf6\x77\x37\x5c\xcf\x4d\xdc\x26\xa3\xa3\x2b\x0c\xf3\x1e\x41\xdd\xbd\x65\x75\xf2\x4d\x50\x5e\x47\xe0\x7f\x9b\x27\x62\xf4\x13\xa3\x31\x4d\x74\xc3\x15\x8e\x54\x04\x5f\x01\x00\x00\xff\xff\xac\x8d\x83\x2a\x95\x01\x00\x00")

func dataBashBytes() ([]byte, error) {
	return bindataRead(
		_dataBash,
		"data/bash",
	)
}

func dataBash() (*asset, error) {
	bytes, err := dataBashBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/bash", size: 405, mode: os.FileMode(420), modTime: time.Unix(1496826836, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dataFish = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x59\x7b\x6f\xdb\x38\x12\xff\xdf\x9f\x62\xea\x06\xb0\xbd\x80\xe2\x6b\xee\xbf\x5d\xe4\x80\x74\x13\xb4\x01\xda\x26\x48\xdb\x5d\x1c\x8a\xae\x40\x4b\x63\x99\x08\x45\x6a\x49\xca\x89\x17\xf9\xf0\x07\x92\x7a\x90\x12\xdd\x07\xaa\x9e\x81\x45\xb6\xe4\xbc\x38\x8f\xdf\x0c\xa9\xe7\xcf\xd6\x1b\xca\xd7\x5b\xaa\x76\xeb\xd9\x6c\x5b\xf3\x4c\x53\xc1\x61\x5f\xcc\x00\x00\xe8\x16\xb8\xd0\x90\x89\x9a\x6b\x38\x21\xb2\xd8\xc3\x7f\x60\x9d\xe3\x7e\xcd\x6b\xc6\x2c\x89\xf9\x65\xa2\x2c\x09\xcf\x5b\x2e\xf3\x93\xa8\x6b\xc9\xed\x3f\x91\xe7\xf6\xef\x73\xc0\x6c\x27\x9c\x98\x4f\x2f\x3e\xdb\x35\xf5\x40\x75\xb6\x0b\xd7\xac\x44\xa2\x10\x28\xa7\xba\x5b\x31\xbf\x74\x5f\xa4\x66\xb1\xa5\x3f\x3d\x4d\x86\x3c\x24\xd3\x74\x4f\x34\x8e\xf8\xda\x8d\x2f\xf0\xe6\x78\x94\x3b\xb2\x65\x59\xb2\xbc\x22\xd9\x3d\x29\x50\x8d\x58\x22\x5b\x96\xa5\x60\x62\x43\xd8\xd5\x23\x66\x23\x96\x7e\xeb\x0b\x56\xd6\x55\x21\x49\x3e\x36\x71\xb8\x6e\x89\x17\xbf\x2c\x02\xc2\x3e\x52\x4e\x43\x17\x20\xf3\x9f\x17\xff\xd6\xd3\xb3\xb8\xfb\xec\xf2\xbe\x00\x46\xf9\xbd\x63\x0d\x38\x07\x47\xa7\x5b\xf8\x04\xc9\x3f\x30\x3f\xf9\xe3\xfa\xee\xc3\xc7\x8b\x37\xaf\x6e\xe6\xd0\x9f\xca\x66\xc5\x95\x94\x42\xfe\x0a\xef\x04\xec\xa9\xd4\x35\x61\x85\x80\x07\x21\xef\x55\x45\x32\x04\xaa\x5c\x60\x11\x88\x06\xbd\x43\x28\x45\x89\xbc\x4f\x8e\x7d\x01\x3b\x64\x55\x2c\x1c\x2e\x11\xe1\x45\x90\x8a\x59\x0e\xbd\x2d\xe9\xed\xc5\x87\xd7\x6b\x25\xb3\x75\xe4\x28\x41\xd8\xdd\x41\xf8\xb1\x83\x44\x12\xc5\x0a\xec\x58\x97\x5e\x21\xad\x20\x61\x1a\xce\x3c\x76\x85\x1a\x6c\xd4\xcf\x3e\xc3\x72\x43\x14\x72\x52\x22\x9c\xdc\xfe\x79\xb9\x0a\x65\x19\xc2\xde\x37\xcb\x7d\x61\x2b\xe5\x3d\x6a\x4d\x79\xa1\x9a\xd4\x39\xb3\xa9\xb3\xea\x39\x92\xe2\x11\x3a\xb3\x5b\x22\xa7\xdd\x46\xe0\xc2\xd9\x4d\x79\xe1\xf9\x26\xce\x6d\x3d\x06\x27\xaf\x6f\xde\x5e\xad\x4f\xbb\x78\xad\x3d\xb6\x80\x2f\xed\x19\x6f\xde\x5c\xbe\xba\x71\xdc\xee\xef\xc0\xaf\xaf\x6e\x5e\x5e\xbf\x9b\x0f\x9c\x12\x93\xf1\xf2\xfa\x1d\x38\x6a\x77\x02\xa6\xf0\xdb\x78\x62\xc6\x18\x08\x1c\xbb\x78\x2c\xc2\x19\x6e\xcd\xee\xa8\x5c\xcd\xa6\x5b\xc2\xd8\x86\x64\xf7\x36\x1c\x7f\xd7\x28\x0f\x6d\x3c\x56\x5e\x02\x9c\x0c\xa9\xcf\x41\xcb\x1a\x23\xe7\x6d\xbd\x14\xfa\xfc\x57\xdf\x6b\xd1\x33\xc7\xf9\xe2\x87\x6b\x3c\x32\x28\x04\xe3\x0b\x4b\x56\xde\xe7\x54\x42\x52\xc5\x4b\x25\xa4\x70\x91\x08\xa4\xb7\x51\xb6\x3a\x7a\x9f\x45\xc0\x20\xbd\xbc\x7e\x7f\xf1\xf2\xcd\x55\x7a\x7b\x77\xf3\xf6\xf6\x83\x1f\xfe\xb6\x16\x15\x24\x19\x98\x26\x95\x56\x52\x94\x95\x86\x54\xb0\x3c\xf5\x16\x46\x1c\x10\xdb\xec\x92\xdd\x24\xdb\xb2\xb7\x60\x05\xf3\x10\x4f\x8f\x49\x6f\x01\xa4\x83\xcc\x00\x2c\x06\xc5\xff\xff\xc3\xbd\x48\x7b\x1a\xe1\x5e\x5f\xe9\x97\x2d\xf5\xa0\xd6\xfb\xf0\x75\xb9\x11\x29\xa1\x90\xca\x45\x39\x56\x54\x3e\x5d\x8c\x6a\x82\x94\xc0\x68\x90\x83\x9c\x19\x46\xf2\x6b\x1c\x18\x8f\x7d\x58\x3b\x08\x11\x80\x44\x88\xd4\x5c\xb3\x13\x39\xf9\xd1\xdd\x81\x93\xc7\xdb\x96\x7d\x9c\x7c\x7e\xff\xb7\x9c\x0c\xcc\x49\xfa\xb4\x1a\xa0\xfa\x0f\x34\x33\x9b\x45\x1f\xab\xdc\x65\x50\xa7\xc1\xee\x15\x02\x0a\xa3\xbd\x86\x82\xea\x5d\xbd\x39\xcd\x44\xb9\x7e\x85\xfa\xbd\x96\x48\xca\xf5\xbe\xf0\x44\xdc\x21\x13\x24\x37\x32\x4c\x6e\xab\x1d\x32\x06\x15\x91\x1a\xc4\xf6\xbb\xa5\x36\xe5\x80\x7b\xc2\x20\x49\x9c\x2c\x13\x45\x78\x02\x25\x6a\x99\xe1\x6c\x70\xea\xc0\x39\x73\x48\x48\xe8\x89\x67\xe7\x63\x9a\x41\xed\xde\x05\x75\x14\xd0\xfa\x15\xda\x8f\x4e\x63\x92\x38\x8e\x0c\x46\xc4\x6f\x98\x1d\x7c\x2c\xf9\xaf\xa8\x41\xed\x44\xcd\x72\x50\x15\x66\x74\x7b\x00\xd2\x0d\x7e\x7a\x47\x74\xbb\xbb\x41\xc0\x47\xcc\x6a\x8d\xf9\x08\x51\x22\x53\x6a\x1c\x51\x7e\x76\xaa\x99\x78\xfa\x13\xcd\xa4\x61\xfc\xfe\xe8\xd8\x99\xb0\xd6\xa2\x1f\x88\x93\x44\xf0\x64\x4f\x24\x25\x1b\x86\x70\xfb\xe7\x65\x6b\xa0\xd2\x44\xd7\x0a\x92\x84\xaa\xa4\x71\x7f\xa2\xea\x8d\xd2\x54\xd7\x46\xd8\xb1\xbb\x52\x73\x59\xda\x52\x9e\x03\xe1\x60\xb4\x25\x6d\xa2\xd9\xbe\xc6\x82\x22\xef\xb7\x52\x29\x84\xb6\x03\xa3\xdd\x7f\xd8\x51\x86\xa0\x51\x99\xac\x19\x10\x19\xb7\xf4\x2d\x8f\x6e\x1d\x59\xb2\x85\xf9\x90\xd4\x1b\xf0\xc2\x1e\x69\xd4\x73\x7c\xe8\x5d\x96\xda\x79\x75\x99\x11\xfd\x65\x29\xab\x40\x0c\xdd\xc2\x06\x0b\xca\xed\x45\xd3\x1e\xe9\xef\x1e\x48\x7f\x03\x21\x9b\x13\x44\x54\x3d\x3b\xf7\x12\xed\xb7\xae\x37\xfb\xbf\x20\xc0\x63\x11\xe1\x60\x30\xe0\xf7\xc2\x32\xdc\x7e\x0e\x7a\x47\x15\x28\x2d\x69\xa5\x2c\x78\x31\xa2\x34\x54\x44\xef\x4c\xa5\x55\x82\x23\xd7\xb0\x95\xa2\xb4\x9b\x66\xfd\x34\x9c\xf3\x07\xf1\x58\x36\x57\xe3\xc1\xf2\x13\x28\xcc\x61\xa1\x9e\xd6\x9f\xfe\x5a\x7f\xfe\xe5\xe4\xe9\x69\xb1\x0a\x33\x73\x9c\x90\xb3\x99\xb1\x80\xa1\x49\xcd\xcc\x38\x20\x49\xf0\x31\x63\xb5\xa2\xfb\xc8\x1e\x87\xb9\xf5\xef\x32\x75\x3d\x8f\xd7\xe5\x06\x65\x2a\xb6\x69\x56\xe6\x29\x91\x85\x4a\x1f\x44\x2a\x2a\xad\x56\x70\x0e\x2f\xe6\x90\x3c\x92\xde\xa5\xf6\x97\xe4\x30\xbf\x68\x57\x48\x83\x39\x34\x8b\x0d\x34\xf3\x69\xf4\xf7\xd7\xbc\x56\xff\xef\x3b\xc2\x0b\xb4\xce\x36\xca\x0c\x20\xe7\x54\x62\xa6\x85\x3c\x80\x16\xae\xbf\xc8\xcc\x5b\x14\x5b\xbb\xd8\x8c\x57\x53\x5b\xd8\xc3\x59\x6b\xe1\x65\xbf\x62\xf4\x66\xb5\x94\x26\x47\x7e\x9e\x93\x72\x54\x5a\x8a\x03\xf4\x41\xba\xc3\x52\xec\x51\x81\xe0\x68\xea\xaa\xac\x99\xa6\x15\xf3\x0e\x0f\xa6\x43\x10\xc6\x8c\x85\x54\x42\x26\xb8\x46\xae\xd5\x44\x06\x21\x57\xb5\x44\xf0\x0c\xba\x80\x07\x49\xaa\x0a\x25\x6c\x85\x84\x1c\x2b\xd7\x9e\x28\x57\x9a\x30\xe6\x0a\x2b\xc7\x0a\x79\x8e\x3c\xa3\xa8\x80\x72\xbb\x16\x1d\x95\xb9\xd2\x48\x72\x3b\x3a\x20\xcf\x85\x9c\xca\x6a\xd3\x80\xc0\xb7\xfa\xea\xb1\x12\xca\x45\x51\x65\x92\x56\x7a\xd4\x54\xf7\x84\x25\x98\xb7\xd6\xb6\xa1\xb6\x33\xc9\x44\x56\x79\x2f\x45\x9d\x55\xb6\x93\x7b\xad\xde\xd1\xb0\x03\x2c\x45\xad\x15\xcd\x31\x9a\xf0\xab\x89\x2c\xb2\x53\x43\xe0\xa7\xd7\x66\x85\x6c\x44\xad\x81\xf0\x43\x6b\xd6\x44\xea\xec\x33\x60\xa0\xee\x77\x89\x16\x80\x78\x0e\xc8\x6d\x2f\x26\x83\xcc\x66\x94\xdf\x03\xd5\x2d\x1e\xb4\x71\xe9\x30\x61\x42\xd3\xba\xc7\x18\x6b\xda\x07\xd3\x2c\xda\xb0\x98\x6d\x4a\x18\xfd\x07\x5d\x82\xab\x96\xd4\x34\x76\x5b\x08\x04\x32\x94\x9a\x50\x3e\x39\x26\x58\x0f\x04\x5e\x7b\x63\x56\xbe\x02\x48\xad\xc3\x6c\xcf\x9b\xda\x5b\x8c\x2a\x3d\x34\x49\x69\x8b\x42\xf8\x48\x95\x9d\xad\x3b\x4b\xa6\x82\x22\x26\x32\xc2\xae\x1d\xcc\x38\xa5\xd7\x2d\xe6\x10\x68\x7a\x8b\x6b\xdf\x07\x51\x4b\x1b\x19\x75\x50\x1a\x4b\x03\x33\x6d\x25\x4d\x1d\x1c\x03\xce\x7f\x58\xe8\x6a\x3d\xf1\xd6\xc2\xb5\x45\x3c\xb7\x3e\xea\x69\x53\xdb\x10\xbc\x5c\x79\xb9\x4b\xb8\x01\xb6\xda\x4c\x23\x5a\x38\xaa\x2e\x49\xfb\x04\xb6\xb9\x3b\xb5\x49\x0a\x75\x5d\x85\x00\x6c\xcb\x5b\x79\xa9\x4a\xb9\x0b\xd4\x94\x10\xdb\x8c\xef\x9e\xde\xf7\x3b\xf1\x00\x94\x6f\x45\x83\x6a\x56\x65\x5b\x37\x53\x1f\xbb\xe6\xb4\xcd\xcf\x46\xfd\xc7\x6e\x65\x90\xa3\x3f\x71\x8e\xa9\xb9\x0f\x19\xce\x0a\xb3\xa2\xe2\x08\xda\xdb\xe3\x35\x66\xbd\x50\x16\x78\x6c\xee\x4c\x65\x97\x7b\xec\xf0\xed\x6a\x56\xc2\xf9\x60\x43\x39\x91\x07\x8b\xfe\xd2\x3e\x36\x18\xfc\xff\x89\x7d\x79\x8f\x52\x99\x4b\x5a\x6f\xd7\xad\xa4\x5c\xb7\x45\xec\x36\xfd\xd7\x0d\x37\x3b\x1c\x44\x0d\x44\x9a\x02\xa3\xbc\x98\x07\xf7\x4d\xab\x7b\x5f\xa4\x76\x2b\x6d\x9a\x48\x77\xff\xcb\xca\x1c\x96\xcd\x22\xa3\x1c\x21\x11\x55\xb6\x8a\x7c\x72\xc8\xca\x7c\x05\x49\xa1\xe1\x85\x77\xfd\x75\x6f\xd2\xcd\xb7\x37\x38\xb7\x54\x9f\xce\x3e\x7b\x14\xfd\x3d\x08\xfe\x15\x7d\x04\x9d\x81\xff\x32\x10\xb9\x2c\x17\xa8\xd3\xe0\xa3\x8c\xbd\xd7\xb6\x4a\xed\x5b\x32\x24\x25\x79\xcc\xb1\xd2\x3b\xf8\x37\x24\x25\xe5\xdd\xff\xeb\x43\x85\x90\x37\xd7\x20\x75\x1e\x70\x9d\x9f\xc3\x13\x14\x12\x2b\x48\xf6\xb0\xf8\xab\x10\xa2\x60\x78\x5a\x08\x46\x78\x71\x2a\x64\xb1\x5e\x7c\x5d\xdd\x99\xa7\xee\xec\x5b\xd5\xc5\x75\x1d\x39\x7a\x57\x07\xa1\x13\x46\xae\x19\xbd\xd8\x1f\x91\xd7\x7c\x40\xf8\x9a\xb0\xe3\x6f\x86\x91\x34\x3f\x92\x64\xdd\x35\xcf\xa5\xf6\x7c\x69\x3f\xf8\x29\x1d\x9f\x18\x8f\x09\x69\xae\x21\x3f\x24\xa3\xc3\x42\x27\x65\xb1\xfc\x82\x7b\x57\x8b\xef\x91\xec\x0f\x02\x23\xe1\x03\x5f\xaf\x16\xb3\xff\x05\x00\x00\xff\xff\x94\x9f\x82\xf7\x2e\x1f\x00\x00")

func dataFishBytes() ([]byte, error) {
	return bindataRead(
		_dataFish,
		"data/fish",
	)
}

func dataFish() (*asset, error) {
	bytes, err := dataFishBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/fish", size: 7982, mode: os.FileMode(420), modTime: time.Unix(1500243560, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dataSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x56\xdf\x6f\xdb\x36\x10\x7e\xd7\x5f\x71\x55\x05\xd4\x2e\xe0\x08\xde\x63\x02\x03\x49\x90\x20\x0d\x90\xc6\x46\x9c\x6e\x18\x8a\x4e\xa0\xa5\xb3\x44\x58\x26\x35\x92\x52\xd2\x35\xfd\xdf\x07\x51\x16\xf5\x8b\xb6\xb3\x35\x7a\x31\xcc\xbb\xfb\xc8\xbb\xfb\xee\x23\xdf\xbf\xf3\x57\x94\xf9\x32\x71\x9c\x22\x1e\x8d\xe1\x87\x03\x00\x10\x12\x89\xe0\x7a\x53\x17\x28\xd3\x0b\xe5\x47\x19\x55\x63\xf3\xaf\xfc\x82\xa0\x88\x83\x72\x19\x5c\xef\xc7\xf9\xe9\x6f\x3f\xdd\x8e\xf9\xec\xcc\xfc\x25\xa1\xa2\x05\x51\x68\x89\xaf\x4d\x47\x31\x22\x3c\x80\xd2\x18\x8f\xe2\x84\x51\x46\xc2\x0d\x89\x51\x5a\x70\x1a\xe3\x51\x9c\x38\xe5\x2b\x92\x5e\x3f\x63\x68\xc1\x69\x8c\x47\x71\xf2\x2c\x16\x24\xb2\x25\xb5\xb3\x1c\x45\xf8\xd8\x8d\x0d\xf9\x76\x4b\x58\x04\x45\x0c\xde\xb9\xb6\xa0\x24\xa1\xe3\xfc\x74\x1c\xd3\x31\xd3\xea\x22\x36\xcd\xa9\xbd\x8b\x18\x52\xca\x36\xc6\xbf\xa9\x89\x89\xa2\x6b\xf8\x0a\x93\x7f\xc0\xf5\x7e\xbf\x7d\x78\xfc\x72\x71\x77\x33\x77\xe1\xdb\x19\xa8\x04\x1b\xc2\x60\x98\x70\xb8\x16\x82\x8b\x53\xb8\xe7\x50\x50\xa1\x72\x92\xc6\x1c\x9e\xb8\xd8\xc8\x8c\x84\x08\x54\x56\xbb\x23\x10\x55\x06\xc3\x96\x6f\x91\x29\x03\x51\xc4\x90\x60\x9a\xb5\x7a\x66\x4c\x02\x55\x2e\x18\x4c\xf5\xc2\x9a\x56\xbc\x8d\xa0\x39\x50\xb0\xb8\x78\xfc\xe4\x4b\x11\xfa\x65\x26\x4e\x87\x6c\xbd\x44\xd8\xe1\x44\x8a\xb8\x45\xbe\x7a\xbf\xaa\xb0\xcf\x19\x17\x0a\x4c\xec\xcc\x1b\x15\xb1\x9e\x94\x25\x2a\x45\x59\x2c\xc1\x3b\x1f\xef\x7c\xcb\x72\x5c\x54\x28\x94\xc5\xad\x93\xda\xb1\xf4\xf9\x67\xde\xa7\xf9\xe7\x6b\xff\xc4\x14\xcf\xdf\x13\x16\x34\x71\xf3\xbb\xab\xc5\x72\x3a\xf3\x16\xcb\xe9\x7e\x8f\x9b\x79\x05\x5f\xfd\x1e\x40\xd2\x5e\xda\xa7\x57\xb0\x9b\xf9\xe5\xed\xbd\xa5\xeb\xf6\xdd\x2e\x6f\xef\x67\x55\x48\xb5\x57\x2a\xf1\x55\x21\xb6\x53\x97\x92\xd5\x69\x43\x35\x6d\xc1\x9a\xa4\xe9\x8a\x84\x9b\xaa\x0b\x7f\xe7\x28\xbe\xd7\x6d\x18\xb7\x4e\xef\x7a\x3d\x7f\x17\x66\xe0\x2a\x91\xe3\xde\x6c\xea\x6a\x75\x9b\x73\xda\xa9\x9e\x25\x23\x7b\x98\x8d\x41\xbb\x6c\x7b\xe4\xad\xf3\xdc\x39\x69\xb0\x5d\xe1\x4f\x75\x4f\xdc\x7d\xd3\x18\x5c\xdd\x2e\x2f\x2e\xef\xae\x83\xc5\xc3\xfc\xf3\xe2\x71\x6f\x62\x25\x51\xdc\x51\x13\x37\x86\x92\x36\x6e\xe7\x88\xdb\x4d\x44\x05\x4c\x32\xfb\x6c\x75\x3d\xaa\x0e\xd7\xca\xd1\x4c\xcd\x2f\x29\xc7\xff\x92\x8c\xde\xc4\xda\x24\xa3\x19\xcb\xab\xda\xf9\xc0\x60\xea\x99\x1a\x8c\x99\xa5\x89\x16\x1e\x3b\x16\x52\xd8\x98\x3d\xe8\x76\x7f\x43\x33\x87\x39\x93\xd8\x12\x8b\xe1\xfc\x5b\xe6\xd8\x36\x61\x56\x59\xe8\x89\x90\xe9\x27\xc9\x15\x1f\x6a\x68\x5d\x3a\xce\x02\xc1\xb9\x9a\x79\x8b\x3f\xae\xb4\xe1\x29\xa1\x29\xea\x81\xeb\xb9\xb8\xf0\x6e\x06\xae\xee\x7b\xc4\x9b\x07\x86\xa6\xc6\x7a\xe8\xde\x92\xbf\x21\x57\xca\x8f\xe1\x53\x60\x5c\x02\x46\xb6\x38\xf3\x46\x21\x51\x87\xa1\xba\xb7\x66\x25\x0d\x5a\x0a\xda\xd4\x9c\x70\x70\xbd\x21\x7e\x95\xc1\x21\x0e\xb7\xf8\xd8\x5c\xb0\x43\xa0\x8e\xff\xee\x1a\xeb\xf2\xd5\xb1\x58\xdf\x83\x4a\xa8\x04\xa9\x04\xcd\xa4\x9e\x82\x94\x48\x05\x19\x51\x49\x79\xfd\x67\x9c\x21\x53\xb0\x16\x7c\xab\x8d\xe5\xfa\x49\xff\x31\xd6\xf4\x6b\xa4\x67\xa0\x5f\x2a\x78\x01\x89\x11\x7c\x90\x2f\xfe\xd7\xbf\xfc\x6f\x1f\xbd\x97\x97\x0f\x55\xc9\x22\xce\xd0\x70\x62\xf7\x48\x31\x6c\xe0\x69\xd4\xa4\xd8\x52\xb4\xd7\x5f\xb7\x3d\xe9\x18\x4e\xeb\x97\x2c\xaa\x26\xd5\xec\x53\xdd\x02\x1c\x62\x54\x30\xc9\x21\xa6\x2a\xc9\x57\x27\x21\xdf\xfa\x37\xa8\x96\x4a\x20\xd9\xfa\x45\xdc\x82\x78\xc0\x94\x93\xa8\xc4\x28\xeb\x23\x13\x4c\x53\xc8\x88\x50\xc0\xd7\xff\x19\xd5\x9c\x59\x68\xd0\xfe\x3d\xd9\x29\x88\x0b\x13\xd2\xcd\x5e\xf3\xa8\xe7\x63\xd5\xc3\x87\x8e\x46\x75\x22\x1c\x2b\xd9\x86\x2e\x6b\x6a\xda\xd6\x3c\x50\x2d\xd2\x3c\x3d\x2c\xc9\x7f\xf2\x1c\x64\xc2\xf3\x34\x02\x99\x61\x48\xd7\xdf\x81\x98\x57\xa7\x4a\x88\xaa\xad\x2b\x04\x7c\xc6\x30\x57\x18\x0d\xf4\xb9\x39\xc0\x61\x7d\xde\xcb\xa7\x37\x22\x54\x41\xd2\xf2\xe1\xfb\xf6\x4d\x7b\x55\x2f\x6c\xc2\xea\xfc\x1b\x00\x00\xff\xff\x58\xdd\xce\x90\x95\x0d\x00\x00")

func dataShBytes() ([]byte, error) {
	return bindataRead(
		_dataSh,
		"data/sh",
	)
}

func dataSh() (*asset, error) {
	bytes, err := dataShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/sh", size: 3477, mode: os.FileMode(420), modTime: time.Unix(1500231182, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dataZsh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x94\xc1\x8a\xdb\x30\x10\x86\xef\x7e\x8a\xa9\x76\xa1\xc9\x21\xec\x13\x14\x5a\xda\x85\x1e\x7a\x2d\x14\x4a\x09\x63\x69\x6c\x8b\x28\x92\xd0\x8c\x9c\xcd\x2e\xfb\xee\xc5\x4a\x9c\x44\x6e\xa0\xb7\xf0\x69\xf4\xcf\x8c\xfe\x3f\x7e\xf8\xf0\xd4\x5a\xff\xf4\xca\x43\xd3\x6c\xb7\x63\xbf\x4d\xe4\x02\x9a\xd5\x1a\xde\x1a\x00\x00\x1a\xd1\x81\x7a\x5c\x8d\xfd\xe9\xe7\x66\xc3\x03\x39\x07\xaf\x3c\xac\x55\xf3\xde\x34\x7a\x88\x07\xb3\xed\xb2\xd7\x62\x83\xe7\x4f\xab\xc7\xb7\x05\xfa\xfd\xf9\xcf\x3b\xa8\x22\x8e\x59\xc2\x16\xb5\xd8\x11\x85\xd4\xfa\xdc\x52\x87\x7d\x74\x54\x6a\x2f\x7d\x27\x86\xc6\xc0\x5c\x5c\x41\x6d\x22\xea\x1d\xf6\xc4\x15\x36\x74\xb7\xda\x10\x4b\x0a\xc7\x8a\x91\xe7\x9c\xea\xb2\x69\xbd\x0a\xf4\x2e\xb4\xe8\x9e\x5f\x48\x57\x78\x20\x17\x2b\x60\xbd\x95\x0a\x38\xeb\x77\x0b\xc0\x75\x05\x93\xe4\x5a\x84\x05\x25\xd7\xeb\x64\x6f\x3d\x0b\x3a\xb7\xa0\xff\xc8\xe7\xd8\x27\x34\xa7\x6d\x1e\x2e\x74\xf3\x0b\xd4\x97\xf3\x7b\x00\x02\x47\xd2\xb6\xb3\x1a\x46\x9b\x24\xa3\xeb\x03\x1c\x42\xda\x71\x44\x4d\xaa\x7e\xe6\x5a\xe3\xdb\xe5\x55\x41\x06\x02\x9d\x53\x22\x2f\xf7\x65\x16\x0e\x2c\x86\x81\x43\xc2\x18\x29\x41\x17\x12\x18\x8a\x20\x03\x0a\x9c\x97\xe4\xa2\x6e\x28\x92\x37\xe4\xb5\x25\x06\xeb\x0b\xbb\xd3\xa9\x5c\x22\x34\x10\x3a\x18\xc9\x9b\x90\xd4\xad\xa5\x75\xdf\xe7\x97\x18\xf8\x34\x3c\xeb\x64\xa3\x9c\xfa\xf2\x10\xb2\x33\xd0\x52\x71\x7e\x43\x66\xee\x37\x6f\x58\x82\xae\xae\xb9\xa8\x55\xbf\x93\x8b\x80\x6d\xc8\x02\xe8\x8f\xd3\xd1\x1e\xbd\x51\xd7\x78\xd4\xe5\x5f\x13\x15\x1f\xfc\x94\x3d\x6c\xdd\x64\xc9\x75\x9d\x09\x4f\xb6\x82\x15\x90\x50\x4d\x61\x6c\x22\x2d\x21\x1d\xd5\x35\x68\xb5\xf4\x8f\xe9\xe2\x7f\xac\x99\x55\x65\xb0\x7c\x2b\x79\xc9\x52\x2d\xf9\xb3\x84\x8c\xef\x0f\x02\x5d\x0a\xfb\x72\x74\xe3\x87\x7c\xe4\x22\x46\x06\x24\xa8\x39\xa5\xe5\x03\x11\xf6\xd1\x50\x07\xcb\xbf\x3a\x8c\x7d\xf3\x37\x00\x00\xff\xff\x27\x7d\x68\x13\x7e\x04\x00\x00")

func dataZshBytes() ([]byte, error) {
	return bindataRead(
		_dataZsh,
		"data/zsh",
	)
}

func dataZsh() (*asset, error) {
	bytes, err := dataZshBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/zsh", size: 1150, mode: os.FileMode(420), modTime: time.Unix(1496592103, 0)}
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
	"data/bash": dataBash,
	"data/fish": dataFish,
	"data/sh": dataSh,
	"data/zsh": dataZsh,
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
		"bash": &bintree{dataBash, map[string]*bintree{}},
		"fish": &bintree{dataFish, map[string]*bintree{}},
		"sh": &bintree{dataSh, map[string]*bintree{}},
		"zsh": &bintree{dataZsh, map[string]*bintree{}},
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

