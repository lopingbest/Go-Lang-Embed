package main

import (
	"embed"
	"fmt"
	"io/fs"
	"io/ioutil"
)

//go:embed files/version.txt
var version string

//go:embed files/google.png
var logo []byte

//go:embed files/*.txt
var path embed.FS

func main() {
	fmt.Println(version)
	err := ioutil.WriteFile("logo_new.png", logo, fs.ModePerm)
	if err != nil {
		panic(err)
	}
	//membaca directory
	dirEntries, _ := path.ReadDir("files")
	for _, entry := range dirEntries {
		//kalau bukan directory
		if !entry.IsDir() {
			//mengambil nama
			fmt.Println(entry.Name())
			content, _ := path.ReadFile("files/" + entry.Name())
			fmt.Println("Content:", string(content))
		}
	}

}
