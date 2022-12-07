package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	url := "https://git.sr.ht/~gioverse/skel"

	userinfo = url.User(url)

	// split to parts

	err := filepath.Walk(".", visit)
	if err != nil {
		panic(err)
	}
}

func visit(path string, fi os.FileInfo, err error) error {

	matched, err := filepath.Match("makefile.template", fi.Name())

	if err != nil {
		panic(err)

	}

	if matched {
		/*
			read, err := ioutil.ReadFile(path)
			if err != nil {
				panic(err)
			}
		*/

		read, err := os.ReadFile(path)
		if err != nil {
			fmt.Println("File reading error", err)
			panic(err)
		}

		fmt.Println(string(read))
		fmt.Println(path)

		newContents := strings.Replace(string(read), "old", "new", -1)

		fmt.Println(newContents)

		// remove .template file extension
		pathNew := strings.TrimSuffix(path, filepath.Ext(path))

		fmt.Println(pathNew)

		/*
			err = ioutil.WriteFile(pathNew, []byte(newContents), 0)
			if err != nil {
				panic(err)
			}
		*/

		file, err := os.Create(pathNew)
		if err != nil {
			log.Fatal(err)
			panic(err)
		}
		defer file.Close()

		_, err2 := file.WriteString(newContents)

		if err2 != nil {
			log.Fatal(err2)
			panic(err)
		}

		fmt.Println(file.Name() + "File created successfully")

	}

	return nil
}
