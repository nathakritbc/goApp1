package main

import "go_app1/app"

func main() {
	var a app.App
	a.CreateConnection()
	a.Routes()
	a.Run()
}
