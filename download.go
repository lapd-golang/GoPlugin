package main

import (
	"fmt"
	"os"
	"plugin"
)

type Greeter interface {
	Greet()
}

type DownloadIDE interface {
	DownloadFile(u string)
}

func main() {
	fmt.Printf("url %s\n", os.Args)
	// load module
	// 1. open the so file to load the symbols
	plug, err := plugin.Open("./download/download.so")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// 2. look up a symbol (an exported function or variable)
	downloadIDE, err := plug.Lookup("DownloadIDE")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// 3. Assert that loaded symbol is of a desired type

	var download DownloadIDE
	download, ok := downloadIDE.(DownloadIDE)
	if !ok {
		fmt.Printf("----unexpected type from module symbol %s %s\n", ok, downloadIDE)
		os.Exit(1)
	}
	// 4. use the module
	download.DownloadFile(os.Args[1])

}
