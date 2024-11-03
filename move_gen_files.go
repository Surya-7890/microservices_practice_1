package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"sync"
)

const source = "proto/gen/"

func main() {

	cmd := exec.Command("ls")
	cmd.Dir = source

	bytes, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	data := strings.Fields(string(bytes))
	wg := sync.WaitGroup{}
	for _, file := range data {
		if strings.Contains(file, "admin") {
			wg.Add(1)
			go moveFile(file, "cmd/admin/gen/", &wg)
		} else if strings.Contains(file, "user") {
			wg.Add(1)
			go moveFile(file, "cmd/user/gen/", &wg)
		} else if strings.Contains(file, "books") {
			wg.Add(1)
			go moveFile(file, "cmd/books/gen/", &wg)
		}
	}
	wg.Wait()
	// err = os.RemoveAll(source)
	// if err != nil {
	// 	panic(err)
	// }

}

func moveFile(filename, path string, wg *sync.WaitGroup) {
	file, err := os.ReadFile(source + filename)
	fmt.Println("file read: ", source + filename)
	if err != nil {
		panic(err)
	}

	newFile, err := os.OpenFile(path + filename, os.O_CREATE | os.O_RDWR, os.ModeAppend)
	if err != nil {
		panic(err)
	}
	defer newFile.Close()
	err = os.Chmod(newFile.Name(), 0644)
	if err != nil {
		panic(err)
	}

	_, err = newFile.Write(file)
	if err != nil {
		panic(err)
	}
	defer wg.Done()
}