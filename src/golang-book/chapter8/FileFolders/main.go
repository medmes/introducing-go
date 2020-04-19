package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

const (
	FILE_NAME     = "test.txt"
	NEW_FILE_NAME = "test2.txt"
)

func main() {
	file, err := os.Open(FILE_NAME)
	if err != nil {
		// handle the error here
		return
	}
	defer file.Close()

	// get the file size
	stat, err := file.Stat()
	if err != nil {
		return
	}

	// read the file
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		return
	}

	str := string(bs)
	fmt.Println(str)

	bs2, err := ioutil.ReadFile(FILE_NAME)
	if err != nil {
		return
	}
	str2 := string(bs2)
	fmt.Println(str2)

	file2, err := os.Create(NEW_FILE_NAME)
	if err != nil {
		// handle the error here
		return
	}
	defer file.Close()

	file2.WriteString("GoodBye Gophers!")

	dir, err := os.Open(".")
	if err != nil {
		return
	}

	defer dir.Close()

	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		return
	}

	for _, fi := range fileInfos {
		fmt.Println(fi.Name())
	}

	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		return nil
	})
}
