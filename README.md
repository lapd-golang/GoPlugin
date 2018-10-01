# Go Plugin Example


## A Pluggable Greeting System
The demo in this repository implements a simple greeting and download files system.  Each plugin package (directories `./download` and `./helloworld`) implements code that prints a greeting meesage and donwload files url.  File `./helloPlugin.go` uses the new Go `plugin` package to load the pluggable modules and run the proper method using passed command-line parameters.

For instance, when the program is executed it prints a greeting message 
using the passed parameter to select the plugin to load for the appropriate language.
```
> go run helloPlugin.go hellow
Hi there welcome to Golang
```
Or to download a file url
```
> go run helloPlugin.go download url
Start download
Download at____%s /Users/nhatle/go/src/go-plugin-example/resourcesimage2018-10-01 11:15:59.637373 +0700 +07 m=+0.060956138.jpg

---->>>Done Download url:http://images.fanpop.com/images/image_uploads/New-York-new-york-186897_1024_768.jpg
```
As you can see, the capability of the driver program is dynamically expanded by the plugins allowing it to display a greeting message or download a file without the need to recompile the program.
