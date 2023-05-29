package main

import (
	"flag"
	"io/ioutil"
	"log"
	"strings"
)

// "ls" command create
func main () {
	showAll := flag.Bool("a", false, "all files")
	// パースさせる
	flag.Parse()

	dir := "."
	if len(flag.Args()) > 0 {
		dir = flag.Args()[0]
	}
	items, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range items {
		name := item.Name()
		if item.IsDir() {
			name += "/"
		}
		if strings.HasPrefix(name, ".") {
			if *showAll {
				println(name)
				continue
			}
			continue
		}
		println(item.Name())
	}
}
