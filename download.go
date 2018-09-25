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
	// in this case, variable Greeter
	symGreeter, err := plug.Lookup("DownloadIDE")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// 3. Assert that loaded symbol is of a desired type
	// in this case interface type Greeter (defined above)

	var download DownloadIDE
	download, ok := symGreeter.(DownloadIDE)
	if !ok {
		fmt.Printf("----unexpected type from module symbol %s %s\n", ok, symGreeter)
		os.Exit(1)
	}
	// 4. use the module
	// greeter.Greet()
	download.DownloadFile(os.Args[1])

}
