package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	ret := 0
	err := run()
	if err != nil {
		fmt.Println(err.Error())
		ret = 1
	}
	os.Exit(ret)
}

func run() error {
	sectionData, err := ioutil.ReadFile("sections/_index")
	if err != nil {
		return fmt.Errorf("Cannot read sections: %s", err)
	}
	lines := strings.Split(string(sectionData), "\n")
	secs := make([]string, 0)
	for _, l := range lines {
		if l != "" {
			secs = append(secs, l)
		}
	}

	layout, err := ioutil.ReadFile("layout")
	if err != nil {
		return fmt.Errorf("Cannot read layouts: %s", err)
	}

	t, err := template.New("cheatsheet").Parse(string(layout))
	if err != nil {
		return fmt.Errorf("Cannot parse layout template: %s", err)
	}

	buff := bytes.NewBuffer([]byte{})
	err = t.Execute(buff, secs)
	if err != nil {
		return fmt.Errorf("Cannot execute layout template: %s", err)
	}
	html := buff.String()

	for _, s := range secs {
		content, err := ioutil.ReadFile("sections/" + s)
		if err != nil {
			return fmt.Errorf("Cannot replace text content for %s: %s", s, err)
		}
		html = strings.Replace(html, "[[ "+s+" ]]", string(content), 1)
	}

	fmt.Println(html)
	return nil
}
