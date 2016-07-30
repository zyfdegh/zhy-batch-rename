package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	base := "/root/_alllayers"
	_, err := os.Stat(base)
	if err != nil {
		log.Printf("path[%s] not found: %v", base, err)
		return
	}

	// L1 L2...
	dirs, err := ioutil.ReadDir(base)
	if err != nil {
		log.Printf("read dir error: %v", err)
		return
	}

	for _, dir := range dirs {
		// R0001D1D4 R0001D1D5...
		subdirs, err := ioutil.ReadDir(filepath.Join(base, dir.Name()))
		if err != nil {
			log.Printf("read dir error: %v", err)
			return
		}
		for _, subdir := range subdirs {
			if subdir.IsDir() {
				before := filepath.Join(base, dir.Name(), subdir.Name())
				after := filepath.Join(base, dir.Name(), stringfy(subdir.Name()))
				rename(before, after)
			}
		}
	}
	return
}

// use absolute path
func rename(before, after string) {
	err := os.Rename(before, after)
	if err != nil {
		log.Printf("rename error: %v", err)
		return
	}
	fmt.Printf("%s -> %s\n", before, after)
}

// R0001D1D5 -> R0001d1d5
func stringfy(in string) (out string) {
	lower := strings.ToLower(in)
	return strings.Title(lower)
}
