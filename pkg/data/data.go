// Code generated by bindata. DO NOT EDIT.
// sources:
// data/graphql-playground.html

package data


import (
	"fmt"
	"os"
	"strings"
	"time"
)


type gzipAsset struct {
	bytes []byte
	info  gzipFileInfoEx
}

type gzipFileInfoEx interface {
	os.FileInfo
	MD5Checksum() string
}

type gzipBindataFileInfo struct {
	name        string
	size        int64
	mode        os.FileMode
	modTime     time.Time
	md5checksum string
}

func (fi gzipBindataFileInfo) Name() string {
	return fi.name
}
func (fi gzipBindataFileInfo) Size() int64 {
	return fi.size
}
func (fi gzipBindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi gzipBindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi gzipBindataFileInfo) MD5Checksum() string {
	return fi.md5checksum
}
func (fi gzipBindataFileInfo) IsDir() bool {
	return false
}
func (fi gzipBindataFileInfo) Sys() interface{} {
	return nil
}

var _gzipBindataDataGraphqlplaygroundhtml = []byte(
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5c\x69\x6f\xdb\xb8\xd6\xfe\x9e\x5f\xc1\x57\x45\xd1\x14\x88\x68\x92" +
	"\xda\x3d\x76\xf0\x76\xc9\x4d\x8b\xc9\xdc\x66\x16\xb4\x77\xee\x37\x45\xa2\x6c\x26\xb2\xa4\x4a\xf4\x96\xc1\xfc\xf7" +
	"\x0b\x52\x8b\x65\xcb\x8e\x15\xc7\x76\x33\x40\x0c\x24\x92\xc8\xb3\x3e\xe7\x1c\x5a\x24\x2d\xf5\xfe\xef\xe3\x97\x0f" +
	"\x7f\xfc\x79\x7d\x01\x86\x7c\x14\x9e\x9f\x9c\xf4\xaa\x23\x75\xfd\xf3\x13\x00\x7a\x23\xca\x5d\xe0\x0d\xdd\x34\xa3" +
	"\xbc\x3f\xe6\x81\x6a\x83\xce\xa2\x23\x72\x47\xb4\xaf\x4c\x18\x9d\x26\x71\xca\x15\xe0\xc5\x11\xa7\x11\xef\x2b\xe3" +
	"\x8c\xa6\x6a\xe6\xb9\xa1\x7b\x13\xd2\x7e\x14\x9f\x01\x16\x31\xce\xdc\x50\x36\xd2\x3e\x86\xe8\x0c\x8c\x58\xc4\x46" +
	"\xe3\xd1\x52\x93\x3b\x6b\x34\x09\x2a\x37\x54\xc7\x4c\x91\x8a\x39\xe3\x21\x3d\xbf\x4c\xdd\x64\xf8\xeb\x15\xb8\x0e" +
	"\xdd\xf9\x20\x8d\xc7\x91\xdf\xeb\xe4\x3d\x82\x26\x64\xd1\x1d\x48\x69\xd8\x57\x32\x3e\x0f\x69\x36\xa4\x94\x2b\x60" +
	"\x98\xd2\xa0\xaf\x74\x3a\x9e\x1f\xc1\xdb\xcc\xa7\x21\x9b\xa4\x30\xa2\xbc\x13\x25\xa3\xce\x40\x08\xfc\x1e\xaa\x49" +
	"\x25\x50\x4d\xa9\xeb\xf1\xce\xcd\x98\x85\x7e\x27\xe3\x2e\x67\x5e\xc7\xcb\xb2\x0e\x8b\x7c\x3a\x83\x5e\x96\x29\x05" +
	"\x14\x35\x6d\xc3\x38\xe5\xde\x98\x03\xe6\xc5\xd1\xd3\x14\x06\xee\x44\x08\x81\x49\x34\x28\xf5\x64\x5e\xca\x12\x0e" +
	"\xb2\xd4\x7b\xa2\x17\xb7\x59\x67\xc4\x7c\x3f\xa4\x53\x37\xa5\xf0\x36\x53\xce\x7b\x9d\x5c\xb8\x88\x7d\x27\x0f\xfe" +
	"\x49\xef\x26\xf6\xe7\xb9\x62\x01\x22\xe0\xf3\x84\xf6\x15\x4e\x67\x5c\xc0\x20\x83\x01\x64\xe2\x80\xbf\xe4\x29\x00" +
	"\x41\x1c\x71\x35\x70\x47\x2c\x9c\x77\x81\xf2\x25\xa1\x11\xf8\xdd\x8d\x32\xe5\x0c\x64\x6e\x94\xa9\x19\x4d\x59\xf0" +
	"\x53\x41\x1b\x4f\x68\x1a\x84\xf1\xb4\x0b\x86\xcc\xf7\x69\x94\xb7\xff\x7d\x22\x0f\x42\x73\x25\x75\xe4\xa6\x03\x16" +
	"\x75\x01\x2a\x59\x6f\x5c\xef\x2e\x77\xae\x0b\x5e\x61\x8b\xb8\x9a\xbb\xc4\x0d\x17\xde\x7f\x8e\x2a\x31\xea\x94\xde" +
	"\xdc\x31\xae\xba\x22\x9d\x38\x8b\xa3\x2e\x58\xa2\x43\xd0\xc8\x00\x75\x33\xaa\xc6\x63\x0e\x82\x38\x9d\xba\xa9\x9f" +
	"\x95\x3a\x77\xe1\x2a\xcc\xf9\xff\x52\xf3\x1d\x9d\x07\xa9\x3b\xa2\x19\x58\x6b\x60\x90\xc6\xa3\xea\x02\x80\x38\x71" +
	"\x3d\xc6\xe7\x35\xbf\x17\x3e\xf0\xd4\x8d\xb2\x20\x4e\x47\x5d\x20\x4f\x43\x97\xd3\x3f\x4f\x31\x4a\x66\x6f\x6b\xc4" +
	"\xa3\xac\x1d\x61\x0b\xa2\xbf\x8b\x23\x8f\xd7\x59\x88\xdb\x5a\x88\x5a\x99\x87\xb6\xda\x86\x56\x0c\x2b\x91\x7e\x41" +
	"\xf8\x70\x08\x03\xd0\xeb\xc8\x81\xe0\xfc\x64\xcb\x98\x00\x03\xd7\xa7\x5f\xc6\xfc\xa1\xd2\x2b\x49\x5a\x57\x5d\x1b" +
	"\x86\xcd\x05\x27\xb8\xf7\x99\x08\x6a\xeb\x4c\x58\xa5\x6c\x43\xf5\xec\x73\x61\xb5\xda\x5e\xe0\x3d\x00\xbc\xeb\xb3" +
	"\xb8\x5e\x57\x9b\x70\xfe\xe1\x2e\x6c\xc2\xf6\x59\x25\xc0\xfa\x34\x7e\xc1\x77\xdf\xf8\x36\xf3\xd8\x4d\x12\xea\xa6" +
	"\xfb\x1c\x30\xda\x7a\xd3\xc6\x97\x7f\xe0\x50\xf1\x02\xed\xc1\xa0\x6d\x66\xaf\x9c\x9b\x6e\x46\x78\x8d\xb9\x92\xe3" +
	"\x21\x4b\x9b\x04\x0f\x74\xae\x05\x6e\xa3\x56\xbc\x4d\x2b\x7e\x48\x2b\xde\x9a\x70\x2f\x68\xac\xcb\x11\x16\x45\x34" +
	"\xfd\x98\xba\xd3\x1a\x32\xe8\x75\xcd\xc2\x8c\xa7\xf1\x1d\x55\x7d\x37\x1b\xc6\x41\x90\x51\xde\x05\x16\x5a\xf5\xc9" +
	"\xd8\xc6\x82\xf5\x06\x0f\x46\xdb\x98\x08\x46\xdb\x82\xfa\x8f\x34\xbf\x19\x85\x78\xcc\x1f\xef\x86\xf9\x78\x93\xb0" +
	"\x41\xb6\x21\xfa\xa3\x4d\x81\xc3\x4f\xdf\x6e\xef\x26\x8d\x29\x59\x95\xe3\x6a\x9c\xb2\x7c\x99\x25\x99\x89\xbf\x52" +
	"\xca\x52\x95\x6c\x22\xda\x2e\x65\xfb\x30\xb0\x65\x10\x78\x88\xb3\x39\xbd\x2c\x47\x25\x04\x89\x91\x81\x90\x45\xd4" +
	"\x4d\xab\xc9\xa2\x68\x5d\xf9\xac\x9b\x71\xee\x28\xa3\x04\x7c\xf0\xe1\xe3\x97\x7b\xff\x05\xf0\x1c\x2c\x7d\x0f\x80" +
	"\x6f\x90\x51\x65\xf8\xe8\x83\x37\x63\x2f\x80\xe7\x60\x99\x7b\x00\x7c\x83\x8c\x12\x70\xfa\xc9\x1d\xfd\xfa\x02\x78" +
	"\x01\x96\xbd\x8c\x95\xb6\x0b\xe0\x1b\x64\x94\x80\xdf\xcc\x87\x83\xcb\xf1\x0b\xe0\x02\x2c\x0c\xd1\x32\x58\xf8\xf1" +
	"\x80\x6f\x94\x51\x02\x1e\x86\xef\x7e\xbe\x7e\xc1\x3b\xc7\x8a\x3c\x39\xc1\x37\xca\xa8\x12\x7c\x10\x7e\xbe\xfc\xa5" +
	"\x05\xe0\xa6\x9e\xcc\x00\xb1\xb7\x41\xde\x20\x6b\x23\xe9\x59\xc1\x7e\x94\x5b\x95\xbb\x6c\xf6\xdb\xe7\xcf\x2d\x60" +
	"\x77\x0c\xe8\xd8\x06\x42\xc8\xc4\x48\x33\xb0\x61\x26\x33\xa0\x9b\xd0\xc0\x08\x21\x44\x88\x6d\xdb\xd8\xd6\xcc\x6d" +
	"\x51\x79\xac\x94\x3d\xd8\xf1\xcc\x62\x7a\x84\xbb\x21\xef\x5b\xfa\x7e\x74\xd3\x32\xa6\x16\x36\x89\x49\x6c\x6c\x69" +
	"\xb6\x28\x05\x60\x6b\x50\x77\x10\x42\xd8\xb4\x6c\xdd\x34\xb7\xd6\xd9\xe3\x64\x3c\xd9\x86\x67\x16\xcd\x23\xdc\x6a" +
	"\x7d\x8b\xc6\x59\x9b\x60\xca\xd1\x0c\x23\x0c\x1d\xcb\x71\x1c\xc3\x22\x96\xa1\x39\x68\x6b\x45\x6e\x65\xdb\x45\xd3" +
	"\x33\x8b\xd2\x31\xee\xcf\x82\xeb\xef\x41\x8b\x28\x69\x04\x22\xcd\xb1\x09\x36\x89\x6e\x19\x86\x6d\xee\x50\x72\x8f" +
	"\x93\xf1\x64\x1b\x9e\x55\x30\x8f\x72\xef\x47\xfd\xdf\x3e\xfc\xf1\xef\xb6\xd1\xd4\x0c\x83\x60\x53\xd4\x82\x8d\x8c" +
	"\x9d\xbe\x15\x1f\x2d\x66\x1f\x96\x3c\xb3\xb0\x92\x27\xd7\xe8\xd6\x5b\x4c\x76\x71\xf9\xf5\xdb\x62\x0d\xae\xb9\x1f" +
	"\x52\x5b\x54\x73\xd3\xd4\x9d\xd7\xd7\xe1\xd6\xb8\x54\x5f\xd8\x5b\xff\xdb\x04\x80\xa0\xb6\xf2\xc9\xce\x16\x1b\x35" +
	"\x08\xe2\x96\x4c\x6b\xc0\xf8\x01\xda\x1b\x18\xa8\x8c\xd3\x34\x3f\xf3\xe2\x71\xc4\xbb\x00\x9f\x2d\xf6\x6f\x5a\x91" +
	"\x55\x03\x68\x16\x7b\xfe\xec\xc8\xc1\x31\x76\x81\xa7\xc1\xb4\x6b\x70\xf6\xab\xfd\x90\xc1\xb9\x7d\xf7\xdf\xff\x8c" +
	"\xae\x8f\x1c\x1c\x6b\xd9\x51\xbd\x15\x3c\x0d\xa6\x5d\x83\xb3\x5f\xed\x87\x0c\xce\xf0\x77\xfa\x2e\x3d\x76\xe5\x38" +
	"\xbb\xe4\x6e\x83\x69\xd7\xe0\xec\x57\xfb\x41\x87\xb5\xaf\x83\xef\x97\x77\x47\x0d\x0e\x86\xf8\xf1\xf0\xac\x61\xda" +
	"\x2d\x38\xfb\xd6\x7e\xd0\xca\xb9\xf8\xd7\xf7\xf7\xfc\xc8\xc1\xd1\x76\x81\xa7\xc1\xb4\x6b\x70\xf6\xab\xfd\x90\xc1" +
	"\xf1\xef\x2f\x7e\xfe\xf0\xcb\xe3\x82\x83\x1e\x08\x4e\x7d\x63\xbb\xf4\x93\x45\x0d\x57\xcd\xa5\x8f\xd5\xc0\xa7\x79" +
	"\xd7\xd9\x60\x59\x13\x9b\xe3\x2b\x6f\x11\x1a\x16\x05\x2c\x62\x9c\xb6\x8b\x50\x83\xba\x0c\xd4\xc7\x3f\xa3\xeb\x47" +
	"\x7e\xfd\x3c\x3d\x4e\x2b\xb7\x41\xa4\x05\x54\x0d\x96\x9d\xe3\xb4\x4f\xe5\xc7\x8b\xd3\xf0\xf6\xfa\xe2\xdd\xaf\xc7" +
	"\x0e\x94\x85\x96\x3e\x6d\xb0\x6a\xb0\xec\x1c\xa8\x7d\x2a\x3f\x5a\xa0\x5e\x85\xb1\xeb\xb3\x68\xa0\x4e\x53\x61\x6c" +
	"\x5a\x45\x2c\x89\x33\x96\x7b\xef\xde\x64\x71\x38\xe6\xb4\xd4\x32\x65\x3e\x1f\x76\x01\x46\x68\x32\x2d\xdb\x86\x94" +
	"\x0d\x86\x3c\x6f\x1c\x96\x8d\x3e\xcb\x92\x50\x44\xb6\xf4\xe6\x26\x9e\x6d\xec\x0b\x42\xba\xa6\x73\x94\xc9\x8e\x75" +
	"\x8c\x75\x86\x0a\xae\x90\x0d\xa4\xf3\xa3\xac\x0b\x3c\x1a\x71\x9a\xae\x92\xdc\xc4\xb3\x9c\xac\x41\x50\xe8\x5a\xdf" +
	"\xdb\x52\x72\xe2\x7a\x77\x9b\xfa\x6f\xc7\x19\x67\xc1\x5c\x2d\x1e\x5e\xdb\xa8\x7f\x9d\x8c\x6d\xbc\x35\x10\x55\x9f" +
	"\xa5\xd4\xcb\x43\xe7\xc5\xe1\x78\x14\x35\x34\x6c\xa4\x78\xa8\xb7\xda\xab\x8c\x07\x8b\x9f\xc0\x15\xb9\x60\x19\x8b" +
	"\xb5\x9e\x32\x15\xea\x6d\xf9\xa3\x55\xea\x4d\xcc\x79\x3c\xea\x02\x52\xdb\x7d\x6c\x8e\x0c\x1b\x9e\xe2\x78\xcc\xa3" +
	"\x53\x2d\xe8\x4b\x77\x38\x9d\xf1\xe5\x07\xcb\x32\x76\x4f\xbb\x40\x23\x0b\x1b\x65\xf3\xb4\xf0\x8b\xa0\xca\x52\xc1" +
	"\xbb\x3e\x5d\xbc\x38\x8c\xd3\x2e\x48\x07\x37\xee\x29\x31\x8c\x33\xb0\xf8\x87\xa0\xf9\xf6\xb9\x38\xef\x5f\x06\x9f" +
	"\x02\x6f\xd9\xfd\xd2\x4f\xbd\xf4\x73\xe9\x91\x1c\x00\x7a\x3e\x9b\x00\xe6\xf7\x95\x95\x91\xa3\x78\x28\xa7\x97\x4d" +
	"\x06\xc0\x0b\xdd\x2c\x13\x14\x83\x58\x01\x13\x46\xa7\xef\xe3\x59\x5f\x41\x00\x01\x4c\x6c\xf1\xa7\x80\xd9\x28\x8c" +
	"\xb2\xee\x2c\x64\xd1\x5d\x5f\x19\x72\x9e\x74\x3b\x9d\xe9\x74\x0a\xa7\x1a\x8c\xd3\x41\x07\x3b\x8e\xd3\x91\xbd\x85" +
	"\xdc\x07\x9e\xc8\x04\x57\xf1\x20\xae\x3d\x96\x29\x89\x7d\x1a\x64\xe5\x45\xfe\xec\x24\x75\xd3\xcb\xd4\xf5\x19\x8d" +
	"\x78\x6e\xff\x52\x93\x8a\x15\x30\xc3\x7d\x45\x87\xb6\xf9\x5a\x01\x33\xd2\x57\x1c\x13\x12\xfc\x5a\x01\x73\xdc\x57" +
	"\x90\x38\x8a\x36\x07\x9a\xe6\x6b\x65\x21\x59\x3e\xa2\x14\x27\x40\xfc\x53\x65\xd8\xfb\xca\xab\x0b\x84\x90\x4d\x94" +
	"\xbc\xb1\x08\x72\x5f\x81\xb6\x02\xf2\xdf\xcf\x49\x79\xe7\x02\xd4\x38\x69\x29\xaa\x64\xc4\x68\x1d\x6b\xaf\xb3\xec" +
	"\x4d\x05\x43\xa7\x8e\x43\x6f\x50\xe3\x10\x45\x2e\x71\x28\x79\x94\xbc\x9a\xfb\x0a\x26\x16\x74\x4c\xa5\xa8\xe4\xc5" +
	"\xf5\xbc\xaf\x60\x05\x04\x2c\x0c\xfb\xca\x38\x0d\x4f\x5f\xad\x22\xf8\x56\x01\xe9\xac\xaf\xe8\xc2\x3e\x21\xbe\xa6" +
	"\x2d\x71\xf9\x50\x6a\x7b\x1f\xa7\x3e\x4d\x4b\x31\x95\x7b\xe2\x52\x4d\xc7\x21\xed\x2b\x51\x1c\xdd\xd3\x34\x56\x80" +
	"\xdf\x57\x7e\xd1\xa1\x05\x08\xb4\x75\x4f\xc5\xd0\xb0\x01\x52\x09\xb4\x4d\x80\x21\xb1\xf3\x33\x02\x6d\x63\x82\xb1" +
	"\x09\x0d\xcb\x43\xe2\xd6\xc7\x92\x9d\x92\x47\x76\xca\xb3\x61\x41\x21\xfb\x91\x6c\x52\x31\x24\x66\x7e\x46\xa0\xad" +
	"\x7d\x35\xa0\x69\x79\x48\x68\x31\x64\x97\x6c\x5d\xfc\xfb\xa4\x43\xd3\xba\x97\xe6\xa0\x5c\x9a\xed\x69\x10\xeb\x00" +
	"\x01\x03\x9a\x42\x9f\x61\xe4\x67\x06\xb4\x26\x05\x01\x02\x82\x44\x25\xd0\xd0\x65\x9f\x5a\x10\x98\xb6\x10\x67\x7b" +
	"\xaa\x06\xb1\x06\x90\x6c\x96\x54\x6a\x45\x25\xec\xb1\x3f\xa8\x58\x08\x16\xfe\x1a\x06\x40\x40\x6a\xbf\x17\xe8\x0a" +
	"\x38\x57\xd1\x2d\x8a\x2e\xff\x6d\x83\x02\x66\x7d\xc5\xd4\x65\xd0\x44\xb9\x15\x70\x07\x41\x90\xe3\x6a\xea\x40\x33" +
	"\x3d\x55\x87\x3a\x01\x48\xb5\x55\x0d\x1a\xb6\x6a\xab\x76\x96\x9f\x00\xf9\x07\xc4\x05\x10\x17\xf9\x89\x68\x13\x49" +
	"\x3d\x17\x61\x5a\xf7\xf3\xf6\x53\x8c\x50\x32\x3b\x03\xf2\xf0\xf6\xa7\x2d\x96\xe6\x3f\x07\x90\x96\x36\x36\xda\xa5" +
	"\xe1\xcd\x2d\x85\xa6\x23\xb6\x03\x91\x0e\x0c\x04\x0d\xe2\xa9\x04\x12\x55\x13\x21\x85\x8e\x6a\x43\x4b\x03\x04\x3a" +
	"\xba\x8a\x11\x74\x4c\xa0\xe5\xa1\x24\xc0\x86\x16\x51\xa1\x03\x44\xb3\x21\x29\x80\x68\x16\x7c\xd0\x11\xbd\x82\xcc" +
	"\xd1\x65\xbf\x29\xc4\x09\x22\x22\xe4\x99\xd0\x91\xc2\x2c\x49\xa0\xd4\x0a\x77\x6f\x98\x78\x72\x3b\xbd\xc2\x64\x69" +
	"\xb3\x5a\x62\xb2\xba\x77\xd6\x44\x04\x23\x22\xfc\xb0\xa0\x21\x01\x91\x8e\xa9\x16\xc4\xc0\x80\xd8\x90\xf6\xeb\x39" +
	"\x2e\x5a\x9e\xfc\x44\x35\x44\x96\x5a\x10\x13\xb5\x86\x57\x89\x25\x90\x1d\x82\x37\x47\x2c\x07\x47\xcb\x8b\x8f\x68" +
	"\x42\xaa\x29\x69\xc0\x02\xb5\x83\x40\x23\xf7\xa6\xeb\x79\xdd\xdc\x15\x5e\x9b\xe7\x18\x23\x99\xe8\x5a\x91\xe8\xa6" +
	"\xc8\x73\x88\x44\xde\xeb\x50\x17\xae\x18\x56\xde\x90\xb7\x67\x8b\xbc\x87\x88\x78\xa2\xec\x04\x56\x86\x25\xaf\xd5" +
	"\xbc\xf9\x20\x1e\xca\x7d\x5d\xe9\x61\x63\xd7\xb4\x65\xec\x89\x01\xf1\x22\xf4\x8d\x5a\xd0\xf6\x57\x0b\x7a\x51\x0b" +
	"\xc6\xe1\x6a\x21\xdf\x19\x5d\xe0\xb1\xb4\xe3\xd8\x7a\x80\xd0\x6c\xe1\xf1\x62\x80\xa8\xea\x81\x2c\x0a\xc2\xa8\x0a" +
	"\x82\x94\x05\x41\xaa\x82\x20\x6d\x0b\x42\xab\x0a\x42\x3f\x4a\x41\xe4\x5b\x8c\x4a\x31\x95\x2e\x7d\x2e\x26\xd6\xc5" +
	"\x77\xb9\x5e\x35\x88\x2f\x6b\xcf\x4d\xfa\x8a\xbc\x67\x5a\x6a\xbe\x8d\x59\x54\xb5\xcb\xaa\xd1\xc4\x17\x0e\xb1\xa0" +
	"\x11\x6a\x62\x30\xc1\x8e\x2a\x8e\x2a\x76\xb6\x7e\xfb\xc8\x9d\xb5\x43\xd8\xe4\x98\x40\x37\x27\x9a\xad\x6a\xf6\x36" +
	"\x23\xf2\x1d\xa4\xc3\x18\x01\x75\x03\xd8\x3a\x34\x42\xb5\x40\x06\xb4\x44\x26\xdf\x39\x39\x48\xb4\x74\x61\x14\x46" +
	"\x5a\x69\x95\x5a\x58\x05\x5a\xc4\x4b\x6e\x19\x1c\xc2\x2a\x8d\x00\x5b\xff\x2a\x43\xb6\x15\x1a\xb9\x34\x7e\x10\x23" +
	"\xe4\x9d\x93\x18\x26\xc2\x02\x18\xb5\x25\x30\xf9\x8a\xb0\x22\xef\x59\xff\x48\x99\x1b\x0d\x42\xaa\xbe\x97\x13\xd8" +
	"\xad\x86\x4a\xcd\x08\xd8\xfa\xd0\x42\x9b\xcc\x7e\x58\xbb\x5c\xe6\x5c\x51\x7e\x45\x83\xed\x18\xc9\x7c\x30\x00\x31" +
	"\xaf\x84\x01\xd6\x6e\xda\xf3\xc5\xbb\x15\xf5\xbf\x89\xc9\x40\x2b\xfd\x8e\x0d\x6c\xeb\xca\xd4\x00\x31\x5b\xea\xef" +
	"\x75\x8a\x99\x49\xaf\x93\x4d\xca\x53\x31\xcf\x2c\x0c\x12\x13\x6d\xe5\xfc\x2a\x9f\x70\x96\x3c\x59\xe2\x46\x55\xb4" +
	"\xe4\x44\x56\x59\xfb\xba\x1e\x41\x57\x4a\xf7\xd9\x44\xce\x61\xf3\x93\xfa\x6c\x36\x8d\x63\xbe\xf2\x02\x9c\xda\x4b" +
	"\x27\x6e\xdd\x89\x9b\xb7\x16\x53\xbf\x29\x8b\xfc\x78\x0a\x5d\xdf\xbf\x98\xd0\x88\x5f\xb1\x8c\xd3\x88\xa6\xa7\x6f" +
	"\xc4\xac\xf8\xcd\x19\x08\xc6\x91\x5c\x43\x01\xa7\x54\xf4\xbf\x05\x7f\x9d\x54\x4b\x03\x51\xc6\x41\x31\x79\xfe\x56" +
	"\xac\xba\xf5\x81\x1f\x7b\xe3\x11\x8d\x38\x1c\x50\x7e\x11\x52\x71\xfa\x7e\xfe\xd9\xcf\x05\xd6\xa6\xd9\x6f\xaa\xc5" +
	"\x83\x65\x11\x50\x22\x21\xec\x10\x46\x9d\xbe\x29\x9e\x31\x17\xe4\x4b\x8a\x85\x9f\x0f\xa9\x13\xfd\x0b\x1d\xe2\x6a" +
	"\x55\x72\xfd\x95\x27\x35\xf1\x05\xf4\x0b\xe4\x21\x8b\x18\x3f\x15\x12\xce\x6a\xcf\x92\xd1\xc8\x4f\x62\x16\xf1\x2e" +
	"\x78\x53\xbe\x45\xe8\xcd\x59\xd1\xfb\x77\xa1\x37\x3f\x2e\xde\x15\xd4\xeb\xe4\xaf\x08\xea\x75\xf2\xf7\x46\xfd\x2f" +
	"\x00\x00\xff\xff\x14\x62\x43\xc8\x50\x4a\x00\x00")

