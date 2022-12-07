package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	// Get the Git details from the user...
	// https://github.com/create-go-app/cli

	// split it out

	fmt.Println("Enter your url here:")

	input := bufio.NewReader(os.Stdin)
	line, _ := input.ReadString('\n')

	// Parse the URL and ensure there are no errors.
	u, err := url.Parse(line)
	if err != nil {
		panic(err)
	}

	// Accessing the scheme is straightforward.
	fmt.Println(u.Scheme)

	// `User` contains all authentication info; call
	// `Username` and `Password` on this for individual
	// values.
	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	p, _ := u.User.Password()
	fmt.Println(p)

	// The `Host` contains both the hostname and the port,
	// if present. Use `SplitHostPort` to extract them.
	fmt.Println(u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println(host)
	fmt.Println(port)

	// Here we extract the `path` and the fragment after
	// the `#`.
	fmt.Println(u.Path)
	fmt.Println(u.Fragment)

	// To get query params in a string of `k=v` format,
	// use `RawQuery`. You can also parse query params
	// into a map. The parsed query param maps are from
	// strings to slices of strings, so index into `[0]`
	// if you only want the first value.
	fmt.Println(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	fmt.Println(m["k"][0])

	data := []Data{
		Data{
			FileType: "makefile",
			Replaces: []Replace{
				Replace{Old: "<<GIT_NAME>>", New: "<<GIT_NEW_NAME>>"},
				Replace{Old: "<<GIT_ORG>>", New: "<<GIT_NEW_ORG>>"},
				Replace{Old: "<<GIT_SERVER>>", New: "<<GIT_NEW_SERVER>>"},
			},
		},
		Data{
			FileType: ".gitignore",
			Replaces: []Replace{
				Replace{Old: "<<GIT_NAME>>", New: "<<GIT_NEW_NAME>>"},
			},
		},
	}

	buf, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("data.json", buf, 0644)
	if err != nil {
		panic(err)
	}

	err = filepath.Walk(".", visit)
	if err != nil {
		panic(err)
	}
}

var (
	OLD string = "<<GIT_NAME>>"
	NEW string = "<<NEW_GIT_NAME>>"
)

type Replace struct {
	Old string `json:"old"`
	New string `json:"new"`
}

type Data struct {
	FileType string    `json:"filetype"`
	Replaces []Replace `json:"replaces"`
}

// visit checks each file.
func visit(path string, fi os.FileInfo, err error) error {

	if err != nil {
		return err
	}

	if !!fi.IsDir() {
		return nil //
	}

	log.Println("visiting: " + path)

	err = checkMakeFile(path, fi, OLD, NEW)
	if err != nil {
		return err
	}
	err = checkMkFile(path, fi, OLD, NEW)
	if err != nil {
		return err
	}

	return nil
}

func replace(path string, old string, new string) (err error) {

	log.Println("matched: " + path)

	read, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	//fmt.Println(string(read))
	//log.Printf("replacing  " + path)

	newContents := strings.Replace(string(read), old, new, -1)

	//fmt.Println(newContents)

	err = ioutil.WriteFile(path, []byte(newContents), 0)
	if err != nil {
		panic(err)
	}
	return nil
}

func checkMkFile(path string, fi os.FileInfo, old string, new string) (err error) {

	matched, err := filepath.Match("*.mk", fi.Name())
	if err != nil {
		panic(err)
	}

	if matched {
		err := replace(path, old, new)

		if err != nil {
			panic(err)
		}

	}

	return nil
}

func checkMakeFile(path string, fi os.FileInfo, old string, new string) (err error) {

	/*
		log.Println("makefile checking filepath: " + path)
		log.Println("old: " + old)
		log.Println("new: " + new)
	*/

	matched, err := filepath.Match("makefile", fi.Name())
	if err != nil {
		panic(err)
	}

	if matched {
		err := replace(path, old, new)
		if err != nil {
			panic(err)
		}
	}

	return nil
}
