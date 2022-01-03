package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func scanDirectory(path string) {
	fmt.Println(path)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		filePath := filepath.Join(path, file.Name())
		if file.IsDir() == true {
			scanDirectory(filePath)
			if err != nil {
				panic(err)
			}
		} else {
			fmt.Println(filePath)
		}
	}
}

func reportPanic() {
	p := recover()
	if p == nil {
		return
	}
	err, ok := p.(error)
	if ok {
		fmt.Println(err)
	} else {
		panic(p)
	}
}

func main() {
	defer reportPanic()
	// panic("some other issue")
	path := os.Args[1]
	scanDirectory(path)
}