func gzipBindataDataGraphqlplaygroundhtml() (*gzipAsset, error) {
	bytes := _gzipBindataDataGraphqlplaygroundhtml
	info := gzipBindataFileInfo{
		name: "data/graphql-playground.html",
		size: 19024,
		md5checksum: "",
		mode: os.FileMode(420),
		modTime: time.Unix(1566281766, 0),
	}

	a := &gzipAsset{bytes: bytes, info: info}

	return a, nil
}



// GzipAsset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func GzipAsset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _gzipbindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("GzipAsset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, &os.PathError{Op: "open", Path: name, Err: os.ErrNotExist}
}

// MustGzipAsset is like GzipAsset but panics when GzipAsset would return an error.
// It simplifies safe initialization of global variables.
// nolint: deadcode
func MustGzipAsset(name string) []byte {
	a, err := GzipAsset(name)
	if err != nil {
		panic("asset: GzipAsset(" + name + "): " + err.Error())
	}

	return a
}

// GzipAssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or could not be loaded.
func GzipAssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _gzipbindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("GzipAssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, &os.PathError{Op: "open", Path: name, Err: os.ErrNotExist}
}

// GzipAssetNames returns the names of the assets.
// nolint: deadcode
func GzipAssetNames() []string {
	names := make([]string, 0, len(_gzipbindata))
	for name := range _gzipbindata {
		names = append(names, name)
	}
	return names
}

