package main

import (
	"io/ioutil"
	"log"
	"github.com/fatih/color"
	"flag"
)

func main() {
	isFilesNeeded := flag.Bool("f", false, "a bool")
	flag.Parse()
	dirTree(flag.Args()[0], *isFilesNeeded, 0)
}

func dirTree(path string, isFilesNeeded bool, step int) {
	step++
	files, err := ioutil.ReadDir(path)

	if err != nil {
		log.Fatal(err)
	}
	var tabs string
	for i := 1; i < step; i++ {
		tabs = tabs + "│\x20"
	}
	for _, f := range files {
		if f.IsDir() || (!f.IsDir() && isFilesNeeded) {
			if files[len(files)-1].Name() == f.Name() {
				color.New(color.FgCyan).Print(tabs + "└───")
			} else {
				color.New(color.FgCyan).Print(tabs + "├───")
			}
		}

		if f.IsDir() {
			color.Magenta(f.Name())
			dirTree(path + "/" + f.Name(), isFilesNeeded, step)
		} else {
			if isFilesNeeded {
				color.Green(f.Name())
			}
		}
	}
}
