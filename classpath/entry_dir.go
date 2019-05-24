package classpath

import (
	"path/filepath"
	"io/ioutil"
)

/**
	目录形式的的类路径
 */
type DirEntry struct {
	absDir	string
}

func newDirEntry(path string) *DirEntry {
	absDir, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &DirEntry{absDir}
}

func (self *DirEntry) readClass(className string) ([]byte, Entry, error) {
	filename := filepath.Join(self.absDir, className)
	data, err := ioutil.ReadFile(filename)
	return data, self, err
}

func (self *DirEntry) String() string {
	return self.absDir
}