//
// _gzipbindata is a table, holding each asset generator, mapped to its name.
//
var _gzipbindata = map[string]func() (*gzipAsset, error){
	"data/graphql-playground.html": gzipBindataDataGraphqlplaygroundhtml,
}


// GzipAssetDir returns the file names below a certain
// directory embedded in the file by bindata.
// For example if you run bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then GzipAssetDir("data") would return []string{"foo.txt", "img"}
// GzipAssetDir("data/img") would return []string{"a.png", "b.png"}
// GzipAssetDir("foo.txt") and GzipAssetDir("notexist") would return an error
// GzipAssetDir("") will return []string{"data"}.
func GzipAssetDir(name string) ([]string, error) {
	node := _gzipbintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, &os.PathError{
					Op: "open",
					Path: name,
					Err: os.ErrNotExist,
				}
			}
		}
	}
	if node.Func != nil {
		return nil, &os.PathError{
			Op: "open",
			Path: name,
			Err: os.ErrNotExist,
		}
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}


type gzipBintree struct {
	Func     func() (*gzipAsset, error)
	Children map[string]*gzipBintree
}

var _gzipbintree = &gzipBintree{Func: nil, Children: map[string]*gzipBintree{
	"data": {Func: nil, Children: map[string]*gzipBintree{
		"graphql-playground.html": {Func: gzipBindataDataGraphqlplaygroundhtml, Children: map[string]*gzipBintree{}},
	}},
}}