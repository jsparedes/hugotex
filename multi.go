package main

import (
  "bufio"
	"fmt"
	"os"
  "strings"
  "regexp"
  "reflect"
  "os/exec"
  "log"
)

func main() {

	f, err := os.Open("config.toml")
	if err != nil {
		fmt.Println("Error reading config.toml") //return 0, err
	}

	// Splits on newlines by default.
	scanner := bufio.NewScanner(f)

  cont := 0
  lang := make([]string, 0)

  defaultLanguage := ""

	for scanner.Scan() {
    line := scanner.Text()
		if strings.Contains(line, "DefaultContentLanguage") {
      re := regexp.MustCompile(`(".*?")`)
      defaultLanguage = re.FindString(line)
      defaultLanguage = defaultLanguage[1:len(defaultLanguage)-1]
		}

    matched, _ := regexp.MatchString(`\[languages\..+\]`, line)
    if matched {
      lang = append(lang, line)
      cont++
    }
	}

  lv := len("languages") + 4
  rv := ""

  i := 0
  source := ""
  destiny := ""
  for _, v := range lang {
    long := len(v)
    rv = v[lv:long-1]
    i++
    fmt.Println(reflect.TypeOf(defaultLanguage))
    if (rv == defaultLanguage) {
      fmt.Println(v)
      source = "public/index.html"
      destiny = "public/" + rv + ".tex"
    } else {
      source = "public/" + rv + "/index.html"
      destiny = "public/" + rv + ".tex"
    }
    cmd := exec.Command("mv", source, destiny)
	  err5 := cmd.Run()
	  if err5 != nil {
		    log.Fatal(err5)
	  }
  }

	if err := scanner.Err(); err != nil {
		// Handle the erro
	}
}
