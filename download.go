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
	fmt.Printf("url---> %s\n", os.Args)
	task := "Plugin"
	if len(os.Args) >= 2 {
		task = os.Args[1]
	}
	var mod string
	switch task {
	case "download":
		mod = "./download/download.so"
	case "HelloPlugin":
		mod = "./helloworld/greeter.so"
	default:
		fmt.Println("don't speak that language")
		os.Exit(1)
	}

	// load module
	// 1. open the so file to load the symbols
	plug, err := plugin.Open(mod)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if task == "download" {
		downloadFile(os.Args[2], plug)
	} else {
		printHelloWord(plug)
	}

}

func downloadFile(url string, plug *plugin.Plugin) {
	downloadIDE, err := plug.Lookup("DownloadIDE")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var download DownloadIDE
	download, ok := downloadIDE.(DownloadIDE)
	if !ok {
		fmt.Printf("----unexpected type from module symbol %s %s\n", ok, downloadIDE)
		os.Exit(1)
	}
	// 4. use the module
	download.DownloadFile(url)
}

func printHelloWord(plug *plugin.Plugin) {
	symGreeter, err := plug.Lookup("Greeter")
	if err != nil {
		fmt.Println("Error--->%s", err)
		os.Exit(1)
	}
	// 3. Assert that loaded symbol is of a desired type
	// in this case interface type Greeter (defined above)
	var greeter Greeter
	greeter, ok := symGreeter.(Greeter)
	if !ok {
		fmt.Println("unexpected type from module symbol")
		os.Exit(1)
	}
	// 4. use the module
	greeter.Greet()
}
