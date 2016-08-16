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
				// C0005793A.png C0005793B.png...
				pngs, err := ioutil.ReadDir(filepath.Join(base, dir.Name(), subdir.Name()))
				if err != nil {
					log.Printf("read dir error: %v", err)
					return
				}
				for _, png := range pngs {
					if !png.IsDir() {
						// C0005793A.png->C0005793a.png
						before := filepath.Join(base, dir.Name(), subdir.Name(), png.Name())
						after := filepath.Join(base, dir.Name(), subdir.Name(), stringfy(png.Name()))
						rename(before, after)
					}
				}
				// R0001D1D4->R0001d1d4
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
// C0005793A.png -> C0005793a.png
func stringfy(in string) string {
	lower := strings.ToLower(in)
	if len(lower) > 0 {
		bytes := []byte(lower)
		return strings.Title(string(bytes[0])) + string(bytes[1:])
	}
	return ""
}
