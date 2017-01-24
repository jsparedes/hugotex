package main

import (
  "bufio"
	"fmt"
	"os"
  "strings"
  "regexp"
  // "path/filepath"
)

func main() {

	f, err := os.Open("config.toml")
	if err != nil {
		fmt.Println("gg") //return 0, err
	}

	// Splits on newlines by default.
	scanner := bufio.NewScanner(f)

	line := 1
	// https://golang.org/pkg/bufio/#Scanner.Scan
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "DefaultContentLanguage") {
			fmt.Println(line) // return line, nil
      fmt.Println(scanner.Text())
      re := regexp.MustCompile(`(".*?")`)
      defaultLanguage := re.FindString(scanner.Text())
      fmt.Println(defaultLanguage)
      // fmt.Printf("%q", rm)
		}

    matched, err2 := regexp.MatchString(`languages\..*`, scanner.Text())
    fmt.Println(matched, err2)
    // if re2 {
    //   fmt.Println(re2.FindString(scanner.Text()))
    // }
		line++
	}

	if err := scanner.Err(); err != nil {
		// Handle the erro
	}

  // Move and name index.html to Lang.tex
  // fileList := []string{}
  // err2 := filepath.Walk("." + string(filepath.Separator), func(path string, f os.FileInfo, err error) error {
  //       fileList = append(fileList, path)
  //       return nil
  // })
  //
  // fmt.Println(err2)
  //
  // for _, file := range fileList {
  //     fmt.Println(file)
  // }
}
