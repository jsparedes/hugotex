package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"log"
	"path/filepath"
)

func main() {

	files, err := ioutil.ReadDir("public/")
	if err != nil {
		log.Fatal(err)
	}

	x := make([]string, 3)
	i := 0
	for _, file := range files {
		if file.Mode().IsRegular() {
			if filepath.Ext(file.Name()) == ".tex" {
				fmt.Println(file.Name())
				x[i] = file.Name()
				i = i + 1
				// x = append(x, file.Name())
			}
		}
	}
	fmt.Println(x)

	for _, file:= range x {
		file_tex := "public/" + file
		fmt.Println(file_tex)

		input, err := ioutil.ReadFile(file_tex)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Removing empty spaces related to "{ " or " }"
		output := bytes.Replace(input, []byte("{ "), []byte("{"), -1)
		output = bytes.Replace(output, []byte(" }"), []byte("}"), -1)
		output = bytes.Replace(output, []byte("&amp;"), []byte("\\&"), -1)
		output = bytes.Replace(output, []byte("&rsquo;"), []byte("'"), -1)

		if err = ioutil.WriteFile(file_tex, output, 0666); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println("=================")
	}
}
