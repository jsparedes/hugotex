package main

import (
  "bufio"
	"fmt"
	"os"
  "strings"
  "regexp"
  // "path/filepath"
  "reflect"
  "os/exec"
  "log"
)

func main() {

	f, err := os.Open("config.toml")
	if err != nil {
		fmt.Println("gg") //return 0, err
	}

	// Splits on newlines by default.
	scanner := bufio.NewScanner(f)

	line := 1

  cont := 0
  lang := make([]string, 0)

  defaultLanguage := ""

	// https://golang.org/pkg/bufio/#Scanner.Scan
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "DefaultContentLanguage") {
			fmt.Println(line) // return line, nil
      fmt.Println(scanner.Text())
      re := regexp.MustCompile(`(".*?")`)
      defaultLanguage = re.FindString(scanner.Text())
      defaultLanguage = defaultLanguage[1:len(defaultLanguage)-1]
      fmt.Println(defaultLanguage)
      // fmt.Printf("%q", rm)
		}

    matched, _ := regexp.MatchString(`\[languages\..+\]`, scanner.Text())
    if matched {
      // lang[cont] = scanner.Text()
      lang = append(lang, scanner.Text())
      cont++
    }
    // if re2 {
    //   fmt.Println(re2.FindString(scanner.Text()))
    // }
		line++
	}

  fmt.Println("================")
  // lg := make([]string, 3)
  lv := len("languages") + 4
  rv := ""

  i := 0
  source := ""
  destiny := ""
  for _, v := range lang {
    long := len(v)
    // lg[i] = v[lv:long-1]
    rv = v[lv:long-1]
    i++
    fmt.Println(reflect.TypeOf(defaultLanguage))
    if (rv == defaultLanguage) {
      fmt.Println(v)
      source = "public/index.html"
      destiny = "public/" + rv + ".tex"
    } else {
      fmt.Println(5)
      source = "public/" + rv + "/index.html"
      destiny = "public/" + rv + ".tex"
    }
    fmt.Println("*****************")
    fmt.Println(source)
    fmt.Println(destiny)
    fmt.Println("*****************")
    cmd := exec.Command("mv", source, destiny)
	  err5 := cmd.Run()
	  if err5 != nil {
		    log.Fatal(err5)
	  }
  }
  // fmt.Println(lg)

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
